package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	var gormDB *gorm.DB

	connection := viper.GetString("DB_CONNECTION")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	database := viper.GetString("DB_DATABASE")
	idleConnection := viper.GetInt("DB_IDLE_CONNECTION")
	maxConnection := viper.GetInt("DB_MAX_CONNECTION")
	maxLifeTimeConnection := viper.GetInt("DB_MAX_LIFE_TIME_CONNECTION")

	// mysql
	if connection == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
			database,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
				SlowThreshold:             time.Second * 5,
				LogLevel:                  logger.Info,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
			}),
		})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		conn, err := db.DB()
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		conn.SetMaxIdleConns(idleConnection)
		conn.SetMaxOpenConns(maxConnection)
		conn.SetConnMaxLifetime(time.Duration(maxLifeTimeConnection) * time.Second)

		gormDB = db
	}

	return gormDB
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
