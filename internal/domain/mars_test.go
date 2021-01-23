package domain

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"strings"
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
		{
			name: "X is too big",
			args: args{
				line: "51 10",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Y is too big",
			args: args{
				line: "50 999",
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "max values are ok",
			args: args{
				line: "50 50",
			},
			want: &Surface{
				MaxX: 50,
				MaxY: 50,
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
		{
			name: "instructions can't exceed a 100",
			args: args{
				lines: []string{
					"1 1 E",
					strings.Repeat("RFRFRF", 17),
					"",
					"1 2 N",
					"RFRFRFRF",
					"",
					"5 3 S",
					"RFRFRFRF",
					"",
				},
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

func TestMarsExplorer_SendInstructions(t *testing.T) {
	type fields struct {
		Surface *Surface
		Robots  []Robot
	}
	tests := []struct {
		name   string
		fields fields
		want   []Robot
	}{
		{
			name: "robot can navigate on mars successfully",
			fields: fields{
				Surface: &Surface{
					MaxX: 4,
					MaxY: 4,
				},
				Robots: []Robot{
					{PosX: 0, PosY: 0, Direction: "E", Instructions: []string{"F", "L", "F", "R", "F"}},
				},
			},
			want: []Robot{
				{PosX: 2, PosY: 1, Direction: "E", Instructions: []string{"F", "L", "F", "R", "F"}},
			},
		},

		{
			name: "1 robot can get lost",
			fields: fields{
				Surface: &Surface{
					MaxX: 5,
					MaxY: 3,
				},
				Robots: []Robot{
					{PosX: 3, PosY: 2, Direction: "N", Lost: false, Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"}},
				},
			},
			want: []Robot{
				{PosX: 3, PosY: 3, Direction: "N", Lost: true, Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"}},
			},
		},
		{
			name: "2 robots can't get lost at the same spot",
			fields: fields{
				Surface: &Surface{
					MaxX: 5,
					MaxY: 3,
				},
				Robots: []Robot{
					{PosX: 3, PosY: 2, Direction: "N", Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"}},
					{PosX: 3, PosY: 3, Direction: "N", Instructions: []string{"F", "R", "F"}},
				},
			},
			want: []Robot{
				{PosX: 3, PosY: 3, Direction: "N", Lost: true, Instructions: []string{"F", "R", "R", "F", "L", "L", "F", "F", "R", "R", "F", "L", "L"}},
				{PosX: 4, PosY: 3, Direction: "E", Lost: false, Instructions: []string{"F", "R", "F"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarsExplorer{
				Surface: tt.fields.Surface,
				Robots:  tt.fields.Robots,
			}

			m.SendInstructions()

			for i, r := range m.Robots {
				if !reflect.DeepEqual(r, tt.want[i]) {
					t.Errorf("SendInstructions() got %v, want %v", r, tt.want[i])
				}
			}
		})
	}
}
