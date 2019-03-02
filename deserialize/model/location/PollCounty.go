package location

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollCounties(data []byte, cursor *int64, dataLen int64, err error) (models.PollsCountySlice, error) {
	var pollsCounties models.PollsCountySlice
	var numPollsCounties int64

	if err != nil {
		return pollsCounties, err
	}

	numPollsCounties, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsCounties == 0 {
		return pollsCounties, err
	}

	pollsCounties = make(models.PollsCountySlice, numPollsCounties)

	for i := int64(0); i < numPollsCounties; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(data, cursor, dataLen, err)

		if err != nil {
			return pollsCounties, err
		}

		pollsCounties[i] = &models.PollsCounty{
			CountyID: continentId,
		}
	}

	return pollsCounties, err
}
