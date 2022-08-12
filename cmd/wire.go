//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pierbin/screechr/cmd/router"
	"github.com/pierbin/screechr/internal/config"
	"github.com/pierbin/screechr/internal/controller"
	"github.com/pierbin/screechr/internal/repo"
)

func Initializing() *router.Router {
	panic(
		wire.Build(
			router.NewRouter,
			router.NewHandler,
			controller.NewScreechrCtl,
			wire.NewSet(repo.NewScreechrRepo, wire.Bind(new(repo.IScreechrRepo), new(*repo.ScreechrRepo))),
			config.NewConfig,
		),
	)
}
