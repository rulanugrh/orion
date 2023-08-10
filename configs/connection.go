package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Server struct {
		Host  string
		GPort string
		Hport string
	}

	Database struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}
	Secret string
}

var app *App
var DB *gorm.DB

func GetConnection() *gorm.DB {
	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", conf.Database.User, conf.Database.Pass, conf.Database.Host, conf.Database.Port, conf.Database.Name)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Found error %v", err)
	}

	DB = db
	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}
	if err := godotenv.Load(); err != nil {
		conf.Server.GPort = ""
		conf.Server.Hport = ""
		conf.Server.Host = "localhost"

		conf.Database.Host = "localhost"
		conf.Database.Name = ""
		conf.Database.Port = "3306"
		conf.Database.User = ""
		conf.Database.Pass = ""
		conf.Secret = ""

		return &conf
	}

	conf.Server.GPort = os.Getenv("GRPC_PORT")
	conf.Server.Hport = os.Getenv("HTTP_PORT")
	conf.Server.Host = os.Getenv("SERVER_HOST")
	conf.Database.Name = os.Getenv("MYSQLDB_NAME")
	conf.Database.Host = os.Getenv("MYSQLDB_HOST")
	conf.Database.User = os.Getenv("MYSQLDB_USER")
	conf.Database.Pass = os.Getenv("MYSQLDB_PASS")
	conf.Database.Port = os.Getenv("MYSQLDB_PORT")
	conf.Secret = os.Getenv("JWT_SECRET")
	
	return &conf
}