package domain

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReporter_Print(t *testing.T) {
	type fields struct {
		Explorer *MarsExplorer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full report with lost robot",
			fields: fields{Explorer: &MarsExplorer{
				Surface: &Surface{
					MaxX: 4,
					MaxY: 4,
				},
				Robots: []Robot{
					{
						PosX:      3,
						PosY:      1,
						Direction: "S",
					},
					{
						PosX:      0,
						PosY:      3,
						Direction: "E",
					},
					{
						PosX:      4,
						PosY:      1,
						Direction: "N",
						Lost:      true,
					},
					{
						PosX:      2,
						PosY:      3,
						Direction: "W",
					},
				},
			}},
			want: `3 1 S
0 3 E
4 1 N LOST
2 3 W
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// better way would be to inject an io.Reader into the printer itself
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			rep := Reporter{
				Explorer: tt.fields.Explorer,
			}

			rep.Print()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout

			if tt.want != string(out) {
				t.Errorf("Print() got %s, want %s", out, tt.want)
			}
		})
	}
}
