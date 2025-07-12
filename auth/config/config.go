package config

import (
	"log"

	"github.com/archon42x/clavis-sdk-go/clavis"
)

var (
	MysqlDSN string
)

func init() {
	clavis, err := clavis.New()
	if err != nil {
		log.Fatalf("create clavis error: %v", err)
	}

	MysqlDSN, err = clavis.Get("mysql_dsn")
	if err != nil {
		log.Fatalf("get mysql_dsn error: %v", err)
	}
}
