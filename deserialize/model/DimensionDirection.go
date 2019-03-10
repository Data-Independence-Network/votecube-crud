package model

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimDir(ctx *deserialize.CreatePollDeserializeContext, err error) (models.DimensionDirection, int64, error) {
	var dimensionDirection models.DimensionDirection

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var dimDirId int64
		dimDirId, err = deserialize.RNum(ctx, err)

		_, dimDirAlreadySpecifiedInRequest := ctx.IdRefs.DimDirIdRefs[dimDirId][ctx.Request.Index]
		if dimDirAlreadySpecifiedInRequest {
			return dimensionDirection, 0, fmt.Errorf("multiple referenes to a DimensionDirection in same Create Poll CreatePollRequest")
		}

		ctx.IdRefs.DimDirIdRefs[dimDirId][ctx.Request.Index] = ctx.Request

		return dimensionDirection, dimDirId, nil
	}

	var dimension models.Dimension
	var dimId int64
	dimension, dimId, err = DeserializeDimension(ctx, err)

	var direction models.Direction
	var dirId int64
	direction, dirId, err = DeserializeDirection(ctx, err)

	if err != nil {
		return dimensionDirection, 0, err
	}

	dimensionDirection = models.DimensionDirection{
		UserAccountID: ctx.Request.UserAccountId,
	}

	if dimId != 0 {
		dimensionDirection.DimensionID = dimId
	} else {
		dimensionDirection.R.Dimension = &dimension
	}

	if dirId != 0 {
		dimensionDirection.DirectionID = dirId
	} else {
		dimensionDirection.R.Direction = &direction
	}

	return dimensionDirection, 0, nil
}
