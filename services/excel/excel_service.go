package excel

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"strings"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
)

type excelService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (f excelService) StudyProgramDistributionDownload(ctx context.Context, studyProgramIds string, semesterID string) ([]byte, *constants.ErrorResponse) {
	var result = []byte{}

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}
	studyProgramIdsArray := []string{}
	studyProgramNames := []string{}
	if studyProgramIds == "" {
		studyPrograms, _, _ := f.StudyProgramRepo.GetList(ctx, tx, common.PaginationRequest{Page: 1, Limit: 1000}, "", "", "")
		for _, v := range studyPrograms {
			studyProgramIdsArray = append(studyProgramIdsArray, v.Id)
			studyProgramNames = append(studyProgramNames, v.Name)
		}
	} else {
		studyProgramIdsArray = strings.Split(studyProgramIds, ",")
		for _, v := range studyProgramIdsArray {
			studyProgram, err := f.StudyProgramRepo.GetDetail(ctx, tx, v, "", "")
			if err == nil {
				studyProgramNames = append(studyProgramNames, studyProgram.Name)
			}
		}
	}

	if semesterID == "" {
		return result, constants.Error(http.StatusBadRequest, codes.Internal, "", "Semester ID mandatory")
	}

	semester, errs := f.SemesterRepo.GetById(ctx, tx, semesterID)
	if errs != nil {
		return result, errs
	}

	modelResult, errs := f.ExcelRepo.StudyProgramDistributionDownload(ctx, tx, studyProgramIdsArray, semesterID)
	if errs != nil {
		return result, errs
	}

	var row int = 1

	sheet1Name := "Sheet1"
	xls := excelize.NewFile()
	index, err := xls.NewSheet(sheet1Name)
	if err != nil {
		return result, constants.ErrUnknown
	}
	xlsStyle := utils.NewExcelStyle(xls)
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("Semester: %s", fmt.Sprintf("%s %d", semester.SemesterType, semester.SemesterStartYear)))
	xls.MergeCell(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row))
	xls.SetCellStyle(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row), xlsStyle.Bold)
	row++
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("Program Studi: %s", strings.Join(studyProgramNames, ", ")))
	xls.MergeCell(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row))
	xls.SetCellStyle(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row), xlsStyle.Bold)
	row++
	row++
	row++
	xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", row), "No")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("B%d", row), "NIM")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("C%d", row), "Nama")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("D%d", row), "Kode MK")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("E%d", row), "Mata Kuliah")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("F%d", row), "Kelas")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("G%d", row), "Nilai")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("H%d", row), "Jumlah Hadir")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("I%d", row), "Jumlah Terlaksana")
	xls.SetCellValue(sheet1Name, fmt.Sprintf("J%d", row), "Persentase")
	xls.SetCellStyle(sheet1Name, fmt.Sprintf("A%d", row), fmt.Sprintf("J%d", row), xlsStyle.Bold)

	for i, v := range modelResult {
		row++

		var gradeCode string = ""
		if v.GradeCode != nil {
			gradeCode = *v.GradeCode
		}
		xls.SetCellValue(sheet1Name, fmt.Sprintf("A%d", row), i+1)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("B%d", row), v.StudentNimNumber)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("C%d", row), v.StudentName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("D%d", row), v.SubjectCode)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("E%d", row), v.SubjectName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("F%d", row), v.ClassName)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("G%d", row), gradeCode)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("H%d", row), v.TotalAttendance)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("I%d", row), v.TotalLectureDone)
		xls.SetCellValue(sheet1Name, fmt.Sprintf("J%d", row), fmt.Sprintf("%.0f %%", v.Percentage*100))
	}

	xls.SetActiveSheet(index)
	var b *bytes.Buffer
	b, err = xls.WriteToBuffer()
	if err != nil {
		return result, constants.ErrUnknown
	}
	result = b.Bytes()

	return result, nil
}
