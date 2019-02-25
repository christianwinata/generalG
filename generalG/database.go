package generalG

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConnect(driver string, user string, pass string, host, string, port string, dbname string, sslmode string) {
	var connstring string

	switch driver {
	case "postgres":
		connstring = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, dbname, sslmode)
		if db, err := gorm.Open("postgres", connstring); err != nil {
			panic(err)
		}
	case "mysql":
		connstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname)
		if db, err := gorm.Open("mysql", connstring); err != nil {
			panic(err)
		}
	default:
		panic(fmt.Errorf("undefined driver \"%s\"", driver))
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	// db.LogMode(true)
	// db.Exec(fmt.Sprintf("SET TIMEZONE = '%s'", timezone))
	// db.DB().SetConnMaxLifetime(time.Minute * time.Duration(App.DBConfig.Int("maxlifetime")))
	// db.DB().SetMaxIdleConns(App.DBConfig.Int("idle_conns"))
	// db.DB().SetMaxOpenConns(App.DBConfig.Int("open_conns"))
}
