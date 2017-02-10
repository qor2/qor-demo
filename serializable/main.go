package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/serializable_meta"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"fmt"
)

type QorJob struct {
	gorm.Model

	Name string
	serializable_meta.SerializableMeta
}

func (qorJob QorJob) GetSerializableArgumentResource() *admin.Resource {
	v := jobsArgumentsMap[qorJob.Kind]
	return v
}



type sendNewsletterArgument struct {
	Subject string
	Content string
}

type importProductArgument struct {

}

var(
	config = &qor.Config{}
	Admin = admin.New(config)
	jobsArgumentsMap = map[string]*admin.Resource {
		"newsletter": Admin.NewResource(&sendNewsletterArgument{}),
		"import_products": Admin.NewResource(&importProductArgument{}),
	}
)


func main(){
	db, _ := gorm.Open("sqlite3", "demo.db")
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&QorJob{})

	var qorJob QorJob

	qorJob.Name = "sending newsletter"
	qorJob.Kind = "newsletter"

	qorJob.SetSerializableArgumentValue(&sendNewsletterArgument{
		Subject: "subject",
		Content: "content",
	})

	db.Create(&qorJob)


	result := &QorJob{}
	db.First(result, "name = ?", "sending newsletter")
	fmt.Println("found result", result)

	argument := result.GetSerializableArgument(result)

	if newsletterArgument, ok := argument.(*sendNewsletterArgument); ok {
		fmt.Println("subject", newsletterArgument.Subject)
		fmt.Println("content", newsletterArgument.Content)
	}
}
