// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package eino

import (
	fmt "fmt"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *SearchOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchOrderReq[number], err)
}

func (x *SearchOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Question, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchOrderResp[number], err)
}

func (x *SearchOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v order.Order
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Order = append(x.Order, &v)
	return offset, nil
}

func (x *PlaceOrderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PlaceOrderReq[number], err)
}

func (x *PlaceOrderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PlaceOrderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Question, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PlaceOrderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PlaceOrderResp[number], err)
}

func (x *PlaceOrderResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Resp, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SearchOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.Question == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetQuestion())
	return offset
}

func (x *SearchOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SearchOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.Order == nil {
		return offset
	}
	for i := range x.GetOrder() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetOrder()[i])
	}
	return offset
}

func (x *PlaceOrderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PlaceOrderReq) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *PlaceOrderReq) fastWriteField2(buf []byte) (offset int) {
	if x.Question == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetQuestion())
	return offset
}

func (x *PlaceOrderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *PlaceOrderResp) fastWriteField1(buf []byte) (offset int) {
	if x.Resp == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetResp())
	return offset
}

func (x *SearchOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SearchOrderReq) sizeField1() (n int) {
	if x.Question == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetQuestion())
	return n
}

func (x *SearchOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SearchOrderResp) sizeField1() (n int) {
	if x.Order == nil {
		return n
	}
	for i := range x.GetOrder() {
		n += fastpb.SizeMessage(1, x.GetOrder()[i])
	}
	return n
}

func (x *PlaceOrderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PlaceOrderReq) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetUid())
	return n
}

func (x *PlaceOrderReq) sizeField2() (n int) {
	if x.Question == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetQuestion())
	return n
}

func (x *PlaceOrderResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *PlaceOrderResp) sizeField1() (n int) {
	if x.Resp == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetResp())
	return n
}

var fieldIDToName_SearchOrderReq = map[int32]string{
	1: "Question",
}

var fieldIDToName_SearchOrderResp = map[int32]string{
	1: "Order",
}

var fieldIDToName_PlaceOrderReq = map[int32]string{
	1: "Uid",
	2: "Question",
}

var fieldIDToName_PlaceOrderResp = map[int32]string{
	1: "Resp",
}

var _ = order.File_order_proto
