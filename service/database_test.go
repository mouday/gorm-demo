package service

import (
	"fmt"
	"log"
	"testing"
)

func TestConnectDB(t *testing.T) {

	db, err := GetDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("db: %v\n", db)
}
