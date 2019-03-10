package location

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollTowns(ctx *deserialize.DeserializeContext, err error) (models.PollsTownSlice, error) {
	var pollsTowns models.PollsTownSlice
	var numPollsTowns int64

	if err != nil {
		return pollsTowns, err
	}

	numPollsTowns, err = deserialize.RNum(ctx, err)

	if numPollsTowns == 0 {
		return pollsTowns, err
	}

	pollsTowns = make(models.PollsTownSlice, numPollsTowns)

	for i := int64(0); i < numPollsTowns; i++ {
		var townId int64
		townId, err = deserialize.RNum(ctx, err)

		if err != nil {
			return pollsTowns, err
		}

		_, townExists := ctx.LocMaps.ContinentMap[townId]

		if !townExists {
			return pollsTowns, fmt.Errorf("invalid town id: %v", townId)
		}

		_, townAlreadySpecifiedInRequest := ctx.ReqLocSets.TownSet[townId]
		if townAlreadySpecifiedInRequest {
			return pollsTowns, fmt.Errorf("town specified more than once, id: %v", townId)
		}
		ctx.ReqLocSets.TownSet[townId] = true

		town := ctx.LocMaps.TownMap[townId]
		_, stateSpecifiedInRequest := ctx.ReqLocSets.StateSet[town.StateID]
		if !stateSpecifiedInRequest {
			return pollsTowns, fmt.Errorf("no matching State specified in request for Town Id %v", townId)
		}

		pollsTowns[i] = &models.PollsTown{
			TownID: townId,
		}
	}

	return pollsTowns, err
}
