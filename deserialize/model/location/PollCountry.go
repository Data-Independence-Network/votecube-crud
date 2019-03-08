package location

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollCountries(ctx *deserialize.DeserializeContext, err error) (models.PollsCountrySlice, error) {
	var pollsCountries models.PollsCountrySlice
	var numPollsCountries int64

	if err != nil {
		return pollsCountries, err
	}

	numPollsCountries, err = deserialize.RNum(ctx, err)

	if numPollsCountries == 0 {
		return pollsCountries, err
	}

	pollsCountries = make(models.PollsCountrySlice, numPollsCountries)

	for i := int64(0); i < numPollsCountries; i++ {
		var countryId int64
		countryId, err = deserialize.RNum(ctx, err)

		if err != nil {
			return pollsCountries, err
		}

		_, countryExists := ctx.LocMaps.ContinentMap[countryId]

		if !countryExists {
			return pollsCountries, fmt.Errorf("invalid country id: %v", countryId)
		}

		pollsCountries[i] = &models.PollsCountry{
			CountryID: countryId,
		}
	}

	return pollsCountries, err
}
