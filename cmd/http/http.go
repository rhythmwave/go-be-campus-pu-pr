package http

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	v8 "github.com/go-redis/redis/v8"
	"github.com/pressly/goose/v3"
	"github.com/rs/cors"
	config "github.com/sccicitb/pupr-backend/config"
	appConstants "github.com/sccicitb/pupr-backend/constants"
	middleware "github.com/sccicitb/pupr-backend/infra/base_middleware"
	"github.com/sccicitb/pupr-backend/infra/context/handler"
	infra "github.com/sccicitb/pupr-backend/infra/context/infra"
	repository "github.com/sccicitb/pupr-backend/infra/context/repository"
	service "github.com/sccicitb/pupr-backend/infra/context/service"
	db "github.com/sccicitb/pupr-backend/infra/db"
	fbs "github.com/sccicitb/pupr-backend/infra/fbs"
	"github.com/sccicitb/pupr-backend/infra/file"
	"github.com/sccicitb/pupr-backend/infra/mail"
	otp "github.com/sccicitb/pupr-backend/infra/otp"
	redis "github.com/sccicitb/pupr-backend/infra/redis"
	notification "github.com/sccicitb/pupr-backend/notification"
	"github.com/sccicitb/pupr-backend/routers"
	logrus "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"

	appMiddleware "github.com/sccicitb/pupr-backend/infra/middleware"

	fileService "github.com/sccicitb/pupr-backend/services/file"

	adminAcademicGuidanceHandler "github.com/sccicitb/pupr-backend/handlers/admin/academic_guidance"
	adminAccreditationHandler "github.com/sccicitb/pupr-backend/handlers/admin/accreditation"
	adminAnnouncementHandler "github.com/sccicitb/pupr-backend/handlers/admin/announcement"
	adminAuthenticationHandler "github.com/sccicitb/pupr-backend/handlers/admin/authentication"
	adminBuildingHandler "github.com/sccicitb/pupr-backend/handlers/admin/building"
	adminClassHandler "github.com/sccicitb/pupr-backend/handlers/admin/class"
	adminClassDiscussionHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_discussion"
	adminClassEventHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_event"
	adminClassExamHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_exam"
	adminClassGradeComponentHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_grade_component"
	adminClassMaterialHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_material"
	adminClassWorkHandler "github.com/sccicitb/pupr-backend/handlers/admin/class_work"
	adminCreditQuotaHandler "github.com/sccicitb/pupr-backend/handlers/admin/credit_quota"
	adminCurriculumHandler "github.com/sccicitb/pupr-backend/handlers/admin/curriculum"
	adminDiktiStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/admin/dikti_study_program"
	adminDocumentActionHandler "github.com/sccicitb/pupr-backend/handlers/admin/document_action"
	adminDocumentTypeHandler "github.com/sccicitb/pupr-backend/handlers/admin/document_type"
	adminExamSupervisorHandler "github.com/sccicitb/pupr-backend/handlers/admin/exam_supervisor"
	adminExamSupervisorRoleHandler "github.com/sccicitb/pupr-backend/handlers/admin/exam_supervisor_role"
	adminExpertiseGroupHandler "github.com/sccicitb/pupr-backend/handlers/admin/expertise_group"
	adminFacultyHandler "github.com/sccicitb/pupr-backend/handlers/admin/faculty"
	adminGradeComponentHandler "github.com/sccicitb/pupr-backend/handlers/admin/grade_component"
	adminGradeTypeHandler "github.com/sccicitb/pupr-backend/handlers/admin/grade_type"
	adminGraduationHandler "github.com/sccicitb/pupr-backend/handlers/admin/graduation"
	adminGraduationPredicateHandler "github.com/sccicitb/pupr-backend/handlers/admin/graduation_predicate"
	adminGraduationSessionHandler "github.com/sccicitb/pupr-backend/handlers/admin/graduation_session"
	adminLearningAchievementHandler "github.com/sccicitb/pupr-backend/handlers/admin/learning_achievement"
	adminLearningAchievementCategoryHandler "github.com/sccicitb/pupr-backend/handlers/admin/learning_achievement_category"
	adminLectureHandler "github.com/sccicitb/pupr-backend/handlers/admin/lecture"
	adminLecturerHandler "github.com/sccicitb/pupr-backend/handlers/admin/lecturer"
	adminLecturerLeaveHandler "github.com/sccicitb/pupr-backend/handlers/admin/lecturer_leave"
	adminLecturerMutationHandler "github.com/sccicitb/pupr-backend/handlers/admin/lecturer_mutation"
	adminLecturerResignationHandler "github.com/sccicitb/pupr-backend/handlers/admin/lecturer_resignation"
	adminLessonPlanHandler "github.com/sccicitb/pupr-backend/handlers/admin/lesson_plan"
	adminMajorHandler "github.com/sccicitb/pupr-backend/handlers/admin/major"
	adminOfficerHandler "github.com/sccicitb/pupr-backend/handlers/admin/officer"
	adminOfficerActionHandler "github.com/sccicitb/pupr-backend/handlers/admin/officer_action"
	adminReportHandler "github.com/sccicitb/pupr-backend/handlers/admin/report"
	adminRoomHandler "github.com/sccicitb/pupr-backend/handlers/admin/room"
	adminSemesterHandler "github.com/sccicitb/pupr-backend/handlers/admin/semester"
	adminSharedFileHandler "github.com/sccicitb/pupr-backend/handlers/admin/shared_file"
	adminStudentHandler "github.com/sccicitb/pupr-backend/handlers/admin/student"
	adminStudentActivityHandler "github.com/sccicitb/pupr-backend/handlers/admin/student_activity"
	adminStudentClassHandler "github.com/sccicitb/pupr-backend/handlers/admin/student_class"
	adminStudentLeaveHandler "github.com/sccicitb/pupr-backend/handlers/admin/student_leave"
	adminStudentSkpiHandler "github.com/sccicitb/pupr-backend/handlers/admin/student_skpi"
	adminStudyLevelHandler "github.com/sccicitb/pupr-backend/handlers/admin/study_level"
	adminStudyPlanHandler "github.com/sccicitb/pupr-backend/handlers/admin/study_plan"
	adminStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/admin/study_program"
	adminSubjectHandler "github.com/sccicitb/pupr-backend/handlers/admin/subject"
	adminSubjectCategoryHandler "github.com/sccicitb/pupr-backend/handlers/admin/subject_category"
	adminSubjectGradeComponentHandler "github.com/sccicitb/pupr-backend/handlers/admin/subject_grade_component"
	adminThesisHandler "github.com/sccicitb/pupr-backend/handlers/admin/thesis"
	adminThesisExaminerRoleHandler "github.com/sccicitb/pupr-backend/handlers/admin/thesis_examiner_role"
	adminThesisSupervisorRoleHandler "github.com/sccicitb/pupr-backend/handlers/admin/thesis_supervisor_role"
	adminTranscriptHandler "github.com/sccicitb/pupr-backend/handlers/admin/transcript"
	adminYudiciumHandler "github.com/sccicitb/pupr-backend/handlers/admin/yudicium"
	adminYudiciumSessionHandler "github.com/sccicitb/pupr-backend/handlers/admin/yudicium_session"
	adminYudiciumTermHandler "github.com/sccicitb/pupr-backend/handlers/admin/yudicium_term"

	careerStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/career/study_program"

	generalAuthHandler "github.com/sccicitb/pupr-backend/handlers/general/auth"
	generalFileHandler "github.com/sccicitb/pupr-backend/handlers/general/file"
	generalLocationHandler "github.com/sccicitb/pupr-backend/handlers/general/location"

	lecturerAcademicGuidanceHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/academic_guidance"
	lecturerAnnouncementHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/announcement"
	lecturerClassHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class"
	lecturerClassAnnouncementHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_announcement"
	lecturerClassDiscussionHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_discussion"
	lecturerClassEventHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_event"
	lecturerClassExamHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_exam"
	lecturerClassGradeComponentHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_grade_component"
	lecturerClassMaterialHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_material"
	lecturerClassWorkHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/class_work"
	lecturerGeneralHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/general"
	lecturerLectureHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/lecture"
	lecturerSemesterHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/semester"
	lecturerSharedFileHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/shared_file"
	lecturerStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/lecturer/study_program"

	pmbStudentHandler "github.com/sccicitb/pupr-backend/handlers/pmb/student"

	rootAdminHandler "github.com/sccicitb/pupr-backend/handlers/root/admin"
	rootAdminActivityLogHandler "github.com/sccicitb/pupr-backend/handlers/root/admin_activity_log"
	rootDiktiStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/root/dikti_study_program"
	rootFacultyHandler "github.com/sccicitb/pupr-backend/handlers/root/faculty"
	rootLecturerHandler "github.com/sccicitb/pupr-backend/handlers/root/lecturer"
	rootLecturerStudentActivityLogHandler "github.com/sccicitb/pupr-backend/handlers/root/lecturer_student_activity_log"
	rootMajorHandler "github.com/sccicitb/pupr-backend/handlers/root/major"
	rootPermissionHandler "github.com/sccicitb/pupr-backend/handlers/root/permission"
	rootRoleHandler "github.com/sccicitb/pupr-backend/handlers/root/role"
	rootStudyProgramHandler "github.com/sccicitb/pupr-backend/handlers/root/study_program"

	excelHandler "github.com/sccicitb/pupr-backend/handlers/admin/excel"
	studentAcademicGuidanceHandler "github.com/sccicitb/pupr-backend/handlers/student/academic_guidance"
	studentAnnouncementHandler "github.com/sccicitb/pupr-backend/handlers/student/announcement"
	studentClassHandler "github.com/sccicitb/pupr-backend/handlers/student/class"
	studentClassAnnouncementHandler "github.com/sccicitb/pupr-backend/handlers/student/class_announcement"
	studentClassDiscussionHandler "github.com/sccicitb/pupr-backend/handlers/student/class_discussion"
	studentClassEventHandler "github.com/sccicitb/pupr-backend/handlers/student/class_event"
	studentClassExamHandler "github.com/sccicitb/pupr-backend/handlers/student/class_exam"
	studentClassMaterialHandler "github.com/sccicitb/pupr-backend/handlers/student/class_material"
	studentClassWorkHandler "github.com/sccicitb/pupr-backend/handlers/student/class_work"
	studentGeneralHandler "github.com/sccicitb/pupr-backend/handlers/student/general"
	studentGradeTypeHandler "github.com/sccicitb/pupr-backend/handlers/student/grade_type"
	studentLectureHandler "github.com/sccicitb/pupr-backend/handlers/student/lecture"
	studentSemesterHandler "github.com/sccicitb/pupr-backend/handlers/student/semester"
	studentSharedFileHandler "github.com/sccicitb/pupr-backend/handlers/student/shared_file"
	studentStudentLeaveHandler "github.com/sccicitb/pupr-backend/handlers/student/student_leave"
	studentStudentSkpiHandler "github.com/sccicitb/pupr-backend/handlers/student/student_skpi"
	studentStudyPlanHandler "github.com/sccicitb/pupr-backend/handlers/student/study_plan"
	studentThesisHandler "github.com/sccicitb/pupr-backend/handlers/student/thesis"
	studentTranscriptHandler "github.com/sccicitb/pupr-backend/handlers/student/transcript"

	academicGuidanceRepository "github.com/sccicitb/pupr-backend/data/repositories/academic_guidance"
	academicGuidanceService "github.com/sccicitb/pupr-backend/services/academic_guidance"

	accreditationRepository "github.com/sccicitb/pupr-backend/data/repositories/accreditation"
	accreditationService "github.com/sccicitb/pupr-backend/services/accreditation"

	adminActivityLogRepository "github.com/sccicitb/pupr-backend/data/repositories/admin_activity_log"
	adminActivityLogService "github.com/sccicitb/pupr-backend/services/admin_activity_log"

	adminRepository "github.com/sccicitb/pupr-backend/data/repositories/admin"
	adminService "github.com/sccicitb/pupr-backend/services/admin"

	announcementRepository "github.com/sccicitb/pupr-backend/data/repositories/announcement"
	announcementService "github.com/sccicitb/pupr-backend/services/announcement"

	authenticationRepository "github.com/sccicitb/pupr-backend/data/repositories/authentication"
	authenticationService "github.com/sccicitb/pupr-backend/services/authentication"

	buildingRepository "github.com/sccicitb/pupr-backend/data/repositories/building"
	buildingService "github.com/sccicitb/pupr-backend/services/building"

	classRepository "github.com/sccicitb/pupr-backend/data/repositories/class"
	classService "github.com/sccicitb/pupr-backend/services/class"

	classAnnouncementRepository "github.com/sccicitb/pupr-backend/data/repositories/class_announcement"
	classAnnouncementService "github.com/sccicitb/pupr-backend/services/class_announcement"

	classDiscussionRepository "github.com/sccicitb/pupr-backend/data/repositories/class_discussion"
	classDiscussionService "github.com/sccicitb/pupr-backend/services/class_discussion"

	classEventRepository "github.com/sccicitb/pupr-backend/data/repositories/class_event"
	classEventService "github.com/sccicitb/pupr-backend/services/class_event"

	classExamRepository "github.com/sccicitb/pupr-backend/data/repositories/class_exam"
	classExamService "github.com/sccicitb/pupr-backend/services/class_exam"

	classGradeComponentRepository "github.com/sccicitb/pupr-backend/data/repositories/class_grade_component"
	classGradeComponentService "github.com/sccicitb/pupr-backend/services/class_grade_component"

	classLecturerRepository "github.com/sccicitb/pupr-backend/data/repositories/class_lecturer"

	classMaterialRepository "github.com/sccicitb/pupr-backend/data/repositories/class_material"
	classMaterialService "github.com/sccicitb/pupr-backend/services/class_material"

	classWorkRepository "github.com/sccicitb/pupr-backend/data/repositories/class_work"
	classWorkService "github.com/sccicitb/pupr-backend/services/class_work"

	creditQuotaRepository "github.com/sccicitb/pupr-backend/data/repositories/credit_quota"
	creditQuotaService "github.com/sccicitb/pupr-backend/services/credit_quota"

	curriculumRepository "github.com/sccicitb/pupr-backend/data/repositories/curriculum"
	curriculumService "github.com/sccicitb/pupr-backend/services/curriculum"

	diktiStudyProgramRepository "github.com/sccicitb/pupr-backend/data/repositories/dikti_study_program"
	diktiStudyProgramService "github.com/sccicitb/pupr-backend/services/dikti_study_program"

	documentActionRepository "github.com/sccicitb/pupr-backend/data/repositories/document_action"
	documentActionService "github.com/sccicitb/pupr-backend/services/document_action"

	documentTypeRepository "github.com/sccicitb/pupr-backend/data/repositories/document_type"
	documentTypeService "github.com/sccicitb/pupr-backend/services/document_type"

	examSupervisorRepository "github.com/sccicitb/pupr-backend/data/repositories/exam_supervisor"
	examSupervisorService "github.com/sccicitb/pupr-backend/services/exam_supervisor"

	examSupervisorRoleRepository "github.com/sccicitb/pupr-backend/data/repositories/exam_supervisor_role"
	examSupervisorRoleService "github.com/sccicitb/pupr-backend/services/exam_supervisor_role"

	expertiseGroupRepository "github.com/sccicitb/pupr-backend/data/repositories/expertise_group"
	expertiseGroupService "github.com/sccicitb/pupr-backend/services/expertise_group"

	facultyRepository "github.com/sccicitb/pupr-backend/data/repositories/faculty"
	facultyService "github.com/sccicitb/pupr-backend/services/faculty"

	gradeComponentRepository "github.com/sccicitb/pupr-backend/data/repositories/grade_component"
	gradeComponentService "github.com/sccicitb/pupr-backend/services/grade_component"

	gradeTypeRepository "github.com/sccicitb/pupr-backend/data/repositories/grade_type"
	gradeTypeService "github.com/sccicitb/pupr-backend/services/grade_type"

	graduationService "github.com/sccicitb/pupr-backend/services/graduation"

	graduationPredicateRepository "github.com/sccicitb/pupr-backend/data/repositories/graduation_predicate"
	graduationPredicateService "github.com/sccicitb/pupr-backend/services/graduation_predicate"

	graduationSessionRepository "github.com/sccicitb/pupr-backend/data/repositories/graduation_session"
	graduationSessionService "github.com/sccicitb/pupr-backend/services/graduation_session"

	graduationStudentRepository "github.com/sccicitb/pupr-backend/data/repositories/graduation_student"

	learningAchievementRepository "github.com/sccicitb/pupr-backend/data/repositories/learning_achievement"
	learningAchievementService "github.com/sccicitb/pupr-backend/services/learning_achievement"

	learningAchievementCategoryRepository "github.com/sccicitb/pupr-backend/data/repositories/learning_achievement_category"
	learningAchievementCategoryService "github.com/sccicitb/pupr-backend/services/learning_achievement_category"

	lectureRepository "github.com/sccicitb/pupr-backend/data/repositories/lecture"
	lectureService "github.com/sccicitb/pupr-backend/services/lecture"

	lecturerRepository "github.com/sccicitb/pupr-backend/data/repositories/lecturer"
	lecturerService "github.com/sccicitb/pupr-backend/services/lecturer"

	lecturerLeaveRepository "github.com/sccicitb/pupr-backend/data/repositories/lecturer_leave"
	lecturerLeaveService "github.com/sccicitb/pupr-backend/services/lecturer_leave"

	lecturerMutationRepository "github.com/sccicitb/pupr-backend/data/repositories/lecturer_mutation"
	lecturerMutationService "github.com/sccicitb/pupr-backend/services/lecturer_mutation"

	lecturerResignationRepository "github.com/sccicitb/pupr-backend/data/repositories/lecturer_resignation"
	lecturerResignationService "github.com/sccicitb/pupr-backend/services/lecturer_resignation"

	lecturerStudentActivityLogRepository "github.com/sccicitb/pupr-backend/data/repositories/lecturer_student_activity_log"
	lecturerStudentActivityLogService "github.com/sccicitb/pupr-backend/services/lecturer_student_activity_log"

	lessonPlanRepository "github.com/sccicitb/pupr-backend/data/repositories/lesson_plan"
	lessonPlanService "github.com/sccicitb/pupr-backend/services/lesson_plan"

	locationRepository "github.com/sccicitb/pupr-backend/data/repositories/location"
	locationService "github.com/sccicitb/pupr-backend/services/location"

	majorRepository "github.com/sccicitb/pupr-backend/data/repositories/major"
	majorService "github.com/sccicitb/pupr-backend/services/major"

	officerRepository "github.com/sccicitb/pupr-backend/data/repositories/officer"
	officerService "github.com/sccicitb/pupr-backend/services/officer"

	officerActionRepository "github.com/sccicitb/pupr-backend/data/repositories/officer_action"
	officerActionService "github.com/sccicitb/pupr-backend/services/officer_action"

	permissionRepository "github.com/sccicitb/pupr-backend/data/repositories/permission"
	permissionService "github.com/sccicitb/pupr-backend/services/permission"

	reportRepository "github.com/sccicitb/pupr-backend/data/repositories/report"
	reportService "github.com/sccicitb/pupr-backend/services/report"

	roleRepository "github.com/sccicitb/pupr-backend/data/repositories/role"
	roleService "github.com/sccicitb/pupr-backend/services/role"

	roomRepository "github.com/sccicitb/pupr-backend/data/repositories/room"
	roomService "github.com/sccicitb/pupr-backend/services/room"

	schedulerService "github.com/sccicitb/pupr-backend/services/scheduler"

	semesterRepository "github.com/sccicitb/pupr-backend/data/repositories/semester"
	semesterService "github.com/sccicitb/pupr-backend/services/semester"

	sharedFileRepository "github.com/sccicitb/pupr-backend/data/repositories/shared_file"
	sharedFileService "github.com/sccicitb/pupr-backend/services/shared_file"

	studentRepository "github.com/sccicitb/pupr-backend/data/repositories/student"
	studentService "github.com/sccicitb/pupr-backend/services/student"

	studentActivityRepository "github.com/sccicitb/pupr-backend/data/repositories/student_activity"
	studentActivityService "github.com/sccicitb/pupr-backend/services/student_activity"

	studentClassRepository "github.com/sccicitb/pupr-backend/data/repositories/student_class"
	studentClassService "github.com/sccicitb/pupr-backend/services/student_class"

	studentLeaveRepository "github.com/sccicitb/pupr-backend/data/repositories/student_leave"
	studentLeaveService "github.com/sccicitb/pupr-backend/services/student_leave"

	studentSkpiRepository "github.com/sccicitb/pupr-backend/data/repositories/student_skpi"
	studentSkpiService "github.com/sccicitb/pupr-backend/services/student_skpi"

	studentSubjectService "github.com/sccicitb/pupr-backend/services/student_subject"

	studyLevelRepository "github.com/sccicitb/pupr-backend/data/repositories/study_level"
	studyLevelService "github.com/sccicitb/pupr-backend/services/study_level"

	studyPlanRepository "github.com/sccicitb/pupr-backend/data/repositories/study_plan"
	studyPlanService "github.com/sccicitb/pupr-backend/services/study_plan"

	studyProgramRepository "github.com/sccicitb/pupr-backend/data/repositories/study_program"
	studyProgramService "github.com/sccicitb/pupr-backend/services/study_program"

	subjectRepository "github.com/sccicitb/pupr-backend/data/repositories/subject"
	subjectService "github.com/sccicitb/pupr-backend/services/subject"

	subjectCategoryRepository "github.com/sccicitb/pupr-backend/data/repositories/subject_category"
	subjectCategoryService "github.com/sccicitb/pupr-backend/services/subject_category"

	subjectEquivalenceRepository "github.com/sccicitb/pupr-backend/data/repositories/subject_equivalence"

	subjectGradeComponentRepository "github.com/sccicitb/pupr-backend/data/repositories/subject_grade_component"
	subjectGradeComponentService "github.com/sccicitb/pupr-backend/services/subject_grade_component"

	subjectPrerequisiteRepository "github.com/sccicitb/pupr-backend/data/repositories/subject_prerequisite"

	thesisRepository "github.com/sccicitb/pupr-backend/data/repositories/thesis"
	thesisService "github.com/sccicitb/pupr-backend/services/thesis"

	thesisExaminerRoleRepository "github.com/sccicitb/pupr-backend/data/repositories/thesis_examiner_role"
	thesisExaminerRoleService "github.com/sccicitb/pupr-backend/services/thesis_examiner_role"

	thesisSupervisorRoleRepository "github.com/sccicitb/pupr-backend/data/repositories/thesis_supervisor_role"
	thesisSupervisorRoleService "github.com/sccicitb/pupr-backend/services/thesis_supervisor_role"

	yudiciumService "github.com/sccicitb/pupr-backend/services/yudicium"

	yudiciumSessionRepository "github.com/sccicitb/pupr-backend/data/repositories/yudicium_session"
	yudiciumSessionService "github.com/sccicitb/pupr-backend/services/yudicium_session"

	yudiciumStudentRepository "github.com/sccicitb/pupr-backend/data/repositories/yudicium_student"

	yudiciumTermRepository "github.com/sccicitb/pupr-backend/data/repositories/yudicium_term"
	yudiciumTermService "github.com/sccicitb/pupr-backend/services/yudicium_term"

	excelRepository "github.com/sccicitb/pupr-backend/data/repositories/excel"
	excelService "github.com/sccicitb/pupr-backend/services/excel"
)

