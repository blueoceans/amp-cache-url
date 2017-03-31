GO=$(find . -name "*.go" -print)

.SUFFIXES: .go .so
.go.so:
	go build -buildmode=c-shared -o $@ $<
LIBGO=$(wildcard lib/*.go)
LIB=$(LIBGO:.go=.so)

all: $(GO) $(LIB) test
	go build

test:
	go test
