package domain

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMarsBuilder_NewSurface(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *Surface
		wantErr bool
	}{
		{
			name: "input is transformed into surface successfully",
			args: args{
				line: "1 5",
			},
			want: &Surface{
				MaxX: 1,
				MaxY: 5,
			},
			wantErr: false,
		},
		{
			name: "type error for X",
			args: args{
				line: "R 5",
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "type error for Y",
			args: args{
				line: "2 G",
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "not enough values",
			args: args{
				line: "2",
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "too many values",
			args: args{
				line: "1 2 5",
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logrus.New()
			l.SetOutput(ioutil.Discard)

			mb := &MarsBuilder{
				logger: l,
			}
			got, err := mb.NewSurface(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSurface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSurface() got = %v, want %v", got, tt.want)
			}
		})
	}
}
