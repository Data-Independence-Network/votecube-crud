package deserialize

import (
	"fmt"
	"time"
)

var (
	NIL       byte = 0
	RECORD    byte = 1
	REFERENCE byte = 2
)

func RStr(data []byte, cursor *int64, dataLen int64, err error) (string, error) {
	length, err := RNum(data, cursor, dataLen, err)

	if err != nil {
		return "", err
	}

	nextCursor := *cursor + length

	if nextCursor > dataLen {
		return "", fmt.Errorf("Out of range data access")
	}

	theString := string(data[*cursor:nextCursor])

	*cursor = nextCursor

	return theString, nil
}

func RNum(data []byte, cursor *int64, dataLen int64, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	if *cursor+1 >= dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	lengthNumBytes := int64(data[*cursor])
	*cursor++

	maxLengthNumBytes := *cursor + lengthNumBytes
	num := int64(data[*cursor])
	*cursor++

	if *cursor+maxLengthNumBytes >= dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	for i := *cursor; i < maxLengthNumBytes; i++ {
		nextByte := int64(data[*cursor+i])
		shift := uint(8 * i)
		num = nextByte<<shift + num
		*cursor++
	}

	return num, nil
}

func RByte(data []byte, cursor *int64, dataLen int64, err error) (byte, error) {
	if err != nil {
		return 0, err
	}

	if *cursor+1 >= dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	aByte := data[*cursor]
	*cursor++

	return aByte, nil
}

func RTime(data []byte, cursor *int64, dataLen int64, err error) (time.Time, error) {
	millis, err := RNum(data, cursor, dataLen, err)

	if err != nil {
		return time.Now(), err
	}

	return time.Unix(0, millis*int64(time.Millisecond)), nil
}
