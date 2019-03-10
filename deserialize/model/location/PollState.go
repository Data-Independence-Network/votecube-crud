package location

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollStates(ctx *deserialize.CreatePollDeserializeContext, err error) (models.PollsStateSlice, error) {
	var pollsStates models.PollsStateSlice
	var numPollsStates int64

	if err != nil {
		return pollsStates, err
	}

	numPollsStates, err = deserialize.RNum(ctx, err)

	if numPollsStates == 0 {
		return pollsStates, err
	}

	pollsStates = make(models.PollsStateSlice, numPollsStates)

	for i := int64(0); i < numPollsStates; i++ {
		var stateId int64
		stateId, err = deserialize.RNum(ctx, err)

		if err != nil {
			return pollsStates, err
		}

		_, stateExists := ctx.LocMaps.ContinentMap[stateId]

		if !stateExists {
			return pollsStates, fmt.Errorf("invalid state id: %v", stateId)
		}

		_, stateAlreadySpecifiedInRequest := ctx.ReqLocSets.StateSet[stateId]
		if stateAlreadySpecifiedInRequest {
			return pollsStates, fmt.Errorf("state specified more than once, id: %v", stateId)
		}
		ctx.ReqLocSets.StateSet[stateId] = true

		state := ctx.LocMaps.StateMap[stateId]
		_, countrySpecifiedInRequest := ctx.ReqLocSets.CountrySet[state.CountryID]
		if !countrySpecifiedInRequest {
			return pollsStates, fmt.Errorf("no matching Country specified in request for State Id %v", stateId)
		}

		pollsStates[i] = &models.PollsState{
			StateID: stateId,
		}
	}

	return pollsStates, err
}
