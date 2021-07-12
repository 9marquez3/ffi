// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../ffi.pc
#include "../ffi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// FilAggregate function as declared in ffi/ffi.h:117
func FilAggregate(flattenedSignaturesPtr []byte, flattenedSignaturesLen uint) *FilAggregateResponse {
	cflattenedSignaturesPtr, cflattenedSignaturesPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&flattenedSignaturesPtr)))
	cflattenedSignaturesLen, cflattenedSignaturesLenAllocMap := (C.size_t)(flattenedSignaturesLen), cgoAllocsUnknown
	__ret := C.fil_aggregate(cflattenedSignaturesPtr, cflattenedSignaturesLen)
	runtime.KeepAlive(cflattenedSignaturesLenAllocMap)
	runtime.KeepAlive(cflattenedSignaturesPtrAllocMap)
	__v := NewFilAggregateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilCreateZeroSignature function as declared in ffi/ffi.h:125
func FilCreateZeroSignature() *FilZeroSignatureResponse {
	__ret := C.fil_create_zero_signature()
	__v := NewFilZeroSignatureResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilDestroyAggregateResponse function as declared in ffi/ffi.h:127
func FilDestroyAggregateResponse(ptr *FilAggregateResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_aggregate_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyGpuDeviceResponse function as declared in ffi/ffi.h:129
func FilDestroyGpuDeviceResponse(ptr *FilGpuDeviceResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_gpu_device_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyHashResponse function as declared in ffi/ffi.h:131
func FilDestroyHashResponse(ptr *FilHashResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_hash_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyInitLogFdResponse function as declared in ffi/ffi.h:133
func FilDestroyInitLogFdResponse(ptr *FilInitLogFdResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_init_log_fd_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyPrivateKeyGenerateResponse function as declared in ffi/ffi.h:135
func FilDestroyPrivateKeyGenerateResponse(ptr *FilPrivateKeyGenerateResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_private_key_generate_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyPrivateKeyPublicKeyResponse function as declared in ffi/ffi.h:137
func FilDestroyPrivateKeyPublicKeyResponse(ptr *FilPrivateKeyPublicKeyResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_private_key_public_key_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyPrivateKeySignResponse function as declared in ffi/ffi.h:139
func FilDestroyPrivateKeySignResponse(ptr *FilPrivateKeySignResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_private_key_sign_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDestroyZeroSignatureResponse function as declared in ffi/ffi.h:141
func FilDestroyZeroSignatureResponse(ptr *FilZeroSignatureResponse) {
	cptr, cptrAllocMap := ptr.PassRef()
	C.fil_destroy_zero_signature_response(cptr)
	runtime.KeepAlive(cptrAllocMap)
}

// FilDropSignature function as declared in ffi/ffi.h:146
func FilDropSignature(sig []byte) {
	csig, csigAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&sig)))
	C.fil_drop_signature(csig)
	runtime.KeepAlive(csigAllocMap)
}

// FilGetGpuDevices function as declared in ffi/ffi.h:151
func FilGetGpuDevices() *FilGpuDeviceResponse {
	__ret := C.fil_get_gpu_devices()
	__v := NewFilGpuDeviceResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilHash function as declared in ffi/ffi.h:161
func FilHash(messagePtr []byte, messageLen uint) *FilHashResponse {
	cmessagePtr, cmessagePtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&messagePtr)))
	cmessageLen, cmessageLenAllocMap := (C.size_t)(messageLen), cgoAllocsUnknown
	__ret := C.fil_hash(cmessagePtr, cmessageLen)
	runtime.KeepAlive(cmessageLenAllocMap)
	runtime.KeepAlive(cmessagePtrAllocMap)
	__v := NewFilHashResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilHashVerify function as declared in ffi/ffi.h:175
func FilHashVerify(signaturePtr []byte, flattenedMessagesPtr []byte, flattenedMessagesLen uint, messageSizesPtr []uint, messageSizesLen uint, flattenedPublicKeysPtr []byte, flattenedPublicKeysLen uint) int32 {
	csignaturePtr, csignaturePtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&signaturePtr)))
	cflattenedMessagesPtr, cflattenedMessagesPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&flattenedMessagesPtr)))
	cflattenedMessagesLen, cflattenedMessagesLenAllocMap := (C.size_t)(flattenedMessagesLen), cgoAllocsUnknown
	cmessageSizesPtr, cmessageSizesPtrAllocMap := copyPSizeTBytes((*sliceHeader)(unsafe.Pointer(&messageSizesPtr)))
	cmessageSizesLen, cmessageSizesLenAllocMap := (C.size_t)(messageSizesLen), cgoAllocsUnknown
	cflattenedPublicKeysPtr, cflattenedPublicKeysPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&flattenedPublicKeysPtr)))
	cflattenedPublicKeysLen, cflattenedPublicKeysLenAllocMap := (C.size_t)(flattenedPublicKeysLen), cgoAllocsUnknown
	__ret := C.fil_hash_verify(csignaturePtr, cflattenedMessagesPtr, cflattenedMessagesLen, cmessageSizesPtr, cmessageSizesLen, cflattenedPublicKeysPtr, cflattenedPublicKeysLen)
	runtime.KeepAlive(cflattenedPublicKeysLenAllocMap)
	runtime.KeepAlive(cflattenedPublicKeysPtrAllocMap)
	runtime.KeepAlive(cmessageSizesLenAllocMap)
	runtime.KeepAlive(cmessageSizesPtrAllocMap)
	runtime.KeepAlive(cflattenedMessagesLenAllocMap)
	runtime.KeepAlive(cflattenedMessagesPtrAllocMap)
	runtime.KeepAlive(csignaturePtrAllocMap)
	__v := (int32)(__ret)
	return __v
}

