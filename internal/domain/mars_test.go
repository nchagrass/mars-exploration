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
			want:    nil,
			wantErr: true,
		},
		{
			name: "type error for Y",
			args: args{
				line: "2 G",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not enough values",
			args: args{
				line: "2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "too many values",
			args: args{
				line: "1 2 5",
			},
			want:    nil,
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

func TestMarsBuilder_LoadRobotInstructions(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Robot
		wantErr bool
	}{
		{
			name: "successfully load n robots",
			args: args{
				lines: []string{
					"1 1 E",
					"RFRFRFRF",
					"",
					"1 2 N",
					"RFRFRFRF",
					"",
					"5 3 S",
					"RFRFRFRF",
					"",
				},
			},
			want: []Robot{
				{
					PosX:         1,
					PosY:         1,
					Direction:    "E",
					Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
				},
				{
					PosX:         1,
					PosY:         2,
					Direction:    "N",
					Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
				},
				{
					PosX:         5,
					PosY:         3,
					Direction:    "S",
					Instructions: []string{"R", "F", "R", "F", "R", "F", "R", "F"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logrus.New()
			l.SetOutput(ioutil.Discard)

			mb := &MarsBuilder{
				logger: l,
			}
			got, err := mb.LoadRobotInstructions(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRobotInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadRobotInstructions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
