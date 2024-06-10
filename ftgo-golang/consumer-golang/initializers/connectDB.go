package initializers

import (
	"consumer-golang/utils"
	"encoding/base64"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	var passwordDB string
	//
	var base64DeCoded []byte
	base64DeCoded, _ = base64.StdEncoding.DecodeString(config.DBUserPassword)
	plaintext := utils.Decrypt(base64DeCoded, "password")
	passwordDB = string(plaintext)
	fmt.Println(passwordDB)
	//
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, passwordDB, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
