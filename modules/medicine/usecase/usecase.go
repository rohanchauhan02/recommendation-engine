package usecase

import (
	"github.com/rohanchauhan02/recommendation-engine/dto/medicine"
	"github.com/rohanchauhan02/recommendation-engine/modules/medicine"
)

type medicineUsecase struct {
	medicineRepo medicine.Repository
}

func NewMedicineUsecase(mr medicine.Repository) medicine.Usecase {
	return &medicineUsecase{
		medicineRepo: mr,
	}
}

func (mu *medicineUsecase) AddMedicine(req *dto.CreateMedicineRequest) error {
	return mu.medicineRepo.AddMedicine(req)
}