GO=$(find . -name "*.go" -print)

.SUFFIXES: .go .so
.go.so:
	go build -buildmode=c-shared -o $@ $<
LIBGO=$(wildcard lib/*.go)
LIB=$(LIBGO:.go=.so)

all: $(GO) $(LIB) test
	go build
	cd cmd/amp-cache-url && go build

test:
	go test

clean:
	rm -f cmd/amp-cache-url/amp-cache-url $(wildcard lib/*.h) $(wildcard lib/*.so)
