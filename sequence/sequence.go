package sequence

import (
	"database/sql"
	"fmt"
)

type Sequence struct {
	CurrentValue int64
	Db           *sql.DB
	IncrementBy  int64
	Max          int64
	Name         string
}

type SequenceBlock struct {
	Start  int64
	Length int64
}

type SequenceError struct {
}

func (e *SequenceError) Error() string {
	return fmt.Sprintf("at %v", 0)
}

func (seq *Sequence) Get() (int64, error) {

}

func (seq *Sequence) selectFromDb() (int64, error) {
	query := "select nextval('votecube." + seq.Name + "');"
	//fmt.Println(query)
	rows, err := seq.Db.Query(query)
	var newMax int64
	if err != nil {
		return newMax, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&newMax)
		if err != nil {
			return newMax, err
		}
	}

	err = rows.Err()
	if err != nil {
		return newMax, err
	}
}

func (seq *Sequence) GetBlocks(numVals int64) ([]SequenceBlock, error) {
	var seqBlocks []SequenceBlock

	if seq.CurrentValue+numVals > seq.Max {
		newMax, err := seq.selectFromDb()

		if err != nil {
			return nil, err
		}

		var existingSeqBlock *SequenceBlock = nil

		var acquiredRange int64 = 0

		if seq.Max > seq.CurrentValue {
			acquiredRange = seq.Max - seq.CurrentValue
			existingSeqBlock = &SequenceBlock{Start: seq.CurrentValue + 1, Length: acquiredRange}
		}

		seq.Max = newMax
		seq.CurrentValue = newMax - seq.IncrementBy

		newVals := numVals - acquiredRange

		newSeqBlock := SequenceBlock{Start: seq.CurrentValue + 1, Length: newVals}

		seq.CurrentValue += newVals

		if existingSeqBlock == nil {
			return []SequenceBlock{newSeqBlock}, nil
		} else {
			return []SequenceBlock{*existingSeqBlock, newSeqBlock}, nil
		}

	} else {
		newSeqBlock := SequenceBlock{Start: seq.CurrentValue + 1, Length: numVals}

		seq.CurrentValue += numVals

		return []SequenceBlock{newSeqBlock}, nil
	}

	return seqBlocks, nil
}
