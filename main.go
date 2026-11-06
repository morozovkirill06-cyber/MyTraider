package main

import (
	"time"
	"fmt"
)



func main(){
	
	for {
		candle, err := Parser()
		fmt.Println(candle, err)
		time.Sleep(1 * time.Minute)
	}
}