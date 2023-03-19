// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/frchandra/chatin/app"
	"github.com/frchandra/chatin/app/controller"
	"github.com/frchandra/chatin/app/messenger"
	"github.com/frchandra/chatin/app/middleware"
	"github.com/frchandra/chatin/app/repository"
	"github.com/frchandra/chatin/app/service"
	"github.com/frchandra/chatin/app/util"
	"github.com/frchandra/chatin/config"
	"github.com/frchandra/chatin/database"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeServer() *app.Server {
	appConfig := config.NewAppConfig()
	client := app.NewCache(appConfig)
	tokenUtil := util.NewTokenUtil(client, appConfig)
	logger := app.NewLogger(appConfig)
	userMiddleware := middleware.NewUserMiddleware(tokenUtil, logger)
	database := app.NewDatabase(appConfig, logger)
	logUtil := util.NewLogUtil(logger)
	userRepository := repository.NewUserRepository(database, logUtil)
	userService := service.NewUserService(userRepository, tokenUtil)
	userController := controller.NewUserController(userService, tokenUtil, appConfig, logUtil)
	roomRepository := repository.NewRoomRepository(database, logUtil)
	roomService := service.NewRoomService(roomRepository)
	hub := messenger.NewHub()
	roomController := controller.NewRoomController(roomService, hub)
	server := app.NewRouter(userMiddleware, userController, roomController)
	return server
}

func InitializeMigrator() *database.Migrator {
	appConfig := config.NewAppConfig()
	logger := app.NewLogger(appConfig)
	mongoDatabase := app.NewDatabase(appConfig, logger)
	migrator := database.NewMigrator(mongoDatabase, logger)
	return migrator
}

// injector.go:

var MiddlewareSet = wire.NewSet(middleware.NewUserMiddleware)

var UserSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

var RoomSet = wire.NewSet(repository.NewRoomRepository, service.NewRoomService, controller.NewRoomController)

var UtilSet = wire.NewSet(util.NewTokenUtil, util.NewLogUtil)

var MessengerSet = wire.NewSet(messenger.NewHub)
