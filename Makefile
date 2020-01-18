build:
	go build -i -o ${GOPATH}/bin/looker

download:
	go get -u -d github.com/chrismaher/looker-cli
	
install: download build
