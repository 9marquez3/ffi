package ffi

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/stretchr/testify/require"
)

func TestGenerateWinningPoStSectorChallenge(t *testing.T) {
	WorkflowGenerateWinningPoStSectorChallenge(newTestingTeeHelper(t))
}

func TestGenerateWinningPoStSectorChallengeEdgeCase(t *testing.T) {
	WorkflowGenerateWinningPoStSectorChallengeEdgeCase(newTestingTeeHelper(t))
}

func newTestingTeeHelper(t *testing.T) *testingTeeHelper {
	return &testingTeeHelper{t: t}
}

type TestHelper interface {
	AssertEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool
	AssertNoError(err error, msgAndArgs ...interface{}) bool
	AssertTrue(value bool, msgAndArgs ...interface{}) bool
	RequireEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	RequireNoError(err error, msgAndArgs ...interface{})
	RequireTrue(value bool, msgAndArgs ...interface{})
}

type testingTeeHelper struct {
	t *testing.T
}

func (tth *testingTeeHelper) RequireTrue(value bool, msgAndArgs ...interface{}) {
	require.True(tth.t, value, msgAndArgs)
}

func (tth *testingTeeHelper) RequireNoError(err error, msgAndArgs ...interface{}) {
	require.NoError(tth.t, err, msgAndArgs)
}

func (tth *testingTeeHelper) RequireEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	require.Equal(tth.t, expected, actual, msgAndArgs)
}

func (tth *testingTeeHelper) AssertNoError(err error, msgAndArgs ...interface{}) bool {
	return assert.NoError(tth.t, err, msgAndArgs)
}

func (tth *testingTeeHelper) AssertEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return assert.Equal(tth.t, expected, actual, msgAndArgs)
}

func (tth *testingTeeHelper) AssertTrue(value bool, msgAndArgs ...interface{}) bool {
	return assert.True(tth.t, value, msgAndArgs)
}

func WorkflowGenerateWinningPoStSectorChallengeEdgeCase(t TestHelper) {
	for i := 0; i < 10000; i++ {
		var randomnessFr32 [32]byte
		_, err := io.ReadFull(rand.Reader, randomnessFr32[0:31]) // last byte of the 32 is always NUL
		t.RequireNoError(err)

		minerID := randActorID()
		eligibleSectorsLen := uint64(1)

		indices2, err := GenerateWinningPoStSectorChallenge(abi.RegisteredPoStProof_StackedDrgWinning2KiBV1, minerID, randomnessFr32[:], eligibleSectorsLen)
		t.RequireNoError(err)
		t.AssertEqual(1, len(indices2))
		t.AssertEqual(0, int(indices2[0]))
	}
}

func WorkflowGenerateWinningPoStSectorChallenge(t TestHelper) {
	for i := 0; i < 10000; i++ {
		var randomnessFr32 [32]byte
		_, err := io.ReadFull(rand.Reader, randomnessFr32[0:31]) // last byte of the 32 is always NUL
		t.RequireNoError(err)

		minerID := randActorID()
		eligibleSectorsLen := randUInt64()

		if eligibleSectorsLen == 0 {
			continue // no fun
		}

		indices, err := GenerateWinningPoStSectorChallenge(abi.RegisteredPoStProof_StackedDrgWinning2KiBV1, minerID, randomnessFr32[:], eligibleSectorsLen)
		t.AssertNoError(err)

		max := uint64(0)
		for idx := range indices {
			if indices[idx] > max {
				max = indices[idx]
			}
		}

		t.AssertTrue(max < eligibleSectorsLen, "out of range value - max: ", max, "eligibleSectorsLen: ", eligibleSectorsLen)
		t.AssertTrue(uint64(len(indices)) <= eligibleSectorsLen, "should never generate more indices than number of eligible sectors")
	}
}

func randActorID() abi.ActorID {
	bID, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	return abi.ActorID(bID.Uint64())
}

func randUInt64() uint64 {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	return binary.LittleEndian.Uint64(buf)
}