// FilInitLogFd function as declared in ffi/ffi.h:192
func FilInitLogFd(logFd int32) *FilInitLogFdResponse {
	clogFd, clogFdAllocMap := (C.int)(logFd), cgoAllocsUnknown
	__ret := C.fil_init_log_fd(clogFd)
	runtime.KeepAlive(clogFdAllocMap)
	__v := NewFilInitLogFdResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyGenerate function as declared in ffi/ffi.h:197
func FilPrivateKeyGenerate() *FilPrivateKeyGenerateResponse {
	__ret := C.fil_private_key_generate()
	__v := NewFilPrivateKeyGenerateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyGenerateWithSeed function as declared in ffi/ffi.h:210
func FilPrivateKeyGenerateWithSeed(rawSeed Fil32ByteArray) *FilPrivateKeyGenerateResponse {
	crawSeed, crawSeedAllocMap := rawSeed.PassValue()
	__ret := C.fil_private_key_generate_with_seed(crawSeed)
	runtime.KeepAlive(crawSeedAllocMap)
	__v := NewFilPrivateKeyGenerateResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeyPublicKey function as declared in ffi/ffi.h:221
func FilPrivateKeyPublicKey(rawPrivateKeyPtr []byte) *FilPrivateKeyPublicKeyResponse {
	crawPrivateKeyPtr, crawPrivateKeyPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&rawPrivateKeyPtr)))
	__ret := C.fil_private_key_public_key(crawPrivateKeyPtr)
	runtime.KeepAlive(crawPrivateKeyPtrAllocMap)
	__v := NewFilPrivateKeyPublicKeyResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilPrivateKeySign function as declared in ffi/ffi.h:234
func FilPrivateKeySign(rawPrivateKeyPtr []byte, messagePtr []byte, messageLen uint) *FilPrivateKeySignResponse {
	crawPrivateKeyPtr, crawPrivateKeyPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&rawPrivateKeyPtr)))
	cmessagePtr, cmessagePtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&messagePtr)))
	cmessageLen, cmessageLenAllocMap := (C.size_t)(messageLen), cgoAllocsUnknown
	__ret := C.fil_private_key_sign(crawPrivateKeyPtr, cmessagePtr, cmessageLen)
	runtime.KeepAlive(cmessageLenAllocMap)
	runtime.KeepAlive(cmessagePtrAllocMap)
	runtime.KeepAlive(crawPrivateKeyPtrAllocMap)
	__v := NewFilPrivateKeySignResponseRef(unsafe.Pointer(__ret))
	return __v
}

// FilVerify function as declared in ffi/ffi.h:249
func FilVerify(signaturePtr []byte, flattenedDigestsPtr []byte, flattenedDigestsLen uint, flattenedPublicKeysPtr []byte, flattenedPublicKeysLen uint) int32 {
	csignaturePtr, csignaturePtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&signaturePtr)))
	cflattenedDigestsPtr, cflattenedDigestsPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&flattenedDigestsPtr)))
	cflattenedDigestsLen, cflattenedDigestsLenAllocMap := (C.size_t)(flattenedDigestsLen), cgoAllocsUnknown
	cflattenedPublicKeysPtr, cflattenedPublicKeysPtrAllocMap := copyPUint8TBytes((*sliceHeader)(unsafe.Pointer(&flattenedPublicKeysPtr)))
	cflattenedPublicKeysLen, cflattenedPublicKeysLenAllocMap := (C.size_t)(flattenedPublicKeysLen), cgoAllocsUnknown
	__ret := C.fil_verify(csignaturePtr, cflattenedDigestsPtr, cflattenedDigestsLen, cflattenedPublicKeysPtr, cflattenedPublicKeysLen)
	runtime.KeepAlive(cflattenedPublicKeysLenAllocMap)
	runtime.KeepAlive(cflattenedPublicKeysPtrAllocMap)
	runtime.KeepAlive(cflattenedDigestsLenAllocMap)
	runtime.KeepAlive(cflattenedDigestsPtrAllocMap)
	runtime.KeepAlive(csignaturePtrAllocMap)
	__v := (int32)(__ret)
	return __v
}
