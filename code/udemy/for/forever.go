package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
