package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePollLabels(ctx *deserialize.DeserializeContext, err error) (models.PollsLabelSlice, error) {
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
		var label models.Label
		var labelId int64

		label, labelId, err = DeserializeLabel(ctx, err)

		if err != nil {
			return pollsLabels, err
		}

		if labelId != 0 {
			pollsLabels[i] = &models.PollsLabel{
				LabelID:       labelId,
				UserAccountID: ctx.Request.UserAccountId,
			}
		} else {
			pollsLabel := &models.PollsLabel{
				UserAccountID: ctx.Request.UserAccountId,
			}

			pollsLabel.R.Label = &label

			pollsLabels[i] = pollsLabel
		}

	}

	return pollsLabels, err
}
