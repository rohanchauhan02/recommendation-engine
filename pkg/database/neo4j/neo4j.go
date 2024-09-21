package database

import (
	"context"
	"fmt"

	driver "github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rohanchauhan02/recommendation-engine/pkg/config"
)

type Neo4jSess interface {
	Init(context.Context) (driver.DriverWithContext, error)
}

type neo4j struct {
	conf config.ImmutableConfig
}

func (db *neo4j) Init(ctx context.Context) (driver.DriverWithContext, error) {
	neo4jConf := db.conf.GetDatabase().Neo4j
	neo4jSess, err := driver.NewDriverWithContext(
		neo4jConf.Uri,
		driver.BasicAuth(
			neo4jConf.User,
			neo4jConf.Password,
			"",
		))
	if err != nil {
		return nil, err
	}

	err = neo4jSess.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Neo4j connection established.")
	return neo4jSess, nil
}

func NewNeo4j(conf config.ImmutableConfig) Neo4jSess {
	return &neo4j{
		conf: conf,
	}
}
