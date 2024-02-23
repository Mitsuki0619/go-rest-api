package main

import (
	"context"

	"github.com/Mitsuki0619/go-rest-api/api/internal/presenter"
)

func main() {
	srv := presenter.NewServer()

	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}