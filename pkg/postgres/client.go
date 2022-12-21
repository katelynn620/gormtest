package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	client *gorm.DB
}

func (m *DBClient) Connect() {
	config := getDbConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Addr,
		config.Port,
		config.Username,
		config.Database,
		config.Password,
	)
	client, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
	m.client = client

	// migration
	m.client.AutoMigrate(&Player{})
}
func (m *DBClient) Disconnect() {
	sqlDB, err := m.client.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
