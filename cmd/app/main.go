package main

import (
	"context"
	"log"
	"os"

	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkpostgres "github.com/rookie-ninja/rk-db/postgres"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"

	"developerProfile/internal/adapters/repository/postgres/developer_profile"
	resthandlers "developerProfile/internal/adapters/rest"
	"developerProfile/internal/core/service"
)

func main() {
	raw, err := os.ReadFile("boot.yaml")
	if err != nil {
		log.Fatalf("failed to read boot.yaml: %v", err)
	}

	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw([]byte(os.ExpandEnv(string(raw)))))
	boot.Bootstrap(context.TODO())

	gin := rkgin.GetGinEntry("developer_profile")

	pgEntry := rkpostgres.GetPostgresEntry("developer_profile_postgres")
	if pgEntry == nil {
		log.Fatal("postgres entry not found")
	}

	dbEntry := pgEntry.GetDB("developer_profile")
	db, err := dbEntry.DB()
	if err != nil {
		panic(err)
	}

	repository := developer_profile.New(db)
	svc := service.NewService(repository)
	handler := resthandlers.NewHandler(svc)

	gin.Router.POST("/profiles", handler.CreateProfile)
	gin.Router.GET("/profiles", handler.ListProfiles)
	gin.Router.GET("/profiles/:id", handler.GetProfile)
	gin.Router.PUT("/profiles/:id", handler.UpdateProfile)
	gin.Router.DELETE("/profiles/:id", handler.DeleteProfile)

	boot.WaitForShutdownSig(context.TODO())
}
