package ffi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"sort"

	bls "github.com/cnc-project/cnc-bls"
	"github.com/filecoin-project/go-state-types/abi"
	big2 "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"
	"github.com/ipfs/go-cid"
)

// BLS

// SignatureBytes is the length of a BLS signature
const SignatureBytes = 96

// PrivateKeyBytes is the length of a BLS private key
const PrivateKeyBytes = 32

// PublicKeyBytes is the length of a BLS public key
const PublicKeyBytes = 48

// DigestBytes is the length of a BLS message hash/digest
const DigestBytes = 96

// Signature is a compressed affine
type Signature [SignatureBytes]byte

// PrivateKey is a compressed affine
type PrivateKey [PrivateKeyBytes]byte

// PublicKey is a compressed affine
type PublicKey [PublicKeyBytes]byte

// Message is a byte slice
type Message []byte

// Digest is a compressed affine
type Digest [DigestBytes]byte

// Used when generating a private key deterministically
type PrivateKeyGenSeed [32]byte

// Proofs

// SortedPublicSectorInfo is a slice of publicSectorInfo sorted
// (lexicographically, ascending) by sealed (replica) CID.
type SortedPublicSectorInfo struct {
	f []publicSectorInfo
}

// SortedPrivateSectorInfo is a slice of PrivateSectorInfo sorted
// (lexicographically, ascending) by sealed (replica) CID.
type SortedPrivateSectorInfo struct {
	f []PrivateSectorInfo
}

func newSortedPublicSectorInfo(sectorInfo ...publicSectorInfo) SortedPublicSectorInfo {
	fn := func(i, j int) bool {
		return bytes.Compare(sectorInfo[i].SealedCID.Bytes(), sectorInfo[j].SealedCID.Bytes()) == -1
	}

	sort.Slice(sectorInfo[:], fn)

	return SortedPublicSectorInfo{
		f: sectorInfo,
	}
}

// Values returns the sorted publicSectorInfo as a slice
func (s *SortedPublicSectorInfo) Values() []publicSectorInfo {
	return s.f
}

// MarshalJSON JSON-encodes and serializes the SortedPublicSectorInfo.
func (s SortedPublicSectorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.f)
}

// UnmarshalJSON parses the JSON-encoded byte slice and stores the result in the
// value pointed to by s.f. Note that this method allows for construction of a
// SortedPublicSectorInfo which violates its invariant (that its publicSectorInfo are sorted
// in some defined way). Callers should take care to never provide a byte slice
// which would violate this invariant.
func (s *SortedPublicSectorInfo) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.f)
}

// NewSortedPrivateSectorInfo returns a SortedPrivateSectorInfo
func NewSortedPrivateSectorInfo(sectorInfo ...PrivateSectorInfo) SortedPrivateSectorInfo {
	fn := func(i, j int) bool {
		return bytes.Compare(sectorInfo[i].SealedCID.Bytes(), sectorInfo[j].SealedCID.Bytes()) == -1
	}

	sort.Slice(sectorInfo[:], fn)

	return SortedPrivateSectorInfo{
		f: sectorInfo,
	}
}

// Values returns the sorted PrivateSectorInfo as a slice
func (s *SortedPrivateSectorInfo) Values() []PrivateSectorInfo {
	return s.f
}

// MarshalJSON JSON-encodes and serializes the SortedPrivateSectorInfo.
func (s SortedPrivateSectorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.f)
}

func (s *SortedPrivateSectorInfo) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.f)
}

type publicSectorInfo struct {
	PoStProofType abi.RegisteredPoStProof
	SealedCID     cid.Cid
	SectorNum     abi.SectorNumber
}

type PrivateSectorInfo struct {
	proof.SectorInfo
	CacheDirPath     string
	PoStProofType    abi.RegisteredPoStProof
	SealedSectorPath string
}

// AllocationManager is an interface that provides Free() capability.
type AllocationManager interface {
	Free()
}

type BigInt = big2.Int

// vdf and proof
type ConsensusConstants struct {
	SlotBlocksTarget               uint64 // How many blocks to target per sub-slot
	MinBlocksPerChallengeBlock     uint8  // How many blocks must be created per slot (to make challenge sb)
	MaxSubSlotBlocks               uint64 //
	NumSpsSubSlot                  uint64 // The number of signage points per sub-slot (including the 0th sp at the sub-slot start)
	SubSlotItersStarting           uint64
	DifficultyConstantFactor       BigInt
	DifficultyStarting             uint64
	DifficultyChangeMaxFactor      uint64
	SubEpochBlocks                 uint64
	EpochBlocks                    uint64
	SignificantBits                int64
	DiscriminantSizeBits           int64
	NumberZeroBitsPlotFilter       int64
	MinPlotSize                    int64
	MaxPlotSize                    int64
	SubSlotTimeTarget              int64
	NumSpIntervalsExtra            int64
	MaxFutureTime                  int64
	NumberOfTimestamps             int64
	GenesisChallenge               [32]byte
	AggSigMeAdditionalData         []byte
	GenesisPreFarmPoolPuzzleHash   [32]byte
	GenesisPreFarmFarmerPuzzleHash [32]byte
	MaxVdfWitnessSize              int64
	MempoolBlockBuffer             int64
	MaxCoinAmount                  uint64
	MaxBlockCostClvm               int64
	CostPerByte                    int64
	WeightProofThreshold           uint8
	WeightProofRecentBlocks        uint64
	MaxBlockCountPerRequests       uint64
	//	RustConditionChecker           uint64
	BlocksCacheSize         uint64
	NetworkType             int64
	MaxGeneratorSize        uint64
	MaxGeneratorRefListSize uint64
	PoolSubSlotIters        uint64
}

// ClassGroupElement
type ClassGroupElement struct {
	Data [100]byte
}

func (c ClassGroupElement) FromBytes(d []byte) ClassGroupElement {
	var data ClassGroupElement
	copy(data.Data[:], d)
	return data
}

func (c ClassGroupElement) GetDefaultElement() ClassGroupElement {
	return c.FromBytes([]byte{0x08})
}

func (c ClassGroupElement) GetSize() int {
	return 100
}

// VDFInfo
type VDFInfo struct {
	// challenge: bytes32  # Used to generate the discriminant (VDF group)
	Challenge [32]byte
	// number_of_iterations: uint64
	NumberOfIterations uint64
	// output: ClassgroupElement {data: bytes100}
	Output ClassGroupElement
}

// VDFProof
type VDFProof struct {
	// witnessType: uint8
	WitnessType uint8
	// witness: bytes
	Witness []byte
	// normalizedToIdentity: bool
	NormalizedToIdentity bool
}

type HashData struct {
	Data bls.HashDigest256
}

func (h *HashData) GetHashDate() bls.HashDigest256 {
	return h.Data
}

func (h *HashData) IsZero() bool {
	return h.Data.IsZero()
}

func (h *HashData) Bytes() []byte {
	return h.Data.Bytes()
}

func hexString32(s string) [32]byte {
	if len(s) != 64 {
		panic("len s != 64")
	}
	bytes, _ := hex.DecodeString(s)
	a := [32]byte{}
	copy(a[:], bytes)
	return a
}

func hexString(s string) []byte {
	if len(s) != 64 {
		panic("len s != 64")
	}
	bytes, _ := hex.DecodeString(s)
	a := make([]byte, 32)
	copy(a, bytes)
	return a
}
