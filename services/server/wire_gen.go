// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/google/wire"
	"grpc_go_blog/config/db"
	"grpc_go_blog/services/user/handler"
	"grpc_go_blog/services/user/infra/sqlite"
	"grpc_go_blog/services/user/usecase"
)

// Injectors from wire.go:

func InjectAPIServer() (*API, error) {
	dbDB := db.NewDB()
	userRepositoryIF := sqlite.NewUserRepository(dbDB)
	userUsecaseIF := usecase.NewUserUsecase(userRepositoryIF)
	userHandler := handler.NewUserHandler(userUsecaseIF)
	api := NewAPI(userHandler)
	return api, nil
}

// wire.go:

var userSuperSet = wire.NewSet(handler.NewUserHandler, usecase.NewUserUsecase, sqlite.NewUserRepository)
