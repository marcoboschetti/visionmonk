package data

import (
	"crypto/tls"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

var pgConnection *pg.DB

func SetDbConnection() {
	password := os.Getenv("PGPASSWORD")

	db := pg.Connect(&pg.Options{
		Addr:      "dpg-cg76fr82qv28u2rq5p00-a.oregon-postgres.render.com:5432",
		Database:  "visionmonk",
		User:      "visionmonk_user",
		Password:  password,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})

	err := createSchema(db)
	if err != nil {
		panic(err.Error())
	}
	pgConnection = db
}

// psql -h ec2-3-209-61-239.compute-1.amazonaws.com -p 5432 -U vwqingapvwreed dfj7hg7pjkheuf

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*entities.User)(nil),
		(*entities.Shop)(nil),
		(*entities.Client)(nil),
		(*entities.CatalogProduct)(nil),
		(*entities.ShopProduct)(nil),
		(*entities.CalendarEvent)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
