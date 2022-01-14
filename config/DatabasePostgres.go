package config

type DatabasePostgresConfig struct {
	User     string
	Password string
	Db       string
	Host     string
	Port     int
}

var databasePostgresConfig *DatabasePostgresConfig

func GetDatabasePostgresConfig() *DatabasePostgresConfig {
	if databasePostgresConfig != nil {
		return databasePostgresConfig
	}
	conf := GetViper()
	databasePostgresConfig = &DatabasePostgresConfig{
		User:     conf.GetString("DATABASE.POSTGRESQL_USER"),
		Password: conf.GetString("DATABASE.POSTGRESQL_PASSWORD"),
		Db:       conf.GetString("DATABASE.POSTGRESQL_DB"),
		Host:     conf.GetString("DATABASE.POSTGRESQL_HOST"),
		Port:     conf.GetInt("DATABASE.POSTGRESQL_PORT"),
	}
	return databasePostgresConfig
}
