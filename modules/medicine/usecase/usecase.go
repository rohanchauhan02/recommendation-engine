type MedicineUsecase interface {
    AddMedicine(medicine *domain.Medicine) error
}