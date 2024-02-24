package student

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type studentHandler struct {
	*service.ServiceCtx
}

func (d studentHandler) BulkCreate(w http.ResponseWriter, r *http.Request) {
	var result BulkCreateResponse

	ctx := r.Context()
	var in BulkCreateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	var req []objects.BulkCreateStudent
	for _, v := range in.GetData() {
		birthDate, errs := utils.StringToDate(v.GetBirthDate())
		if errs != nil {
			utils.PrintError(*errs)
			result = BulkCreateResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		fatherBirthDate, errs := utils.StringToDate(v.GetFatherBirthDate())
		if errs != nil {
			utils.PrintError(*errs)
			result = BulkCreateResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		motherBirthDate, errs := utils.StringToDate(v.GetMotherBirthDate())
		if errs != nil {
			utils.PrintError(*errs)
			result = BulkCreateResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}
		guardianBirthDate, errs := utils.StringToDate(v.GetGuardianBirthDate())
		if errs != nil {
			utils.PrintError(*errs)
			result = BulkCreateResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		req = append(req, objects.BulkCreateStudent{
			NimNumber:                       v.GetNimNumber(),
			Name:                            v.GetName(),
			Sex:                             v.GetSex(),
			MaritalStatus:                   v.GetMaritalStatus(),
			BirthRegencyId:                  v.GetBirthRegencyId(),
			BirthDate:                       birthDate,
			Religion:                        v.GetReligion(),
			Address:                         v.GetAddress(),
			Rt:                              v.GetRt(),
			Rw:                              v.GetRw(),
			VillageId:                       v.GetVillageId(),
			PostalCode:                      v.GetPostalCode(),
			IdNumber:                        v.GetIdNumber(),
			NisnNumber:                      v.GetNisnNumber(),
			MobilePhoneNumber:               v.GetMobilePhoneNumber(),
			Nationality:                     v.GetNationality(),
			DiktiStudyProgramCode:           v.GetDiktiStudyProgramCode(),
			SchoolName:                      v.GetSchoolName(),
			SchoolAddress:                   v.GetSchoolAddress(),
			SchoolProvinceId:                v.GetSchoolProvinceId(),
			SchoolMajor:                     v.GetSchoolMajor(),
			SchoolType:                      v.GetSchoolType(),
			SchoolGraduationYear:            v.GetSchoolGraduationYear(),
			FatherName:                      v.GetFatherName(),
			FatherIdNumber:                  v.GetFatherIdNumber(),
			FatherBirthDate:                 fatherBirthDate,
			FatherFinalAcademicBackground:   v.GetFatherFinalAcademicBackground(),
			FatherOccupation:                v.GetFatherOccupation(),
			MotherName:                      v.GetMotherName(),
			MotherIdNumber:                  v.GetMotherIdNumber(),
			MotherBirthDate:                 motherBirthDate,
			MotherFinalAcademicBackground:   v.GetMotherFinalAcademicBackground(),
			MotherOccupation:                v.GetMotherOccupation(),
			GuardianName:                    v.GetGuardianName(),
			GuardianIdNumber:                v.GetGuardianIdNumber(),
			GuardianBirthDate:               guardianBirthDate,
			GuardianFinalAcademicBackground: v.GetGuardianFinalAcademicBackground(),
			GuardianOccupation:              v.GetGuardianOccupation(),
		})
	}
	data, errs := d.StudentService.BulkCreate(ctx, req)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkCreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*BulkCreateResponseData{}
	for _, v := range data {
		nimNumber, err := strconv.Atoi(v.Username)
		if err != nil {
			errs := constants.ErrorInternalServer(err.Error())
			utils.PrintError(*errs)
			result = BulkCreateResponse{
				Meta: &Meta{
					Message: errs.Err.Error(),
					Status:  uint32(errs.HttpCode),
					Code:    errs.CustomCode,
				},
			}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		resultData = append(resultData, &BulkCreateResponseData{
			Name:      v.Name,
			NimNumber: int64(nimNumber),
			Password:  v.Password,
		})
	}

	result = BulkCreateResponse{
		Meta: &Meta{
			Message: "Bulk Create Student",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
