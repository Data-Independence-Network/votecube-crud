package deserialize

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	NIL       byte = 0
	RECORD    byte = 1
	REFERENCE byte = 2
)

type IdsToLookUp struct {
	dimensionIds []int64
	dimDirIds    []int64
	directionIds []int64
	labelIds     []int64
}

type LocationMaps struct {
	ContinentMap map[int64]bool
	CountryMap   map[int64]bool
	StateMap     map[int64]bool
	TownMap      map[int64]bool
}

type DeserializeContext struct {
	cursor        *int64
	data          *[]byte
	dataLen       int64
	IdRefs        *IdReferences
	LocMaps       *LocationMaps
	Lookups       *IdsToLookUp
	UserAccountId int64
}

type Request struct {
	ctx  *fasthttp.RequestCtx
	done chan bool
}

type IdReferences struct {
	dimIdRefs    map[int64][]Request
	dimDirIdRefs map[int64][]Request
	dirIdRefs    map[int64][]Request
	labelIdRefs  map[int64][]Request
}

func RStr(ctx *DeserializeContext, err error) (string, error) {
	length, err := RNum(ctx, err)

	if err != nil {
		return "", err
	}

	nextCursor := *ctx.cursor + length

	if nextCursor > ctx.dataLen {
		return "", fmt.Errorf("out of range data access")
	}

	theString := string((*ctx.data)[*ctx.cursor:nextCursor])

	*ctx.cursor = nextCursor

	return theString, nil
}

func RNum(ctx *DeserializeContext, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	if *ctx.cursor+1 >= ctx.dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	lengthNumBytes := int64((*ctx.data)[*ctx.cursor])
	*ctx.cursor++

	maxLengthNumBytes := *ctx.cursor + lengthNumBytes
	num := int64((*ctx.data)[*ctx.cursor])
	*ctx.cursor++

	if *ctx.cursor+maxLengthNumBytes >= ctx.dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	for i := *ctx.cursor; i < maxLengthNumBytes; i++ {
		nextByte := int64((*ctx.data)[*ctx.cursor+i])
		shift := uint(8 * i)
		num = nextByte<<shift + num
		*ctx.cursor++
	}

	return num, nil
}

func RByte(ctx *DeserializeContext, err error) (byte, error) {
	if err != nil {
		return 0, err
	}

	if *ctx.cursor+1 >= ctx.dataLen {
		return 0, fmt.Errorf("Out of range data access")
	}

	aByte := (*ctx.data)[*ctx.cursor]
	*ctx.cursor++

	return aByte, nil
}

func RTime(ctx *DeserializeContext, err error) (time.Time, error) {
	millis, err := RNum(ctx, err)

	if err != nil {
		return time.Now(), err
	}

	return time.Unix(0, millis*int64(time.Millisecond)), nil
}
