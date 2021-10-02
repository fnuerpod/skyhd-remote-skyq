GOCMD=go
EXE=skyhd-remote-skyq

${EXE}: *.go */*.go
	${GOCMD} build -o ${EXE}

tidy:
	gofmt -s -w .
	go mod tidy
	go vet

clean:
	rm ${EXE}
