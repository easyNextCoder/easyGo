package main

import "testing"

func TestCancelWhenError(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "有一个任务失败的时候就取消所有任务"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CancelWhenError()
		})
	}
}

func TestReturnFirstError(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "任务全部执行，并返回遇到的第一个错误"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnFirstError()
		})
	}
}
