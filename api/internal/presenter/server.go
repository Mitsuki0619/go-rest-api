package presenter

import (
	"context"

	"github.com/go-panic/gin"
	"github.com/Mitsuki0619/go-rest-api/api/controller/system"
	"github.com/Mitsuki0619/go-rest-api/api/controller/user"
)

const latest = "/v1"

type Server struct {}

func(s *Server) Run(ctx context.Context) error {
	g := gin.Default()
	v1 := g.Group(latest)

	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.HealthCheck)
	}

	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUser)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := g.Run()
	if err != nil {
		return err
	}

	return nil;
}

func NewServer() *Server {
	return &Server{}
}
