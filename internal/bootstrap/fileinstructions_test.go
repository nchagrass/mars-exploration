package bootstrap

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_contentToStringArray(t *testing.T) {
	type args struct {
		c io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "source is converted into an array of strings",
			args: args{c: strings.NewReader("abcde\r\nhello")},
			want: []string{"abcde", "hello"},
		},
		{
			name: "source is converted into an array of strings including empty lines",
			args: args{c: strings.NewReader("abcde\r\nhello\r\n\r\n")},
			want: []string{"abcde", "hello", ""},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := contentToStringArray(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentToStringArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contentToStringArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}
