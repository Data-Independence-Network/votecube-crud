package deserialize

import (
	"fmt"
	"github.com/diapco/votecube-crud/models"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	NIL       byte = 0
	RECORD    byte = 1
	REFERENCE byte = 2
)

type LocationMaps struct {
	ContinentMap map[int64]models.Continent
	CountryMap   map[int64]models.Country
	CountyMap    map[int64]models.County
	StateMap     map[int64]models.State
	TownMap      map[int64]models.Town
}

type LocationSets struct {
	ContinentSet map[int64]bool
	CountrySet   map[int64]bool
	CountySet    map[int64]bool
	StateSet     map[int64]bool
	TownSet      map[int64]bool
}

type RequestInput struct {
	Cursor  *int64
	Data    *[]byte
	DataLen int64
}

type CreatePollDeserializeContext struct {
	RequestInput
	IdRefs     *CreatePollIdReferences
	LocMaps    *LocationMaps
	ReqLocSets *LocationSets
	Request    *CreatePollRequest
}

type CreatePollRequest struct {
	Ctx           *fasthttp.RequestCtx
	Done          chan bool
	Index         int
	Poll          models.Poll
	UserAccountId int64
}

type CreatePollIdReferences struct {
	DimDirIdRefs map[int64]map[int]*CreatePollRequest
	DimIdRefs    map[int64]map[int]*CreatePollRequest
	DirIdRefs    map[int64]map[int]*CreatePollRequest
	LabelIdRefs  map[int64]map[int]*CreatePollRequest
}

func RStr(ctx *CreatePollDeserializeContext, err error) (string, error) {
	length, err := RNum(ctx, err)

	if err != nil {
		return "", err
	}

	nextCursor := *ctx.Cursor + length

	if nextCursor > ctx.DataLen {
		return "", fmt.Errorf("out of range data access")
	}

	theString := string((*ctx.Data)[*ctx.Cursor:nextCursor])

	*ctx.Cursor = nextCursor

	return theString, nil
}

func RNum(ctx *CreatePollDeserializeContext, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	if *ctx.Cursor+1 >= ctx.DataLen {
		return 0, fmt.Errorf("out of range data access")
	}

	lengthNumBytes := int64((*ctx.Data)[*ctx.Cursor])
	*ctx.Cursor++

	maxLengthNumBytes := *ctx.Cursor + lengthNumBytes
	num := int64((*ctx.Data)[*ctx.Cursor])
	*ctx.Cursor++

	if *ctx.Cursor+maxLengthNumBytes >= ctx.DataLen {
		return 0, fmt.Errorf("out of range data access")
	}

	for i := *ctx.Cursor; i < maxLengthNumBytes; i++ {
		nextByte := int64((*ctx.Data)[*ctx.Cursor+i])
		shift := uint(8 * i)
		num = nextByte<<shift + num
		*ctx.Cursor++
	}

	return num, nil
}

func RByte(ctx *CreatePollDeserializeContext, err error) (byte, error) {
	if err != nil {
		return 0, err
	}

	if *ctx.Cursor+1 >= ctx.DataLen {
		return 0, fmt.Errorf("out of range data access")
	}

	aByte := (*ctx.Data)[*ctx.Cursor]
	*ctx.Cursor++

	return aByte, nil
}

func RTime(ctx *CreatePollDeserializeContext, err error) (time.Time, error) {
	millis, err := RNum(ctx, err)

	if err != nil {
		return time.Now(), err
	}

	return time.Unix(0, millis*int64(time.Millisecond)), nil
}
