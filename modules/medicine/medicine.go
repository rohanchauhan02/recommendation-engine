package medicine

type Medicine struct {
	ID          int64
	Name        string
	Description string
}

func (m *Medicine) AddMedicine() error {
	if m.Name == "" || m.Description == "" {
		return errors.New("medicine name and description cannot be empty")
	}
	return nil
}