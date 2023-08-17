package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	fmt.Println(token)
	fmt.Println("fdgdgfd")
}
