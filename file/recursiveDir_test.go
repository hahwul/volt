package file

import (
	"testing"
)

func TestGetFiles(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				path: ".",
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				path: "@$&*^#*",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFiles(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFiles() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("%s", got)
				return
			}
		})
	}
}
