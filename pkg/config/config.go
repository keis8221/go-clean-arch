package config

import "os"

type appConfig struct {
	HTTPInfo *HTTPInfo
	DBInfo   *DBInfo
}

type HTTPInfo struct {
	Addr string
}

type DBInfo struct {
	DBUser     string
	DBPassword string
	DBAddr     string
	DBName     string
}

func LoadConfig() *appConfig {
	addr := ":" + os.Getenv("DB_PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	dbInfo := &DBInfo{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddr:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

	conf := &appConfig{
		HTTPInfo: httpInfo,
		DBInfo:   dbInfo,
	}

	return conf
}