var (
	routerCMD = &cobra.Command{
		Use:   "serve-http",
		Short: "Run http server",
		Long:  "PUPR Backend",
		RunE:  runHTTP,
	}
)

var EmbedMigration embed.FS
var EmbedMigrationLog embed.FS

func migrate(dbAppConfig config.DBConfig, dbLogConfig config.DBConfig) error {
	goose.SetBaseFS(EmbedMigration)

	dbApp, err := sql.Open("postgres", dbAppConfig.Host)
	if err != nil {
		return err
	}
	defer dbApp.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(dbApp, "data/migrations"); err != nil {
		return err
	}

	goose.SetBaseFS(EmbedMigrationLog)

	dbLog, err := sql.Open("postgres", dbLogConfig.Host)
	if err != nil {
		return err
	}
	defer dbLog.Close()

	if err := goose.Up(dbLog, "data/migrations/logs"); err != nil {
		return err
	}

	return nil
}

func initStorageInstance(cfg config.Config) (file.FileCtx, error) {
	var result file.FileCtx

	localStorage := file.NewLocalStorage(&cfg)

	urlStorage := file.NewUrlStorage(&cfg)

	result = file.FileCtx{
		Local: localStorage,
		Url:   urlStorage,
	}

	return result, nil
}

func initInfraCtx(database *db.DB, databaseLog *db.DB, cfg config.Config, rdb *v8.Client, mail mail.MailInterface, otpInterface otp.OTPInterface) *infra.InfraCtx {
	jwtSvc := middleware.NewJWT(&cfg.JWTConfig, rdb, &cfg.Firebase, false)

	fileSvc, err := initStorageInstance(cfg)
	if err != nil {
		logrus.Fatalln(err)
	}

	notificationSvc := fbs.NewFCM(database, &cfg.Firebase)

	notificationTemplate := notification.NotificationTemplate{}

	return &infra.InfraCtx{
		Config:               &cfg,
		DB:                   database,
		DBLog:                databaseLog,
		Jwt:                  jwtSvc,
		Mail:                 mail,
		Storage:              fileSvc,
		Notification:         notificationSvc,
		NotificationTemplate: notificationTemplate,
		Otp:                  otpInterface,
		RedisClient:          rdb,
	}
}

