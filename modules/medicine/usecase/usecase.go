func (m *MedicineUsecase) AddMedicine(medicine *Medicine) error {
    err := m.MedicineRepo.Insert(medicine)
    if err != nil {
        return err
    }
    return nil
}