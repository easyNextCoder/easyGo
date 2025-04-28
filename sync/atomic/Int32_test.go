package atomic

import "testing"

func Test_workInt32(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "1万个go程同时加减int型的原子变量和普通变量"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atomicInt32()
		})
	}
}

func Test_workInt32Value(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试到底atomic.Value有什么作用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atomicValue()
		})
	}
}

func Test_atomicValue2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "atomicValue2 存储不同类型变量会panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atomicValue2()
		})
	}
}
