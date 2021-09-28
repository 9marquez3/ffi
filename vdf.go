package ffi

/*
#cgo LDFLAGS: -L . -lcgo -lstdc++ -lm -lgmp -lzcnt
#cgo CFLAGS: -I ./
#include "c_interface.h"
*/
import "C"
import (
	"unsafe"
)
	
func CreateDiscriminant(challengeHash  string, size int32) string {
	result := C.c_create_discriminant(C.CString(challengeHash), C.int(size))
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

func Prove(challengeHash, initialEl string, discriminantSizeBits int32, numIterations uint64) string {

	result := C.c_prove(C.CString(challengeHash), C.CString(initialEl), C.int(discriminantSizeBits), C.uint64_t(numIterations))
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

func GetBFromNWesolowski(discriminant, x, y_result, y_proof string, iters uint64) string {
	result := C.c_get_b_from_n_wesolowski(C.CString(discriminant), C.CString(x), C.CString(y_result+y_proof), C.uint64_t(iters), C.uint64_t(0))

	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

func VerifyNWesolowskiWithB(discriminant, b_hex, x, y_proof string, iters uint64) (bool, string) {
	result := C.c_verify_n_wesolowski_with_b(C.CString(discriminant), C.CString(b_hex), C.CString(x), C.CString(y_proof), C.uint64_t(iters), C.uint64_t(0))

	defer C.free(unsafe.Pointer(result.y_from_compression))

	return bool(result.is_valid), C.GoString(result.y_from_compression)
}

func VerifyWesolowski(discriminant, initial_el, result_y, proof string, iters uint64) bool {
	is_valid := C.c_verify_wesolowski(C.CString(discriminant), C.CString(initial_el), C.CString(result_y), C.CString(proof), C.uint64_t(iters))

	return bool(is_valid)
}

func VerifyNWesolowski(discriminant, initial_el, recursion string, iters, discriminant_size uint64) bool {
	is_valid := C.c_verify_n_wesolowski(C.CString(discriminant), C.CString(initial_el), C.CString(recursion), C.uint64_t(iters), C.uint64_t(discriminant_size), C.uint64_t(5))
	return bool(is_valid)
}