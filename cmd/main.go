package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ultra-hattrick/players-back/docs"
	playersHttp "github.com/ultra-hattrick/players-back/internal/adapters/handler/http"
	"github.com/ultra-hattrick/players-back/internal/adapters/repository/postgres"
	"github.com/ultra-hattrick/players-back/internal/core/domain"
	"github.com/ultra-hattrick/players-back/internal/core/services"
	"gitlab.com/uchile1/helper/helperDB"
	"gitlab.com/uchile1/helper/helperLog"
	"gorm.io/gorm"
)

// @title Players BACK API
// @version 1.0
// @description Para obtener los jugadores de los equipos desde Hattick y la base de datos.
// @host localhost:8082
// @BasePath /api/v1/team/{team_id}/player

func main() {
	r := gin.Default()
	// r.Use(helperMiddleware.RequestLogger())
	db, err := helperDB.GetDBPostgres()
	if err != nil {
		helperLog.Logger.Panic().AnErr("NO se pudo conectar a la BD!", err)
	}
	// Crear el esquema si no existe
	if err := createSchema(db, os.Getenv("DB_POSTGRES_SCHEMA")); err != nil {
		helperLog.Logger.Panic().AnErr("error al crear el esquema!", err)
	}
	// AutoMigrate ejecutará las migraciones automáticamente
	if err := db.AutoMigrate(&domain.Player{}); err != nil {
		helperLog.Logger.Panic().AnErr("error al ejecutar automigrate!", err)
	}
	repository := postgres.NewPlayersRepository(db)
	service := services.NewPlayersService(repository)
	handler := playersHttp.NewPlayersHandler(service)
	group := r.Group("/api/v1/team/:team_id/player")
	group.GET("/", handler.GetPlayersByTeamID)
	group.POST("/", handler.CreatePlayersByTeamID)
	group.GET("/:player_id", handler.GetPlayersByIDAndTeamID)
	group.GET("/match/:match_id", handler.FetchPlayersByTeamIDAndMatchID)
	group.POST("/match/:match_id", handler.GetPlayersByTeamIDAndMatchID)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/health", func(ctx *gin.Context) {
		if err := helperDB.HealthheckPostgresHandler(); err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error conectarse a la DB %v", err.Error()))
			return
		}
		ctx.Status(http.StatusOK)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createSchema(db *gorm.DB, schemaName string) error {
	return db.Exec("CREATE SCHEMA IF NOT EXISTS " + schemaName).Error
}
