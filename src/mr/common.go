package mr

import "fmt"

type TaskPhase int

const (
	MapPhase TaskPhase = iota
	ReducePhase
)

type Task struct {
	FileName string
	NReduce  int
	NMaps    int
	Seq      int
	Phase    TaskPhase
	Alive    bool // worker should exit when alive is false
}

func reduceName(mapIdx, reduceIdx int) string {
	return fmt.Sprintf("mr-%d-%d", mapIdx, reduceIdx)
}

func mergeName(reduceIdx int) string {
	return fmt.Sprintf("mr-out-%d", reduceIdx)
}