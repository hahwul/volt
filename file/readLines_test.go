package file

import (
	"testing"
)

func TestReadLinesOrLiteral(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				arg: "",
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				arg: "ASFASFASFASFASFSDFSDFSDFHSD(F*H",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadLinesOrLiteral(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLinesOrLiteral() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