func runScheduler(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) {
	tz, err := time.LoadLocation(appConstants.DefaultTimezone)
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	schedulerSvc := schedulerService.NewSchedulerService(repoCtx, infraCtx)

	s := gocron.NewScheduler(tz)

	_, err = s.Cron("0 0 * * *").Do(schedulerSvc.AutoSetLeaveActive)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	// _, err = s.Cron("0 0 * * *").Do(schedulerSvc.AutoSetActiveSemester)
	// if err != nil {
	// 	logrus.Errorln(err)
	// 	return
	// }

	s.StartAsync()
}

// initRepoCtx for context repository
func initRepoCtx(dbApp *db.DB, dbLog *db.DB) *repository.RepoCtx {
	adminActivityLogRepo := adminActivityLogRepository.NewAdminActivityLogRepository(dbLog)
	lecturerStudentActivityLogRepo := lecturerStudentActivityLogRepository.NewLecturerStudentActivityLogRepository(dbLog)

	academicGuidanceRepo := academicGuidanceRepository.NewAcademicGuidanceRepository(dbApp)
	accreditationRepo := accreditationRepository.NewAccreditationRepository(dbApp)
	adminRepo := adminRepository.NewAdminRepository(dbApp)
	announcementRepo := announcementRepository.NewAnnouncementRepository(dbApp)
	authenticationRepo := authenticationRepository.NewAuthenticationRepository(dbApp)
	buildingRepo := buildingRepository.NewBuildingRepository(dbApp)
	classRepo := classRepository.NewClassRepository(dbApp)
	classAnnouncementRepo := classAnnouncementRepository.NewClassAnnouncementRepository(dbApp)
	classDiscussionRepo := classDiscussionRepository.NewClassDiscussionRepository(dbApp)
	classEventRepo := classEventRepository.NewClassEventRepository(dbApp)
	classExamRepo := classExamRepository.NewClassExamRepository(dbApp)
	classGradeComponentRepo := classGradeComponentRepository.NewClassGradeComponentRepository(dbApp)
	classLecturerRepo := classLecturerRepository.NewClassLecturerRepository(dbApp)
	classMaterialRepo := classMaterialRepository.NewClassMaterialRepository(dbApp)
	classWorkRepo := classWorkRepository.NewClassWorkRepository(dbApp)
	creditQuotaRepo := creditQuotaRepository.NewCreditQuotaRepository(dbApp)
	curriculumRepo := curriculumRepository.NewCurriculumRepository(dbApp)
	diktiStudyProgramRepo := diktiStudyProgramRepository.NewDiktiStudyProgramRepository(dbApp)
	documentActionRepo := documentActionRepository.NewDocumentActionRepository(dbApp)
	documentTypeRepo := documentTypeRepository.NewDocumentTypeRepository(dbApp)
	examSupervisorRepo := examSupervisorRepository.NewExamSupervisorRepository(dbApp)
	examSupervisorRoleRepo := examSupervisorRoleRepository.NewExamSupervisorRoleRepository(dbApp)
	expertiseGroupRepo := expertiseGroupRepository.NewExpertiseGroupRepository(dbApp)
	facultyRepo := facultyRepository.NewFacultyRepository(dbApp)
	gradeComponentRepo := gradeComponentRepository.NewGradeComponentRepository(dbApp)
	gradeTypeRepo := gradeTypeRepository.NewGradeTypeRepository(dbApp)
	graduationPredicateRepo := graduationPredicateRepository.NewGraduationPredicateRepository(dbApp)
	graduationSessionRepo := graduationSessionRepository.NewGraduationSessionRepository(dbApp)
	graduationStudentRepo := graduationStudentRepository.NewGraduationStudentRepository(dbApp)
	learningAchievementRepo := learningAchievementRepository.NewLearningAchievementRepository(dbApp)
	learningAchievementCategoryRepo := learningAchievementCategoryRepository.NewLearningAchievementCategoryRepository(dbApp)
	lectureRepo := lectureRepository.NewLectureRepository(dbApp)
	lecturerRepo := lecturerRepository.NewLecturerRepository(dbApp)
	lecturerLeaveRepo := lecturerLeaveRepository.NewLecturerLeaveRepository(dbApp)
	lecturerMutationRepo := lecturerMutationRepository.NewLecturerMutationRepository(dbApp)
	lecturerResignationRepo := lecturerResignationRepository.NewLecturerResignationRepository(dbApp)
	lessonPlanRepo := lessonPlanRepository.NewLessonPlanRepository(dbApp)
	locationRepo := locationRepository.NewLocationRepository(dbApp)
	majorRepo := majorRepository.NewMajorRepository(dbApp)
	officerRepo := officerRepository.NewOfficerRepository(dbApp)
	officerActionRepo := officerActionRepository.NewOfficerActionRepository(dbApp)
	permissionRepo := permissionRepository.NewPermissionRepository(dbApp)
	reportRepo := reportRepository.NewReportRepository(dbApp)
	roleRepo := roleRepository.NewRoleRepository(dbApp)
	roomRepo := roomRepository.NewRoomRepository(dbApp)
	semesterRepo := semesterRepository.NewSemesterRepository(dbApp)
	sharedFileRepo := sharedFileRepository.NewSharedFileRepository(dbApp)
	studentRepo := studentRepository.NewStudentRepository(dbApp)
	studentActivityRepo := studentActivityRepository.NewStudentActivityRepository(dbApp)
	studentClassRepo := studentClassRepository.NewStudentClassRepository(dbApp)
	studentLeaveRepo := studentLeaveRepository.NewStudentLeaveRepository(dbApp)
	studentSkpiRepo := studentSkpiRepository.NewStudentSkpiRepository(dbApp)
	studyLevelRepo := studyLevelRepository.NewStudyLevelRepository(dbApp)
	studyPlanRepo := studyPlanRepository.NewStudyPlanRepository(dbApp)
	studyProgramRepo := studyProgramRepository.NewStudyProgramRepository(dbApp)
	subjectRepo := subjectRepository.NewSubjectRepository(dbApp)
	subjectCategoryRepo := subjectCategoryRepository.NewSubjectCategoryRepository(dbApp)
	subjectEquivalenceRepo := subjectEquivalenceRepository.NewSubjectEquivalenceRepository(dbApp)
	subjectGradeComponentRepo := subjectGradeComponentRepository.NewSubjectGradeComponentRepository(dbApp)
	subjectPrerequisiteRepo := subjectPrerequisiteRepository.NewSubjectPrerequisiteRepository(dbApp)
	thesisRepo := thesisRepository.NewThesisRepository(dbApp)
	thesisExaminerRoleRepo := thesisExaminerRoleRepository.NewThesisExaminerRoleRepository(dbApp)
	thesisSupervisorRoleRepo := thesisSupervisorRoleRepository.NewThesisSupervisorRoleRepository(dbApp)
	yudiciumSessionRepo := yudiciumSessionRepository.NewYudiciumSessionRepository(dbApp)
	yudiciumStudentRepo := yudiciumStudentRepository.NewYudiciumStudentRepository(dbApp)
	yudiciumTermRepo := yudiciumTermRepository.NewYudiciumTermRepository(dbApp)
	excelRepo := excelRepository.NewExcelRepository(dbApp)

	return &repository.RepoCtx{
		AdminActivityLogRepo:           adminActivityLogRepo,
		LecturerStudentActivityLogRepo: lecturerStudentActivityLogRepo,

		AcademicGuidanceRepo:            academicGuidanceRepo,
		AccreditationRepo:               accreditationRepo,
		AdminRepo:                       adminRepo,
		AnnouncementRepo:                announcementRepo,
		AuthenticationRepo:              authenticationRepo,
		BuildingRepo:                    buildingRepo,
		ClassRepo:                       classRepo,
		ClassAnnouncementRepo:           classAnnouncementRepo,
		ClassDiscussionRepo:             classDiscussionRepo,
		ClassEventRepo:                  classEventRepo,
		ClassExamRepo:                   classExamRepo,
		ClassGradeComponentRepo:         classGradeComponentRepo,
		ClassLecturerRepo:               classLecturerRepo,
		ClassMaterialRepo:               classMaterialRepo,
		ClassWorkRepo:                   classWorkRepo,
		CreditQuotaRepo:                 creditQuotaRepo,
		CurriculumRepo:                  curriculumRepo,
		DiktiStudyProgramRepo:           diktiStudyProgramRepo,
		DocumentActionRepo:              documentActionRepo,
		DocumentTypeRepo:                documentTypeRepo,
		ExamSupervisorRepo:              examSupervisorRepo,
		ExamSupervisorRoleRepo:          examSupervisorRoleRepo,
		ExpertiseGroupRepo:              expertiseGroupRepo,
		FacultyRepo:                     facultyRepo,
		GradeComponentRepo:              gradeComponentRepo,
		GradeTypeRepo:                   gradeTypeRepo,
		GraduationPredicateRepo:         graduationPredicateRepo,
		GraduationSessionRepo:           graduationSessionRepo,
		GraduationStudentRepo:           graduationStudentRepo,
		LearningAchievementRepo:         learningAchievementRepo,
		LearningAchievementCategoryRepo: learningAchievementCategoryRepo,
		LectureRepo:                     lectureRepo,
		LecturerRepo:                    lecturerRepo,
		LecturerLeaveRepo:               lecturerLeaveRepo,
		LecturerMutationRepo:            lecturerMutationRepo,
		LecturerResignationRepo:         lecturerResignationRepo,
		LessonPlanRepo:                  lessonPlanRepo,
		LocationRepo:                    locationRepo,
		MajorRepo:                       majorRepo,
		OfficerRepo:                     officerRepo,
		OfficerActionRepo:               officerActionRepo,
		PermissionRepo:                  permissionRepo,
		ReportRepo:                      reportRepo,
		RoleRepo:                        roleRepo,
		RoomRepo:                        roomRepo,
		SemesterRepo:                    semesterRepo,
		SharedFileRepo:                  sharedFileRepo,
		StudentRepo:                     studentRepo,
		StudentActivityRepo:             studentActivityRepo,
		StudentClassRepo:                studentClassRepo,
		StudentLeaveRepo:                studentLeaveRepo,
		StudentSkpiRepo:                 studentSkpiRepo,
		StudyLevelRepo:                  studyLevelRepo,
		StudyPlanRepo:                   studyPlanRepo,
		StudyProgramRepo:                studyProgramRepo,
		SubjectRepo:                     subjectRepo,
		SubjectCategoryRepo:             subjectCategoryRepo,
		SubjectEquivalenceRepo:          subjectEquivalenceRepo,
		SubjectGradeComponentRepo:       subjectGradeComponentRepo,
		SubjectPrerequisiteRepo:         subjectPrerequisiteRepo,
		ThesisRepo:                      thesisRepo,
		ThesisExaminerRoleRepo:          thesisExaminerRoleRepo,
		ThesisSupervisorRoleRepo:        thesisSupervisorRoleRepo,
		YudiciumSessionRepo:             yudiciumSessionRepo,
		YudiciumStudentRepo:             yudiciumStudentRepo,
		YudiciumTermRepo:                yudiciumTermRepo,
		ExcelRepo:                       excelRepo,
	}
}

