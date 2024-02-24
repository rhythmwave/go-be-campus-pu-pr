package excel

const (
	getStudyProgramDistributionData = `
		SELECT 
			students.id id,  
			students.nim_number student_nim_number,  
			students.name student_name,
			subjects.code subject_code,
			subjects.name subject_name,
			classes.name class_name,
			student_classes.grade_code,
			student_classes.total_attendance,
			classes.total_lecture_done,
			(CASE WHEN classes.total_lecture_done > 0 AND student_classes.total_attendance > 0 THEN (student_classes.total_attendance::decimal / classes.total_lecture_done::decimal) 
			ELSE 0 END) percentage
		from students
		join student_subjects on student_subjects.student_id = students.id
		join subjects on subjects.id = student_subjects.subject_id
		join student_classes on student_classes.student_id = students.id
		join classes on classes.id = student_classes.class_id
		where  students.study_program_id in (SELECT UNNEST($1::uuid[]))
		and classes.semester_id = $2
		order by students.name
	`
)
