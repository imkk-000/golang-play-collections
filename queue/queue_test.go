package queue

import (
	"reflect"
	"testing"
)

func Test_New_Should_Be_Queue(t *testing.T) {
	expectedQueue := Queue{}

	actualQueue := New()

	if !reflect.DeepEqual(expectedQueue, actualQueue) {
		t.Errorf("expect %v but it got %v", expectedQueue, actualQueue)
	}
}

func Test_Enqueue_Input_String_Hello_And_String_World_Should_Be_Exist_String_Hello_And_World_In_Queue(t *testing.T) {
	expectedQueue := Queue{
		values: []interface{}{"hello", "world"},
	}

	queue := New()
	queue.Enqueue("hello")
	queue.Enqueue("world")
	actualQueue := queue

	if !reflect.DeepEqual(expectedQueue, actualQueue) {
		t.Errorf("expect %v but it got %v", expectedQueue, actualQueue)
	}
}
