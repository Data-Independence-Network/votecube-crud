package model

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeFactorPosition(ctx *deserialize.CreatePollDeserializeContext, err error) (models.FactorPosition, int64, error) {
	var factorPosition models.FactorPosition

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var factorPositionId int64
		factorPositionId, err = deserialize.RNum(ctx, err)

		_, factorPositionAlreadySpecifiedInRequest := ctx.IdRefs.FactorPositionIdRefs[factorPositionId][ctx.Request.Index]
		if factorPositionAlreadySpecifiedInRequest {
			return factorPosition, 0, fmt.Errorf("multiple referenes to a FactorPosition in same Create Poll CreatePollRequest")
		}

		ctx.IdRefs.FactorPositionIdRefs[factorPositionId][ctx.Request.Index] = ctx.Request

		return factorPosition, factorPositionId, nil
	}

	var factor models.Factor
	var dimId int64
	factor, dimId, err = DeserializeFactor(ctx, err)

	var position models.Position
	var dirId int64
	position, dirId, err = DeserializePosition(ctx, err)

	if err != nil {
		return factorPosition, 0, err
	}

	factorPosition = models.FactorPosition{
		UserAccountID: ctx.Request.UserAccountId,
	}

	if dimId != 0 {
		factorPosition.FactorID = dimId
	} else {
		factorPosition.R.Factor = &factor
	}

	if dirId != 0 {
		factorPosition.PositionID = dirId
	} else {
		factorPosition.R.Position = &position
	}

	return factorPosition, 0, nil
}
