// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	baguwen "github.com/ecodeclub/webook/internal/question"
	"github.com/ecodeclub/webook/internal/question/internal/web"
	testioc "github.com/ecodeclub/webook/internal/test/ioc"
)

// Injectors from wire.go:

func InitHandler() (*web.Handler, error) {
	db := testioc.InitDB()
	cache := testioc.InitCache()
	handler, err := baguwen.InitHandler(db, cache)
	if err != nil {
		return nil, err
	}
	return handler, nil
}
