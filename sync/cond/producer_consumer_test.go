package cond

import "testing"

func TestConsumeProduce(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "经典的生产者消费者模型"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConsumeProduce()
		})
	}
}
