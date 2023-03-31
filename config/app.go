package config

import "gorm.io/gorm"

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()
	app.DB = DatabaseConnection()

	return *app
}
