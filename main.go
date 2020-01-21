package main

type queue struct {
	firstQueue *queueNode
	lastQueue  *queueNode
	size       int
}

type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
}

func (receiver *queue) lenQueue() int {
	return receiver.size
}

func (receiver *queue) addElementToQueue(elementPtr interface{}) {
	if receiver.lenQueue() == 0 {
		receiver.firstQueue = &queueNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
		}
		receiver.lastQueue = receiver.firstQueue
		receiver.size++
		return
	}
	receiver.size++
}

func (receiver *queue) addLastQueue(elementPtr interface{}) {
	if receiver.lenQueue() == 0 {
		receiver.firstQueue = &queueNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
		}
		receiver.lastQueue = receiver.firstQueue
		receiver.size++
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
}

func main() {

}
