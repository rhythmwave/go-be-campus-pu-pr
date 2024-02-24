package room

const (
	getListQuery = `
		SELECT 
			r.id,
			r.code,
			r.name,
			r.capacity,
			r.is_laboratory
		FROM rooms r
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM rooms r
		%s
	`

	getDetailQuery = `
		SELECT 
			r.id,
			b.id AS building_id,
			b.code AS building_code,
			b.name AS building_name,
			r.code,
			r.name,
			r.capacity,
			r.exam_capacity,
			r.purpose,
			r.is_usable,
			r.area,
			r.phone_number,
			r.facility,
			r.remarks,
			r.owner,
			r.location,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			r.is_laboratory
		FROM rooms r
		JOIN buildings b ON b.id = r.building_id
		LEFT JOIN study_programs sp ON sp.id = r.study_program_id
		WHERE r.id = $1
	`

	getDetailByRoomIdsQuery = `
		SELECT 
			r.id,
			b.id AS building_id,
			b.code AS building_code,
			b.name AS building_name,
			r.code,
			r.name,
			r.capacity,
			r.exam_capacity,
			r.purpose,
			r.is_usable,
			r.area,
			r.phone_number,
			r.facility,
			r.remarks,
			r.owner,
			r.location,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			r.is_laboratory
		FROM rooms r
		JOIN buildings b ON b.id = r.building_id
		LEFT JOIN study_programs sp ON sp.id = r.study_program_id
		WHERE r.id IN (SELECT UNNEST($1::uuid[]))
	`

	createQuery = `
		INSERT INTO rooms (
			building_id,
			code,
			name,
			capacity,
			exam_capacity,
			is_usable,
			area,
			phone_number,
			facility,
			remarks,
			purpose,
			owner,
			location,
			study_program_id,
			is_laboratory,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`

	updateQuery = `
		UPDATE rooms SET
			code = $1,
			name = $2,
			capacity = $3,
			exam_capacity = $4,
			is_usable = $5,
			area = $6,
			phone_number = $7,
			facility = $8,
			remarks = $9,
			purpose = $10,
			owner = $11,
			location = $12,
			study_program_id = $13,
			updated_by = $14
		WHERE id = $15
	`

	deleteQuery = `
		DELETE FROM rooms WHERE id = $1
	`

	getScheduleQuery = `
		SELECT
			DISTINCT(r.id) AS room_id,
			r.name AS room_name
		FROM rooms r
		LEFT JOIN lectures l ON l.room_id = r.id
		LEFT JOIN classes c ON c.id = l.class_id
	`
	countScheduleQuery = `
		WITH d AS (
			SELECT
				DISTINCT(r.id) AS room_id,
				r.name AS room_name
			FROM rooms r
			LEFT JOIN lectures l ON l.room_id = r.id
			LEFT JOIN classes c ON c.id = l.class_id
			%s
		) SELECT COUNT(1) FROM d
	`

	getScheduleByRoomIdsQuery = `
		SELECT
			r.id AS room_id,
			r.name AS room_name,
			l.lecture_plan_date,
			l.lecture_plan_start_time AS start_time,
			l.lecture_plan_end_time AS end_time,
			su.name AS subject_name,
			c.name AS class_name,
			sp.name AS study_program_name
		FROM rooms r
		LEFT JOIN lectures l ON l.room_id = r.id
		LEFT JOIN classes c ON c.id = l.class_id
		LEFT JOIN subjects su ON su.id = c.subject_id
		LEFT JOIN curriculums cu ON cu.id = su.curriculum_id
		LEFT JOIN study_programs sp ON sp.id = cu.study_program_id
		%s
		ORDER BY r.id ASC, l.lecture_plan_date ASC, l.lecture_plan_start_time ASC
	`
)
