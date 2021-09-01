package ffi

import (
	bls "github.com/cnc-project/cnc-bls"
)

// Hash computes the digest of a message
func Hash(publicKey PublicKey, message Message) Digest {
	pk, _ := bls.NewPublicKey(publicKey[:])
	// HashToCurve
	digest := (&bls.AugSchemeMPL{}).HashToCurve(pk, message)

	var out Digest
	copy(out[:], digest)
	return out
}

// Verify verifies that a signature is the aggregated signature of digests - pubkeys
func Verify(signature *Signature, digests []Digest, publicKeys []PublicKey) bool {
	// type to [][]byte
	dgs := make([][]byte, len(digests))
	for i, _ := range digests {
		dgs[i] = digests[i][:]
	}

	pks := make([][]byte, len(publicKeys))
	for i, _ := range publicKeys {
		pks[i] =  publicKeys[i][:]
	}
	return (&bls.AugSchemeMPL{}).AggregateHashVerify(pks, dgs, signature[:])
}

// HashVerify verifies that a signature is the aggregated signature of hashed messages.
func HashVerify(signature *Signature, messages []Message, publicKeys []PublicKey) bool {
	// type to [][]byte
	msgs := make([][]byte, len(messages))
	for i, _ := range messages {
		msgs[i] = messages[i][:]
	}

	pks := make([][]byte, len(publicKeys))
	for i, _ := range publicKeys {
		pks[i] = publicKeys[i][:]
	}
	return (&bls.AugSchemeMPL{}).AggregateVerify(pks, msgs, signature[:])
}

// Aggregate aggregates signatures together into a new signature. If the
// provided signatures cannot be aggregated (due to invalid input or an
// an operational error), Aggregate will return nil.
func Aggregate(signatures []Signature) *Signature {
	if len(signatures) == 1 {
		return &signatures[0]
	}

	// type to [][]byte
	ss := make([][]byte, len(signatures))
	for i, _ := range signatures {
		ss[i] = signatures[i][:]
	}

	aggrs, err := (&bls.AugSchemeMPL{}).Aggregate(ss...)
	if err != nil {
		return nil
	}

	var out Signature
	copy(out[:], aggrs)
	return &out
}

// PrivateKeyGenerate generates a private key
func PrivateKeyGenerate() PrivateKey {
	entropy, _ := bls.NewEntropy()
	mnemonic, _ := bls.NewMnemonic(entropy)
	seed := bls.NewSeed(mnemonic, "")
	pk := bls.KeyGen(seed[:])

	var out PrivateKey
	copy(out[:], pk.Bytes())
	return out
}

// PrivateKeyGenerateWithSeed generates a private key in a predictable manner
func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	pk := bls.KeyGen(seed[:])
	var out PrivateKey
	copy(out[:], pk.Bytes())
	return out
}

// PrivateKeySign signs a message
func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {

	pk := bls.KeyFromBytes(privateKey[:])
	sign := (&bls.AugSchemeMPL{}).Sign(pk, message)

	var signature Signature
	copy(signature[:], sign)
	return &signature
}

// PrivateKeyPublicKey gets the public key for a private key
func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {

	pubK := bls.KeyFromBytes(privateKey[:]).GetPublicKey()

	var publicKey PublicKey
	copy(publicKey[:], pubK.Bytes())
	return publicKey
}

// CreateZeroSignature creates a zero signature, used as placeholder in filecoin.
func CreateZeroSignature() Signature {
	sign := bls.CreateZeroSign()
	var sig Signature
	copy(sig[:], sign[:])

	return sig
}
