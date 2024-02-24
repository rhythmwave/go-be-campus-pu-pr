package scheduler

import (
	"context"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type schedulerService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

type classParticipants struct {
	Id                 string
	MaximumParticipant *uint32
	TotalParticipant   uint32
}

func (s schedulerService) AutoSetLeaveActive() {
	ctx := context.Background()

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		utils.PrintError(*constants.ErrUnknown)
		return
	}

	lecturerData, errs := s.LecturerLeaveRepo.AutoSetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		utils.PrintError(*errs)
		return
	}

	if len(lecturerData) != 0 {
		lecturerIds := []string{}
		for _, v := range lecturerData {
			lecturerIds = append(lecturerIds, v.LecturerId)
		}

		errs = s.LecturerRepo.UpdateStatus(ctx, tx, lecturerIds, appConstants.LecturerStatusCuti)
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		utils.PrintError(*constants.ErrUnknown)
		return
	}
}

func (s schedulerService) AutoSetActiveSemester() {
	ctx := context.Background()

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		utils.PrintError(*constants.ErrUnknown)
		return
	}

	semesterId, errs := s.SemesterRepo.AutoSetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		utils.PrintError(*errs)
		return
	}

	curriculumData, errs := s.CurriculumRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		utils.PrintError(*errs)
		return
	}

	curriculums := make(map[string]string)
	for _, v := range curriculumData {
		curriculums[v.StudyProgramId] = v.Id
	}

	if semesterId != "" {
		errs = s.StudentRepo.UpdateActiveSemesterPackage(ctx, tx, nil)
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}

		previousSemesterData, errs := s.SemesterRepo.GetPreviousSemester(ctx, tx, semesterId)
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}

		activeStudents, errs := s.StudentRepo.GetActive(ctx, tx, previousSemesterData.Id, nil)
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}

		classes, errs := s.ClassRepo.GetActiveBySemesterId(ctx, tx, semesterId)
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}

		creditQuotaData, _, errs := s.CreditQuotaRepo.GetList(ctx, tx, common.PaginationRequest{Page: constants.DefaultPage, Limit: constants.DefaultUnlimited})
		if errs != nil {
			_ = tx.Rollback()
			utils.PrintError(*errs)
			return
		}

		studyPlanData := []models.CreateStudyPlan{}
		studentClassData := []models.CreateStudentClass{}
		subjectParticipantData := make(map[string]map[string]map[uint32][]classParticipants)
		for _, v := range classes {
			if subjectParticipantData[v.StudyProgramId] == nil {
				subjectParticipantData[v.StudyProgramId] = make(map[string]map[uint32][]classParticipants)
			}
			if subjectParticipantData[v.StudyProgramId][v.SubjectId] == nil {
				subjectParticipantData[v.StudyProgramId][v.SubjectId] = make(map[uint32][]classParticipants)
			}
			subjectParticipantData[v.StudyProgramId][v.SubjectId][v.SubjectSemesterPackage] = append(subjectParticipantData[v.StudyProgramId][v.SubjectId][v.SubjectSemesterPackage], classParticipants{
				Id:                 v.Id,
				MaximumParticipant: v.MaximumParticipant,
			})
		}

		for _, v := range activeStudents {
			if curriculums[v.StudyProgramId] == "" {
				continue
			}

			studyPlanId := uuid.New().String()

			studyPlanData = append(studyPlanData, models.CreateStudyPlan{
				Id:              studyPlanId,
				StudentId:       v.Id,
				SemesterId:      semesterId,
				SemesterPackage: v.CurrentSemesterPackage,
				MaximumCredit:   appUtils.GetMaximumCredit(creditQuotaData, v.PreviousSemesterGradePoint),
			})

			for _, w := range subjectParticipantData[v.StudyProgramId] {
				var subjectClassId string
				leastParticipants := ^uint32(0)
				for _, x := range w[v.CurrentSemesterPackage] {
					maximumParticipant := utils.NullUint32Scan(x.MaximumParticipant)
					if x.TotalParticipant < leastParticipants && (x.TotalParticipant < maximumParticipant || maximumParticipant == 0) {
						leastParticipants = x.TotalParticipant
						subjectClassId = x.Id
					}
				}
				if subjectClassId != "" {
					studentClassData = append(studentClassData, models.CreateStudentClass{
						StudyPlanId:         studyPlanId,
						CurriculumId:        curriculums[v.StudyProgramId],
						StudentCurriculumId: v.CurriculumId,
						ClassId:             subjectClassId,
					})
					for l, x := range w[v.CurrentSemesterPackage] {
						if x.Id == subjectClassId {
							x.TotalParticipant++
							w[v.CurrentSemesterPackage][l] = x
						}
					}
				}
			}
		}

		if len(studyPlanData) != 0 {
			errs = s.StudyPlanRepo.BulkCreate(ctx, tx, studyPlanData)
			if errs != nil {
				_ = tx.Rollback()
				utils.PrintError(*errs)
				return
			}
		}
		if len(studentClassData) != 0 {
			errs = s.StudentClassRepo.BulkCreate(ctx, tx, studentClassData)
			if errs != nil {
				_ = tx.Rollback()
				utils.PrintError(*errs)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		utils.PrintError(*constants.ErrUnknown)
		return
	}
}
