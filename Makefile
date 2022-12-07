
ifeq ($(OS),Windows_NT)
	GOOS=windows
else
	GOOS=linux
endif

OBJ=./hls

all: clean ${OBJ}

${OBJ}:
	go mod tidy
	go build

make: ${OBJ}

clean:
	rm -f ${OBJ}

install: ${OBJ}
	sudo cp $(OBJ) /bin/hls
	
.PHONY: clean
