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
	var dimId int64
	dimension, dimId, err = DeserializeDimension(data, cursor, dataLen, err)

	var direction models.Direction
	var dirId int64
	direction, dirId, err = DeserializeDirection(data, cursor, dataLen, err)

	if err != nil {
		return dimensionDirection, 0, err
	}

	dimensionDirection = models.DimensionDirection{}

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
