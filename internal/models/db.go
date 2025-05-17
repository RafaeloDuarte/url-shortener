package models

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=pg-shortener port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar com o banco: ", err)
	}

	// AutoMigrate cria a tabela se não existir
	err = database.AutoMigrate(&URL{})
	if err != nil {
		log.Fatal("Erro ao fazer migrate: ", err)
	}

	DB = database
	fmt.Println("Conectado ao banco com sucesso!")
}
