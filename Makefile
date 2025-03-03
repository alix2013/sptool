mac:
	go build -o ./bin/sptool_Darwin_x86_64 *.go
linux:
	env GOOS=linux GOARCH=amd64 go build   -o ./bin/sptool_Linux_x86_64 *.go
win:
	env GOOS=windows GOARCH=amd64 go build   -o ./bin/sptool.exe *.go
aix:
	env GOOS=aix GOARCH=ppc64 go build   -o ./bin/sptool_AIX_ppc64 *.go
linuxppc64le:
	env GOOS=linux GOARCH=ppc64le go build   -o ./bin/sptool_Linux_ppc64le *.go
linuxs390x:
	env GOOS=linux GOARCH=s390x go build   -o ./bin/sptool_Linux_s390x *.go
all: mac linux win aix linuxppc64le linuxs390x

deploy: 
	cp bin/* ~/Downloads/sptooltest/
