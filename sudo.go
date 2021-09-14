package main

import (
	"fmt"
	"sudo/sudo"
)

func main() {
	sd := sudo.FromStr("030070000020040060000830520004008600000000738180700040090057000008090400003200010")
	fmt.Println(sd.PrintStr())

}
