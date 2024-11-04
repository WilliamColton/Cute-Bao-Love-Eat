package main

import (
	"backend/controller"
	"backend/dao"
	"backend/db"
	"backend/middleware"
	"backend/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	databaseName := viper.GetString("database_path")
	jwtSecretKey := viper.GetString("jwt_secret_key")
	fmt.Println("databaseName:")
	fmt.Println(databaseName)
	fmt.Println("jwtSecretKey:")
	fmt.Println(jwtSecretKey)

	jwt := middleware.NewJWT(jwtSecretKey)

	DB, err := db.InitDB(databaseName)
	if err != nil {
		return
	}
	d := dao.InitDao(DB)
	c := controller.NewController(jwt, d)
	r := router.NewRouter(c)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
