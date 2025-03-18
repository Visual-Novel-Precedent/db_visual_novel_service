package atlas

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"db_novel_service/internal/models"
	"io"
	"os"
)

func StartAtlasSchemaValidation() bool {
	stmts, err := gormschema.New("postgres").Load(
		&models.Media{},
		&models.Request{},
		&models.Character{},
		&models.Node{},
		&models.Admin{},
		&models.Chapter{},
		&models.Player{},
	)

	if err != nil {
		io.WriteString(os.Stderr, "failed to load gorm schema: "+err.Error()+"\n")
		return false
	}

	io.WriteString(os.Stdout, stmts)
	return true
}
