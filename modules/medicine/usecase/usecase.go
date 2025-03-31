func (m *MedicineUsecase) AddMedicine(ctx context.Context, medicine *domain.Medicine) error {
    err := m.MedicineRepo.Insert(ctx, medicine)
    if err != nil {
        return err
    }
    return nil
}