package domain

import "testing"

func TestRobot_Execute(t *testing.T) {
	type fields struct {
		PosX         int
		PosY         int
		Direction    string
		Instructions []string
	}

	type want struct {
		directions []string
		posX, posY []int
	}

	tests := []struct {
		name    string
		fields  fields
		want    want
		wantErr bool
	}{
		{
			name: "robot can turn right multiple times",
			fields: fields{
				Direction:    "E",
				Instructions: []string{"R", "R", "R", "R"},
			},
			want: want{
				[]string{"S", "W", "N", "E"},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
		},
		{
			name: "robot can turn right and left multiple times",
			fields: fields{
				Direction:    "N",
				Instructions: []string{"R", "R", "L", "L"},
			},
			want: want{
				[]string{"E", "S", "E", "N"},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
		},
		{
			name: "robot fails for unsupported direction",
			fields: fields{
				Direction: "R",
			},
			want:    want{},
			wantErr: true,
		},
		{
			name: "robot can go forward and move around",
			fields: fields{
				PosX:         0,
				PosY:         0,
				Direction:    "E",
				Instructions: []string{"F", "L", "F"},
			},
			want: want{
				directions: []string{"E", "N", "N"},
				posX:       []int{1, 1, 1},
				posY:       []int{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Robot{
				PosX:         tt.fields.PosX,
				PosY:         tt.fields.PosY,
				Direction:    tt.fields.Direction,
				Instructions: tt.fields.Instructions,
			}

			for i, c := range r.Instructions {
				err := r.Execute(c)
				if (err != nil) != tt.wantErr {
					t.Errorf("Robot.Execute() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if r.Direction != tt.want.directions[i] {
					t.Errorf("Robot.Execute() got direction %s, want %s", r.Direction, tt.want.directions[i])
				}
				if r.PosY != tt.want.posY[i] {
					t.Errorf("Robot.Execute() got pos Y %d, want %d", r.PosY, tt.want.posY[i])
				}
				if r.PosX != tt.want.posX[i] {
					t.Errorf("Robot.Execute() got pos X %d, want %d", r.PosX, tt.want.posX[i])
				}
			}
		})
	}
}
