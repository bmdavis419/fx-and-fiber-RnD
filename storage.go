package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

// FAKE STORAGE
type Storage struct {
	data string
}

func NewStorage(lc fx.Lifecycle) *Storage {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// do something on start
			fmt.Println("STORAGE HAS BEEN STARTED!!")
			return nil
		},
	})
	return &Storage{
		data: "THIS IS FROM STORAGE",
	}
}

func (s *Storage) GetData() string {
	return s.data
}
