demo: dummyprocess dummychild generateTracer startDummyProcessGetPID setCPUMask
	echo "runDemo.sh is ready to be run."

startDummyProcessGetPID: startDummyProcessGetPID.go
	go build startDummyProcessGetPID.go

generateTracer: generateTracer.go
	go build generateTracer.go

dummyprocess: dummyprocess.go 
	go build dummyprocess.go

dummychild: dummychild.go
	go build dummychild.go

setCPUMask: setCPUMask.go
	go build setCPUMask.go

clean:
	rm driver dummyprocess dummychild generateTracer startDummyProcessGetPID *_tracer.bt setCPUMask
