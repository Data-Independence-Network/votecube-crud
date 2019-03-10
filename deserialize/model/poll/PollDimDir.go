package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/deserialize/model"
	"github.com/diapco/votecube-crud/models"
)

type PollDimDirRefs struct {
	tempIds map[int64]int64
	dimDirs models.PollsDimensionsDirectionSlice
}

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePollDimDirs(ctx *deserialize.CreatePollDeserializeContext, err error) (models.PollsDimensionsDirectionSlice, error) {
	var pollsDimDirRefs models.PollsDimensionsDirectionSlice
	var numPollsDimDirs int64

	if err != nil {
		return pollsDimDirRefs, err
	}

	numPollsDimDirs, err = deserialize.RNum(ctx, err)

	if numPollsDimDirs == 0 {
		return pollsDimDirRefs, err
	}

	pollsDimDirRefs = make(models.PollsDimensionsDirectionSlice, numPollsDimDirs)

	for i := int64(0); i < numPollsDimDirs; i++ {
		// PollDimDir is always a RECORD, but byte is still sent (for now)
		var _ byte
		_, err = deserialize.RByte(ctx, err)

		if err != nil {
			return pollsDimDirRefs, err
		}

		var pollsDimDir models.PollsDimensionsDirection

		// PollDimDirs are never sent as references and will always be created
		//if objectType == deserialize.REFERENCE {
		//} else {
		var (
			axis     string
			axisByte byte
			colorId  int64
			dir      bool
			dirByte  byte
		)

		axisByte, err = deserialize.RByte(ctx, err)
		if err != nil {
			return pollsDimDirRefs, err
		}

		switch axisByte {
		case 0:
			axis = "x"
		case 1:
			axis = "y"
		case 2:
			axis = "z"
		}

		colorId, err = deserialize.RNum(ctx, err)

		pollsDimDir = models.PollsDimensionsDirection{
			ColorID:                 colorId,
			DimensionCoordinateAxis: axis,
		}

		var (
			dimensionDirection models.DimensionDirection
			dimDirId           int64
		)
		dimensionDirection, dimDirId, err = model.DeserializeDimDir(ctx, err)

		if err != nil {
			return pollsDimDirRefs, err
		}

		if dimDirId != 0 {
			pollsDimDir.DimensionDirectionID = dimDirId
		} else {
			pollsDimDir.R.DimensionDirection = &dimensionDirection
		}

		dirByte, err = deserialize.RByte(ctx, err)
		if err != nil {
			return pollsDimDirRefs, err
		}

		switch dirByte {
		case 0:
			dir = false
		case 1:
			dir = true
		}

		pollsDimDir.DirectionOrientation = dir

		//}

		pollsDimDirRefs[i] = &pollsDimDir

	}

	return pollsDimDirRefs, err
}
