package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDimension(ctx *deserialize.DeserializeContext, err error) (models.Dimension, int64, error) {
	var dimension models.Dimension

	if err != nil {
		return dimension, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	var dimensionId int64
	dimensionId, err = deserialize.RNum(ctx, err)

	if objectType == deserialize.REFERENCE {
		return dimension, dimensionId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		var parentDimensionId int64
		var parentDimIdReference null.Int64
		objectType, err = deserialize.RByte(ctx, err)
		if objectType == deserialize.REFERENCE {
			parentDimensionId, err = deserialize.RNum(ctx, err)
			parentDimIdReference = null.Int64{
				Int64: parentDimensionId,
				Valid: true,
			}
		}

		if err != nil {
			return dimension, 0, err
		}

		return models.Dimension{
			DimensionName:     name,
			ParentDimensionID: parentDimIdReference,
			UserAccountID:     ctx.UserAccountId,
		}, 0, nil
	}

}
