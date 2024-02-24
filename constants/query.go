package constants

const (
	DefaultLimit     = uint32(20)
	DefaultPage      = uint32(1)
	DefaultUnlimited = ^uint32(0)
)

const (
	Ascending          = "ASC"
	Descending         = "DESC"
	DescendingNullLast = "DESC NULLS LAST"
	Random             = "RAND()"
)

const (
	ID                                 = "id"
	Name                               = "name"
	SignatureID                        = "s.id"
	Date                               = "date"
	Sort                               = "sort"
	SignatureCreatedAt                 = "s.created_at"
	CreatedAt                          = "created_at"
	UpdatedAt                          = "updated_at"
	EmployeeCreatedAt                  = "e.created_at"
	EmployeeInvitationRequestCreatedAt = "eir.created_at"
	EmployeeUpdatedAt                  = "e.updated_at"
	EmployeeJoinDate                   = "e.join_date"
	PositionCreatedAt                  = "p.created_at"
	PositionUpdatedAt                  = "p.updated_at"
	UserName                           = "u.name"
	UserCreatedAt                      = "u.created_at"
	WorkspaceCreatedAt                 = "w.created_at"
	UserUpdatedAt                      = "u.updated_at"
	ClockinTime                        = "a.clockin_time"
	DueDate                            = "due_date"
	// LastUpdatedMessage          = "last_updated_message"
	// StartDate                   = "start_date"
	// DateCreated                 = "date_created"
	// DepartureDate               = "tgl_berangkat"
	// EventMonth                  = "event.bulan"
	// EventYear                   = "event.tahun"
	// EventYearMonth              = "event.tahun, event.bulan"
	// EventReimbursementCreatedAt = "r.created_at"
	// EventReimbursementUpdatedAt = "r.updated_at"
	// Order                       = "`order`"
	// NameOnUsers                 = "user.nama"
	// CreatedAtOnUsers            = "user.created_at"
	// IdOnlogUser                 = "log.id"
	// CreatedAtOnSubInstance      = "sub_instance.created_at"
	// CreatedAtOnInstance         = "instance.created_at"
	// IdOnUnit                    = "unit.id"
	// IdOnSubUnit                 = "sub_unit.id"
	// IdOnBanner                  = "banner.id"
	// DateCreatedOnEvent          = "event.date_created"
	// DateCreatedOnOfficialMemo   = "official_memo.date_created"
	// DateCreatedOnInvitation     = "invitation.date_created"
	// CreatedAtOnAttendancePolicy = "attendance.created_at"
	// CreatedAtOnDoc              = "doc.created_at"
)

func ValidOrderValue() []string {
	return []string{
		Ascending,
		Descending,
		DescendingNullLast,
		Random,
	}
}

func ValidOrderKey() []string {
	return []string{
		ID,
		Name,
		Date,
		UpdatedAt,
		CreatedAt,
		EmployeeCreatedAt,
		EmployeeUpdatedAt,
		EmployeeJoinDate,
		PositionCreatedAt,
		PositionUpdatedAt,
		UserCreatedAt,
		UserUpdatedAt,
		ClockinTime,
		SignatureID,
		SignatureCreatedAt,
		EmployeeInvitationRequestCreatedAt,
		UserName,
		WorkspaceCreatedAt,
		Sort,
		DueDate,
		Name,
		// LastUpdatedMessage,
		// StartDate,
		// DateCreated,
		// DepartureDate,
		// EventMonth,
		// EventYear,
		// EventYearMonth,
		// EventReimbursementCreatedAt,
		// EventReimbursementUpdatedAt,
		// Order,
		// NameOnUsers,
		// CreatedAtOnUsers,
		// IdOnlogUser,
		// CreatedAtOnInstance,
		// CreatedAtOnSubInstance,
		// IdOnUnit,
		// IdOnSubUnit,
		// IdOnBanner,
		// DateCreatedOnEvent,
		// DateCreatedOnOfficialMemo,
		// DateCreatedOnInvitation,
		// CreatedAtOnAttendancePolicy,
		// CreatedAtOnDoc,
	}
}
