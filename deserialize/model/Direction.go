package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDirection(ctx *deserialize.DeserializeContext, err error) (models.Direction, int64, error) {
	var direction models.Direction

	if err != nil {
		return direction, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var directionId int64
		directionId, err = deserialize.RNum(ctx, err)

		return direction, directionId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		var parentDirectionId int64
		var parentDirIdReference null.Int64
		objectType, err = deserialize.RByte(ctx, err)
		if objectType == deserialize.REFERENCE {
			parentDirectionId, err = deserialize.RNum(ctx, err)
			parentDirIdReference = null.Int64{
				Int64: parentDirectionId,
				Valid: true,
			}
		}

		if err != nil {
			return direction, 0, err
		}

		return models.Direction{
			DirectionDescription: name,
			ParentDirectionID:    parentDirIdReference,
			UserAccountID:        ctx.UserAccountId,
		}, 0, nil
	}
}
