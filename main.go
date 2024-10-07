package main

import (
	"demo_ecommerce/api"
	"demo_ecommerce/internal/sqlclient"
	"demo_ecommerce/repository"
)

func init() {
	sqlClientConfig := sqlclient.SqlConfig{
		Host:         "localhost",
		Port:         5432,
		Database:     "demo_ecommerce",
		Username:     "admin",
		Password:     "ManhToan0123",
		DialTimeout:  20,
		ReadTimeout:  30,
		WriteTimeout: 30,
		Timeout:      30,
		PoolSize:     10,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	}
	repository.SqlClient = sqlclient.NewSqlClient(sqlClientConfig)
}

func main() {
	server := api.NewServer()
	server.Start("8080")
}
