package common

import (
	"fmt"
<<<<<<< HEAD
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../gorm.db")
=======

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", DbURL(BuildDBConfig()))
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
<<<<<<< HEAD
	//db.LogMode(true)

	//proooof
=======
	db.LogMode(true)
>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
	DB = db
	return DB
}

<<<<<<< HEAD
// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

// Using this function to get a connection, you can create your connection pool here.
=======
// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// func BuildDBConfig() *DBConfig {
// 	dbConfig := DBConfig{
// 		Host:     "localhost",
// 		Port:     3306,
// 		User:     "root",
// 		Password: "laspalmas12",
// 		DBName:   "Kalypso",
// 	}
// 	// fmt.Println(&dbConfig)
// 	return &dbConfig
// }

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "mysql",
		Port:     3306,
		User:     "root",
		Password: "root",
		DBName:   "kalypso",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

>>>>>>> 54f4ab9460419a42b520998c60f9fa0be7b23b8d
func GetDB() *gorm.DB {
	return DB
}