// initServiceCtx for contextService
func initServiceCtx(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx, appMw appMiddleware.MiddlewareInterface) *service.ServiceCtx {
	fileSvc := fileService.NewFileService(infraCtx)

	academicGuidanceSvc := academicGuidanceService.NewAcademicGuidanceService(repoCtx, infraCtx)
	accreditationSvc := accreditationService.NewAccreditationService(repoCtx, infraCtx)
	adminSvc := adminService.NewAdminService(repoCtx, infraCtx)
	announcementSvc := announcementService.NewAnnouncementService(repoCtx, infraCtx)
	authenticationSvc := authenticationService.NewAuthenticationService(repoCtx, infraCtx)
	adminActivityLogSvc := adminActivityLogService.NewAdminActivityLogService(repoCtx, infraCtx, appMw)
	buildingSvc := buildingService.NewBuildingService(repoCtx, infraCtx)
	classSvc := classService.NewClassService(repoCtx, infraCtx)
	classAnnouncementSvc := classAnnouncementService.NewClassAnnouncementService(repoCtx, infraCtx)
	classDiscussionSvc := classDiscussionService.NewClassDiscussionService(repoCtx, infraCtx)
	classEventSvc := classEventService.NewClassEventService(repoCtx, infraCtx)
	classExamSvc := classExamService.NewClassExamService(repoCtx, infraCtx)
	classGradeComponentSvc := classGradeComponentService.NewClassGradeComponentService(repoCtx, infraCtx)
	classMaterialSvc := classMaterialService.NewClassMaterialService(repoCtx, infraCtx)
	classWorkSvc := classWorkService.NewClassWorkService(repoCtx, infraCtx)
	creditQuotaSvc := creditQuotaService.NewCreditQuotaService(repoCtx, infraCtx)
	curriculumSvc := curriculumService.NewCurriculumService(repoCtx, infraCtx)
	diktiStudyProgramSvc := diktiStudyProgramService.NewDiktiStudyProgramService(repoCtx, infraCtx)
	documentActionSvc := documentActionService.NewDocumentActionService(repoCtx, infraCtx)
	documentTypeSvc := documentTypeService.NewDocumentTypeService(repoCtx, infraCtx)
	examSupervisorSvc := examSupervisorService.NewExamSupervisorService(repoCtx, infraCtx)
	examSupervisorRoleSvc := examSupervisorRoleService.NewExamSupervisorRoleService(repoCtx, infraCtx)
	expertiseGroupSvc := expertiseGroupService.NewExpertiseGroupService(repoCtx, infraCtx)
	facultySvc := facultyService.NewFacultyService(repoCtx, infraCtx)
	gradeComponentSvc := gradeComponentService.NewGradeComponentService(repoCtx, infraCtx)
	gradeTypeSvc := gradeTypeService.NewGradeTypeService(repoCtx, infraCtx)
	graduationSvc := graduationService.NewGraduationService(repoCtx, infraCtx)
	graduationPredicateSvc := graduationPredicateService.NewGraduationPredicateService(repoCtx, infraCtx)
	graduationSessionSvc := graduationSessionService.NewGraduationSessionService(repoCtx, infraCtx)
	learningAchievementSvc := learningAchievementService.NewLearningAchievementService(repoCtx, infraCtx)
	learningAchievementCategorySvc := learningAchievementCategoryService.NewLearningAchievementCategoryService(repoCtx, infraCtx)
	lectureSvc := lectureService.NewLectureService(repoCtx, infraCtx)
	lecturerSvc := lecturerService.NewLecturerService(repoCtx, infraCtx)
	lecturerLeaveSvc := lecturerLeaveService.NewLecturerLeaveService(repoCtx, infraCtx)
	lecturerMutationSvc := lecturerMutationService.NewLecturerMutationService(repoCtx, infraCtx)
	lecturerResignationSvc := lecturerResignationService.NewLecturerResignationService(repoCtx, infraCtx)
	lecturerStudentActivityLogSvc := lecturerStudentActivityLogService.NewLecturerStudentActivityLogService(repoCtx, infraCtx, appMw)
	lessonPlanSvc := lessonPlanService.NewLessonPlanService(repoCtx, infraCtx)
	locationSvc := locationService.NewLocationService(repoCtx, infraCtx)
	majorSvc := majorService.NewMajorService(repoCtx, infraCtx)
	officerSvc := officerService.NewOfficerService(repoCtx, infraCtx)
	officerActionSvc := officerActionService.NewOfficerActionService(repoCtx, infraCtx)
	permissionSvc := permissionService.NewPermissionService(repoCtx, infraCtx)
	reportSvc := reportService.NewReportService(repoCtx, infraCtx)
	roleSvc := roleService.NewRoleService(repoCtx, infraCtx)
	roomSvc := roomService.NewRoomService(repoCtx, infraCtx)
	semesterSvc := semesterService.NewSemesterService(repoCtx, infraCtx)
	sharedFileSvc := sharedFileService.NewSharedFileService(repoCtx, infraCtx)
	studentSvc := studentService.NewStudentService(repoCtx, infraCtx)
	studentActivitySvc := studentActivityService.NewStudentActivityService(repoCtx, infraCtx)
	studentClassSvc := studentClassService.NewStudentClassService(repoCtx, infraCtx)
	studentLeaveSvc := studentLeaveService.NewStudentLeaveService(repoCtx, infraCtx)
	studentSkpiSvc := studentSkpiService.NewStudentSkpiService(repoCtx, infraCtx)
	studentSubjectSvc := studentSubjectService.NewStudentSubjectService(repoCtx, infraCtx)
	studyLevelSvc := studyLevelService.NewStudyLevelService(repoCtx, infraCtx)
	studyPlanSvc := studyPlanService.NewStudyPlanService(repoCtx, infraCtx)
	studyProgramSvc := studyProgramService.NewStudyProgramService(repoCtx, infraCtx)
	subjectSvc := subjectService.NewSubjectService(repoCtx, infraCtx)
	subjectCategorySvc := subjectCategoryService.NewSubjectCategoryService(repoCtx, infraCtx)
	subjectGradeComponentSvc := subjectGradeComponentService.NewSubjectGradeComponentService(repoCtx, infraCtx)
	thesisSvc := thesisService.NewThesisService(repoCtx, infraCtx)
	thesisExaminerRoleSvc := thesisExaminerRoleService.NewThesisExaminerRoleService(repoCtx, infraCtx)
	thesisSupervisorRoleSvc := thesisSupervisorRoleService.NewThesisSupervisorRoleService(repoCtx, infraCtx)
	yudiciumSvc := yudiciumService.NewYudiciumService(repoCtx, infraCtx)
	yudiciumSessionSvc := yudiciumSessionService.NewYudiciumSessionService(repoCtx, infraCtx)
	yudiciumTermSvc := yudiciumTermService.NewYudiciumTermService(repoCtx, infraCtx)

	excelSvc := excelService.NewExcelServiceInterface(repoCtx, infraCtx)

	return &service.ServiceCtx{
		FileService: fileSvc,

		AcademicGuidanceService:            academicGuidanceSvc,
		AccreditationService:               accreditationSvc,
		AdminService:                       adminSvc,
		AnnouncementService:                announcementSvc,
		AuthenticationService:              authenticationSvc,
		AdminActivityLogService:            adminActivityLogSvc,
		BuildingService:                    buildingSvc,
		ClassService:                       classSvc,
		ClassAnnouncementService:           classAnnouncementSvc,
		ClassDiscussionService:             classDiscussionSvc,
		ClassEventService:                  classEventSvc,
		ClassExamService:                   classExamSvc,
		ClassGradeComponentService:         classGradeComponentSvc,
		ClassMaterialService:               classMaterialSvc,
		ClassWorkService:                   classWorkSvc,
		CreditQuotaService:                 creditQuotaSvc,
		CurriculumService:                  curriculumSvc,
		DiktiStudyProgramService:           diktiStudyProgramSvc,
		DocumentActionService:              documentActionSvc,
		DocumentTypeService:                documentTypeSvc,
		ExamSupervisorService:              examSupervisorSvc,
		ExamSupervisorRoleService:          examSupervisorRoleSvc,
		ExpertiseGroupService:              expertiseGroupSvc,
		FacultyService:                     facultySvc,
		GradeComponentService:              gradeComponentSvc,
		GradeTypeService:                   gradeTypeSvc,
		GraduationService:                  graduationSvc,
		GraduationPredicateService:         graduationPredicateSvc,
		GraduationSessionService:           graduationSessionSvc,
		LearningAchievementService:         learningAchievementSvc,
		LearningAchievementCategoryService: learningAchievementCategorySvc,
		LectureService:                     lectureSvc,
		LecturerService:                    lecturerSvc,
		LecturerLeaveService:               lecturerLeaveSvc,
		LecturerMutationService:            lecturerMutationSvc,
		LecturerResignationService:         lecturerResignationSvc,
		LecturerStudentActivityLogService:  lecturerStudentActivityLogSvc,
		LessonPlanService:                  lessonPlanSvc,
		LocationService:                    locationSvc,
		MajorService:                       majorSvc,
		OfficerService:                     officerSvc,
		OfficerActionService:               officerActionSvc,
		PermissionService:                  permissionSvc,
		ReportService:                      reportSvc,
		RoleService:                        roleSvc,
		RoomService:                        roomSvc,
		SemesterService:                    semesterSvc,
		SharedFileService:                  sharedFileSvc,
		StudentService:                     studentSvc,
		StudentActivityService:             studentActivitySvc,
		StudentClassService:                studentClassSvc,
		StudentLeaveService:                studentLeaveSvc,
		StudentSkpiService:                 studentSkpiSvc,
		StudentSubjectService:              studentSubjectSvc,
		StudyLevelService:                  studyLevelSvc,
		StudyPlanService:                   studyPlanSvc,
		StudyProgramService:                studyProgramSvc,
		SubjectService:                     subjectSvc,
		SubjectCategoryService:             subjectCategorySvc,
		SubjectGradeComponentService:       subjectGradeComponentSvc,
		ThesisService:                      thesisSvc,
		ThesisExaminerRoleService:          thesisExaminerRoleSvc,
		ThesisSupervisorRoleService:        thesisSupervisorRoleSvc,
		YudiciumService:                    yudiciumSvc,
		YudiciumSessionService:             yudiciumSessionSvc,
		YudiciumTermService:                yudiciumTermSvc,
		ExcelService:                       excelSvc,
	}
}

