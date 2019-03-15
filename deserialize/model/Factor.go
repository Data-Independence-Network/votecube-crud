package model

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeFactor(ctx *deserialize.CreatePollDeserializeContext, err error) (models.Factor, int64, error) {
	var factor models.Factor

	if err != nil {
		return factor, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var factorId int64
		factorId, err = deserialize.RNum(ctx, err)

		_, factorAlreadySpecifiedInRequest := ctx.IdRefs.FactorPositionIdRefs[factorId][ctx.Request.Index]
		if factorAlreadySpecifiedInRequest {
			return factor, 0, fmt.Errorf("multiple referenes to a Factor in same Create Poll CreatePollRequest")
		}

		ctx.IdRefs.FactorIdRefs[factorId][ctx.Request.Index] = ctx.Request

		return factor, factorId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		var parentFactorId int64
		var parentDimIdReference null.Int64
		objectType, err = deserialize.RByte(ctx, err)
		if objectType == deserialize.REFERENCE {
			parentFactorId, err = deserialize.RNum(ctx, err)
			parentDimIdReference = null.Int64{
				Int64: parentFactorId,
				Valid: true,
			}
		}

		if err != nil {
			return factor, 0, err
		}

		return models.Factor{
			FactorName:     name,
			ParentFactorID: parentDimIdReference,
			UserAccountID:  ctx.Request.UserAccountId,
		}, 0, nil
	}

}
