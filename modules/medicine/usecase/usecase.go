package usecase

import (
    "context"
    "modules/medicine/domain"
)

type MedicineUsecase struct {
    medicineRepo domain.MedicineRepository
}

func NewMedicineUsecase(mr domain.MedicineRepository) *MedicineUsecase {
    return &MedicineUsecase{
        medicineRepo: mr,
    }
}

// AddMedicine adds a new medicine to the repository
func (mu *MedicineUsecase) AddMedicine(ctx context.Context, medicine *domain.Medicine) error {
    err := mu.medicineRepo.Store(ctx, medicine)
    if err != nil {
        return err
    }
    return nil
}