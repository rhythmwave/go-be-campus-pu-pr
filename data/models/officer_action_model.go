package models

type GetOfficerAction struct {
	Id                          string  `db:"id"`
	DocumentTypeId              string  `db:"document_type_id"`
	DocumentTypeName            string  `db:"document_type_name"`
	DocumentActionId            string  `db:"document_action_id"`
	DocumentActionAction        string  `db:"document_action_action"`
	DocumentActionEnglishAction string  `db:"document_action_english_action"`
	OfficerId                   string  `db:"officer_id"`
	OfficerName                 string  `db:"officer_name"`
	OfficerTitle                *string `db:"officer_title"`
	OfficerEnglishTitle         *string `db:"officer_english_title"`
	OfficerStudyProgramId       *string `db:"officer_study_program_id"`
	OfficerStudyProgramName     *string `db:"officer_study_program_name"`
}

type CreateOfficerAction struct {
	DocumentTypeId   string `db:"document_type_id"`
	DocumentActionId string `db:"document_action_id"`
	OfficerId        string `db:"officer_id"`
	CreatedBy        string `db:"created_by"`
}

type UpdateOfficerAction struct {
	Id               string `db:"id"`
	DocumentTypeId   string `db:"document_type_id"`
	DocumentActionId string `db:"document_action_id"`
	OfficerId        string `db:"officer_id"`
	UpdatedBy        string `db:"updated_by"`
}
