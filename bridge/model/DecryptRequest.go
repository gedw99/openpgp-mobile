// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DecryptRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsDecryptRequest(buf []byte, offset flatbuffers.UOffsetT) *DecryptRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DecryptRequest{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DecryptRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DecryptRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DecryptRequest) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DecryptRequest) PrivateKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DecryptRequest) Passphrase() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DecryptRequest) Options(obj *KeyOptions) *KeyOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KeyOptions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DecryptRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DecryptRequestAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(message), 0)
}
func DecryptRequestAddPrivateKey(builder *flatbuffers.Builder, privateKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(privateKey), 0)
}
func DecryptRequestAddPassphrase(builder *flatbuffers.Builder, passphrase flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(passphrase), 0)
}
func DecryptRequestAddOptions(builder *flatbuffers.Builder, options flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(options), 0)
}
func DecryptRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}