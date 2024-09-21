package medicine

import dto "github.com/rohanchauhan02/recommendation-engine/dto/medicine"

type Usecase interface {
	AddMedicine(req *dto.CreateMedicineRequest) error
}

type Repository interface {
	AddMedicine(req *dto.CreateMedicineRequest) error
}
