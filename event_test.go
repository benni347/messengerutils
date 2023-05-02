package messengerutils

import (
	"sync"
	"testing"
)

func TestEvent(t *testing.T) {
	event := &Event{}

	var receivedData []string
	var wg sync.WaitGroup

	// A listener function that appends received data to the receivedData slice
	listener := func(data interface{}) {
		strData, ok := data.(string)
		if !ok {
			t.Errorf("Expected data of type string, but got %T", data)
		}
		receivedData = append(receivedData, strData)
		wg.Done()
	}

	event.Subscribe(listener)

	// Emit data to the listeners
	dataToEmit := []string{"data1", "data2", "data3"}
	wg.Add(len(dataToEmit))

	for _, data := range dataToEmit {
		event.Emit(data)
	}

	wg.Wait()

	// Check if receivedData matches dataToEmit
	if len(receivedData) != len(dataToEmit) {
		t.Fatalf(
			"Expected receivedData to have %d elements, but got %d",
			len(dataToEmit),
			len(receivedData),
		)
	}

	for i, data := range receivedData {
		if data != dataToEmit[i] {
			t.Errorf("Expected receivedData[%d] to be '%s', but got '%s'", i, dataToEmit[i], data)
		}
	}
}
