package databases

import (
	"idnatiya/go-auth/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMysqlDB() {
	database, err := gorm.Open(
		mysql.Open(
			os.Getenv("DB_USER")+":"+
				os.Getenv("DB_PASS")+"@tcp("+
				os.Getenv("DB_HOST")+":"+
				os.Getenv("DB_PORT")+")/"+
				os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		panic("Failed to connect MySQL Database")
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Permission{})
	database.AutoMigrate(&models.Role{})
	database.AutoMigrate(&models.File{})

	DB = database
}
