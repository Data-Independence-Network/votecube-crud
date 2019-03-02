package location

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollContinents(data []byte, cursor *int64, dataLen int64, err error) (models.PollsContinentSlice, error) {
	var pollsContinents models.PollsContinentSlice
	var numPollsContinents int64

	if err != nil {
		return pollsContinents, err
	}

	numPollsContinents, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsContinents == 0 {
		return pollsContinents, err
	}

	pollsContinents = make(models.PollsContinentSlice, numPollsContinents)

	for i := int64(0); i < numPollsContinents; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(data, cursor, dataLen, err)

		if err != nil {
			return pollsContinents, err
		}

		pollsContinents[i] = &models.PollsContinent{
			ContinentID: continentId,
		}
	}

	return pollsContinents, err
}
