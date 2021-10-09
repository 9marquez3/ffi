package ffi

import (
	"encoding/hex"
)

func get_vdf_info_and_proof(constants ConsensusConstants, vdf_input ClassGroupElement,
	challenge_hash string, number_iters uint64, normalized_to_identity bool) (VDFInfo, VDFProof) {

	form_size := vdf_input.GetSize()
	result := Prove(challenge_hash, hex.EncodeToString(vdf_input.Data[:]),
		int32(constants.DiscriminantSizeBits), number_iters)
	result_byte, _ := hex.DecodeString(result)

	output := vdf_input.FromBytes(result_byte[:form_size])
	proof_bytes := result_byte[form_size : 2*form_size]

	var challenge_byte32 [32]byte
	challenge_byte, _ := hex.DecodeString(challenge_hash)
	copy(challenge_byte32[:], challenge_byte)
	return VDFInfo{challenge_byte32, number_iters, output}, VDFProof{uint8(0), proof_bytes, normalized_to_identity}
}
