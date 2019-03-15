package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/deserialize/model"
	"github.com/diapco/votecube-crud/models"
)

type PollFactorPositionRefs struct {
	tempIds         map[int64]int64
	factorPositions models.PollsFactorsPositionSlice
}

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePollFactorPositions(ctx *deserialize.CreatePollDeserializeContext, err error) (models.PollsFactorsPositionSlice, error) {
	var pollsFactorPositionRefs models.PollsFactorsPositionSlice
	var numPollsFactorPositions int64

	if err != nil {
		return pollsFactorPositionRefs, err
	}

	numPollsFactorPositions, err = deserialize.RNum(ctx, err)

	if numPollsFactorPositions == 0 {
		return pollsFactorPositionRefs, err
	}

	pollsFactorPositionRefs = make(models.PollsFactorsPositionSlice, numPollsFactorPositions)

	for i := int64(0); i < numPollsFactorPositions; i++ {
		// PollFactorPosition is always a RECORD, but byte is still sent (for now)
		var _ byte
		_, err = deserialize.RByte(ctx, err)

		if err != nil {
			return pollsFactorPositionRefs, err
		}

		var pollsFactorPosition models.PollsFactorsPosition

		// PollFactorPositions are never sent as references and will always be created
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
			return pollsFactorPositionRefs, err
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

		pollsFactorPosition = models.PollsFactorsPosition{
			ColorID:              colorId,
			FactorCoordinateAxis: axis,
		}

		var (
			factorPosition   models.FactorPosition
			factorPositionId int64
		)
		factorPosition, factorPositionId, err = model.DeserializeFactorPosition(ctx, err)

		if err != nil {
			return pollsFactorPositionRefs, err
		}

		if factorPositionId != 0 {
			pollsFactorPosition.FactorPositionID = factorPositionId
		} else {
			pollsFactorPosition.R.FactorPosition = &factorPosition
		}

		dirByte, err = deserialize.RByte(ctx, err)
		if err != nil {
			return pollsFactorPositionRefs, err
		}

		switch dirByte {
		case 0:
			dir = false
		case 1:
			dir = true
		}

		pollsFactorPosition.PositionOrientation = dir

		//}

		pollsFactorPositionRefs[i] = &pollsFactorPosition

	}

	return pollsFactorPositionRefs, err
}
