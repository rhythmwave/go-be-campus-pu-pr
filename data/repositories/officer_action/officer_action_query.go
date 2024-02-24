package officer_action

const (
	getListQuery = `
		SELECT
			oa.id,
			dt.id AS document_type_id,
			dt.name AS document_type_name,
			da.id AS document_action_id,
			da.action AS document_action_action,
			da.english_action AS document_action_english_action,
			o.id AS officer_id,
			o.name AS officer_name,
			o.title AS officer_title,
			o.english_title AS officer_english_title,
			osp.id AS officer_study_program_id,
			osp.name AS officer_study_program_name
		FROM officer_actions oa
		JOIN document_types dt ON dt.id = oa.document_type_id
		JOIN document_actions da ON da.id = oa.document_action_id
		JOIN officers o ON o.id = oa.officer_id
		LEFT JOIN study_programs osp ON osp.id = o.study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM officer_actions oa
		JOIN document_types dt ON dt.id = oa.document_type_id
		JOIN document_actions da ON da.id = oa.document_action_id
		JOIN officers o ON o.id = oa.officer_id
		LEFT JOIN study_programs osp ON osp.id = o.study_program_id
	`

	getDetailQuery = `
		SELECT 
			oa.id,
			dt.id AS document_type_id,
			dt.name AS document_type_name,
			da.id AS document_action_id,
			da.action AS document_action_action,
			da.english_action AS document_action_english_action,
			o.id AS officer_id,
			o.name AS officer_name,
			o.title AS officer_title,
			o.english_title AS officer_english_title,
			osp.id AS officer_study_program_id,
			osp.name AS officer_study_program_name
		FROM officer_actions oa
		JOIN document_types dt ON dt.id = oa.document_type_id
		JOIN document_actions da ON da.id = oa.document_action_id
		JOIN officers o ON o.id = oa.officer_id
		LEFT JOIN study_programs osp ON osp.id = o.study_program_id
		WHERE oa.id = $1
	`

	createQuery = `
		INSERT INTO officer_actions (
			document_type_id,
			document_action_id,
			officer_id,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE officer_actions SET
			document_type_id = $1,
			document_action_id = $2,
			officer_id = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM officer_actions WHERE id = $1
	`
)
