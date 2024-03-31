package main

import (
	"fmt"

	"github.com/wandermaia/pos-golang-apis/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
