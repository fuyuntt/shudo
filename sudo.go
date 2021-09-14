package main

import (
	"fmt"
	"sudo/sudo"
	"time"
)

func main() {
	origin := "030070000020040060000830520004008600000000738180700040090057000008090400003200010"
	sd := sudo.FromStr(origin)
	var start = time.Now()
	resolve := sd.Resolve()
	var stop = time.Now()
	fmt.Printf("cost time %d ns.\n", stop.Nanosecond()-start.Nanosecond())
	fmt.Println(resolve)
	fmt.Println(origin)
	fmt.Println(sd.ToStr())
	fmt.Println(sd.PrintStr())

}
