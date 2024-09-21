package usecase

import (
	dto "github.com/rohanchauhan02/recommendation-engine/dto/medicine"
	"github.com/rohanchauhan02/recommendation-engine/modules/medicine"
)

type usecase struct {
	repository medicine.Repository
}

func NewMedicineUsecase(repository *medicine.Repository) medicine.Usecase {
	return &usecase{
		repository: *repository,
	}
}

func (u *usecase) AddMedicine(req *dto.CreateMedicineRequest) error {
	return u.repository.AddMedicine(req)
}
