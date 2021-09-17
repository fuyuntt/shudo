package main

import (
	"fmt"
	"sudo/sudo"
	"time"
)

func main() {
	origin := "050900000800040307000280190538607940020301000109804623907400000045000209000030070"
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
