package constants

const (
	AnnouncementTypeAcademicInformation = "academic_information"
	AnnouncementTypeStudentActivity     = "student_activity"
	AnnouncementTypeRegistration        = "registration"
	AnnouncementTypeEvent               = "event"
)

func ValidAnnouncementType() []string {
	return []string{
		AnnouncementTypeAcademicInformation,
		AnnouncementTypeStudentActivity,
		AnnouncementTypeRegistration,
		AnnouncementTypeEvent,
	}
}
