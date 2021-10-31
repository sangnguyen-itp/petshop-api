package main

import (
	"petshop/root"
)

func main() {
	r := root.SetupRouting()
	r.Run(":8080")
}
