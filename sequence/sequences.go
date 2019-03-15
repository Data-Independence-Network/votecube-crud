package sequence

import (
	"github.com/diapco/votecube-crud/crud"
	"github.com/diapco/votecube-crud/models"
)

var (
	FactorId             Sequence
	FactorPositionId     Sequence
	PositionId           Sequence
	LabelId              Sequence
	PollContinentId      Sequence
	PollCountryId        Sequence
	PollCountyId         Sequence
	PollFactorPositionId Sequence
	PollId               Sequence
	PollLabelId          Sequence
	PollStateId          Sequence
	PollTownId           Sequence
)

func SetupSequences() {
	FactorId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.FactorColumns.FactorID,
	}

	FactorPositionId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.FactorPositionColumns.FactorPositionID,
	}

	PositionId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PositionColumns.PositionID,
	}

	LabelId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.LabelColumns.LabelID,
	}

	PollContinentId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsContinentColumns.PollContinentID,
	}

	PollCountryId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsCountryColumns.PollCountryID,
	}

	PollCountyId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsCountyColumns.PollCountyID,
	}

	PollFactorPositionId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsFactorsPositionColumns.PollFactorPositionID,
	}

	PollLabelId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsLabelColumns.PollLabelID,
	}

	PollId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollColumns.PollID,
	}

	PollStateId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsStateColumns.PollStateID,
	}

	PollTownId = Sequence{
		CurrentValue: 0,
		Db:           crud.DB,
		IncrementBy:  60000,
		Max:          0,
		Name:         models.PollsTownColumns.PollTownID,
	}

	numSequences := 12
	seqInitsDone := make(chan bool, numSequences)

	go FactorId.Init(seqInitsDone)
	go FactorPositionId.Init(seqInitsDone)
	go PositionId.Init(seqInitsDone)
	go LabelId.Init(seqInitsDone)
	go PollContinentId.Init(seqInitsDone)
	go PollCountryId.Init(seqInitsDone)
	go PollCountyId.Init(seqInitsDone)
	go PollFactorPositionId.Init(seqInitsDone)
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
