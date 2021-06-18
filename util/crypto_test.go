package util

import (
	"fmt"
	"testing"
)

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}

	pwd := "Admin@1234"
	hash, _ := HashPassword(pwd)
	fmt.Println(hash)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{pwd, hash}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
