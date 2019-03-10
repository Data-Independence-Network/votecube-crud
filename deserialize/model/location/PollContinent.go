package location

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollContinents(ctx *deserialize.CreatePollDeserializeContext, err error) (models.PollsContinentSlice, error) {
	var pollsContinents models.PollsContinentSlice
	var numPollsContinents int64

	if err != nil {
		return pollsContinents, fmt.Errorf("no continents specified for the poll")
	}

	numPollsContinents, err = deserialize.RNum(ctx, err)

	if numPollsContinents == 0 {
		return pollsContinents, err
	}

	pollsContinents = make(models.PollsContinentSlice, numPollsContinents)

	for i := int64(0); i < numPollsContinents; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(ctx, err)

		if err != nil {
			return pollsContinents, err
		}

		_, continentExists := ctx.LocMaps.ContinentMap[continentId]
		if !continentExists {
			return pollsContinents, fmt.Errorf("invalid continent id: %v", continentId)
		}

		_, continentAlreadySpecifiedInRequest := ctx.ReqLocSets.ContinentSet[continentId]
		if continentAlreadySpecifiedInRequest {
			return pollsContinents, fmt.Errorf("continent specified more than once, id: %v", continentId)
		}
		ctx.ReqLocSets.ContinentSet[continentId] = true

		pollsContinents[i] = &models.PollsContinent{
			ContinentID: continentId,
		}
	}

	return pollsContinents, err
}
