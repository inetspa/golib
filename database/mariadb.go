package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type Database struct {
	Config Config
	Ctx    *gorm.DB
}

func New(
	host string,
	port string,
	username string,
	password string,
	name string,
) Database {
	return Database{
		Config: Config{
			Host:     host,
			Port:     port,
			Username: username,
			Password: password,
			Name:     name,
		},
	}
}

func (db *Database) Connect(prod bool) error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		db.Config.Username,
		db.Config.Password,
		db.Config.Host,
		db.Config.Port,
		db.Config.Name,
	)
	var err error
	db.Ctx, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}
	if prod {
		db.Ctx.Logger.LogMode(logger.Silent)
	} else {
		db.Ctx.Logger.LogMode(logger.Info)
	}
	return nil
}

func (db *Database) MigrateDatabase(tables []interface{}) error {
	tx := db.Ctx.Begin()
	for _, t := range tables {
		if err := tx.AutoMigrate(t); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
