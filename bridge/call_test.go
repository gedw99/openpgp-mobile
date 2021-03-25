package openPGPBridge

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/jerson/openpgp-mobile/bridge/model"
	"testing"
)

func TestCall(t *testing.T) {

	b := flatbuffers.NewBuilder(0)

	model.KeyOptionsStart(b)
	model.KeyOptionsAddCipher(b, model.CipherAES256)
	model.KeyOptionsAddCompression(b, model.CompressionZLIB)
	model.KeyOptionsAddCompressionLevel(b, 9)
	model.KeyOptionsAddHash(b, model.HashSHA512)
	model.KeyOptionsAddRsaBits(b, 1024)
	keyOptions := model.KeyOptionsEnd(b)

	comment := b.CreateSharedString("this is a sample comment")
	email := b.CreateSharedString("sample@sample.com")
	name := b.CreateSharedString("jhon doe")
	passphrase := b.CreateSharedString("123465")

	model.OptionsStart(b)
	model.OptionsAddComment(b, comment)
	model.OptionsAddEmail(b, email)
	model.OptionsAddName(b, name)
	model.OptionsAddPassphrase(b, passphrase)
	model.OptionsAddKeyOptions(b, keyOptions)
	options := model.OptionsEnd(b)

	model.GenerateRequestStart(b)
	model.GenerateRequestAddOptions(b, options)
	b.Finish(model.GenerateRequestEnd(b))

	data, err := Call("generate", b.Bytes[b.Head():])
	if err != nil {
		t.Fatal(err)
		return
	}
	response := model.GetRootAsKeyPairResponse(data, 0)
	keyPairOutput := response.Output(nil)
	t.Log(string(keyPairOutput.PrivateKey()))
	t.Log(string(keyPairOutput.PublicKey()))
}
