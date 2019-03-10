package poll

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeLabel(ctx *deserialize.CreatePollDeserializeContext, err error) (models.Label, int64, error) {
	var label models.Label

	if err != nil {
		return label, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var labelId int64
		labelId, err = deserialize.RNum(ctx, err)

		_, labelAlreadySpecifiedInRequest := ctx.IdRefs.LabelIdRefs[labelId][ctx.Request.Index]
		if labelAlreadySpecifiedInRequest {
			return label, 0, fmt.Errorf("multiple referenes to a Label in same Create Poll CreatePollRequest")
		}

		ctx.IdRefs.LabelIdRefs[labelId][ctx.Request.Index] = ctx.Request

		return label, labelId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		return models.Label{
			Name:          name,
			UserAccountID: ctx.Request.UserAccountId,
		}, 0, err
	}

}
