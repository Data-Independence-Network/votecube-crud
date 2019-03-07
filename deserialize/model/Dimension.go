package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimension(data []byte, cursor *int64, dataLen int64, err error) (models.Dimension, int64, error) {
	var dimension models.Dimension

	if err != nil {
		return dimension, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(data, cursor, dataLen, err)

	var dimensionId int64
	dimensionId, err = deserialize.RNum(data, cursor, dataLen, err)

	if objectType == deserialize.REFERENCE {
		return dimension, dimensionId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(data, cursor, dataLen, err)

		var parentDimensionId int64
		var parentDimIdReference null.Int64
		objectType, err = deserialize.RByte(data, cursor, dataLen, err)
		if objectType == deserialize.REFERENCE {
			parentDimensionId, err = deserialize.RNum(data, cursor, dataLen, err)
			parentDimIdReference = null.Int64{
				Int64: parentDimensionId,
				Valid: true,
			}
		}

		if err != nil {
			return dimension, 0, err
		}

		dimension = models.Dimension{
			DimensionName:     name,
			ParentDimensionID: parentDimIdReference,
		}
	}

}
