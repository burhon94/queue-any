package main

import "testing"

func Test_empty_queue(t *testing.T) {
	l := queue{}
	if l.lenQueue() != 0 {
		t.Error("queue must be zero, got: ", l.lenQueue())
	}
}

func Test_queue_with_one_element(t *testing.T)  {
	l := queue{}
	l.addElementToQueue(1)
	if l.lenQueue() != 1 {
		t.Error("queue size must be one, got: ", l.lenQueue())
	}
}

func Test_queue_with_multiple_elements(t *testing.T)  {
	l := queue{}
	l.addElementToQueue(1)
	l.addLastQueue(2)
	l.addLastQueue(3)
	l.addLastQueue(4)
	l.addLastQueue(5)
	if l.lenQueue() != 5 {
		t.Error("queue size must be 5, got: ", l.lenQueue())
	}
}

func Test_add_element_to_empty_queue(t *testing.T){
	l := queue{}
	l.addElementToQueue(1)
	if l.lenQueue() != 1{
		t.Error("queue size after adding one element must be one, got: ", l.lenQueue())
	}
}

func Test_add_element_last_in_queue_with_queue_one_element(t *testing.T){
	l := queue{}
	l.addElementToQueue(1)
	l.addLastQueue(2)
	if l.lenQueue() != 2{
		t.Error("queue size is one after adding one element must be 2, got: ", l.lenQueue())
	}
}

func Test_add_elements_last_to_queue_with_multiple_elements(t *testing.T){
	l := queue{}
	l.addElementToQueue(1)
	l.addLastQueue(2)
	l.addLastQueue(1)
	l.addLastQueue(2)
	l.addLastQueue(1)
	l.addLastQueue(2)
	if l.lenQueue() != 6{
		t.Error("queue size is 3 after adding 3 elements must be 6, got: ", l.lenQueue())
	}
}

func Test_remove_first_element_to_empty_queue(t *testing.T){
	l := queue{}
	l.size = 0
	l.removeFirstFromQueue()
	if l.lenQueue() != 0{
		t.Error("queue size is zero after remove first element queue size must be zero(IT'S NOT LOGIC), got: ", l.lenQueue())
	}
}

func Test_remove_first_element_from_queue_with_one_element(t *testing.T){
	l := queue{}
	l.size = 1
	l.removeFirstFromQueue()
	if l.lenQueue() != 0{
		t.Error("queue size is one after remove first element queue size must be zero, got: ", l.lenQueue())
	}
}

func Test_remove_first_element_from_queue_with_multiple_element(t *testing.T){
	l := queue{}
	l.addElementToQueue(1)
	l.addLastQueue(2)
	l.addLastQueue(3)
	l.addLastQueue(4)
	l.addLastQueue(5)
	l.removeFirstFromQueue()
	if l.lenQueue() != 4{
		t.Error("queue size is 5 after remove first element queue size must be 4, got: ", l.lenQueue())
	}
}