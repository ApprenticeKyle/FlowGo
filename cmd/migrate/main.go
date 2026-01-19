package main

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"FLOWGO/internal/domain/entity"
	"FLOWGO/internal/infrastructure/config"
)

func main() {
	// 1. Load Config
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Println("Could not load config.yaml, assuming environment variables or defaults")
	}

	// MySQL DSN
	cfg := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	log.Printf("Connecting to MySQL source: %s...", cfg.Host)
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	log.Println("Connected to MySQL.")

	// SQLite Destination
	sqliteFile := "flowgo.db"
	log.Printf("Connecting to SQLite destination: %s...", sqliteFile)
	sqliteDB, err := gorm.Open(sqlite.Open(sqliteFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}
	log.Println("Connected to SQLite.")

	// 2. AutoMigrate Schema
	log.Println("Migrating schema...")
	err = sqliteDB.AutoMigrate(
		&entity.User{},
		&entity.Project{},
		&entity.Team{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	// 3. Migrate Data

	// Users
	var users []entity.User
	if err := mysqlDB.Find(&users).Error; err != nil {
		log.Printf("Error finding users: %v", err)
	} else {
		log.Printf("Found %d Users", len(users))
		if len(users) > 0 {
			sqliteDB.Create(&users)
			log.Printf("Migrated %d Users", len(users))
		}
	}

	// Teams
	var teams []entity.Team
	if err := mysqlDB.Find(&teams).Error; err != nil {
		log.Printf("Error finding teams: %v", err)
	} else {
		log.Printf("Found %d Teams", len(teams))
		if len(teams) > 0 {
			sqliteDB.Create(&teams)
			log.Printf("Migrated %d Teams", len(teams))
		}
	}

	// Projects
	var projects []entity.Project
	if err := mysqlDB.Find(&projects).Error; err != nil {
		log.Printf("Error finding projects: %v", err)
	} else {
		log.Printf("Found %d Projects", len(projects))
		if len(projects) > 0 {
			sqliteDB.Create(&projects)
			log.Printf("Migrated %d Projects", len(projects))
		}
	}

	log.Println("Migration complete!")
}