func initHandlerCtx(serviceCtx *service.ServiceCtx) *handler.HandlerCtx {
	adminAcademicGuidance := adminAcademicGuidanceHandler.NewAdminAcademicGuidanceHandler(serviceCtx)
	adminAccreditation := adminAccreditationHandler.NewAdminAccreditationHandler(serviceCtx)
	adminAnnouncement := adminAnnouncementHandler.NewAdminAnnouncementHandler(serviceCtx)
	adminAuthentication := adminAuthenticationHandler.NewAdminAuthentication(serviceCtx)
	adminBuilding := adminBuildingHandler.NewAdminBuildingHandler(serviceCtx)
	adminClass := adminClassHandler.NewAdminClassHandler(serviceCtx)
	adminClassDiscussion := adminClassDiscussionHandler.NewAdminClassDiscussionHandler(serviceCtx)
	adminClassEvent := adminClassEventHandler.NewAdminClassEventHandler(serviceCtx)
	adminClassExam := adminClassExamHandler.NewAdminClassExamHandler(serviceCtx)
	adminClassGradeComponent := adminClassGradeComponentHandler.NewAdminClassGradeComponentHandler(serviceCtx)
	adminClassMaterial := adminClassMaterialHandler.NewAdminClassMaterialHandler(serviceCtx)
	adminClassWork := adminClassWorkHandler.NewAdminClassWorkHandler(serviceCtx)
	adminCreditQuota := adminCreditQuotaHandler.NewAdminCreditQuotaHandler(serviceCtx)
	adminCurriculum := adminCurriculumHandler.NewAdminCurriculumHandler(serviceCtx)
	adminDiktiStudyProgram := adminDiktiStudyProgramHandler.NewAdminDiktiStudyProgramHandler(serviceCtx)
	adminDocumentAction := adminDocumentActionHandler.NewAdminDocumentActionHandler(serviceCtx)
	adminDocumentType := adminDocumentTypeHandler.NewAdminDocumentTypeHandler(serviceCtx)
	adminExamSupervisor := adminExamSupervisorHandler.NewAdminExamSupervisorHandler(serviceCtx)
	adminExamSupervisorRole := adminExamSupervisorRoleHandler.NewAdminExamSupervisorRoleHandler(serviceCtx)
	adminExpertiseGroup := adminExpertiseGroupHandler.NewAdminExpertiseGroupHandler(serviceCtx)
	adminFaculty := adminFacultyHandler.NewAdminFacultyHandler(serviceCtx)
	adminGradeComponent := adminGradeComponentHandler.NewAdminGradeComponentHandler(serviceCtx)
	adminGradeType := adminGradeTypeHandler.NewAdminGradeTypeHandler(serviceCtx)
	adminGraduation := adminGraduationHandler.NewAdminGraduationHandler(serviceCtx)
	adminGraduationPredicate := adminGraduationPredicateHandler.NewAdminGraduationPredicateHandler(serviceCtx)
	adminGraduationSession := adminGraduationSessionHandler.NewAdminGraduationSessionHandler(serviceCtx)
	adminLearningAchievement := adminLearningAchievementHandler.NewAdminLearningAchievementHandler(serviceCtx)
	adminLearningAchievementCategory := adminLearningAchievementCategoryHandler.NewAdminLearningAchievementCategoryHandler(serviceCtx)
	adminLecture := adminLectureHandler.NewAdminLectureHandler(serviceCtx)
	adminLecturer := adminLecturerHandler.NewAdminLecturerHandler(serviceCtx)
	adminLecturerLeave := adminLecturerLeaveHandler.NewAdminLecturerLeaveHandler(serviceCtx)
	adminLecturerMutation := adminLecturerMutationHandler.NewAdminLecturerMutationHandler(serviceCtx)
	adminLecturerResignation := adminLecturerResignationHandler.NewAdminLecturerResignationHandler(serviceCtx)
	adminLessonPlan := adminLessonPlanHandler.NewAdminLessonPlanHandler(serviceCtx)
	adminMajor := adminMajorHandler.NewAdminMajorHandler(serviceCtx)
	adminOfficer := adminOfficerHandler.NewAdminOfficerHandler(serviceCtx)
	adminOfficerAction := adminOfficerActionHandler.NewAdminOfficerActionHandler(serviceCtx)
	adminReport := adminReportHandler.NewAdminReportHandler(serviceCtx)
	adminRoom := adminRoomHandler.NewAdminRoomHandler(serviceCtx)
	adminSemester := adminSemesterHandler.NewAdminSemesterHandler(serviceCtx)
	adminSharedFile := adminSharedFileHandler.NewAdminSharedFileHandler(serviceCtx)
	adminStudent := adminStudentHandler.NewAdminStudentHandler(serviceCtx)
	adminStudentActivity := adminStudentActivityHandler.NewAdminStudentActivityHandler(serviceCtx)
	adminStudentClass := adminStudentClassHandler.NewAdminStudentClassHandler(serviceCtx)
	adminStudentLeave := adminStudentLeaveHandler.NewAdminStudentLeaveHandler(serviceCtx)
	adminStudentSkpi := adminStudentSkpiHandler.NewAdminStudentSkpiHandler(serviceCtx)
	adminStudyLevel := adminStudyLevelHandler.NewAdminStudyLevelHandler(serviceCtx)
	adminStudyPlan := adminStudyPlanHandler.NewAdminStudyPlanHandler(serviceCtx)
	adminStudyProgram := adminStudyProgramHandler.NewAdminStudyProgramHandler(serviceCtx)
	adminSubject := adminSubjectHandler.NewAdminSubjectHandler(serviceCtx)
	adminSubjectCategory := adminSubjectCategoryHandler.NewAdminSubjectCategoryHandler(serviceCtx)
	adminSubjectGradeComponent := adminSubjectGradeComponentHandler.NewAdminSubjectGradeComponentHandler(serviceCtx)
	adminThesis := adminThesisHandler.NewAdminThesisHandler(serviceCtx)
	adminThesisExaminerRole := adminThesisExaminerRoleHandler.NewAdminThesisExaminerRoleHandler(serviceCtx)
	adminThesisSupervisorRole := adminThesisSupervisorRoleHandler.NewAdminThesisSupervisorRoleHandler(serviceCtx)
	adminTranscript := adminTranscriptHandler.NewAdminTranscriptHandler(serviceCtx)
	adminYudicium := adminYudiciumHandler.NewAdminYudiciumHandler(serviceCtx)
	adminYudiciumSession := adminYudiciumSessionHandler.NewAdminYudiciumSessionHandler(serviceCtx)
	adminYudiciumTerm := adminYudiciumTermHandler.NewAdminYudiciumTermHandler(serviceCtx)

	careerStudyProgram := careerStudyProgramHandler.NewCareerStudyProgramHandler(serviceCtx)

	generalAuth := generalAuthHandler.NewGeneralAuthHandler(serviceCtx)
	generalFile := generalFileHandler.NewGeneralFileHandler(serviceCtx)
	generalLocation := generalLocationHandler.NewGeneralLocationHandler(serviceCtx)

	lecturerAcademicGuidance := lecturerAcademicGuidanceHandler.NewLecturerAcademicGuidanceHandler(serviceCtx)
	lecturerAnnouncement := lecturerAnnouncementHandler.NewLecturerAnnouncementHandler(serviceCtx)
	lecturerClass := lecturerClassHandler.NewLecturerClassHandler(serviceCtx)
	lecturerClassAnnouncement := lecturerClassAnnouncementHandler.NewLecturerClassAnnouncementHandler(serviceCtx)
	lecturerClassDiscussion := lecturerClassDiscussionHandler.NewLecturerClassDiscussionHandler(serviceCtx)
	lecturerClassEvent := lecturerClassEventHandler.NewLecturerClassEventHandler(serviceCtx)
	lecturerClassExam := lecturerClassExamHandler.NewLecturerClassExamHandler(serviceCtx)
	lecturerClassGradeComponent := lecturerClassGradeComponentHandler.NewLecturerClassGradeComponentHandler(serviceCtx)
	lecturerClassMaterial := lecturerClassMaterialHandler.NewLecturerClassMaterialHandler(serviceCtx)
	lecturerClassWork := lecturerClassWorkHandler.NewLecturerClassWorkHandler(serviceCtx)
	lecturerGeneral := lecturerGeneralHandler.NewLecturerGeneralHandler(serviceCtx)
	lecturerLecture := lecturerLectureHandler.NewLecturerLectureHandler(serviceCtx)
	lecturerSemester := lecturerSemesterHandler.NewLecturerSemesterHandler(serviceCtx)
	lecturerSharedFile := lecturerSharedFileHandler.NewLecturerSharedFileHandler(serviceCtx)
	lecturerStudyProgram := lecturerStudyProgramHandler.NewLecturerStudyProgramHandler(serviceCtx)

	pmbStudent := pmbStudentHandler.NewPmbStudentHandler(serviceCtx)

	rootAdmin := rootAdminHandler.NewRootAdminHandler(serviceCtx)
	rootAdminActivityLog := rootAdminActivityLogHandler.NewRootAdminActivityLogHandler(serviceCtx)
	rootDiktiStudyProgram := rootDiktiStudyProgramHandler.NewRootDiktiStudyProgramHandler(serviceCtx)
	rootFaculty := rootFacultyHandler.NewRootFacultyHandler(serviceCtx)
	rootLecturer := rootLecturerHandler.NewRootLecturerHandler(serviceCtx)
	rootLecturerStudentActivityLog := rootLecturerStudentActivityLogHandler.NewRootLecturerStudentActivityLogHandler(serviceCtx)
	rootMajor := rootMajorHandler.NewRootMajorHandler(serviceCtx)
	rootPermission := rootPermissionHandler.NewRootPermissionHandler(serviceCtx)
	rootRole := rootRoleHandler.NewRootRoleHandler(serviceCtx)
	rootStudyProgram := rootStudyProgramHandler.NewRootStudyProgramHandler(serviceCtx)

	studentAcademicGuidance := studentAcademicGuidanceHandler.NewStudentAcademicGuidanceHandler(serviceCtx)
	studentAnnouncement := studentAnnouncementHandler.NewStudentAnnouncementHandler(serviceCtx)
	studentClass := studentClassHandler.NewStudentClassHandler(serviceCtx)
	studentClassAnnouncement := studentClassAnnouncementHandler.NewStudentClassAnnouncementHandler(serviceCtx)
	studentClassDiscussion := studentClassDiscussionHandler.NewStudentClassDiscussionHandler(serviceCtx)
	studentClassEvent := studentClassEventHandler.NewStudentClassEventHandler(serviceCtx)
	studentClassExam := studentClassExamHandler.NewStudentClassExamHandler(serviceCtx)
	studentClassMaterial := studentClassMaterialHandler.NewStudentClassMaterialHandler(serviceCtx)
	studentClassWork := studentClassWorkHandler.NewStudentClassWorkHandler(serviceCtx)
	studentGeneral := studentGeneralHandler.NewStudentGeneralHandler(serviceCtx)
	studentGradeType := studentGradeTypeHandler.NewStudentGradeTypeHandler(serviceCtx)
	studentLecture := studentLectureHandler.NewStudentLectureHandler(serviceCtx)
	studentSemester := studentSemesterHandler.NewStudentSemesterHandler(serviceCtx)
	studentSharedFile := studentSharedFileHandler.NewStudentSharedFileHandler(serviceCtx)
	studentStudentLeave := studentStudentLeaveHandler.NewStudentStudentLeaveHandler(serviceCtx)
	studentStudentSkpi := studentStudentSkpiHandler.NewStudentStudentSkpiHandler(serviceCtx)
	studentStudyPlan := studentStudyPlanHandler.NewStudentStudyPlanHandler(serviceCtx)
	studentThesis := studentThesisHandler.NewStudentThesisHandler(serviceCtx)
	studentTranscript := studentTranscriptHandler.NewStudentTranscriptHandler(serviceCtx)
	excel := excelHandler.NewExcelHandler(serviceCtx)

	return &handler.HandlerCtx{
		AdminAcademicGuidanceHandler:            adminAcademicGuidance,
		AdminAccreditationHandler:               adminAccreditation,
		AdminAnnouncementHandler:                adminAnnouncement,
		AdminAuthenticationHandler:              adminAuthentication,
		AdminBuildingHandler:                    adminBuilding,
		AdminClassHandler:                       adminClass,
		AdminClassDiscussionHandler:             adminClassDiscussion,
		AdminClassEventHandler:                  adminClassEvent,
		AdminClassExamHandler:                   adminClassExam,
		AdminClassGradeComponentHandler:         adminClassGradeComponent,
		AdminClassMaterialHandler:               adminClassMaterial,
		AdminClassWorkHandler:                   adminClassWork,
		AdminCreditQuotaHandler:                 adminCreditQuota,
		AdminCurriculumHandler:                  adminCurriculum,
		AdminDiktiStudyProgramHandler:           adminDiktiStudyProgram,
		AdminDocumentActionHandler:              adminDocumentAction,
		AdminDocumentTypeHandler:                adminDocumentType,
		AdminExamSupervisorHandler:              adminExamSupervisor,
		AdminExamSupervisorRoleHandler:          adminExamSupervisorRole,
		AdminExpertiseGroupHandler:              adminExpertiseGroup,
		AdminFacultyHandler:                     adminFaculty,
		AdminGradeComponentHandler:              adminGradeComponent,
		AdminGradeTypeHandler:                   adminGradeType,
		AdminGraduationHandler:                  adminGraduation,
		AdminGraduationPredicateHandler:         adminGraduationPredicate,
		AdminGraduationSessionHandler:           adminGraduationSession,
		AdminLearningAchievementHandler:         adminLearningAchievement,
		AdminLearningAchievementCategoryHandler: adminLearningAchievementCategory,
		AdminLectureHandler:                     adminLecture,
		AdminLecturerHandler:                    adminLecturer,
		AdminLecturerLeaveHandler:               adminLecturerLeave,
		AdminLecturerMutationHandler:            adminLecturerMutation,
		AdminLecturerResignationHandler:         adminLecturerResignation,
		AdminLessonPlanHandler:                  adminLessonPlan,
		AdminMajorHandler:                       adminMajor,
		AdminOfficerHandler:                     adminOfficer,
		AdminOfficerActionHandler:               adminOfficerAction,
		AdminReportHandler:                      adminReport,
		AdminRoomHandler:                        adminRoom,
		AdminSemesterHandler:                    adminSemester,
		AdminSharedFileHandler:                  adminSharedFile,
		AdminStudentHandler:                     adminStudent,
		AdminStudentActivityHandler:             adminStudentActivity,
		AdminStudentClassHandler:                adminStudentClass,
		AdminStudentLeaveHandler:                adminStudentLeave,
		AdminStudentSkpiHandler:                 adminStudentSkpi,
		AdminStudyLevelHandler:                  adminStudyLevel,
		AdminStudyPlanHandler:                   adminStudyPlan,
		AdminStudyProgramHandler:                adminStudyProgram,
		AdminSubjectHandler:                     adminSubject,
		AdminSubjectCategoryHandler:             adminSubjectCategory,
		AdminSubjectGradeComponentHandler:       adminSubjectGradeComponent,
		AdminThesisHandler:                      adminThesis,
		AdminThesisExaminerRoleHandler:          adminThesisExaminerRole,
		AdminThesisSupervisorRoleHandler:        adminThesisSupervisorRole,
		AdminTranscriptHandler:                  adminTranscript,
		AdminYudiciumHandler:                    adminYudicium,
		AdminYudiciumSessionHandler:             adminYudiciumSession,
		AdminYudiciumTermHandler:                adminYudiciumTerm,

		CareerStudyProgramHandler: careerStudyProgram,

		GeneralAuthHandler:     generalAuth,
		GeneralFileHandler:     generalFile,
		GeneralLocationHandler: generalLocation,

		LecturerAcademicGuidanceHandler:    lecturerAcademicGuidance,
		LecturerAnnouncementHandler:        lecturerAnnouncement,
		LecturerClassHandler:               lecturerClass,
		LecturerClassAnnouncementHandler:   lecturerClassAnnouncement,
		LecturerClassDiscussionHandler:     lecturerClassDiscussion,
		LecturerClassEventHandler:          lecturerClassEvent,
		LecturerClassExamHandler:           lecturerClassExam,
		LecturerClassGradeComponentHandler: lecturerClassGradeComponent,
		LecturerClassMaterialHandler:       lecturerClassMaterial,
		LecturerClassWorkHandler:           lecturerClassWork,
		LecturerGeneralHandler:             lecturerGeneral,
		LecturerLectureHandler:             lecturerLecture,
		LecturerSemesterHandler:            lecturerSemester,
		LecturerSharedFileHandler:          lecturerSharedFile,
		LecturerStudyProgramHandler:        lecturerStudyProgram,

		PmbStudentHandler: pmbStudent,

		RootAdminHandler:                      rootAdmin,
		RootAdminActivityLogHandler:           rootAdminActivityLog,
		RootDiktiStudyProgramHandler:          rootDiktiStudyProgram,
		RootFacultyHandler:                    rootFaculty,
		RootLecturerHandler:                   rootLecturer,
		RootLecturerStudentActivityLogHandler: rootLecturerStudentActivityLog,
		RootMajorHandler:                      rootMajor,
		RootPermissionHandler:                 rootPermission,
		RootRoleHandler:                       rootRole,
		RootStudyProgramHandler:               rootStudyProgram,

		StudentAcademicGuidanceHandler:  studentAcademicGuidance,
		StudentAnnouncementHandler:      studentAnnouncement,
		StudentClassHandler:             studentClass,
		StudentClassAnnouncementHandler: studentClassAnnouncement,
		StudentClassDiscussionHandler:   studentClassDiscussion,
		StudentClassEventHandler:        studentClassEvent,
		StudentClassExamHandler:         studentClassExam,
		StudentClassMaterialHandler:     studentClassMaterial,
		StudentClassWorkHandler:         studentClassWork,
		StudentGeneralHandler:           studentGeneral,
		StudentGradeTypeHandler:         studentGradeType,
		StudentLectureHandler:           studentLecture,
		StudentSemesterHandler:          studentSemester,
		StudentSharedFileHandler:        studentSharedFile,
		StudentStudentLeaveHandler:      studentStudentLeave,
		StudentStudentSkpiHandler:       studentStudentSkpi,
		StudentStudyPlanHandler:         studentStudyPlan,
		StudentThesisHandler:            studentThesis,
		StudentTranscriptHandler:        studentTranscript,
		ExcelHandler:                    excel,
	}
}

