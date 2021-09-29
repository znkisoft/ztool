package config

import (
	"testing"
)

func Test_validateConfig(t *testing.T) {
	tests := []struct {
		name string
		args *ServerConfig
	}{
		// TODO handling test case result with panic status
		{"test for empty pkg", &ServerConfig{Host: "dummy test"}},
		{"test for loaded pkg", &ServerConfig{Host: "192.168.1.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateConfig(tt.args)
		})
	}
}
