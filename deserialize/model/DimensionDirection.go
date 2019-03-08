package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimDir(ctx *deserialize.DeserializeContext, err error) (models.DimensionDirection, int64, error) {
	var dimensionDirection models.DimensionDirection

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	var dimDirId int64
	dimDirId, err = deserialize.RNum(ctx, err)

	if objectType == deserialize.REFERENCE {
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
		UserAccountID: ctx.UserAccountId,
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
