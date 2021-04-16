package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	Stats()
}

func Stats()  {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.NumGC)
}

func input()  {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Print("Masukkan nama : "); scanner.Scan()
	var name = scanner.Text()
	fmt.Print("Masukkan umur : "); scanner.Scan()
	var age = scanner.Text()

	fmt.Println(name)
	fmt.Println(age)
}