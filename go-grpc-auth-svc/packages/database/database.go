package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nuea/go-grpc-auth-svc/packages/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseHandler struct {
	db       *gorm.DB
	host     string
	port     string
	dbName   string
	user     string
	password string
	timezone string
}

func (h *DatabaseHandler) Connect() error {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", h.host, h.port, h.user, h.password, h.dbName, h.timezone)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			return now
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	h.db = db
	return err
}

func (g *DatabaseHandler) AutoMigrate() {
	g.db.AutoMigrate(&models.User{})
}

func Connect() *DatabaseHandler {
	gorm := &DatabaseHandler{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		dbName:   os.Getenv("DB_NAME"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		timezone: os.Getenv("DB_TZ"),
	}

	if err := gorm.Connect(); err != nil {
		log.Fatalln("Cannot access database: ", err)
	}
	gorm.AutoMigrate()
	log.Println("Successfully connected to the database!")

	return gorm
}
