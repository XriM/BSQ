BINARY_NAME=bsq

BINARY_SRC?=src/main.go 		\
			src/print.go		\
			src/algo.go			\
			src/formating.go	\
			src/error.go

build:
	go build -o ${BINARY_NAME} ${BINARY_SRC}

clean:
	go clean rm ${BINARY_NAME}