package main

import(
	"time"
	"fmt"
)

func main() {
	for {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
	}
}
