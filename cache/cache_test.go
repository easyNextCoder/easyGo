package cache

import "testing"

func Test_work(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试本地cache的过期"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			work()
		})
	}
}
