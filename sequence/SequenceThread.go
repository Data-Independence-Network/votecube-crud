package sequence

type SequenceThread struct {
	Seq Sequence
	Get chan int64
}


func (seqThread *SequenceThread) get() {
	seqThread.Seq.
}