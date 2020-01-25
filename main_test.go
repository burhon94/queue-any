package main

import "testing"

func Test_queue_with_empty_element(t *testing.T) {
	queueList := queue{}
	if queueList.lenQueue() != 0 {
		t.Error("The queue length must be zero, got: ", queueList.lenQueue())
	}
	firstInQueue := queueList.firstInQueue()
	if firstInQueue != nil {
		t.Error("The first element must be nil from empty queue, got: ", firstInQueue)
	}
	lastInQueue := queueList.lastInQueue()
	if lastInQueue != nil {
		t.Error("The last element must be nil from empty queue, got: ", lastInQueue)
	}
	queueList.addLast("Jack")
	if queueList.lenQueue() != 1 {
		t.Error("The queue length must be equal 1, got: ", queueList.lenQueue())
	}
	queueList.removeFirstFromQueue()
	if queueList.lenQueue() != 0 {
		t.Error("The queue after remove element must be empty, got: ", queueList.lenQueue())
	}
}

func Test_queue_with_one_element(t *testing.T) {
	queueList := queue{}
	queueList.addLast("Jack")
	if queueList.lenQueue() != 1 {
		t.Error("The queue length must be equal one, got: ", queueList.lenQueue())
	}
	firstInQueue := queueList.firstInQueue()
	if firstInQueue != "Jack" {
		t.Error("The first element in queue must be \"Jack\", got: ", firstInQueue)
	}
	lastInQueue := queueList.lastInQueue()
	if lastInQueue != "Jack" {
		t.Error("The last element in queue must be \"Jack\", got: ", lastInQueue)
	}
	queueList.addLast("Alex")
	if queueList.lenQueue() != 2 {
		t.Error("The queue length must be equal 2, got: ", queueList.lenQueue())
	}
	queueList.removeFirstFromQueue()
	if queueList.lastInQueue() != queueList.firstInQueue() {
		t.Error("The remove first element in queue \"Jack\",the first and the last must be \"Alex\" , got: ", "first in queue ", queueList.firstInQueue(), ", last in queue ", queueList.lastInQueue())
	}
}

func Test_queue_with_multiple_element(t *testing.T) {
	queueList := queue{}
	queueList.addLast("Jack")
	queueList.addLast("Alex")
	queueList.addLast("Mike")
	firstInQueue := queueList.firstInQueue()
	if firstInQueue != "Jack" {
		t.Error("The first in queue must be \"Jack\", got: ", firstInQueue)
	}
	lastInQueue := queueList.lastInQueue()
	if lastInQueue != "Mike" {
		t.Error("The last in queue must be \"Mike\", got: ", lastInQueue)
	}
	queueList.addLast("Anna")
	if queueList.lastInQueue() != "Anna" {
		t.Error("After add last in queue \"Anna\", the last must be, got: ", queueList.lastInQueue())
	}
	queueList.removeFirstFromQueue()
	if queueList.firstInQueue() == "Jack" {
		t.Error("After remove the first element in queue \"Jack\", the first in queue must be \"Alex\", got: ", queueList.firstInQueue())
	}
}

func Test_remove_first_from_empty_queue(t *testing.T) {
	queueList := queue{}
	queueList.removeFirstFromQueue()
	if queueList.lenQueue() != 0 {
		t.Error("The queue length must be 0, got: ", queueList.lenQueue())
	}
}
