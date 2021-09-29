package crypto

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
		{"case 1", args{pwd, "da0246c88464f6f7a3a68789ea169d5a74eb919246b366d86c1811baf436525e"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
