package file

import (
	"os"
	"testing"
)

func TestReadLinesOrLiteral(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "literal string",
			args: args{
				arg: "literal",
			},
			want:    []string{"literal"},
			wantErr: false,
		},
		{
			name: "empty file",
			args: args{
				arg: "emptyfile.txt",
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "file with content",
			args: args{
				arg: "testfile.txt",
			},
			want:    []string{"line1", "line2", "line3"},
			wantErr: false,
		},
	}
	// Create test files
	_ = os.WriteFile("emptyfile.txt", []byte(""), 0644)
	_ = os.WriteFile("testfile.txt", []byte("line1\nline2\nline3"), 0644)
	defer os.Remove("emptyfile.txt")
	defer os.Remove("testfile.txt")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLinesOrLiteral(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLinesOrLiteral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equal(got, tt.want) {
				t.Errorf("ReadLinesOrLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
