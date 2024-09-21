package repository

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	dto "github.com/rohanchauhan02/recommendation-engine/dto/medicine"
	"github.com/rohanchauhan02/recommendation-engine/modules/medicine"
)

type repository struct {
	neo4jSess neo4j.SessionWithContext
}

func NewMedicineRepository(neo4jSess neo4j.SessionWithContext) medicine.Repository {
	return &repository{
		neo4jSess: neo4jSess,
	}
}

func (r *repository) AddMedicine(req *dto.CreateMedicineRequest) error {
	// ctx := context.Background()
	for _, medicine := range req.Medicines {
		_, err := r.neo4jSess.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
			// Create Medicine node
			_, err := tx.Run(context.Background(),
				"MERGE (m:Medicine {id: $id}) "+
					"ON CREATE SET m.name = $name",
				map[string]interface{}{
					"id":   medicine.ID,
					"name": medicine.Name,
				})
			if err != nil {
				return nil, err
			}

			// Create ActiveIngredient nodes and relationships
			for _, ingredient := range medicine.ActiveIngredients {
				_, err = tx.Run(context.Background(),
					"MERGE (i:ActiveIngredient {name: $name}) "+
						"ON CREATE SET i.dosage = $dosage "+
						"MERGE (m)-[:CONTAINS_INGREDIENT]->(i)",
					map[string]interface{}{
						"name":   ingredient.Name,
						"dosage": ingredient.Dosage,
					})
				if err != nil {
					return nil, err
				}
			}

			// Create Manufacturer node and relationship
			_, err = tx.Run(context.Background(),
				"MERGE (mf:Manufacturer {name: $name}) "+
					"MERGE (m)-[:MANUFACTURED_BY]->(mf)",
				map[string]interface{}{
					"name": medicine.Manufacturer.Name,
				})
			if err != nil {
				return nil, err
			}

			// Create Category node and relationship
			_, err = tx.Run(context.Background(),
				"MERGE (c:Category {name: $name}) "+
					"MERGE (m)-[:BELONGS_TO_CATEGORY]->(c)",
				map[string]interface{}{
					"name": medicine.Category.Name,
				})
			if err != nil {
				return nil, err
			}

			// Create Patient nodes and relationships
			for _, patient := range medicine.Patients {
				_, err = tx.Run(context.Background(),
					"MERGE (p:Patient {name: $name, age: $age}) "+
						"MERGE (m)-[:PRESCRIBED_TO {prescription_date: $prescription_date}]->(p)",
					map[string]interface{}{
						"name":              patient.Name,
						"age":               patient.Age,
						"prescription_date": patient.PrescriptionDate,
					})
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		})

		if err != nil {
			log.Fatal("Error adding medicine data to Neo4j:", err)
		}
	}
	return nil
}
