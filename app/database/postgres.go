package database

import (
	_ "github.com/joho/godotenv"
	"go-cursovic/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func QueryGetAllShurles() ([]models.Shurls, error) {
	var sus []models.Shurls

	tx := db.Find(&sus)
	if tx.Error != nil {
		return []models.Shurls{}, tx.Error
	}

	return sus, nil
}

func QueryGetShurl(id uint64) (models.Shurls, error) {
	var su models.Shurls

	tx := db.Where("id = ?", id).First(&su)

	if tx.Error != nil {
		return models.Shurls{}, tx.Error
	}

	return su, nil
}

func QueryCreateShurl(su models.Shurls) error {
	tx := db.Create(&su)
	return tx.Error
}

func QueryUpdateShurls(su models.Shurls) error {

	tx := db.Save(&su)
	return tx.Error
}

func QueryDeleteShurl(id uint64) error {

	tx := db.Unscoped().Delete(&models.Shurls{}, id)
	return tx.Error
}

func QueryFindByShurlUrl(url string) (models.Shurls, error) {
	var su models.Shurls
	tx := db.Where("shurl = ?", url).First(&su)
	return su, tx.Error
}

func ConnectToDB() {
	var err error

	dsn := "host=postgres user=admin password=test dbname=admin port=5432 sslmode=disable"

	log.Print("Connecting to Postgres DB...")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		log.Fatal("error connects to database. \n", err)
	}

	log.Println("connected")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Running the migrations...")
	err = db.AutoMigrate(&models.Shurls{}, &models.User{})
	if err != nil {
		log.Println(err)
	}

}
