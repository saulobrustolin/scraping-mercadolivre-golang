package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"scrapping-mercadolivre-golang/src/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() (string, string) {
	// Carregar variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obter as variáveis do banco de dados do ambiente carregado
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")

	// Verifica se as variáveis de ambiente foram corretamente carregadas
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing required environment variables")
	}

	// Retornar a string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	log.Println(connStr)

	// Aqui, se necessário, você pode pegar um ambiente de configuração
	// Exemplo de como pegar uma variável de ambiente
	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // Caso não esteja definido, assume como "development"
	}

	return connStr, env
}

func SaveData(products *[]models.Product) error {
	// Carregar configuração de ambiente
	connStr, env := loadEnv()

	// Conectar ao banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inserindo dados na tabela
	stmt, err := db.Prepare("INSERT INTO products(title, quantity_reviews, stars, price, anchor_price, url, picture, free_shipping, installments, installments_amount) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Executando o insert com dados
	for _, product := range *products {
		log.Printf("Inserting Product: %v", product)

		_, err = stmt.Exec(
			product.Title,
			product.QuantityReviews,
			product.Stars,
			product.Price,
			product.AnchorPrice,
			product.URL,
			product.Picture,
			product.IsFreeShipping,
			product.Installments,
			product.InstallmentsAmount,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Confirmando a inserção
	fmt.Printf("Dados inseridos com sucesso no ambiente: %s!\n", env)
	return nil
}
