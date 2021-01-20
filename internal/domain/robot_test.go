package domain

import "testing"

func TestRobot_Execute(t *testing.T) {
	type fields struct {
		PosX         int
		PosY         int
		Direction    string
		Instructions []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "robot can turn right multiple times",
			fields: fields{
				Direction:    "E",
				Instructions: []string{"R", "R", "R", "R"},
			},
			want: []string{"S", "W", "N", "E"},
		},
		{
			name: "robot can turn right and left multiple times",
			fields: fields{
				Direction:    "N",
				Instructions: []string{"R", "R", "L", "L"},
			},
			want: []string{"E", "S", "E", "N"},
		},
		{
			name: "robot fails for unsupported direction",
			fields: fields{
				Direction: "R",
			},
			want:    nil,
			wantErr: true,
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
				if r.Direction != tt.want[i] {
					t.Errorf("Robot.Execute() got %s, want %s", r.Direction, tt.want[i])
				}
			}
		})
	}
}
