TASK ?= ./bin/task

install-task:
	curl -sL https://taskfile.dev/install.sh | sh

run:
	go run main.go

run-task:
	${TASK} -v -w run
