package scheduler

type Scheduler interface {
	SelectCandidateNode()
	Score()
	Pick()
}
