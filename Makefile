all:
	go build dummyprocess.go
	go build driver.go

clean:
	rm driver dummyprocess
