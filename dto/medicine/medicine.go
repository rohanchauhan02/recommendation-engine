package dto

type CreateMedicineRequest struct {
	Medicines  []Medicine `json:"medicines"`
	Categories []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"categories"`
	Manufacturers []struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	} `json:"manufacturers"`
}

type Medicine struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	ActiveIngredients []struct {
		Name   string `json:"name"`
		Dosage string `json:"dosage"`
	} `json:"active_ingredients"`
	Manufacturer struct {
		Name string `json:"name"`
	} `json:"manufacturer"`
	Category struct {
		Name string `json:"name"`
	} `json:"category"`
	Patients []struct {
		Name             string `json:"name"`
		Age              int    `json:"age"`
		PrescriptionDate string `json:"prescription_date"`
	} `json:"patients"`
}
