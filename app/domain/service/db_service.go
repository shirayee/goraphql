package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goraphql/app/domain/model"
	"os"
)

type DbService interface {
	Connect() (*gorm.DB, error)
	Migrate() error
}

type dbService struct{}

func NewDbService() DbService {
	return &dbService{}
}

func (s dbService) Connect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	db, err := gorm.Open(DBMS, CONNECT)
	return db, err
}

func (s dbService) Migrate() error {
	db, err := s.Connect()
	if err != nil {
		return err
	}
	db.AutoMigrate(&model.User{})
	return nil
}
