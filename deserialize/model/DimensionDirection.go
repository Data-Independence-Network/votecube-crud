package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimDir(data []byte, cursor *int64, dataLen int64, err error) (models.DimensionDirection, int64, error) {
	var dimensionDirection models.DimensionDirection

	var objectType byte
	objectType, err = deserialize.RByte(data, cursor, dataLen, err)

	var dimDirId int64
	dimDirId, err = deserialize.RNum(data, cursor, dataLen, err)

	if objectType == deserialize.REFERENCE {
		return dimensionDirection, dimDirId, nil
	}

	var dimension models.Dimension
	var dimIsRef bool
	dimension, dimIsRef, err = DeserializeDimension(data, cursor, dataLen, err)

	var direction models.Direction
	var dirIsRef bool
	direction, dirIsRef, err = DeserializeDirection(data, cursor, dataLen, err)

	if err != nil {
		return dimensionDirection, 0, err
	}

	dimensionDirection = models.DimensionDirection{}

	if dimIsRef {
		dimensionDirection.DimensionID = dimension.DimensionID
	} else {
		dimensionDirection.R.Dimension = &dimension
	}

	if dirIsRef {
		dimensionDirection.DirectionID = direction.DirectionID
	} else {
		dimensionDirection.R.Direction = &direction
	}

	return dimensionDirection, 0, nil

}
