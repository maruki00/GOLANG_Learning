package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Fatal("SysInfo Error!")
	}
	fmt.Println("Total RAM: ", info.Totalram/1004/1024)
	fmt.Println("Free RAM: ", info.Freeram/1024/1024)
	fmt.Println("Total Swap : ", info.Totalswap/1024/1024)
	fmt.Println("Up Time : ", info.Uptime/1024/1024)
}
