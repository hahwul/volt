package file

import (
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLinesOrLiteral(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLinesOrLiteral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLinesOrLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readLines(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readLines(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFile(tt.args.path); got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
