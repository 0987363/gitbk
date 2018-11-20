all: 
	go build -v -ldflags "-X main.BuildVersion=$(BUILD_VERSION).$(BUILD_NUMBER) -X main.BuildCommit=$(BUILD_COMMIT) -X main.BuildDate=$(BUILD_DATE)" 

linux: 
	GOARCH=amd64 GOOS=linux go build -v -ldflags "-X main.BuildVersion=$(BUILD_VERSION).$(BUILD_NUMBER) -X main.BuildCommit=$(BUILD_COMMIT) -X main.BuildDate=$(BUILD_DATE)" 
