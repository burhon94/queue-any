package main

type queue struct {
	firstQueue *queueNode
	lastQueue  *queueNode
	size       int
}

type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value string
}

func (receiver *queue) lenQueue() int {
	return receiver.size
}

func (receiver *queue) firstInQueue() interface{} {
	if receiver.firstQueue == nil {
		return nil
	}
	return receiver.firstQueue.value
}

func (receiver *queue) lastInQueue() interface{} {
	if receiver.lastQueue == nil {
		return nil
	}
	return receiver.lastQueue.value
}

func (receiver *queue) addLast(elementPtr string) {
	if receiver.lenQueue() == 0 {
		receiver.firstQueue = &queueNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
		}
		receiver.size++
		receiver.lastQueue = receiver.firstQueue
		return
	}
	receiver.size++
	current := receiver.firstQueue
	for {
		if current.next == nil {
			current.next = &queueNode{
				next:  nil,
				prev:  current,
				value: elementPtr,
			}
			receiver.lastQueue = current.next
			return
		}
		current = current.next
	}
}

func (receiver *queue) removeFirstFromQueue() {
	if receiver.size == 0 {
		return
	}
	if receiver.size == 1 {
		receiver.firstQueue = nil
		receiver.lastQueue = nil
		receiver.size--
		return
	}
	receiver.firstQueue = receiver.firstQueue.next
	receiver.firstQueue.prev = nil
	receiver.size--
	return
}

func main() {

}
