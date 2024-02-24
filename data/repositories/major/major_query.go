package major

const (
	getListQuery = `
		SELECT
			m.id,
			f.name AS faculty_name,
			m.name
		FROM majors m
		JOIN faculties f ON f.id = m.faculty_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM majors m
		JOIN faculties f ON f.id = m.faculty_id
	`

	getDetailQuery = `
		SELECT
			m.id,
			f.id AS faculty_id,
			f.name AS faculty_name,
			m.name,
			m.short_name,
			m.english_name,
			m.english_short_name,
			m.address,
			m.phone_number,
			m.fax,
			m.email,
			m.experiment_building_area,
			m.contact_person,
			m.lecture_hall_count,
			m.lecture_hall_area,
			m.laboratorium_count,
			m.laboratorium_area,
			m.administration_room_area,
			m.permanent_lecturer_room_area,
			m.book_copy_count,
			m.book_count
		FROM majors m
		JOIN faculties f ON f.id = m.faculty_id
		WHERE m.id = $1
	`

	createQuery = `
		INSERT INTO majors (
			faculty_id,
			name,
			short_name,
			english_name,
			english_short_name,
			address,
			phone_number,
			fax,
			email,
			contact_person,
			experiment_building_area,
			lecture_hall_area,
			lecture_hall_count,
			laboratorium_area,
			laboratorium_count,
			permanent_lecturer_room_area,
			administration_room_area,
			book_count,
			book_copy_count,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);
	`

	updateQuery = `
		UPDATE majors SET
			faculty_id = $1,
			name = $2,
			short_name = $3,
			english_name = $4,
			english_short_name = $5,
			address = $6,
			phone_number = $7,
			fax = $8,
			email = $9,
			contact_person = $10,
			experiment_building_area = $11,
			lecture_hall_area = $12,
			lecture_hall_count = $13,
			laboratorium_area = $14,
			laboratorium_count = $15,
			permanent_lecturer_room_area = $16,
			administration_room_area = $17,
			book_count = $18,
			book_copy_count = $19,
			updated_by = $20
		WHERE id = $21
	`

	deleteQuery = `
		DELETE FROM majors WHERE id = $1
	`
)
