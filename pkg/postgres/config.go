package postgres

type dbConfig struct {
	Addr     string
	Port     int
	Username string
	Database string
	Password string
}

func getDbConfig() *dbConfig {
	return &dbConfig{
		Addr:     "127.0.0.1",
		Port:     5432,
		Username: "playground",
		Password: "1234",
		Database: "playground",
	}
}
