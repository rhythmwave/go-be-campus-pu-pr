package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudentSkpiRequest struct {
	StudyProgramId string
	Name           string
	NimNumber      int64
	NimNumberFrom  int64
	NimNumberTo    int64
	IsApproved     *bool
}

type GetStudentSkpi struct {
	Id                           string
	StudentId                    string
	StudentNimNumber             int64
	StudentName                  string
	StudentStudyProgramId        string
	StudentStudyProgramName      string
	StudentDiktiStudyProgramCode string
	IsApproved                   bool
}

type StudentSkpiListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentSkpi
}

type GetStudentSkpiDetailAchievement struct {
	Id   string
	Name string
	Year uint32
}

type GetStudentSkpiDetailOrganization struct {
	Id            string
	Name          string
	Position      string
	ServiceLength uint32
}

type GetStudentSkpiDetailCertificate struct {
	Id   string
	Name string
}

type GetStudentSkpiDetailCharacterBuilding struct {
	Id   string
	Name string
}

type GetStudentSkpiDetailInternship struct {
	Id   string
	Name string
}

type GetStudentSkpiDetailLanguage struct {
	Id    string
	Name  string
	Score string
	Date  time.Time
}

type GetStudentSkpiDetail struct {
	Id                           string
	StudentId                    string
	StudentNimNumber             int64
	StudentName                  string
	StudentStudyProgramId        string
	StudentStudyProgramName      string
	StudentDiktiStudyProgramCode string
	SkpiNumber                   *string
	IsApproved                   bool
	AchievementPath              *string
	AchievementPathType          *string
	AchievementUrl               string
	OrganizationPath             *string
	OrganizationPathType         *string
	OrganizationUrl              string
	CertificatePath              *string
	CertificatePathType          *string
	CertificateUrl               string
	LanguagePath                 *string
	LanguagePathType             *string
	LanguageUrl                  string
	Achievements                 []GetStudentSkpiDetailAchievement
	Organizations                []GetStudentSkpiDetailOrganization
	Certificates                 []GetStudentSkpiDetailCertificate
	CharacterBuildings           []GetStudentSkpiDetailCharacterBuilding
	Internships                  []GetStudentSkpiDetailInternship
	Languages                    []GetStudentSkpiDetailLanguage
}

type UpsertStudentSkpiAchievement struct {
	Name string
	Year uint32
}

type UpsertStudentSkpiOrganization struct {
	Name          string
	Position      string
	ServiceLength uint32
}

type UpsertStudentSkpiCertificate struct {
	Name string
}

type UpsertStudentSkpiCharacterBuilding struct {
	Name string
}

type UpsertStudentSkpiInternship struct {
	Name string
}

type UpsertStudentSkpiLanguage struct {
	Name  string
	Score string
	Date  time.Time
}

type UpsertStudentSkpi struct {
	StudentId            string
	AchievementPath      string
	AchievementPathType  string
	OrganizationPath     string
	OrganizationPathType string
	CertificatePath      string
	CertificatePathType  string
	LanguagePath         string
	LanguagePathType     string
	Achievements         []UpsertStudentSkpiAchievement
	Organizations        []UpsertStudentSkpiOrganization
	Certificates         []UpsertStudentSkpiCertificate
	CharacterBuildings   []UpsertStudentSkpiCharacterBuilding
	Internships          []UpsertStudentSkpiInternship
	Languages            []UpsertStudentSkpiLanguage
}

type ApproveStudentSkpi struct {
	Id         string
	SkpiNumber string
}
