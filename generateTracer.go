// Give it a PID as a flag. A bpftrace script will be generated to watch for that PID and
// execute the CPU control program if it is seen starting up.
package main

import (
	"fmt"
	"flag"
	"log"
	"os"
)

var pidRef = flag.Int("pid", -1, "PID to generate bpftrace script for.")
var maskRef = flag.String("mask", "", "Hexadecimal representation of CPU mask.")

func main() {
	flag.Parse()
	pid := *pidRef
	mask := *maskRef
	if pid == -1 {
		flag.PrintDefaults()
		return;
	}
	fmt.Printf("Generating bpftrace script for PID %d.\n", pid)

	f, err := os.OpenFile(fmt.Sprintf("%d_tracer.bt", pid), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("#include <linux/sched.h>\n\n")
	f.WriteString(fmt.Sprintf("BEGIN {\nprintf(\"Watching %d...\\n\")", pid))
	f.WriteString("}\n\n")
	f.WriteString(
		fmt.Sprintf(
			"tracepoint:syscalls:sys_enter_execve /pid == %d | curtask->parent->pid == %d/ {\n",
			pid,
			pid,
		))
	f.WriteString("\tprintf(\"CPUs allowed on %s: %d\\n\", comm, curtask->nr_cpus_allowed);\n")

	totalCPUs := 6
	longs := totalCPUs / 64 // number of elements we need from the mask.
	if totalCPUs % 64 != 0 { longs++ }

	f.WriteString("\tprintf(\"CPU_mask: ")
	for i := 0; i < longs; i++ {
		f.WriteString("%x")
	}
	f.WriteString("\\n\", ")
	for i := 0; i < longs-1; i++ {
		f.WriteString(fmt.Sprintf("curtask->cpus_mask.bits[%d], ", i))
	}
	f.WriteString(fmt.Sprintf("curtask->cpus_mask.bits[%d]);\n", longs-1))

	f.WriteString(fmt.Sprintf("\tsystem(\"./setCPUMask -pid=%d -mask=%s\");\n", pid, mask))
	f.WriteString("}\n\n")
}
