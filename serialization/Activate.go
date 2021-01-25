// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serialization

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Activate struct {
	_tab flatbuffers.Table
}

func GetRootAsActivate(buf []byte, offset flatbuffers.UOffsetT) *Activate {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Activate{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Activate) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Activate) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Activate) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Activate) MutateTimestamp(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Activate) ProductId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Activate) OrderId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Activate) UserId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Activate) ProfileId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Activate) Size() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Activate) MutateSize(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

func (rcv *Activate) Price() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Activate) MutatePrice(n float64) bool {
	return rcv._tab.MutateFloat64Slot(16, n)
}

func (rcv *Activate) Side() Side {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return Side(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Activate) MutateSide(n Side) bool {
	return rcv._tab.MutateInt8Slot(18, int8(n))
}

func (rcv *Activate) Funds() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Activate) MutateFunds(n float64) bool {
	return rcv._tab.MutateFloat64Slot(20, n)
}

func (rcv *Activate) StopType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Activate) StopPrice() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Activate) MutateStopPrice(n float64) bool {
	return rcv._tab.MutateFloat64Slot(24, n)
}

func (rcv *Activate) Private() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Activate) MutatePrivate(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func ActivateStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func ActivateAddTimestamp(builder *flatbuffers.Builder, timestamp int64) {
	builder.PrependInt64Slot(0, timestamp, 0)
}
func ActivateAddProductId(builder *flatbuffers.Builder, productId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(productId), 0)
}
func ActivateAddOrderId(builder *flatbuffers.Builder, orderId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(orderId), 0)
}
func ActivateAddUserId(builder *flatbuffers.Builder, userId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(userId), 0)
}
func ActivateAddProfileId(builder *flatbuffers.Builder, profileId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(profileId), 0)
}
func ActivateAddSize(builder *flatbuffers.Builder, size float64) {
	builder.PrependFloat64Slot(5, size, 0.0)
}
func ActivateAddPrice(builder *flatbuffers.Builder, price float64) {
	builder.PrependFloat64Slot(6, price, 0.0)
}
func ActivateAddSide(builder *flatbuffers.Builder, side Side) {
	builder.PrependInt8Slot(7, int8(side), 0)
}
func ActivateAddFunds(builder *flatbuffers.Builder, funds float64) {
	builder.PrependFloat64Slot(8, funds, 0.0)
}
func ActivateAddStopType(builder *flatbuffers.Builder, stopType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(stopType), 0)
}
func ActivateAddStopPrice(builder *flatbuffers.Builder, stopPrice float64) {
	builder.PrependFloat64Slot(10, stopPrice, 0.0)
}
func ActivateAddPrivate(builder *flatbuffers.Builder, private bool) {
	builder.PrependBoolSlot(11, private, false)
}
func ActivateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
