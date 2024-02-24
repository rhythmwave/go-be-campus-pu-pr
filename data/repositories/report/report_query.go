package report

const (
	getActiveSemesterStudentStatusQuery = `
		WITH
			pst AS (
				SELECT UNNEST($1::text[]) AS status
			),
			mst AS (
				SELECT UNNEST($2::text[]) AS status
			),
			d AS (
				SELECT COUNT(s.id) AS total, sp.id AS study_program_id, pst.status 
				FROM pst
				JOIN study_programs sp ON sp.id IN (SELECT UNNEST($4::uuid[]))
				LEFT JOIN students s ON s.status = pst.status AND s.study_program_id = sp.id
				GROUP BY pst.status, sp.id
					UNION
				SELECT COUNT(s.id) AS total, sp.id AS study_program_id, mst.status 
				FROM mst
				JOIN study_programs sp ON sp.id IN (SELECT UNNEST($4::uuid[]))
				LEFT JOIN students s ON s.status = mst.status AND s.study_program_id = sp.id AND s.status_change_semester_id = $3
				GROUP BY mst.status, sp.id
			)
		SELECT
			s.id AS semester_id,
			s.semester_type,
			s.semester_start_year,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.type AS dikti_study_program_type,
			dsp.code AS dikti_study_program_code,
			sl.short_name AS study_level_short_name,
			d.status,
			d.total
		FROM d
		JOIN semesters s ON s.id = $3
		JOIN study_programs sp ON sp.id = d.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		ORDER BY sp.name, d.status
	`

	getStudentStatusQuery = `
		SELECT
			rss.semester_id,
			s.semester_type,
			s.semester_start_year,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.type AS dikti_study_program_type,
			dsp.code AS dikti_study_program_code,
			sl.short_name AS study_level_short_name,
			rss.status,
			rss.total
		FROM report_student_statuses rss
		JOIN semesters s ON s.id = rss.semester_id
		JOIN study_programs sp ON sp.id = rss.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE rss.semester_id = $1 AND rss.study_program_id IN (SELECT UNNEST($2::uuid[]))
		ORDER BY sp.name, rss.status
	`

	getStudentClassGradeQuery = `
		SELECT
			s.id AS subject_id,
			gt.code AS grade_code,
			COALESCE(rscg.total, 0) AS total
		FROM subjects s
		CROSS JOIN grade_types gt
		LEFT JOIN report_student_class_grades rscg ON rscg.semester_id = $1 AND rscg.subject_id = s.id AND rscg.grade_code = gt.code
		WHERE s.id IN (SELECT UNNEST($2::uuid[]))
		ORDER BY gt.code
	`

	getStudentProvinceQuery = `
		WITH d AS (
			SELECT GENERATE_SERIES($2::integer, $3::integer) AS student_force
		)
		SELECT
			p.id AS province_id,
			p.name AS province_name,
			d.student_force,
			COALESCE(SUM(rsp.total), 0) AS total
		FROM provinces p
		CROSS JOIN d
		LEFT JOIN report_student_provinces rsp ON rsp.student_force = d.student_force AND rsp.province_id = p.id AND rsp.study_program_id IN (SELECT UNNEST($1::uuid[]))
		GROUP BY p.id, d.student_force
		ORDER BY p.name, d.student_force
	`

	getStudentSchoolProvinceQuery = `
		WITH d AS (
			SELECT GENERATE_SERIES($2::integer, $3::integer) AS student_force
		)
		SELECT
			p.id AS province_id,
			p.name AS province_name,
			d.student_force,
			COALESCE(SUM(rsp.total), 0) AS total
		FROM provinces p
		CROSS JOIN d
		LEFT JOIN report_student_school_provinces rsp ON rsp.student_force = d.student_force AND rsp.province_id = p.id AND rsp.study_program_id IN (SELECT UNNEST($1::uuid[]))
		GROUP BY p.id, d.student_force
		ORDER BY p.name, d.student_force
	`
)
