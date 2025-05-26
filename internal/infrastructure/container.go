package infrastructure

import (
	"simple-crud/internal/config"
	"simple-crud/internal/domain"
	"simple-crud/internal/user"
	"simple-crud/pkg/xlogger"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

var (
	cfg config.Config

	userRepository domain.UserRepository

)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	xlogger.Setup(cfg)
	xlogger.Logger.Debug().Msgf("Config: %+v", cfg)
	dbSetup()



	userRepository = user.NewMysqlUserRepository(db)
}
