package db

import (
	"github.com/diapco/votecube-crud/crud"
	"github.com/diapco/votecube-crud/sequence"
)

var (
	DimensionId              sequence.Sequence
	DimensionDirectionId     sequence.Sequence
	DirectionId              sequence.Sequence
	LabelId                  sequence.Sequence
	PollContinentId          sequence.Sequence
	PollCountryId            sequence.Sequence
	PollCountyId             sequence.Sequence
	PollDimensionDirectionId sequence.Sequence
	PollId                   sequence.Sequence
	PollLabelId              sequence.Sequence
	PollStateId              sequence.Sequence
	PollTownId               sequence.Sequence
)

func SetupSequences() {
	DimensionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  1000,
		Max:          0,
		Name:         "dimension_id",
	}

	DimensionDirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "dimension_direction_id",
	}

	DirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  1000,
		Max:          0,
		Name:         "direction_id",
	}

	LabelId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  1000,
		Max:          0,
		Name:         "label_id",
	}

	PollContinentId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_continent_id",
	}

	PollCountryId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_country_id",
	}

	PollCountyId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_county_id",
	}

	PollDimensionDirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_dimension_direction_id",
	}

	PollLabelId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_label_id",
	}

	PollId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_id",
	}

	PollStateId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_state_id",
	}

	PollTownId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  10000,
		Max:          0,
		Name:         "poll_town_id",
	}
}
