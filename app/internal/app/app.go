package app

import (
	"github.com/go-telegram/bot"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "github.com/arakhimiy/edu-connect/artifacts/migrations"
	"github.com/arakhimiy/edu-connect/internal/config"
	"github.com/arakhimiy/edu-connect/internal/handler"
	"github.com/arakhimiy/edu-connect/internal/repository"
	"github.com/arakhimiy/edu-connect/internal/service"
)

func NewApp(config *config.Config) *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		db := app.DB()
		logger := app.Logger()
		newBot, err := bot.New(config.Otp.BotToken)
		if err != nil {
			logger.Error("telegram bot init error")
			//panic(err)
		}

		repos := repository.NewRepository(db)
		services := service.NewService(repos, newBot)

		handlers := handler.NewHandler(logger, services, config)
		handlers.Register(e.Router)

		return e.Next()
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	return app
}
