package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v5/neo4j"

	medicieneHandler "github.com/rohanchauhan02/recommendation-engine/modules/medicine/delivery/http"
	medicieneRepository "github.com/rohanchauhan02/recommendation-engine/modules/medicine/repository"
	medicieneUsecase "github.com/rohanchauhan02/recommendation-engine/modules/medicine/usecase"

	"github.com/rohanchauhan02/recommendation-engine/pkg/config"
	database "github.com/rohanchauhan02/recommendation-engine/pkg/database/neo4j"
)

func main() {

	e := echo.New()
	cfg := config.NewImmutableConfig()
	ctx := context.Background()

	neo4j := database.NewNeo4j(cfg)
	neo4jSes, err := neo4j.Init(ctx)
	if err != nil {
		log.Fatal("Error initializing Neo4j session:", err)
	}
	defer neo4jSes.Close(ctx)

	session := neo4jSes.NewSession(ctx, neo4jDriver.SessionConfig{
		AccessMode: neo4jDriver.AccessModeWrite,
		DatabaseName: cfg.GetDatabase().Neo4j.Name,
	})
	defer session.Close(ctx)
	medicieneRepo := medicieneRepository.NewMedicineRepository(session)
	medicineUcase := medicieneUsecase.NewMedicineUsecase(&medicieneRepo)
	medicieneHandler.NewMedicineHandler(e, &medicineUcase)
	// Start the Echo server
	fmt.Println(e.Start(fmt.Sprintf(":%s", cfg.GetPort())))
}
