package main

import (
	"fmt"
	"sudo/sudo"
	"time"
)

func main() {
	origin := "300900000007000250500000010000102079000008100000004000070000000020070045001300006"
	sd := sudo.FromStr(origin)
	var start = time.Now()
	resolve := sd.Resolve()
	var stop = time.Now()
	fmt.Printf("cost time %.4f ms.\n", float64(stop.Nanosecond()-start.Nanosecond())/1000000)
	fmt.Println(resolve)
	fmt.Println(origin)
	fmt.Println(sd.ToStr())
	fmt.Println(sd.PrintStr())

}
