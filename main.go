package main

import (
	"fmt"

	"github.com/skyrocketOoO/gorm-enhance-plugin/columnname"
	"github.com/skyrocketOoO/gorm-enhance-plugin/example/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.Create(&model.Role{Name: "abc"})

	// Migrate the schema
	// db.AutoMigrate(&model.Account{}, &model.Role{}, &model.Store{})

	// account := model.Account{}
	var result []map[string]interface{}
	db.Debug().Model(&model.Role{}).Scan(&result)
	// fmt.Println(result[0])

	// columns, _ := GenColumnCodes.GetColumns(db, "accounts")
	// fmt.Println(GenColumnCodes.GenColumnFieldsCode("Account", columns))

	// fmt.Println(tablename.GenTableNamesCode(db, "gen/tablename/tablename.go"))
	fmt.Println(
		columnname.GenTableColumnNamesCode(
			db,
			[]string{"roles"},
			"gen/columnname/columnname.go",
		),
	)
}
