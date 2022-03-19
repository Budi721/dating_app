package manager

import (
	"context"
	"fmt"
	"github.com/Budi721/dating_app/utils/logger"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"log"
)

const (
	AppName    = "dating-app"
	ConfigName = "config"
)

type InfraManager interface {
	SqlDatabase() *pgx.Conn
	Config() *viper.Viper
}

type infraManager struct {
}

func (i *infraManager) SqlDatabase() *pgx.Conn {
	dbName := i.Config().GetString("datingapp.db.name")
	dbHost := i.Config().GetString("datingapp.db.host")
	dbPort := i.Config().GetString("datingapp.db.port")
	dbUser := i.Config().GetString("datingapp.db.user")
	dbPassword := i.Config().GetString("datingapp.db.password")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		logger.Log.Error().Err(err).Msg("failed to connect database")
		return nil
	}
	return conn
}

func (i *infraManager) Config() *viper.Viper {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("File not found")
		} else {
			log.Fatalln(err)
			log.Fatalln("File config error")
		}
	}

	return viper.GetViper()
}

func NewInfra() InfraManager {
	logger.NewLogger()
	return &infraManager{}
}
