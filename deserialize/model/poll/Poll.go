package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/deserialize/model/location"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePoll(ctx *deserialize.DeserializeContext, err error) (models.Poll, error) {
	var poll models.Poll

	if err != nil {
		return poll, err
	}

	poll.EndDate, err = deserialize.RTime(ctx, err)
	poll.PollTitle, err = deserialize.RStr(ctx, err)

	poll.R.PollsContinents, err = location.DeserializePollContinents(ctx, err)
	poll.R.PollsCountries, err = location.DeserializePollCountries(ctx, err)
	//poll.R.PollsCounties, err = location.DeserializePollCounties(ctx, err)
	poll.R.PollsDimensionsDirections, err = DeserializePollDimDirs(ctx, err)
	poll.R.PollsLabels, err = DeserializePollLabels(ctx, err)
	poll.R.PollsStates, err = location.DeserializePollStates(ctx, err)
	poll.R.PollsTowns, err = location.DeserializePollTowns(ctx, err)

	return poll, err
}
