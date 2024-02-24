package constants

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	ErrUneditableAuthenticationCode            = "500"
	ErrUneditableApprovedStudentSkpiCode       = "501"
	ErrIdenticalSemesterIdCode                 = "502"
	ErrIdenticalLeaveCode                      = "503"
	ErrLeaveIsNotActiveCode                    = "504"
	ErrLeaveActiveUneditableCode               = "505"
	ErrInvalidResignPurposeCode                = "506"
	ErrLecturerIsResignedCode                  = "507"
	ErrLecturerHasNoStudyProgramCode           = "508"
	ErrDifferentSubjectClassesCode             = "509"
	ErrNoActiveCurriculumCode                  = "510"
	ErrUneditableClassActiveSemesterCode       = "511"
	ErrUneditableClassMaterialCode             = "512"
	ErrUneditableClassEventCode                = "513"
	ErrUneditableClassWorkCode                 = "514"
	ErrUneditableClassDiscussionCode           = "515"
	ErrUneditableClassDiscussionCommentCode    = "516"
	ErrUneditableClassExamCode                 = "517"
	ErrInvalidStudentStatusCode                = "518"
	ErrStudentHasNoStudyProgramCode            = "519"
	ErrApprovedStudyPlanCode                   = "520"
	ErrNotGradingTimeCode                      = "521"
	ErrUngradeableClassCode                    = "522"
	ErrUneditableClassAnnouncementCode         = "523"
	ErrUneditableSharedFileCode                = "524"
	ErrClassExceedRoomCapacityCode             = "525"
	ErrInvalidAnnouncementTypeCode             = "526"
	ErrInvalidNimNumberCode                    = "527"
	ErrInvalidDiktiStudyProgramCodeCode        = "528"
	ErrActiveThesisExistsCode                  = "529"
	ErrActiveThesisDefenseRequestExistsCode    = "530"
	ErrActiveThesisDefenseRequestNotExistsCode = "530"
	ErrLecturerNotAssignedCode                 = "531"
	ErrInvalidLectureCode                      = "532"
	ErrClassWorkDeadlineCode                   = "533"
	ErrClassExamDeadlineCode                   = "534"
	ErrSubjectNotMbkmCode                      = "535"
	ErrInsufficientCreditConversionCode        = "536"
	ErrClassNotGradedCode                      = "537"
	ErrUneditableLectureCode                   = "538"
)

