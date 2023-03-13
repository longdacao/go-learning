package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// 编码一个整数到 Scale 格式
func encodeInt(val int64) []byte {
	// 对于 1 ~ 4 字节的整数，直接返回对应字节数组
	if val >= -0x10 && val <= 0x7f {
		return []byte{byte(val)}
	} else if val >= -0x8000 && val <= 0x7fff {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(val))
		return append([]byte{0x80}, buf...)
	} else if val >= -0x80000000 && val <= 0x7fffffff {
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(val))
		return append([]byte{0x81}, buf...)
	} else {
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, uint64(val))
		return append([]byte{0x82}, buf...)
	}
}

// 解码一个 Scale 格式的整数
func decodeInt(data []byte) (int64, []byte, error) {
	if len(data) < 1 {
		return 0, data, errors.New("Invalid data length")
	}
	tag := data[0]
	switch tag {
	case 0x00:
		return 0, data[1:], nil
	case 0x01:
		if len(data) < 2 {
			return 0, data, errors.New("Invalid data length")
		}
		return int64(int8(data[1])), data[2:], nil
	case 0x02:
		if len(data) < 3 {
			return 0, data, errors.New("Invalid data length")
		}
		return int64(int16(binary.LittleEndian.Uint16(data[1:]))), data[3:], nil
	case 0x03:
		if len(data) < 5 {
			return 0, data, errors.New("Invalid data length")
		}
		return int64(int32(binary.LittleEndian.Uint32(data[1:]))), data[5:], nil
	case 0x04:
		if len(data) < 9 {
			return 0, data, errors.New("Invalid data length")
		}
		return int64(binary.LittleEndian.Uint64(data[1:])), data[9:], nil
	default:
		return 0, data, fmt.Errorf("Unsupported integer encoding tag: %d", tag)
	}
}

func main() {
	data := encodeInt(12345)
	fmt.Printf("Encoded data: %x\n", data)

	val, dataLeft, err := decodeInt(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded value: %d, remaining data: %x\n", val, dataLeft)
}
