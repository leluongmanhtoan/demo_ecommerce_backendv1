package main

import (
	api "demo_ecommerce/api"
	apiv1 "demo_ecommerce/api/v1"
	"demo_ecommerce/internal/sqlclient"
	"demo_ecommerce/middleware/auth"
	"demo_ecommerce/repository"
	"demo_ecommerce/repository/database"
	"demo_ecommerce/service"
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
	repository.UserRepo = database.NewUser()

}

func main() {
	jwtservice := service.NewJWTService()
	server := api.NewServer()
	auth.AuthMdw = auth.NewAuthMiddleware(*jwtservice)
	apiv1.NewUser(server.Engine, service.NewUser(*jwtservice))
	server.Start("8080")
}
