package main

import (
	"fmt"
	"log"
	"wbtasksl2/8/timereceiver"
)

func main() {
	address := "0.beevik-ntp.pool.ntp.org"
	timeReceiver := timereceiver.New(address)
	curTime, err := timeReceiver.GetTime()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(curTime)

}
