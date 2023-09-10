package utils

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func Execute_sql_file(db *gorm.DB, path string) {
	b, err := os.ReadFile(path)

	if err != nil {
		fmt.Print(err)
	}

	db.Exec(string(b))
}
