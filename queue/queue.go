package queue

type Quque []int

func (q *Quque) Push(v int) {
	*q = append(*q, v)
}

func (q *Quque) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Quque) IsEmpty() bool {
	return len(*q) == 0
}
