package postgres

import (
	"github.com/Da-max/todo-go/internal/repositories/todo"
	"github.com/Da-max/todo-go/internal/repositories/user"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDsn(gormForm bool) string {
	var (
		res string
	)
	if gormForm {
		res = "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432 sslmode=disable"
	} else {
		res = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":5432" + "/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	}
	return res
}

func New() *gorm.DB {
	if db, err := gorm.Open(postgres.Open(getDsn(true)), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		return db
	}

}

func Migrate() {

	db := New()

	var err = db.AutoMigrate(&user.User{}, &todo.Todo{})

	if err != nil {
		log.Fatal(err)
	}

	// db, err := sql.Open("postgres", getDsn(false))
	// if err != nil {
	// 	log.Fatal(err)
	// 	// }
	//
	// driver, err := postgresDriver.WithInstance(db, &postgresDriver.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// 	// }
	//
	// m, err := migrate.NewWithDatabaseInstance("file://postgres/migrations", "postgres", driver)
	// if err := m.Up(); err != nil {
	// 	log.Println(err)
	// }
}
