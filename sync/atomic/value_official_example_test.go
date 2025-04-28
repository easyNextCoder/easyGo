package atomic

import "testing"

func Test_valueOfficalExampleWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试官方的使用atomic.Value的例子"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valueOfficialExampleWork()
		})
	}
}
