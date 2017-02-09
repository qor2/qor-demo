package main

import (
	"fmt"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
)

type Config struct {
	Root string
}

func main() {
	db, _ := gorm.Open("sqlite3", "demo.db")
	defer db.Close()

	root, _ := filepath.Abs(".")
	config := Config{
		Root: root,
	}

	I18n := i18n.New(
		database.New(db),		// better using one backend for translate
		yaml.New(filepath.Join(config.Root, "config/locales")),
	)

	// found config/locales/zh-CN.yml
	t := I18n.T("zh-CN", "demo.greeting")
	fmt.Println(t)

	t = I18n.T("en-US", "demo.greeting")
	fmt.Println(t)
}
