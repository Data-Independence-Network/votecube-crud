package db

import (
	"github.com/diapco/votecube-crud/crud"
	"github.com/diapco/votecube-crud/models"
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
		IncrementBy:  60000,
		Max:          0,
		Name:         models.DimensionColumns.DimensionID,
	}

	DimensionDirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.DimensionDirectionColumns.DimensionDirectionID,
	}

	DirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.DirectionColumns.DirectionID,
	}

	LabelId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.LabelColumns.LabelID,
	}

	PollContinentId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsContinentColumns.PollContinentID,
	}

	PollCountryId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsCountryColumns.PollCountryID,
	}

	PollCountyId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsCountyColumns.PollCountyID,
	}

	PollDimensionDirectionId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsDimensionsDirectionColumns.PollDimensionDirectionID,
	}

	PollLabelId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsLabelColumns.PollLabelID,
	}

	PollId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollColumns.PollID,
	}

	PollStateId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsStateColumns.PollStateID,
	}

	PollTownId = sequence.Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsTownColumns.PollTownID,
	}

	numSequences := 12
	seqInitsDone := make(chan bool, numSequences)

	go DimensionId.Init(seqInitsDone)
	go DimensionDirectionId.Init(seqInitsDone)
	go DirectionId.Init(seqInitsDone)
	go LabelId.Init(seqInitsDone)
	go PollContinentId.Init(seqInitsDone)
	go PollCountryId.Init(seqInitsDone)
	go PollCountyId.Init(seqInitsDone)
	go PollDimensionDirectionId.Init(seqInitsDone)
	go PollLabelId.Init(seqInitsDone)
	go PollId.Init(seqInitsDone)
	go PollStateId.Init(seqInitsDone)
	go PollTownId.Init(seqInitsDone)

	numInitializedSequences := 0

	for range seqInitsDone {
		numInitializedSequences++
		if numInitializedSequences == numSequences {
			return
		}
	}
}
