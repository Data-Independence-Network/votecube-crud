package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimension(data []byte, cursor *int64, dataLen int64, err error) (models.Dimension, bool, error) {
	var dimension models.Dimension

	if err != nil {
		return dimension, false, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(data, cursor, dataLen, err)

	if objectType == deserialize.REFERENCE {
		var dimensionId int64
		dimensionId, err = deserialize.RNum(data, cursor, dataLen, err)
	} else {

	}

}
