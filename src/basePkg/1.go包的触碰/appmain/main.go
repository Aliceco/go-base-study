package main

import (
	"basePkg/services"
	"fmt"
)

func main() {
	res := services.GetUser()
	fmt.Print(res)
}

