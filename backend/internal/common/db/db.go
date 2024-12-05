package db

import (
	"docsfly/internal/common/constants"
	_ "docsfly/internal/common/startup"
	"docsfly/internal/types"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	if _, err = os.Stat(constants.ConfInst.Database.Resource); os.IsNotExist(err) {
		dir := filepath.Dir(constants.ConfInst.Database.Resource)

		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create directory: %v", err)
		}

		file, err := os.Create(constants.ConfInst.Database.Resource)
		if err != nil {
			panic("failed to create file")
		}
		defer file.Close()
	}

	DB, err = gorm.Open(sqlite.Open(constants.ConfInst.Database.Resource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// modelsToDrop := []interface{}{
	// 	&types.Category{},
	// 	&types.Document{},
	// 	&types.Comment{},
	// 	&types.Announce{},
	// 	&types.User{},
	// }

	// for _, model := range modelsToDrop {
	// 	if err := DB.Migrator().DropTable(model); err != nil {

	// 	}
	// }

	err = DB.AutoMigrate(
		&types.Category{}, &types.Document{},
		&types.Comment{}, &types.Announce{}, &types.Statistics{},
	)

	managers := make([]Manager, 0)
	managers = append(
		managers,
		NewDocumentManager(constants.ConfInst.Resource.Documents),
		NewAnnounceManager(
			filepath.Join(constants.ConfInst.Resource.Announces, constants.ConfInst.Resource.MetaFile),
		),
		NewCommentManager(
			filepath.Join(constants.ConfInst.Resource.Comments, constants.ConfInst.Resource.MetaFile),
		),
		NewUserManager(),
	)

	for _, manager := range managers {
		manager.Read()
		manager.Build()
		manager.Reduce()
	}

}