func runHTTP(cmd *cobra.Command, args []string) error {
	tz, err := time.LoadLocation(appConstants.DefaultTimezone)
	if err != nil {
		logrus.Fatalln(err)
	}

	// initial config
	ctx := context.Background()
	cfg := config.InitConfig(tz)

	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	dbApp, err := db.Open(&cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
	}

	dbLog, err := db.Open(&cfg.DBLog)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = migrate(cfg.DB, cfg.DBLog)
	if err != nil {
		logrus.Fatalln(err)
	}

	// initial redis server
	redisServer := redis.NewRedisServer(&cfg.Redis)
	redisClient, err := redisServer.Connect(ctx)
	if err != nil {
		logrus.Fatalln(err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*time.Duration(cfg.Server.GraceFulTimeout), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// init repo ctx
	repoCtx := initRepoCtx(dbApp, dbLog)

	mailer := mail.NewSmtpMailer(&cfg.Mailer)
	otpInterface := otp.NewOTP(redisClient, cfg.Server.AppName)

	// init infra ctx
	infraCtx := initInfraCtx(dbApp, dbLog, cfg, redisClient, mailer, otpInterface)

	// init scheduler ctx
	runScheduler(repoCtx, infraCtx)

	appMw := appMiddleware.AccessTokenMiddleware(repoCtx, infraCtx)

	// init service ctx
	serviceCtx := initServiceCtx(repoCtx, infraCtx, appMw)

	// cors handler
	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		AllowedMethods: []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedOrigins: []string{
			"*",
		},
		OptionsPassthrough: false,
		AllowCredentials:   true,
	})

	// init handler
	handlerCtx := initHandlerCtx(serviceCtx)

	// initial router
	r := routers.InitialRouter(handlerCtx, appMw, &cfg)

	// server conf
	srv := &http.Server{
		Handler: corsHandler.Handler(r),
		Addr:    cfg.Server.Addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}

	logrus.Println(fmt.Sprintf("API Listening on %s", cfg.Server.Addr))
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logrus.Println("shutting down")
	os.Exit(0)

	return nil
}

// ServeHTTP return instance of serve HTTP command object
func ServeHTTP() *cobra.Command {
	return routerCMD
}
