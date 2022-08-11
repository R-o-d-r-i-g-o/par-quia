package configs

var (
	DBase  DBaseConfig
	Server ServerConfig
)

type DBaseConfig struct {
	DB       string `env:"DB_NAME"`
	PORT     string `env:"DB_PORT" envDefault:"3306"`
	HOST     string `env:"DB_HOST"`
	USER     string `env:"DB_USER"`
	PASSWORD string `env:"DB_PASSWORD"`
}

type ServerConfig struct {
	PORT string `env:"SERVER_PORT"`
	HOST string `env:"SERVER_HOST"`
}

func Load() {

	loadStructWithEnvVars(&DBase)
	loadStructWithEnvVars(&Server)
}
