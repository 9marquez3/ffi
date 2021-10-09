package ffi

import (
	"encoding/hex"
	"testing"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/stretchr/testify/assert"
)

func TestGetVdfInfoAndProof(t *testing.T) {
	constants := ConsensusConstants{
		SlotBlocksTarget:               32,
		MinBlocksPerChallengeBlock:     12,
		MaxSubSlotBlocks:               50,
		NumSpsSubSlot:                  16,
		SubSlotItersStarting:           1024,
		DifficultyConstantFactor:       big.Lsh(big.NewInt(2), 24), //33554432,
		DifficultyStarting:             4096,
		DifficultyChangeMaxFactor:      3,
		SubEpochBlocks:                 170,
		EpochBlocks:                    340,
		SignificantBits:                8,
		DiscriminantSizeBits:           16,
		NumberZeroBitsPlotFilter:       1,
		MinPlotSize:                    18,
		MaxPlotSize:                    50,
		SubSlotTimeTarget:              600,
		NumSpIntervalsExtra:            3,
		MaxFutureTime:                  864000,
		NumberOfTimestamps:             11,
		GenesisChallenge:               hexString32("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
		AggSigMeAdditionalData:         hexString("ccd5bb71183532bff220ba46c268991a3ff07eb358e8255a65c30a2dce0e5fbb"),
		GenesisPreFarmPoolPuzzleHash:   hexString32("d23da14695a188ae5708dd152263c4db883eb27edeb936178d4d988b8f3ce5fc"),
		GenesisPreFarmFarmerPuzzleHash: hexString32("3d8765d3a597ec1d99663f6c9816d915b9f68613ac94009884c4addaefcce6af"),
		MaxVdfWitnessSize:              64,
		MempoolBlockBuffer:             6,
		MaxCoinAmount:                  18446744073709551615,
		MaxBlockCostClvm:               11000000000,
		CostPerByte:                    1337,
		WeightProofThreshold:           2,
		WeightProofRecentBlocks:        380,
		MaxBlockCountPerRequests:       32,
		//	RustConditionChecker:
		BlocksCacheSize:         490,
		NetworkType:             1,
		MaxGeneratorSize:        1000000,
		MaxGeneratorRefListSize: 512,
		PoolSubSlotIters:        37600000000,
	}
	var classGroupElement ClassGroupElement
	VDFInfo, VDFProof := get_vdf_info_and_proof(constants, classGroupElement.GetDefaultElement(), hex.EncodeToString(constants.GenesisChallenge[:]), 231, false)

	assert.True(t, hex.EncodeToString(VDFInfo.Challenge[:]) == "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.True(t, VDFInfo.NumberOfIterations == 231)
	assert.True(t, hex.EncodeToString(VDFInfo.Output.Data[:]) == "03004e00010100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	assert.True(t, VDFProof.NormalizedToIdentity == false)
	assert.True(t, hex.EncodeToString(VDFProof.Witness) == "04000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	assert.True(t, VDFProof.WitnessType == 0)
}