var (
	ErrUneditableAuthentication            = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableAuthenticationCode, "You cannot edit this authentication data.")
	ErrUneditableApprovedStudentSkpi       = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableApprovedStudentSkpiCode, "You cannot edit approved skpi.")
	ErrIdenticalSemesterId                 = Error(http.StatusBadRequest, codes.InvalidArgument, ErrIdenticalSemesterIdCode, "Identical semester ID.")
	ErrIdenticalLeave                      = Error(http.StatusBadRequest, codes.InvalidArgument, ErrIdenticalLeaveCode, "leave for this person on this date is exist.")
	ErrLeaveIsNotActive                    = Error(http.StatusBadRequest, codes.InvalidArgument, ErrLeaveIsNotActiveCode, "this leave is not active")
	ErrLeaveActiveUneditable               = Error(http.StatusBadRequest, codes.InvalidArgument, ErrLeaveActiveUneditableCode, "cannot edit/delete active leave")
	ErrInvalidResignPurpose                = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidResignPurposeCode, "invalid resignation purpose")
	ErrLecturerIsResigned                  = Error(http.StatusBadRequest, codes.InvalidArgument, ErrLecturerIsResignedCode, "lecturer already resigned")
	ErrLecturerHasNoStudyProgram           = Error(http.StatusBadRequest, codes.InvalidArgument, ErrLecturerHasNoStudyProgramCode, "lecturer has no study program")
	ErrDifferentSubjectClasses             = Error(http.StatusBadRequest, codes.InvalidArgument, ErrDifferentSubjectClassesCode, "some classes have different subjects")
	ErrNoActiveCurriculum                  = Error(http.StatusBadRequest, codes.InvalidArgument, ErrNoActiveCurriculumCode, "no active curriculum")
	ErrUneditableClassActiveSemester       = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassActiveSemesterCode, "class in active semester cannot be edited")
	ErrUneditableClassMaterial             = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassMaterialCode, "materials for this class is not yours to edit")
	ErrUneditableClassEvent                = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassEventCode, "events for this class is not yours to edit")
	ErrUneditableClassWork                 = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassWorkCode, "works for this class is not yours to edit")
	ErrUneditableClassDiscussion           = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassDiscussionCode, "discussions for this class is not yours to edit")
	ErrUneditableClassDiscussionComment    = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassDiscussionCommentCode, "this comment is not yours to edit")
	ErrUneditableClassExam                 = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassExamCode, "exams for this class is not yours to edit")
	ErrInvalidStudentStatus                = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidStudentStatusCode, "invalid student status")
	ErrStudentHasNoStudyProgram            = Error(http.StatusBadRequest, codes.InvalidArgument, ErrStudentHasNoStudyProgramCode, "student has no study program")
	ErrNotGradingTime                      = Error(http.StatusBadRequest, codes.InvalidArgument, ErrNotGradingTimeCode, "it's not time to grade the students yet")
	ErrUngradeableClass                    = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUngradeableClassCode, "this class is not yours to grade")
	ErrApprovedStudyPlan                   = Error(http.StatusBadRequest, codes.InvalidArgument, ErrApprovedStudyPlanCode, "student already has approved study plan for this semester")
	ErrUneditableClassAnnouncement         = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableClassAnnouncementCode, "announcements for this class is not yours to edit")
	ErrUneditableSharedFile                = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableSharedFileCode, "this shared file is not yours to edit")
	ErrClassExceedRoomCapacity             = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClassExceedRoomCapacityCode, "participant for this class exceeds the room's capacity")
	ErrInvalidAnnouncementType             = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidAnnouncementTypeCode, "invalid announcement type")
	ErrInvalidNimNumber                    = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidNimNumberCode, "invalid nim number")
	ErrInvalidDiktiStudyProgramCode        = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidDiktiStudyProgramCodeCode, "invalid dikti study program code")
	ErrActiveThesisExists                  = Error(http.StatusBadRequest, codes.InvalidArgument, ErrActiveThesisExistsCode, "active thesis is exist.")
	ErrActiveThesisDefenseRequestExists    = Error(http.StatusBadRequest, codes.InvalidArgument, ErrActiveThesisDefenseRequestExistsCode, "active thesis defense request is exist.")
	ErrActiveThesisDefenseRequestNotExists = Error(http.StatusBadRequest, codes.InvalidArgument, ErrActiveThesisDefenseRequestNotExistsCode, "active thesis defense request is not exist.")
	ErrLecturerNotAssigned                 = Error(http.StatusBadRequest, codes.InvalidArgument, ErrLecturerNotAssignedCode, "lecturer is not assigned to this class.")
	ErrInvalidLecture                      = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInvalidLectureCode, "invalid lecture.")
	ErrClassWorkDeadline                   = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClassWorkDeadlineCode, "you have exceed deadline for this work.")
	ErrClassExamDeadline                   = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClassExamDeadlineCode, "you have exceed deadline for this exam.")
	ErrSubjectNotMbkm                      = Error(http.StatusBadRequest, codes.InvalidArgument, ErrSubjectNotMbkmCode, "subject is not mbkm.")
	ErrClassNotGraded                      = Error(http.StatusBadRequest, codes.InvalidArgument, ErrClassNotGradedCode, "class is not graded")
	ErrInsufficientCreditConversion        = Error(http.StatusBadRequest, codes.InvalidArgument, ErrInsufficientCreditConversionCode, "insufficient credit for conversion.")
	ErrUneditableLecture                   = Error(http.StatusBadRequest, codes.InvalidArgument, ErrUneditableLectureCode, "this lecture is uneditable.")
)
