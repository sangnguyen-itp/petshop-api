package main

import (
	"fmt"
	"petshop/assets/ldap"
)

func main() {
	server := ldap.NewServer()
	fmt.Println("LDAP Server is running...")
	server.Run("127.0.0.1:389")
}
