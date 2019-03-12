package poll

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeLabel(pollLabel *models.PollsLabel, ctx *deserialize.CreatePollDeserializeContext, err error) error {
	if err != nil {
		return err
	}

	var objectType byte
	objectType, err = deserialize.RByte(ctx, err)

	if objectType == deserialize.REFERENCE {
		var labelId int64
		labelId, err = deserialize.RNum(ctx, err)

		_, labelAlreadySpecifiedInRequest := ctx.IdRefs.LabelIdRefs[labelId][ctx.Request.Index]
		if labelAlreadySpecifiedInRequest {
			return fmt.Errorf("multiple references to a Label in same Create Poll CreatePollRequest")
		}

		pollLabel.LabelID = labelId

		ctx.IdRefs.LabelIdRefs[labelId][ctx.Request.Index] = ctx.Request

		return nil
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(ctx, err)

		if err != nil {
			return err
		}

		nameLength := len(name)

		if nameLength < 3 {
			return fmt.Errorf("the Label Name is less than 3 characters long")
		} else if nameLength > 40 {
			return fmt.Errorf("the Label Name is greater than 40 characters long")
		}

		_, requestAlreadyHasLabel := ctx.RequestNewLabelMapByName[name]

		if requestAlreadyHasLabel {
			return fmt.Errorf("the Label Name is specified more than once in the Request")
		}
		ctx.RequestNewLabelMapByName[name] = pollLabel

		ctxForName, pollLabelsForNameExist := ctx.CtxMapByLabelName[name]

		if pollLabelsForNameExist {
			pollLabel.R.Label = ctxForName[0].RequestNewLabelMapByName[name].R.Label
		} else {
			pollLabel.R.Label = &models.Label{
				Name:          name,
				UserAccountID: ctx.Request.UserAccountId,
			}
			ctxForName = make([]*deserialize.CreatePollDeserializeContext, 1)
			ctxForName[0] = ctx
			ctx.CtxMapByLabelName[name] = ctxForName
		}

		return nil
	}

}
