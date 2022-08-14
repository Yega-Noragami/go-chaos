package database

import (
	"fmt"

	"github.com/Yega-Noragami/go-chaos/pkg/config"
	"github.com/Yega-Noragami/go-chaos/pkg/constants"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getDSN() string {
	// Get configs
	config := config.NewConfig()
	// Build DSN
	var DSN string = fmt.Sprintf(
		"%s@%s(%s:%s)/%s?%s",
		config.DBUser,
		constants.SQL_CONN_METHOD,
		config.DBHost,
		config.DBPort,
		config.DBDatabase,
		constants.SQL_ADDITIONAL_PROPS,
	)
	return DSN
}

func Setup() *gorm.DB {
	fmt.Println("Trying to connect to MYSQL instance")
	// Setup Connection
	d, err := gorm.Open("mysql", getDSN())
	if err != nil {
		panic(err)
	}
	fmt.Println("Success! Connected to MySQL instance")
	// Need to setup the table first
	setupTable()
	fmt.Println("Success! Created Table `lists`")
	return d
}
