package util

import (
	"fmt"
	"testing"
)

func Test_genUUID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case 1"},
		{"case 2"},
		{"case 3"},
		{"case 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenUUID()
			fmt.Println(got)
			if len(got) != 36 {
				t.Errorf("gen uuid error")
			}
		})
	}
}

func Test_genShortId(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"case 1", false},
		{"case 2", false},
		{"case 3", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenShortId()
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenShortId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
