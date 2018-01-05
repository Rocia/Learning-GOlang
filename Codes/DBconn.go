package main

import "github.com/garyburd/redigo/redis"
import "fmt"

func main() {
	//Connect
	c, err := redis.Dial("tcp", "http://ip:port")
	if err != nil {
		panic(err)
	}

	response, err := c.Do("AUTH", "YOUR_PASSWORD")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected! ", response)
	defer c.Close()
}

