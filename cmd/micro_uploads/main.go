package main

import (
	"log"
	"micro_uploads/internal/config"
	"micro_uploads/internal/controllers"
	"micro_uploads/internal/server"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.InitConfig("micro_uploads")
	db, err := config.ConnectDB(conf.GetString("database.path"))

	if err != nil {
		log.Fatalf("unable to connect to database: %s\n", err.Error())
	}

	router := gin.Default()
	store := cookie.NewStore([]byte(conf.GetString("secret.key")))

	router.Use(sessions.Sessions("micro_uploads", store))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // allow all origins to access data
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"}, // allow only the given headers
		AllowCredentials: false,                                                // disallow credentials
	}))

	v1 := router.Group("/api/v1")
	root := router.Group("/")

	uploadCon := controllers.NewUploadController(v1, db, conf.GetString("uploads.path"))
	authCon := controllers.NewAuthController(v1, db)
	fCon := controllers.NewFrontControllers(root, db)

	server := server.New(conf, router, &uploadCon, &authCon, &fCon)

	server.Start()
}
