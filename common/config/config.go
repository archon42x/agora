package config

import (
	"log"

	"github.com/archon42x/clavis-sdk-go/clavis"
)

var (
	JWTKey string
)

func init() {
	clavis, err := clavis.New()
	if err != nil {
		log.Fatalf("create clavis error: %v", err)
	}

	JWTKey, err = clavis.Get("jwt_key")
	if err != nil {
		log.Fatalf("get jwt_key error: %v", err)
	}
}
