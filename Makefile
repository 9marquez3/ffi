DEPS:=ffi.h ffi.pc libffi.a

all: $(DEPS)
.PHONY: all

# Create a file so that parallel make doesn't call `./install-ffi` for
# each of the deps
$(DEPS): .install-ffi  chiapos;

.install-ffi: rust
	go clean -cache -testcache .
	./install-ffi
	@touch $@

chiapos:
	cd $(dir $@); git clone https://github.com/cnc-project/chiapos.git
	cd $@; git checkout dev_cgo; mkdir -p build && cd build; cmake ../; make

clean:
	go clean -cache -testcache .
	rm -rf $(DEPS) .install-ffi
	rm -f ./runner
	cd rust && cargo clean && cd ..
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
