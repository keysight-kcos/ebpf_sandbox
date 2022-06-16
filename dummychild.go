package main

import "time"

func main() {
	z := 664
	for i := 0; i < 1 << 31; i++ {
		z *= z
	}
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
	}
}
