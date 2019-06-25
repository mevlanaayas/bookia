package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.LogMode(true)

	// to create table names as singular. word -> word if true, word -> words otherwise
	db.SingularTable(true)

	// in development mode drop all tables if already exist
	if os.Getenv("dev") == "true" {
		db.Debug().DropTableIfExists(&Word{}, &Book{})
	}

	//Database migration
	db.Debug().AutoMigrate(&Word{}, &Book{})

	// creating Foreign keys relations
	db.Model(&Word{}).AddForeignKey("book_id", "book(id)", "CASCADE", "CASCADE")
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
