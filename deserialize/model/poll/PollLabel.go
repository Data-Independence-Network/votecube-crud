package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePollLabels(ctx *deserialize.CreatePollDeserializeContext, err error) (models.PollsLabelSlice, error) {
	var pollsLabels models.PollsLabelSlice
	var numPollsLabels int64

	if err != nil {
		return pollsLabels, err
	}

	numPollsLabels, err = deserialize.RNum(ctx, err)

	if numPollsLabels == 0 {
		return pollsLabels, err
	}

	pollsLabels = make(models.PollsLabelSlice, numPollsLabels)

	for i := int64(0); i < numPollsLabels; i++ {
		pollLabel := models.PollsLabel{
			UserAccountID: ctx.Request.UserAccountId,
		}

		err = DeserializeLabel(&pollLabel, ctx, err)

		if err != nil {
			return pollsLabels, err
		}

		pollLabel.UserAccountID = ctx.Request.UserAccountId
		pollsLabels[i] = &pollLabel
	}

	return pollsLabels, err
}
