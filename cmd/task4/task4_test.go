package main

import (
	"reflect"
	"testing"
)

func Test_someFunction(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		want    interface{}
		wantErr bool
	}{
		{
			"temp 1",
			[6]int{17, 18, 5, 4, 6, 1},
			[]int64{18, 6, 6, 6, 1, -1},
			false,
		},
		{
			"temp 2",
			[1]int{400},
			[]int64{-1},
			false,
		},
		{
			"temp 3",
			[10001]int{},
			nil,
			true,
		},
		{
			"temp 4",
			[12]int{17, 1, 0, -10, 18, -5, 100000, 100001, 0, -10, 0, -5},
			nil,
			true,
		},
		{
			"temp 4",
			[]int{0, 001, 00001, 0000000001},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := someFunction(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("someFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("someFunction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
