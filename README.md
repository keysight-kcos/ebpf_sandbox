## Running the demo
Run make demo.

Run sudo ./runDemo.sh mask, where mask is the hexadecimal representation of the CPU mask you want to enforce.

A mask with the value 3d will set permissions to run a process on CPUs 1, 3, 4, 5, and 6 (0011 1101).

dummyprocess will spawn 5 child processes over the course of 5 seconds.

The generated bpftrace script will watch dummyprocess and all child processes
it spawns based on PID (it is possible to do this based on command name instead). When it sees
these processes being created it will print out information about the CPUs that process is
allowed to run on. 

The script will then call a Go program that will set the CPU mask of the 
parent process based on the mask that was passed as an argument to runDemo.sh.

Conditional execution of the CPU enforcement Go program is possible, but for now it's hardcoded to always run when it sees dummyprocess or its child processes being executed.

## Further steps
It would be much more efficient to bundle everything into one program rather than a set of glued-together programs.
To do this, I would use an eBPF library for Go to run the eBPF programs I need from a single Go program as opposed to using bpftrace as a middleman.

This video looks like a good resource for getting started on this implementation using Cilium's eBPF Go package:
https://www.youtube.com/watch?v=eZp_3EjJdnA
