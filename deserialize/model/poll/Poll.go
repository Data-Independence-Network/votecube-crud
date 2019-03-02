package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/deserialize/model/location"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func deserializePoll(data []byte, cursor *int64, dataLen int64, err error) (models.Poll, error) {
	var poll models.Poll

	poll.EndDate, err = deserialize.RTime(data, cursor, dataLen, err)
	poll.PollTitle, err = deserialize.RStr(data, cursor, dataLen, err)

	poll.R.PollsContinents, err = location.DeserializePollContinents(data, cursor, dataLen, err)
	poll.R.PollsCountries, err = location.DeserializePollCountries(data, cursor, dataLen, err)
	poll.R.PollsCounties, err = location.DeserializePollCounties(data, cursor, dataLen, err)
	poll.R.PollsDimensionsDirections, err = DeserializePollDimDirs(data, cursor, dataLen, err)
	poll.R.PollsLabels, err = DeserializePollLabels(data, cursor, dataLen, err)
	poll.R.PollsStates, err = location.DeserializePollStates(data, cursor, dataLen, err)
	poll.R.PollsTowns, err = location.DeserializePollTowns(data, cursor, dataLen, err)

	return poll, err
}
