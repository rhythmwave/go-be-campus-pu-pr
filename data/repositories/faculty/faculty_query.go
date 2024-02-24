package faculty

const (
	getListQuery = `
		SELECT
			f.id,
			f.name,
			f.short_name
		FROM faculties f
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM faculties f
	`

	getDetailQuery = `
		SELECT
			f.id,
			f.name,
			f.short_name,
			f.english_name,
			f.english_short_name,
			f.address,
			f.phone_number,
			f.fax,
			f.email,
			f.experiment_building_area,
			f.contact_person,
			f.lecture_hall_count,
			f.lecture_hall_area,
			f.laboratorium_count,
			f.laboratorium_area,
			f.administration_room_area,
			f.permanent_lecturer_room_area,
			f.book_copy_count,
			f.book_count
		FROM faculties f
		WHERE f.id = $1
	`

	createQuery = `
		INSERT INTO faculties (
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
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);
	`

	updateQuery = `
		UPDATE faculties SET
			name = $1,
			short_name = $2,
			english_name = $3,
			english_short_name = $4,
			address = $5,
			phone_number = $6,
			fax = $7,
			email = $8,
			contact_person = $9,
			experiment_building_area = $10,
			lecture_hall_area = $11,
			lecture_hall_count = $12,
			laboratorium_area = $13,
			laboratorium_count = $14,
			permanent_lecturer_room_area = $15,
			administration_room_area = $16,
			book_count = $17,
			book_copy_count = $18,
			updated_by = $19
		WHERE id = $20
	`

	deleteQuery = `
		DELETE FROM faculties WHERE id = $1
	`
)
