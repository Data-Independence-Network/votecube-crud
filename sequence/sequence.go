package sequence

import (
	"database/sql"
	"fmt"
	"log"
)

type Sequence struct {
	CurrentValue int64
	Db           *sql.DB
	IncrementBy  int64
	Max          int64
	Name         string
	nextBlock    *SequenceBlock
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

	return newMax, err
}

func (seq *Sequence) Init(done chan bool) {
	seq.getNextBlock()

	done <- true
}

func (seq *Sequence) getNextBlock() {
	newMax, err := seq.selectFromDb()

	if err != nil {
		log.Fatalf("Could not obtain sequences for %s", seq.Name)
		log.Fatal(err)
		panic(err)
	}

	seq.nextBlock = &SequenceBlock{
		Start:  newMax - seq.IncrementBy,
		Length: seq.IncrementBy,
	}

}

func (seq *Sequence) GetBlocks(numVals int64) []SequenceBlock {
	if seq.CurrentValue+numVals > seq.Max {
		if seq.nextBlock == nil {
			log.Fatalf("Could not obtain sequences for %s in time", seq.Name)
			panic(nil)
		}

		if seq.Max == 0 {
			seq.CurrentValue = seq.nextBlock.Start
			seq.Max = seq.nextBlock.Start + seq.nextBlock.Length
		}

		var existingSeqBlock *SequenceBlock = nil

		var acquiredRange int64 = 0

		if seq.Max > seq.CurrentValue {
			acquiredRange = seq.Max - seq.CurrentValue
			existingSeqBlock = &SequenceBlock{Start: seq.CurrentValue + 1, Length: acquiredRange}
		}

		seq.CurrentValue = seq.nextBlock.Start
		seq.Max = seq.nextBlock.Start + seq.nextBlock.Length
		seq.nextBlock = nil

		go seq.getNextBlock()

		newVals := numVals - acquiredRange

		newSeqBlock := SequenceBlock{Start: seq.CurrentValue + 1, Length: newVals}

		seq.CurrentValue += newVals

		if existingSeqBlock == nil {
			return []SequenceBlock{newSeqBlock}
		} else {
			return []SequenceBlock{*existingSeqBlock, newSeqBlock}
		}
	} else {
		newSeqBlock := SequenceBlock{Start: seq.CurrentValue + 1, Length: numVals}

		seq.CurrentValue += numVals

		return []SequenceBlock{newSeqBlock}
	}
}
