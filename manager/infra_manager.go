package manager

import (
	"log"

	"fmt"
	"warung_nasi_padang/config"

	"database/sql"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	SqlDB() *sql.DB
}

type infraManager struct {
	db *sql.DB
	cfg config.Config
}

func (i *infraManager) initDb() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to run", err)
		}
	}()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	i.db = db
	fmt.Println("DB Connected")
}
func (i *infraManager) SqlDB() *sql.DB {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}