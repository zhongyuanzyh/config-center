package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
)

var MySqlPool *sql.DB

func SetUp() {
	h := "192.168.1.2"
	pa := "nacos"
	u := "nacos"
	p := 3307
	d := "nacos_devtest"
	conn := strings.Join([]string{u, ":", pa, "@tcp(", h, ":", strconv.Itoa(p), ")/", d, "?charset=utf8"}, "")

	log.Println(conn)
	MySqlPool, _ = sql.Open("mysql", conn)
	MySqlPool.SetConnMaxLifetime(100)
	MySqlPool.SetMaxIdleConns(10)
	if err := MySqlPool.Ping(); err != nil {
		log.Println("MySQL initial failed!")
	} else {
		log.Println("MySQL initial successful!")
	}
}
