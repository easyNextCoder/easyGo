package cache

import "testing"

func Test_work(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试本地cache value值的过期"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheValue()
		})
	}
}

func Test_cachePointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试本地cache pointer值的过期"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cachePointer()
		})
	}
}

func Test_cacheCountDelete(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试本地cache ItemCount() Delete() 函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheItemCountDelete()
		})
	}
}
