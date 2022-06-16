// Sets a process's CPU mask.
package main

import (
	"flag"
	"strconv"
	"log"
	"golang.org/x/sys/unix"
)

var pidRef = flag.Int("pid", -1, "PID of process who's CPU mask will be set.")
var maskRef = flag.String("mask", "", "Hexadecimal representation of CPU mask.")

func parseHex(hex string) []uint64 {
	longs := make([]uint64, 0)
	index := 0
	for index+16 <= len(hex) {
		val, err := strconv.ParseUint(hex[index:index+16], 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		longs = append(longs, val)
		index += 16
	}
	if index < len(hex) {
		val, err := strconv.ParseUint(hex[index:len(hex)], 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		longs = append(longs, val)
	}
	return longs
}

func main() {
	flag.Parse()

	var set *unix.CPUSet = &unix.CPUSet{}
	longs := parseHex(*maskRef)
	offset := 0
	for _, bits := range longs {
		for i := 0; i < 64; i++ {
			var temp uint64 = 1
			temp = temp << i
			if temp & bits != 0 {
				set.Set(i+offset)
			}
		}
		offset += 64
	}
	/*
	fmt.Printf("CPU Mask of started process: %x\n", set)
	fmt.Printf("Number of CPUs in mask: %d\n\n", set.Count())
	*/


	unix.SchedSetaffinity(*pidRef, set)
	err := unix.SchedGetaffinity(*pidRef, set)
	if err != nil && err.Error() != "no such process" {
		log.Fatal(err)
	}
}
