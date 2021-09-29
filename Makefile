# this target will build our project by providing the basic
# commands to get any github library dependencies, 'clean'
# our golang code in our go files and create an executable
build:
	go get github.com/Robbie08/SmartSafe/pkg/piUtils
	go get github.com/sirupsen/logrus
	gofmt -w pkg/piUtils/piUtils.go
	gofmt -w main.go
	go build -o bin/main main.go

# This target does what it says, once we have everything setup
# we can just run our program.
run:
	go run main.go
