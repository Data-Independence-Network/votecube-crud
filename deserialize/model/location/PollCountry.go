package location

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollCountries(data []byte, cursor *int64, dataLen int64, err error) (models.PollsCountrySlice, error) {
	var pollsCountries models.PollsCountrySlice
	var numPollsCountries int64

	if err != nil {
		return pollsCountries, err
	}

	numPollsCountries, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsCountries == 0 {
		return pollsCountries, err
	}

	pollsCountries = make(models.PollsCountrySlice, numPollsCountries)

	for i := int64(0); i < numPollsCountries; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(data, cursor, dataLen, err)

		if err != nil {
			return pollsCountries, err
		}

		pollsCountries[i] = &models.PollsCountry{
			CountryID: continentId,
		}
	}

	return pollsCountries, err
}
