// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package casbin

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UpgradeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpgradeReq[number], err)
}

func (x *UpgradeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpgradeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Certificate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Empty) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *OkReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OkReq[number], err)
}

func (x *OkReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpgradeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpgradeReq) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *UpgradeReq) fastWriteField2(buf []byte) (offset int) {
	if x.Certificate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCertificate())
	return offset
}

func (x *Empty) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *OkReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *OkReq) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *UpgradeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpgradeReq) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

func (x *UpgradeReq) sizeField2() (n int) {
	if x.Certificate == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCertificate())
	return n
}

func (x *Empty) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *OkReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *OkReq) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

var fieldIDToName_UpgradeReq = map[int32]string{
	1: "Email",
	2: "Certificate",
}

var fieldIDToName_Empty = map[int32]string{}

var fieldIDToName_OkReq = map[int32]string{
	1: "Email",
}
