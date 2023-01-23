package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"test/internal/config"
)

func ConnectDb(c *config.Config) *sql.DB {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB)
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("Ошибка подключения к бд")
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Ошибка пинга")
		log.Fatal(err)
	}

	fmt.Println("Успешное подключение к бд")
	return db
}
