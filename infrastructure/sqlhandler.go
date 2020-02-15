package infrastructure

import (
	"github.com/IkezawaYuki/pictweet-go/interface/adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() adapter.SQLHandler {
	conn, err := Connect()

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	conn.Close()
	return sqlHandler
}

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	db, err = gorm.Open("mysql",
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"@tcp("+
			os.Getenv("DB_HOST")+":"+
			os.Getenv("DB_PORT")+")/"+
			os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) error {
	db := handler.Conn.Find(out, where...)
	if db.RecordNotFound() {
		return nil
	}
	return db.Error
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) error {
	return handler.Conn.Exec(sql, values...).Error
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) error {
	db := handler.Conn.First(out, where...)
	if db.RecordNotFound() {
		return nil
	}
	return db.Error
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) error {
	return handler.Conn.Raw(sql, values...).Error
}

func (handler *SqlHandler) Create(value interface{}) error {
	return handler.Conn.Create(value).Error
}

func (handler *SqlHandler) Save(value interface{}) error {
	return handler.Conn.Save(value).Error
}

func (handler *SqlHandler) Delete(value interface{}) error {
	return handler.Conn.Delete(value).Error
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) error {
	db := handler.Conn.Where(query, args...)
	if db.RecordNotFound() {
		return nil
	}
	return db.Error
}

func (handler *SqlHandler) Close() error {
	return handler.Conn.Close()
}
