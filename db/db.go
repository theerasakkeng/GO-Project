package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%v://%v:%v@%v?database=%v",
		viper.GetString("db.driver"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.database"),
	)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
