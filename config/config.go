package config

type Config struct {
	DbHost          string
	DbPort          int
	DbUserName      string
	DbPassword      string
	DbDatabase      string
	DbRunMigrations bool
}
