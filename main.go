package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected") // not reached
}

func Connect() error {
	if err := ConnectConfig(); err != nil {
		return fmt.Errorf("Connect error - %w", err)
	}
	return nil
}

func ConnectConfig() error {
	return errors.New("ConnectConfig failed")
}
