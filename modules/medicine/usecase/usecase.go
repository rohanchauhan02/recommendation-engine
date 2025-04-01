package usecase

import (
    "fmt"
    "github.com/your_project/models"
)

var id int = 0

var medicines []models.Medicine

func AddMedicine(medicine models.Medicine) models.Medicine {
    id++
    medicine.ID = id
    medicines = append(medicines, medicine)
    fmt.Println("Medicine added successfully")
    return medicine
}