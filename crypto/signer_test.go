package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/bls"
	"github.com/ethereum/go-ethereum/crypto/bls/blst"
	"github.com/ethereum/go-ethereum/crypto/bls/common"
	"github.com/stretchr/testify/assert"
)


func TestSigner(t *testing.T) {

	data := []byte("data to be signed")
	hash := Keccak256Hash(data)

	key, err := ecdsa.GenerateKey(S256(), rand.Reader)
	assert.NoError(t, err, "failed to generate key")
	sig, err := Sign(hash.Bytes(), key)
	assert.NoError(t, err, "failed to sign data")
	t.Logf("ecdsa pub key %x size %d", FromECDSAPub(&key.PublicKey), len(FromECDSAPub(&key.PublicKey)) )
	t.Logf("ecdsa sig %x, size %d", sig, len(sig))


	priv, err := bls.RandKey()
	assert.NoError(t, err)
	blsSig := priv.Sign(hash.Bytes())
	t.Logf("bls pub key %x size %d", priv.PublicKey().Marshal(), len(priv.PublicKey().Marshal()) )
	t.Logf("bls sig %x, size %d", blsSig.Marshal(), len(blsSig.Marshal()))
	assert.True(t, blst.VerifyCompressed(blsSig.Marshal(), priv.PublicKey().Marshal(), hash.Bytes()))


	priv2, err := bls.RandKey()
	assert.NoError(t, err)
	blsSig2 := priv2.Sign(hash.Bytes())
	assert.True(t, blst.VerifyCompressed(blsSig2.Marshal(), priv2.PublicKey().Marshal(), hash.Bytes()))

	
	aggPub := bls.AggregateMultiplePubkeys([]common.PublicKey{priv.PublicKey(), priv2.PublicKey()})
	aggSig := bls.AggregateSignatures([]common.Signature{blsSig, blsSig2})
	assert.True(t, blst.VerifyCompressed(aggSig.Marshal(), aggPub.Marshal(), hash.Bytes()))

	t.Logf("agg pub key %x size %d", aggPub.Marshal(), len(aggPub.Marshal()) )
	t.Logf("agg sig %x, size %d", aggSig.Marshal(), len(aggSig.Marshal()))
}