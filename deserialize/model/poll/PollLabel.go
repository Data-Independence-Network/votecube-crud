package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePollLabels(data []byte, cursor *int64, dataLen int64, err error) (models.PollsLabelSlice, error) {
	var pollsLabels models.PollsLabelSlice
	var numPollsLabels int64

	if err != nil {
		return pollsLabels, err
	}

	numPollsLabels, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsLabels == 0 {
		return pollsLabels, err
	}

	pollsLabels = make(models.PollsLabelSlice, numPollsLabels)

	for i := int64(0); i < numPollsLabels; i++ {
		var objectType byte
		objectType, err = deserialize.RByte(data, cursor, dataLen, err)

		var dimDirId int64
		dimDirId, err = deserialize.RNum(data, cursor, dataLen, err)

		if objectType == deserialize.REFERENCE {
			var labelId int64
			labelId, err = deserialize.RNum(data, cursor, dataLen, err)

			if err != nil {
				return pollsLabels, err
			}

			pollsLabels[i] = &models.PollsLabel{
				LabelID: labelId,
			}
		} else {

		}

	}

	return pollsLabels, err
}
