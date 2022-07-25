package rpc

import (
	context "context"

	"github.com/rog-golang-buddies/go-automatic-apps/database"
)

type databaseService struct {
    UnimplementedDatabaseServiceServer
}


func NewDatabaseService() DatabaseServiceServer {
	return &databaseService{}
}

func (s *databaseService) ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error) {
    tables := database.GetTables()
	return &ListTablesResponse{
		Tables:        tables,
	}, nil
}
