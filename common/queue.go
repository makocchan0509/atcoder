package common

//please change me
type address struct {
	x int
	y int
}

type queue struct {
	q []address
}

func (qu *queue) dequeue() address {
	a := qu.q[0]
	qu.q = qu.q[1:]
	return a
}

func (qu *queue) enqueue(a address) {
	qu.q = append(qu.q, a)
}

func (qu *queue) len() int {
	return len(qu.q)
}
