package config

type Config struct {
	AppPort         int
	DbHost          string
	DbPort          int
	DbUserName      string
	DbPassword      string
	DbDatabase      string
	DbRunMigrations bool
}
