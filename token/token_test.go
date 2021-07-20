package token

import "testing"

func TestMakeRandomTokenSha256(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			name: "Test",
			args: "volt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeRandomTokenSha256(tt.args); got == "" {
				t.Errorf("MakeRandomTokenSha256() = %v, want %v", got, "")
			}
		})
	}
}
