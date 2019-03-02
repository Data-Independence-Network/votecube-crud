package location

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollStates(data []byte, cursor *int64, dataLen int64, err error) (models.PollsStateSlice, error) {
	var pollsStates models.PollsStateSlice
	var numPollsStates int64

	if err != nil {
		return pollsStates, err
	}

	numPollsStates, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsStates == 0 {
		return pollsStates, err
	}

	pollsStates = make(models.PollsStateSlice, numPollsStates)

	for i := int64(0); i < numPollsStates; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(data, cursor, dataLen, err)

		if err != nil {
			return pollsStates, err
		}

		pollsStates[i] = &models.PollsState{
			StateID: continentId,
		}
	}

	return pollsStates, err
}
