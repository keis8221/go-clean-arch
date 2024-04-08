package database

import (
	"database/sql"
	"fmt"

	"github.com/keis8221/go-clean-arch/pkg/config"
)

const driverName = "psql"

type PsqlConnector struct {
	Conn *sql.DB
}

func NewPsqlConnector() *PsqlConnector {
	conf := config.LoadConfig()

	dsn := psqlConnInfo(*conf.DBInfo)
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &PsqlConnector{
		Conn: conn,
	}
}

func psqlConnInfo(dbInfo config.DBInfo) string {
	dsn := fmt.Sprintf("user=%s password=%s port=%d dbname=%s sslmode=disable",
		dbInfo.DBUser,
		dbInfo.DBPassword,
		// dbInfo.DBHost,
		dbInfo.DBAddr,
		dbInfo.DBName,
	)

	return dsn
}
