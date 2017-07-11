package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	redisCommand string        = "SET info \"Your Redis instance is insecured. You should add a password to it !\"\r\n"
	timeout      time.Duration = time.Second * 2
)

var (
	maxWorker = getEnv("GGRF_MAX_WORKER", 10)
	//maxQueue  = getEnv("GGRF_MAX_QUEUE", 1)
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func getEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}

func createIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:6379", rand.Intn(256), rand.Intn(256),
		rand.Intn(256), rand.Intn(256))
}

func work(order chan string) {
	for {
		ip := <-order
		conn, err := net.DialTimeout("tcp", ip, timeout)
		if err == nil {
			conn.SetDeadline(time.Now().Add(time.Second * 5))
			fmt.Fprintf(conn, redisCommand)
			var line string
			_, err := fmt.Fscanln(conn, &line)
			if err == nil {
				fmt.Println(ip)
				fmt.Println(line)
			}
		}
	}
}

func main() {
	fmt.Println(createIP())
	fmt.Println(redisCommand)
	c := make(chan string, maxWorker)
	fmt.Println("Starting worker")
	go work(c)
	fmt.Println("Worker running... Starting to send orders.")
	for {
		//ip := createIP()
		//fmt.Println("Sending order: ", ip)
		ip := "localhost:6379"
		c <- ip
		fmt.Scanln()
	}
}
