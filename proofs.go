package ffi

import (
	"github.com/filecoin-project/go-state-types/abi"
	proof5 "github.com/filecoin-project/specs-actors/v5/actors/runtime/proof"
)

// VerifyWinningPoSt returns true if the Winning PoSt-generation operation from which its
// inputs were derived was valid, and false if not.
func VerifyWinningPoSt(info proof5.WinningPoStVerifyInfo) (bool, error) {

	return false, nil
}

// VerifyWindowPoSt returns true if the Winning PoSt-generation operation from which its
// inputs were derived was valid, and false if not.
func VerifyWindowPoSt(info proof5.WindowPoStVerifyInfo) (bool, error) {

	return false, nil
}

// GenerateWinningPoStSectorChallenge
func GenerateWinningPoStSectorChallenge(
	proofType abi.RegisteredPoStProof,
	minerID abi.ActorID,
	randomness abi.PoStRandomness,
	eligibleSectorsLen uint64,
) ([]uint64, error) {
	// copy from C memory space to Go
	out := make([]uint64, 0)
	return out, nil
}

// GenerateWinningPoSt
func GenerateWinningPoSt(
	minerID abi.ActorID,
	privateSectorInfo SortedPrivateSectorInfo,
	randomness abi.PoStRandomness,
) ([]proof5.PoStProof, error) {

	return nil, nil
}

// GenerateWindowPoSt
func GenerateWindowPoSt(
	minerID abi.ActorID,
	privateSectorInfo SortedPrivateSectorInfo,
	randomness abi.PoStRandomness,
) ([]proof5.PoStProof, []abi.SectorNumber, error) {

	return nil, nil, nil
}
