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

Conditional execution of this script is possible but for now it's hardcoded to always run.
