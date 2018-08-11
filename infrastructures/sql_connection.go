package infrastructures

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/spf13/viper"
)

// SQLConnection - Struct sql connection
type SQLConnection struct{}

func (s *SQLConnection) Connect() (*sql.DB, error) {
	address := fmt.Sprintf("tcp(%s:%s)",
		config.GetString("database.host"),
		config.GetString("database.port"),
	)
	// DB Source
	datasource := fmt.Sprintf("%s:%s@%s/%s?charset=%s",
		config.GetString("database.user"),
		config.GetString("database.password"),
		address,
		config.GetString("database.name"),
		config.GetString("database.charset"),
	)
	// SQL Connection
	conn, err := sql.Open("mysql", datasource)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %v", err.Error()))
		return nil, err
	}

	return conn, nil
}
