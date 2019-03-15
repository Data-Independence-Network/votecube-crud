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

func DeserializePosition(ctx *deserialize.CreatePollDeserializeContext, err error) (models.Position, int64, error) {
	var position models.Position

	if err != nil {
		return position, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var positionId int64
		positionId, err = deserialize.RNum(ctx, err)

		_, positionAlreadySpecifiedInRequest := ctx.IdRefs.FactorPositionIdRefs[positionId][ctx.Request.Index]
		if positionAlreadySpecifiedInRequest {
			return position, 0, fmt.Errorf("multiple referenes to a Position in same Create Poll CreatePollRequest")
		}

		ctx.IdRefs.FactorPositionIdRefs[positionId][ctx.Request.Index] = ctx.Request

		return position, positionId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		var parentPositionId int64
		var parentDirIdReference null.Int64
		objectType, err = deserialize.RByte(ctx, err)
		if objectType == deserialize.REFERENCE {
			parentPositionId, err = deserialize.RNum(ctx, err)
			parentDirIdReference = null.Int64{
				Int64: parentPositionId,
				Valid: true,
			}
		}

		if err != nil {
			return position, 0, err
		}

		return models.Position{
			PositionDescription: name,
			ParentPositionID:    parentDirIdReference,
			UserAccountID:       ctx.Request.UserAccountId,
		}, 0, nil
	}
}
