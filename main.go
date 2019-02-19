package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"votecube-crud/models"
	"votecube-crud/sequence"

	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

var (
	addr        = flag.String("addr", ":8082", "TCP address to listen to")
	compress    = flag.Bool("compress", false, "Whether to enable transparent response compression")
	db          *sql.DB
	directionId sequence.Sequence
)

func SetupDb() *sql.DB {
	db, err := sql.Open("postgres", `postgresql://root@localhost:26257/votecube?sslmode=disable`)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	directionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           db,
		IncrementBy:  10,
		Max:          0,
		Name:         "direction_id",
	}

	return db
}

func main() {
	flag.Parse()

	db = SetupDb()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	if ctx.IsPut() {
		if bytes.Equal(ctx.RequestURI(), []byte("/direction")) {
			//var json = jsoniter.ConfigCompatibleWithStandardLibrary

			var data models.Direction

			//json.Unmarshal(ctx.PostBody(), &data)

			rawData := ctx.PostBody()

			cursorPosition := 2

			var cursor *int = &cursorPosition

			name, err := readString(rawData, cursor)

			if err != nil {
				fmt.Fprintf(ctx, err.Error())
			}

			data.DirectionDescription = name

			//seqBlocks, err := directionId.GetBlocks(9)
			//
			//if err != nil {
			//	fmt.Fprintf(ctx, "Error")
			//}
			//
			//for _, seqBlock := range seqBlocks {
			//	fmt.Fprintf(ctx, "Block, Start: %v, Length: %v\n", seqBlock.Start, seqBlock.Length)
			//}
		}
	} else {
		fmt.Fprintf(ctx, "Hello, world!\n\n")

		fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())

		fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
		fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
		fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
		fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
		fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
		fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
		fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
		fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

		fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
	}

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "PUT")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}

func readString(data []byte, cursor *int) (string, error) {
	dataLen := len(data)

	if *cursor+1 >= dataLen {
		return "", fmt.Errorf("Out of range data access")
	}

	lengthNumBytes := int(data[*cursor])
	*cursor++

	maxLengthNumBytes := *cursor + lengthNumBytes
	length := int(data[*cursor])
	*cursor++

	if *cursor+maxLengthNumBytes >= dataLen {
		return "", fmt.Errorf("Out of range data access")
	}

	for i := *cursor; i < maxLengthNumBytes; i++ {
		nextByte := int(data[*cursor+i])
		shift := uint(8 * i)
		length = nextByte<<shift + length
		*cursor++
	}

	nextCursor := *cursor + length

	if nextCursor > dataLen {
		return "", fmt.Errorf("Out of range data access")
	}

	theString := string(data[*cursor:nextCursor])

	*cursor = nextCursor

	return theString, nil
}
