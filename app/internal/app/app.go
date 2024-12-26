package app

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "gitlab.saidoff.uz/company/muslim-administration/mosque/back/artifacts/migrations"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/config"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/handler"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/repository"
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/service"
)

func NewApp(config *config.Config) *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		db := app.DB()
		logger := app.Logger()

		repos := repository.NewRepository(db)
		services := service.NewService(repos)

		handlers := handler.NewHandler(logger, services, config)
		handlers.Register(e.Router)

		return e.Next()
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	return app
}
