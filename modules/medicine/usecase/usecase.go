package usecase

import (
	"context"
	"modules/medicine"
)

type MedicineUsecase struct {
	medicineRepo medicine.Repository
}

func NewMedicineUsecase(mr medicine.Repository) *MedicineUsecase {
	return &MedicineUsecase{
		medicineRepo: mr,
	}
}

func (mu *MedicineUsecase) AddMedicine(ctx context.Context, m *medicine.Medicine) error {
	return mu.medicineRepo.AddMedicine(ctx, m)
}