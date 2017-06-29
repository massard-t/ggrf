package main

import (
	"fmt"
	"time"
	//"net"
	"math/rand"
)

type worker struct {
	ip      string
	timeOut int
	success bool
}

const message string = "Your Redis instance is insecured. You should add a password to it !"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func createIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:6379", rand.Intn(256), rand.Intn(256),
		rand.Intn(256), rand.Intn(256))
}

func main() {
	fmt.Println(createIP())
	fmt.Println(message)
}
