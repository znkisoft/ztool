package config

import (
	"testing"
)

func Test_validateConfig(t *testing.T) {
	type args struct {
		config interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"test for empty pkg", args{config: &ServerConfig{Host: ""}}},
		{"test for empty pkg", args{}},
		{"test for loaded pkg", args{config: &ServerConfig{Host: "192.168.1.1"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateConfig(tt.args)
		})
	}
}
