package location

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollCounties(ctx *deserialize.DeserializeContext, err error) (models.PollsCountySlice, error) {
	var pollsCounties models.PollsCountySlice
	var numPollsCounties int64

	if err != nil {
		return pollsCounties, err
	}

	numPollsCounties, err = deserialize.RNum(ctx, err)

	if numPollsCounties == 0 {
		return pollsCounties, err
	}

	pollsCounties = make(models.PollsCountySlice, numPollsCounties)

	for i := int64(0); i < numPollsCounties; i++ {
		var countyId int64
		countyId, err = deserialize.RNum(ctx, err)

		if err != nil {
			return pollsCounties, err
		}

		_, countyExists := ctx.LocMaps.ContinentMap[countyId]

		if !countyExists {
			return pollsCounties, fmt.Errorf("invalid county id: %v", countyId)
		}

		pollsCounties[i] = &models.PollsCounty{
			CountyID: countyId,
		}
	}

	return pollsCounties, err
}
