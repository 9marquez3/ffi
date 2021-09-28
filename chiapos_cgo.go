package main

// #cgo LDFLAGS: -L chiapos/build -lcgo -lstdc++ -lm
// #cgo CFLAGS: -I chiapos/src
// #include "c_interface.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func validate_proof() bool {
	seed := "05683404333717545b0a6f0c0dde9710e4d3fe2d5cc6cc0a090a0b818bab0f17"
	k := 21
	challenge := "df3f619804a92fdb4057192dc43dd748ea778adc52bc498ce80524c014b81119"
	proof := "7180f6b52a24bf8d35798f6ed7d1aa2ae6c5527646e86cc2225514269b10856a7f919f16aa0871a5bb37b11b7da3ec1558f20d40fafa553e14c7317b6ca6a41d96e90666f6e86dfa97d14af3cc0e84dade8ba4ca1c7b55697fac1b1ba0a4e2c8927fb4fee8a07e01a7be884b37e062d4b62a66c008f3adc16e724f3641af0402a434bbd831954b8fb4eb0b7c1278a6d3e81e70df86147edb379e54cd9b7714145e0aa1a3ce4abd07"
	result := C.c_validate_proof(C.CString(seed), C.uint8_t(k), C.CString(challenge), C.CString(proof))

	defer C.free(unsafe.Pointer(result))
	fmt.Println("discriminant:", C.GoString(result))

	if C.GoString(result) == "d404956ce6b81abd8339c62782be2278f168f7539d92a4ac4b22a9cb8653c6ee" {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(validate_proof())
}
