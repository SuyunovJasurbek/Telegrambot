package postgres

import (
	"fmt"
	"log"
	"telegram_bot/storage"

	"telegram_bot/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StorageI struct {
	db             *sqlx.DB
	dataRepository storage.DataI
}

// Data implements storage.StorageI
func (s *StorageI) Data() storage.DataI {
	if s.dataRepository == nil {
		s.dataRepository = NewData(s.db)
	}
	return s.dataRepository
}

// CloseDb implements storage.StorageI
func (s *StorageI) CloseDb() {
	defer s.db.Close()
}

func NewPostgres(cnf config.Config) storage.StorageI {

	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.PostgresHost, cnf.PostgresPort, cnf.PostgresUser, cnf.PostgresPassword, cnf.PostgresDatabase)

	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	} else {
		log.Printf("Connected to database: %s", cnf.PostgresDatabase)
	}

	return &StorageI{
		db: db,
	}
}
