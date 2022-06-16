#!/usr/bin/bash

PID=$(./startDummyProcessGetPID)
echo Dummy process has PID $PID
./generateTracer -pid=$PID -mask=$1
bpftrace --unsafe "${PID}_tracer.bt"
