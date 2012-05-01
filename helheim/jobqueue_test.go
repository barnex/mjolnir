package helheim

import(
	"testing"
)

func TestJobQueue(test*testing.T){
	var q JobQueue

	q.Push(&Job{3})
	q.Push(&Job{7})
	q.Push(&Job{2})
	q.Push(&Job{9})
	q.Push(&Job{1})

	prev := 10
	if q.Len() != 5{
		test.Error("q.Len(): ", q.Len())
	}
	for q.Len() != 0{
		job := q.Pop()
		if job.priority > prev{test.Error(prev, "<", job.priority)}
		prev = job.priority
	}
}
