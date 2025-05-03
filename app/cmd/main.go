package main

import (
	core "app/core/db"
	"app/services/auth"
	"fmt"
)

func main() {
	fmt.Println("Initializing databases...")
	core.InitDatabase("root:password@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println("Starting server...")
	auth.StartAuthService()	
}