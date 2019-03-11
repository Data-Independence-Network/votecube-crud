package serialize

func WNum(naturalNumber int64, data *[]byte) {
	// we want to represent the input as a 8-bytes array

	var numBytes []byte

	for naturalNumber > 0 {
		byteVal := naturalNumber & 0xff

		numBytes = append(numBytes, byte(byteVal))
		naturalNumber = (naturalNumber - byteVal) / 256
	}

	localData := append(*data, byte(len(numBytes)))

	for _, aByte := range numBytes {
		localData = append(*data, aByte)
	}

	data = &localData
}
