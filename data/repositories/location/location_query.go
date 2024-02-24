package location

const (
	getListCountryQuery = `
		SELECT 
			l.id, 
			l.name
		FROM countries l
	`

	getListProvinceQuery = `
		SELECT 
			l.id, 
			l.name
		FROM provinces l
	`

	getListRegencyQuery = `
		SELECT 
			l.id, 
			l.name
		FROM regencies l
	`

	getListDistrictQuery = `
		SELECT 
			l.id, 
			l.name
		FROM districts l
	`

	getListVillageQuery = `
		SELECT 
			l.id, 
			l.name
		FROM villages l
	`

	countListCountryQuery = `
		SELECT COUNT(1) FROM countries l
	`

	countListProvinceQuery = `
		SELECT COUNT(1) FROM provinces l
	`

	countListRegencyQuery = `
		SELECT COUNT(1) FROM regencies l
	`

	countListDistrictQuery = `
		SELECT COUNT(1) FROM districts l
	`

	countListVillageQuery = `
		SELECT COUNT(1) FROM villages l
	`

	tempCreateDataQuery = `
		INSERT INTO temp_data (
			title,
			body
		) VALUES (
			:title,
			:body
		)
	`

	tempGetDataQuery = `
		SELECT 
			l.id,
			l.title,
			l.body
		FROM temp_data l
	`

	countTempGetDataQuery = `
		SELECT COUNT(1) FROM temp_data l
	`
)
