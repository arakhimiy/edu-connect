package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"path/filepath"
	"sync"
)

type Config struct {
	Otp OtpConfig
}

type OtpConfig struct {
	BotToken  string `env:"OTP_BOT_TOKEN"`
	ChannelId string `env:"OTP_CHANNEL_ID"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		rootPath := flag.String("root_path", "", "Root path")
		flag.Parse()

		envFilePath, err := filepath.Abs(*rootPath + ".env")
		if err != nil {
			fmt.Println("Env file path error: ", err)
		}

		if err := cleanenv.ReadConfig(envFilePath, instance); err != nil {
			helpText := "Saidoff - Mosque project!"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			fmt.Println("Application is starting with default config")
		}
	})
	return instance
}
