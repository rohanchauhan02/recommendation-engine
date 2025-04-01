func (u *usecase) AddMedicine(medicine *Medicine) error {
    err := u.repo.Add(medicine)
    if err != nil {
        return err
    }
    return nil
}