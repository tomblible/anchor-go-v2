package token_interface

import (
	"encoding/binary"
)

func GetTransferFeeExtension(data []byte) float64 {
	index := 166
	for index < len(data) {
		if len(data) < index+2 {
			break
		}
		if binary.LittleEndian.Uint16(data[index:index+2]) == 1 {
			if len(data) >= index+94 {
				fee := binary.LittleEndian.Uint16(data[index+92 : index+94])
				return float64(fee) / 10000
			}
			break
		}
		if len(data) < index+4 {
			break
		}
		offset := binary.LittleEndian.Uint16(data[index+2 : index+4])
		index += 4 + int(offset)
	}
	return 0
}
