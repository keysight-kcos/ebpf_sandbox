package main

import (
	"fmt"
	"runtime"
	"log"
	"os/exec"
	"golang.org/x/sys/unix"
)

func main() {
	fmt.Printf("Number of cores on machine: %d\n", runtime.NumCPU())
	cmd := exec.Command("./dummyprocess")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	pid := cmd.Process.Pid
	fmt.Printf("PID of started process: %d\n", pid)

	var set *unix.CPUSet = &unix.CPUSet{}
	err = unix.SchedGetaffinity(pid, set)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CPU Mask of started process: %x\n", set)
	fmt.Printf("Number of CPUs in mask: %d\n\n", set.Count())


	set.Clear(0)
	unix.SchedSetaffinity(pid, set)
	err = unix.SchedGetaffinity(pid, set)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After changing CPU affinity:\n")
	fmt.Printf("CPU Mask of started process: %x\n", set)
	fmt.Printf("Number of CPUs in mask: %d\n\n", set.Count())

	set.Clear(1)
	unix.SchedSetaffinity(pid, set)
	err = unix.SchedGetaffinity(pid, set)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After changing CPU affinity:\n")
	fmt.Printf("CPU Mask of started process: %x\n", set)
	fmt.Printf("Number of CPUs in mask: %d\n\n", set.Count())
}
