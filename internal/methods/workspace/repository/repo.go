package repository

import (
	"fmt"

	"github.com/gandarfh/httui-repl/external/database"
	"gorm.io/gorm"
)

type WorkspaceModel struct {
	gorm.Model
	Name string `db:"name" json:"name" validate:"required"`
	Uri  string `db:"uri" json:"uri" validate:"required"`
}

type WorkspaceRepo struct {
	Sql *gorm.DB
}

func NewWorkspaceRepo() (*WorkspaceRepo, error) {
	db, err := database.SqliteConnection()
	db.AutoMigrate(&WorkspaceModel{})

	if err != nil {
		fmt.Println("Deu ruim database")
		return nil, err
	}

	return &WorkspaceRepo{
		Sql: db,
	}, nil
}

func (repo *WorkspaceRepo) Create(ws *WorkspaceModel) {
	result := repo.Sql.Create(ws)

	if result.Error != nil {
		fmt.Println("Deu ruim criar dado")
	}
}

func (repo *WorkspaceRepo) List() *[]WorkspaceModel {
	ws := []WorkspaceModel{}
	result := repo.Sql.Find(&ws)

	if result.Error != nil {
		fmt.Println("Deu ruim criar dado")
	}

	return &ws
}
