DEPS:=ffi.h ffi.pc libffi.a c_interface.h libcgo.a libzcnt.a

all: $(DEPS)
.PHONY: all

# Create a file so that parallel make doesn't call `./install-ffi` for
# each of the deps
$(DEPS):build/.update-modules .install-ffi .install-vdf;

.install-vdf:
	cd chiavdf/src && $(MAKE) -f Makefile.vdf-client
	cp chiavdf/src/libcgo.a .
	cp chiavdf/src/libzcnt.a .
	cp chiavdf/src/c_interface.h .

.install-ffi: rust
	go clean -cache -testcache .
	./install-ffi
	@touch $@

clean:
	go clean -cache -testcache .
	rm -rf $(DEPS) .install-ffi
	rm -f ./runner
	cd rust && cargo clean && cd ..
	cd chiavdf/src && $(MAKE) -f Makefile.vdf-client clean
.PHONY: clean

go-lint: $(DEPS)
	golangci-lint run -v --concurrency 2 --new-from-rev origin/master
.PHONY: go-lint

shellcheck:
	shellcheck install-ffi

lint: shellcheck go-lint

cgo-leakdetect: runner
	valgrind --leak-check=full --show-leak-kinds=definite ./runner
.PHONY: cgo-leakdetect

cgo-gen: $(DEPS)
	go run github.com/xlab/c-for-go --nostamp ffi.yml
.PHONY: cgo-gen

runner: $(DEPS)
	rm -f ./runner
	go build -o ./runner ./cgoleakdetect/
.PHONY: runner

build/.update-modules:
	git submodule update --init --recursive