package constants

const (
	AuthenticationAdmin    = "admin"
	AuthenticationLecturer = "lecturer"
	AuthenticationStudent  = "student"
)

func AdminEditableAuth() []string {
	return []string{
		AuthenticationLecturer,
		AuthenticationStudent,
	}
}
