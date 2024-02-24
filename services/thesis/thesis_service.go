package thesis

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/sccicitb/pupr-backend/constants"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type thesisService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (t thesisService) GetList(ctx context.Context, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId, status, supervisorLecturerId string) (objects.GetListThesisWithPagination, *constants.ErrorResponse) {
	var result objects.GetListThesisWithPagination

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := t.ThesisRepo.GetList(ctx, tx, pagination, studyProgramId, nimNumber, startSemesterId, status, supervisorLecturerId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var resultData []objects.GetListThesis
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetListThesis{
			Id:                        v.Id,
			Topic:                     v.Topic,
			Title:                     v.Title,
			Status:                    v.Status,
			StudentId:                 v.StudentId,
			StudentName:               v.StudentName,
			StudentNimNumber:          v.StudentNimNumber,
			StudentStatus:             v.StudentStatus,
			StudyProgramId:            v.StudyProgramId,
			StudyProgramName:          v.StudyProgramName,
			DiktiStudyProgramCode:     v.DiktiStudyProgramCode,
			DiktiStudyProgramType:     v.DiktiStudyProgramType,
			StudyLevelShortName:       v.StudyLevelShortName,
			StudentHasThesisStudyPlan: v.StudentHasThesisStudyPlan,
			StartSemesterId:           v.StartSemesterId,
			StartSemesterType:         v.StartSemesterType,
			StartSemesterSchoolYear:   appUtils.GenerateSchoolYear(v.StartSemesterStartYear),
		})
	}

	result = objects.GetListThesisWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (t thesisService) GetDetail(ctx context.Context, id string) (objects.GetDetailThesis, *constants.ErrorResponse) {
	var result objects.GetDetailThesis

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	var resultData models.GetDetailThesis
	if claims.Role == appConstants.AppTypeStudent {
		resultData, errs = t.ThesisRepo.GetNonCancelled(ctx, tx, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		if resultData.Id == "" {
			_ = tx.Rollback()
			return result, nil
		}
	} else {
		resultData, errs = t.ThesisRepo.GetById(ctx, tx, id)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	fileData, errs := t.ThesisRepo.GetFileByThesisId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	supervisorData, errs := t.ThesisRepo.GetSupervisorByThesisId(ctx, tx, resultData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result, errs = t.mapGetDetail(resultData, fileData, supervisorData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (t thesisService) Create(ctx context.Context, data objects.CreateThesis) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	studentId := data.StudentId
	status := data.Status
	startSemesterId := data.StartSemesterId
	startDate := data.StartDate

	studentData, errs := t.StudentRepo.GetDetail(ctx, tx, studentId, "")
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	activeCurriculumData, errs := t.CurriculumRepo.GetActiveByStudyProgramId(ctx, tx, utils.NullStringScan(studentData.StudyProgramId))
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	activeSemesterData, errs := t.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if claims.Role == appConstants.AppTypeStudent {
		studentId = claims.ID
		status = appConstants.ThesisStatusDiajukan
		startSemesterId = activeSemesterData.Id
		startDate = time.Now()
	}

	activeThesisData, errs := t.ThesisRepo.GetNonCancelled(ctx, tx, studentId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if activeThesisData.Id != "" {
		_ = tx.Rollback()
		return appConstants.ErrActiveThesisExists
	}

	thesisId := uuid.New().String()
	createData := models.CreateThesis{
		Id:                        thesisId,
		StudentId:                 studentId,
		Topic:                     data.Topic,
		Status:                    status,
		Title:                     data.Title,
		EnglishTitle:              utils.NewNullString(data.EnglishTitle),
		StartSemesterId:           startSemesterId,
		StartDate:                 startDate,
		Remarks:                   utils.NewNullString(data.Remarks),
		IsJointThesis:             data.IsJointThesis,
		ProposalSeminarDate:       utils.NewNullTime(data.ProposalSeminarDate),
		ProposalCertificateNumber: utils.NewNullString(data.ProposalCertificateNumber),
		ProposalCertificateDate:   utils.NewNullTime(data.ProposalCertificateDate),
	}

	var fileData []models.UpsertThesisFile
	if data.FilePath != "" && data.FilePathType != "" {
		fileData = append(fileData, models.UpsertThesisFile{
			ThesisId:     thesisId,
			FilePath:     data.FilePath,
			FilePathType: data.FilePathType,
			Description:  utils.NewNullString(data.FileDescription),
		})
	}

	firstSupervisorData, errs := t.ThesisSupervisorRoleRepo.GetFirstOrder(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studyPlanData, errs := t.StudyPlanRepo.GetByStudentIdAndSemesterId(ctx, tx, studentId, activeSemesterData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	thesisSubjectData, errs := t.SubjectRepo.GetThesisByCurriculumId(ctx, tx, utils.NullStringScan(studentData.CurriculumId))
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studyPlanId := studyPlanData.Id
	if studyPlanId == "" {
		studyPlanId = uuid.New().String()
		newStudyPlanData := []models.CreateStudyPlan{
			{
				Id:              studyPlanId,
				StudentId:       studentId,
				SemesterId:      activeSemesterData.Id,
				SemesterPackage: studentData.CurrentSemesterPackage,
				MaximumCredit:   thesisSubjectData.TheoryCredit + thesisSubjectData.PracticumCredit + thesisSubjectData.FieldPracticumCredit,
				IsSubmitted:     true,
				IsThesis:        true,
			},
		}
		errs = t.StudyPlanRepo.BulkCreate(ctx, tx, newStudyPlanData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	var supervisorData []models.UpsertThesisSupervisor
	var firstLecturerId string
	for _, v := range data.ThesisSupervisors {
		if firstSupervisorData.Id == v.ThesisSupervisorRoleId {
			firstLecturerId = v.LecturerId
		}
		supervisorData = append(supervisorData, models.UpsertThesisSupervisor{
			ThesisId:               thesisId,
			LecturerId:             v.LecturerId,
			ThesisSupervisorRoleId: v.ThesisSupervisorRoleId,
		})
	}

	errs = t.ThesisRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	classId, errs := t.ClassRepo.UpsertThesisClass(ctx, tx, thesisSubjectData.Id, firstLecturerId, activeSemesterData.Id, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	studentClassData := []models.CreateStudentClass{
		{
			Id:                  uuid.New().String(),
			StudyPlanId:         studyPlanId,
			CurriculumId:        activeCurriculumData.Id,
			StudentCurriculumId: utils.NullStringScan(studentData.CurriculumId),
			ClassId:             classId,
		},
	}
	errs = t.StudentClassRepo.BulkCreate(ctx, tx, studentClassData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(fileData) != 0 {
		errs = t.ThesisRepo.UpsertFile(ctx, tx, fileData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	if len(supervisorData) != 0 {
		errs = t.ThesisRepo.UpsertSupervisor(ctx, tx, supervisorData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) Update(ctx context.Context, data objects.UpdateThesis) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	id := data.Id
	studentId := data.StudentId
	status := data.Status
	startSemesterId := data.StartSemesterId
	startDate := data.StartDate
	proposalSeminarDate := data.ProposalSeminarDate
	proposalCertificateNumber := data.ProposalCertificateNumber
	proposalCertificateDate := data.ProposalCertificateDate

	if claims.Role == appConstants.AppTypeStudent {
		thesisData, errs := t.ThesisRepo.GetNonCancelled(ctx, tx, claims.ID)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		if thesisData.Id == "" {
			_ = tx.Rollback()
			return utils.ErrDataNotFound("thesis")
		}

		id = thesisData.Id
		studentId = thesisData.StudentId
		status = thesisData.Status
		startSemesterId = thesisData.StartSemesterId
		startDate = thesisData.StartDate
		proposalSeminarDate = utils.NullTimeScan(thesisData.ProposalSeminarDate)
		proposalCertificateNumber = utils.NullStringScan(thesisData.ProposalCertificateNumber)
		proposalCertificateDate = utils.NullTimeScan(thesisData.ProposalCertificateDate)
	}

	_, errs = t.ThesisRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	var filePaths []string
	var upsertFileData []models.UpsertThesisFile
	for _, v := range data.Files {
		filePaths = append(filePaths, v.FilePath)
		upsertFileData = append(upsertFileData, models.UpsertThesisFile{
			ThesisId:     id,
			FilePath:     v.FilePath,
			FilePathType: v.FilePathType,
			Description:  utils.NewNullString(v.FileDescription),
		})
	}
	var supervisorLecturerIds []string
	var upsertSupervisorData []models.UpsertThesisSupervisor
	for _, v := range data.ThesisSupervisors {
		supervisorLecturerIds = append(supervisorLecturerIds, v.LecturerId)
		upsertSupervisorData = append(upsertSupervisorData, models.UpsertThesisSupervisor{
			ThesisId:               id,
			LecturerId:             v.LecturerId,
			ThesisSupervisorRoleId: v.ThesisSupervisorRoleId,
		})
	}

	updateData := models.UpdateThesis{
		Id:                        id,
		StudentId:                 studentId,
		Topic:                     data.Topic,
		Status:                    status,
		Title:                     data.Title,
		EnglishTitle:              utils.NewNullString(data.EnglishTitle),
		StartSemesterId:           startSemesterId,
		StartDate:                 startDate,
		Remarks:                   utils.NewNullString(data.Remarks),
		IsJointThesis:             data.IsJointThesis,
		ProposalSeminarDate:       utils.NewNullTime(proposalSeminarDate),
		ProposalCertificateNumber: utils.NewNullString(proposalCertificateNumber),
		ProposalCertificateDate:   utils.NewNullTime(proposalCertificateDate),
	}

	errs = t.ThesisRepo.DeleteFileExcludingPaths(ctx, tx, id, filePaths)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	errs = t.ThesisRepo.DeleteSupervisorExcludingLecturerIds(ctx, tx, id, supervisorLecturerIds)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(upsertFileData) != 0 {
		errs = t.ThesisRepo.UpsertFile(ctx, tx, upsertFileData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	if len(upsertSupervisorData) != 0 && claims.Role == appConstants.AppTypeAdmin {
		errs = t.ThesisRepo.UpsertSupervisor(ctx, tx, upsertSupervisorData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	errs = t.ThesisRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := t.ThesisRepo.GetById(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = t.ThesisRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) RegisterThesisDefense(ctx context.Context, studentId string) *constants.ErrorResponse {
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	thesisData, errs := t.ThesisRepo.GetByStudentIdStatus(ctx, tx, studentId, appConstants.ThesisStatusSedangDikerjakan)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if thesisData.Id == "" {
		_ = tx.Rollback()
		return utils.ErrDataNotFound("thesis")
	}

	activeRequestData, errs := t.ThesisRepo.GetActiveDefenseRequest(ctx, tx, thesisData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if activeRequestData.Id != "" {
		_ = tx.Rollback()
		return appConstants.ErrActiveThesisDefenseRequestExists
	}

	errs = t.ThesisRepo.CreateDefenseRequest(ctx, tx, thesisData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) GetListThesisDefenseRequest(ctx context.Context, pagination common.PaginationRequest, studyProgramId string, nimNumber int64, startSemesterId string) (objects.GetListThesisDefenseRequestWithPagination, *constants.ErrorResponse) {
	var result objects.GetListThesisDefenseRequestWithPagination
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := t.ThesisRepo.GetListDefenseRequest(ctx, tx, pagination, studyProgramId, nimNumber, startSemesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var thesisDefenseIds []string
	for _, v := range modelResult {
		if v.ThesisDefenseId != nil {
			thesisDefenseIds = append(thesisDefenseIds, *v.ThesisDefenseId)
		}
	}

	var examinerData []models.GetThesisDefenseExaminer
	if len(thesisDefenseIds) != 0 {
		examinerData, errs = t.ThesisRepo.GetThesisDefenseExaminerByThesisDefenseIds(ctx, tx, thesisDefenseIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	examinerMap := make(map[string][]objects.GetListThesisDefenseRequestExaminer)
	for _, v := range examinerData {
		examinerMap[v.ThesisDefenseId] = append(examinerMap[v.ThesisDefenseId], objects.GetListThesisDefenseRequestExaminer{
			Id:                     v.Id,
			LecturerId:             v.LecturerId,
			LecturerName:           v.LecturerName,
			LecturerFrontTitle:     v.LecturerFrontTitle,
			LecturerBackDegree:     v.LecturerBackDegree,
			ThesisExaminerRoleId:   v.ThesisExaminerRoleId,
			ThesisExaminerRoleName: v.ThesisExaminerRoleName,
		})
	}

	var resultData []objects.GetListThesisDefenseRequest
	for _, v := range modelResult {
		var examiners []objects.GetListThesisDefenseRequestExaminer
		if v.ThesisDefenseId != nil {
			examiners = examinerMap[*v.ThesisDefenseId]
		}

		resultData = append(resultData, objects.GetListThesisDefenseRequest{
			Id:                           v.Id,
			StudentId:                    v.StudentId,
			StudentName:                  v.StudentName,
			StudentNimNumber:             v.StudentNimNumber,
			StudentStatus:                v.StudentStatus,
			StudyProgramId:               v.StudyProgramId,
			StudyProgramName:             v.StudyProgramName,
			DiktiStudyProgramCode:        v.DiktiStudyProgramCode,
			DiktiStudyProgramType:        v.DiktiStudyProgramType,
			StudyLevelId:                 v.StudyLevelId,
			StudyLevelShortName:          v.StudyLevelShortName,
			ThesisId:                     v.ThesisId,
			ThesisTitle:                  v.ThesisTitle,
			ThesisStatus:                 v.ThesisStatus,
			ThesisDefenseCount:           v.ThesisDefenseCount,
			ThesisDefenseId:              v.ThesisDefenseId,
			ThesisDefensePlanDate:        v.ThesisDefensePlanDate,
			ThesisDefensePlanStartTime:   v.ThesisDefensePlanStartTime,
			ThesisDefensePlanEndTime:     v.ThesisDefensePlanEndTime,
			ThesisDefenseActualDate:      v.ThesisDefenseActualDate,
			ThesisDefenseActualStartTime: v.ThesisDefenseActualStartTime,
			ThesisDefenseActualEndTime:   v.ThesisDefenseActualEndTime,
			ThesisDefenseRoomId:          v.ThesisDefenseRoomId,
			ThesisDefenseRoomName:        v.ThesisDefenseRoomName,
			ThesisDefenseRevision:        v.ThesisDefenseRevision,
			ThesisDefenseIsPassed:        v.ThesisDefenseIsPassed,
			ThesisGradeCode:              v.ThesisGradeCode,
			CreatedAt:                    v.CreatedAt,
			Examiners:                    examiners,
		})
	}

	result = objects.GetListThesisDefenseRequestWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (t thesisService) CreateThesisDefense(ctx context.Context, data objects.CreateThesisDefense) *constants.ErrorResponse {
	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	activeRequestData, errs := t.ThesisRepo.GetActiveDefenseRequest(ctx, tx, data.ThesisId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}
	if activeRequestData.Id == "" {
		_ = tx.Rollback()
		return appConstants.ErrActiveThesisDefenseRequestNotExists
	}

	thesisDefenseId := uuid.New().String()
	createData := models.CreateThesisDefense{
		Id:            thesisDefenseId,
		ThesisId:      data.ThesisId,
		PlanDate:      data.PlanDate,
		PlanStartTime: data.PlanStartTime,
		PlanEndTime:   data.PlanEndTime,
		RoomId:        data.RoomId,
	}

	var examinerData []models.UpsertThesisDefenseExaminer
	for _, v := range data.Examiners {
		examinerData = append(examinerData, models.UpsertThesisDefenseExaminer{
			ThesisDefenseId:      thesisDefenseId,
			LecturerId:           v.LecturerId,
			ThesisExaminerRoleId: v.ThesisExaminerRoleId,
		})
	}

	errs = t.ThesisRepo.CreateDefense(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(examinerData) != 0 {
		errs = t.ThesisRepo.UpsertDefenseExaminer(ctx, tx, examinerData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) UpdateThesisDefense(ctx context.Context, data objects.UpdateThesisDefense) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	defenseData, errs := t.ThesisRepo.GetDefenseById(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateThesisDefense{
		Id:              data.Id,
		PlanDate:        data.PlanDate,
		PlanStartTime:   data.PlanStartTime,
		PlanEndTime:     data.PlanEndTime,
		RoomId:          data.RoomId,
		ActualDate:      utils.NewNullTime(data.ActualDate),
		ActualStartTime: utils.NewNullInt32(int32(data.ActualStartTime)),
		ActualEndTime:   utils.NewNullInt32(int32(data.ActualEndTime)),
		IsPassed:        data.IsPassed,
		Revision:        utils.NewNullString(data.Revision),
	}
	if defenseData.ActualDate != nil {
		updateData.PlanDate = defenseData.PlanDate
		updateData.PlanStartTime = defenseData.PlanStartTime
		updateData.PlanEndTime = defenseData.PlanEndTime
		updateData.RoomId = defenseData.RoomId
	}

	errs = t.ThesisRepo.UpdateDefense(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if data.IsPassed {
		gradeTypeData, errs := t.GradeTypeRepo.GetByGradeCode(ctx, tx, defenseData.StudyLevelId, data.GradeCode)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
		activeSemesterData, errs := t.SemesterRepo.GetActive(ctx, tx)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}

		classData, errs := t.ClassRepo.GetThesisClass(ctx, tx, defenseData.StudentId, activeSemesterData.Id)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}

		finishData := models.FinishThesisDefense{
			Id:               defenseData.ThesisId,
			FinishSemesterId: activeSemesterData.Id,
			GradePoint:       gradeTypeData.GradePoint,
			GradeCode:        data.GradeCode,
			Status:           appConstants.ThesisStatusBerhasilDiselesaikan,
		}
		errs = t.ThesisRepo.FinishDefense(ctx, tx, finishData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}

		classGradeComponentData, errs := t.ClassGradeComponentRepo.GetByClassId(ctx, tx, classData.Id)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}

		var gradeData []models.GradeStudentClass

		var adminId sql.NullString
		var lecturerId sql.NullString
		if claims.Role == appConstants.AppTypeAdmin {
			adminId = utils.NewNullString(claims.ID)
		}
		if claims.Role == appConstants.AppTypeLecturer {
			lecturerId = utils.NewNullString(claims.ID)
		}

		for _, v := range classGradeComponentData {
			gradeData = append(gradeData, models.GradeStudentClass{
				ClassId:               classData.Id,
				StudentId:             defenseData.StudentId,
				ClassGradeComponentId: v.Id,
				InitialGrade:          gradeTypeData.MinimumGrade,
				GradedByAdminId:       adminId,
				GradedByLecturerId:    lecturerId,
			})
		}

		errs = t.StudentClassRepo.BulkGradeStudentClass(ctx, tx, gradeData)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (t thesisService) GetThesisSupervisorLog(ctx context.Context, pagination common.PaginationRequest, idNationalLecturer, semesterId string) (objects.GetThesisSupervisorLogWithPagination, *constants.ErrorResponse) {
	var result objects.GetThesisSupervisorLogWithPagination

	tx, err := t.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := t.LecturerRepo.GetList(ctx, tx, pagination, objects.GetLecturerRequest{IdNationalLecturer: idNationalLecturer})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var lecturerIds []string
	for _, v := range modelResult {
		lecturerIds = append(lecturerIds, v.Id)
	}

	activeSemesterData, errs := t.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	var logData []models.GetThesisSupervisorLog
	if activeSemesterData.Id == semesterId || semesterId == "" {
		logData, errs = t.ThesisRepo.GetActiveSemesterThesisSupervisorLog(ctx, tx, lecturerIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	} else {
		logData, errs = t.ThesisRepo.GetThesisSupervisorLog(ctx, tx, semesterId, lecturerIds)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	result = objects.GetThesisSupervisorLogWithPagination{
		Pagination: paginationResult,
		Data:       mapGetThesisSupervisorLog(modelResult, logData),
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
