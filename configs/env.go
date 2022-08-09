package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBase  DBaseConfig
	Server ServerConfig
)

type DBaseConfig struct {
	NAME     string
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
}

type ServerConfig struct {
	HOST string
	PORT string
}

// -----------------------------------< Interface >---------------------------------- \\

type FeedConfigs interface {
	builder()
}

// ---------------------------------< Feed configs >--------------------------------- \\

func LoadEnv() {
	godotenv.Load(".env")
	DBase.builder()
	Server.builder()
}

func (DBaseConfig) builder() {

	DBase = DBaseConfig{
		NAME:     os.Getenv("DB_NAME"),
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
	}
}

func (ServerConfig) builder() {

	Server = ServerConfig{
		HOST: os.Getenv("SERVER_HOST"),
		PORT: os.Getenv("SERVER_PORT"),
	}
}
