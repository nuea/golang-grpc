package database

import (
	"fmt"
	"log"
	"time"

	"github.com/nuea/go-grpc-product-svc/packages/config"
	"github.com/nuea/go-grpc-product-svc/packages/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	db       *gorm.DB
	host     string
	port     string
	dbName   string
	user     string
	password string
	timezone string
}

func (h *Handler) Connect() error {
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

func (h *Handler) AutoMigrate() {
	h.db.AutoMigrate(&models.Product{})
	h.db.AutoMigrate(&models.StockDecreaseLog{})
}

func (h *Handler) GetDatabase() *gorm.DB {
	return h.db
}

func Connect(config config.Config) Handler {
	gorm := Handler{
		host:     config.DBHost,
		port:     config.DBPort,
		dbName:   config.DBName,
		user:     config.DBUser,
		password: config.DBPassword,
		timezone: config.DBTZ,
	}

	if err := gorm.Connect(); err != nil {
		log.Fatalln("Cannot access database: ", err)
	}
	gorm.AutoMigrate()
	log.Println("Successfully connected to the database!")

	return gorm
}
