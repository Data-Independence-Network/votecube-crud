package poll

import (
	"fmt"
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/deserialize/model/location"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializePoll(ctx *deserialize.CreatePollDeserializeContext, err error) (models.Poll, error) {
	var poll models.Poll

	if err != nil {
		return poll, err
	}

	poll.EndDate, err = deserialize.RTime(ctx, err)
	poll.PollTitle, err = deserialize.RStr(ctx, err)

	if len(poll.PollTitle) < 3 {
		return poll, fmt.Errorf("poll Title is less than 3 characters long")
	}

	poll.R.PollsContinents, err = location.DeserializePollContinents(ctx, err)
	poll.R.PollsCountries, err = location.DeserializePollCountries(ctx, err)
	poll.R.PollsStates, err = location.DeserializePollStates(ctx, err)
	//poll.R.PollsCounties, err = location.DeserializePollCounties(ctx, err)
	poll.R.PollsTowns, err = location.DeserializePollTowns(ctx, err)
	poll.R.PollsFactorsPositions, err = DeserializePollFactorPositions(ctx, err)
	poll.R.PollsLabels, err = DeserializePollLabels(ctx, err)

	poll.StartDate, err = deserialize.RTime(ctx, err)

	if poll.StartDate.Before(ctx.Tomorrow) {
		return poll, fmt.Errorf("provided Start Date is before tomorrow: %v", poll.StartDate)
	}

	if poll.EndDate.Before(poll.StartDate) {
		return poll, fmt.Errorf("provided End Date is before Start Date: %v", poll.StartDate)
	}

	poll.ThemeID, err = deserialize.RNum(ctx, err)
	_, themeIdExists := (*ctx.ThemeMap)[poll.ThemeID]

	if !themeIdExists {
		return poll, fmt.Errorf("provided Theme ID does not exist: %v", poll.ThemeID)
	}

	poll.UserAccountID = ctx.Request.UserAccountId

	return poll, err
}
