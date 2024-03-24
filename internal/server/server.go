package server

import (
	"log"
	"micro_uploads/internal/controllers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const MAX_MULTIPART_MEM int64 = 200 << 20

type Server struct {
	controllers []controllers.IControllers
	config      *viper.Viper
	router      *gin.Engine
}

func New(config *viper.Viper, router *gin.Engine, controllers ...controllers.IControllers) Server {
	return Server{
		router:      router,
		config:      config,
		controllers: controllers,
	}
}

func (s *Server) Start() {
	s.router.MaxMultipartMemory = MAX_MULTIPART_MEM
	s.multiTemplate()
	s.setupRoutes()
	s.router.Static("/static", "./static")

	if err := s.router.Run(s.config.GetString("http.addr")); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}

func (s *Server) multiTemplate() {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("index", "templates/index.html", "templates/root.html")
	r.AddFromFiles("user", "templates/user.html", "templates/root.html")

	s.router.HTMLRender = r
}

func (s *Server) setupRoutes() {
	for _, i := range s.controllers {
		i.StartRoutes()
	}
}
