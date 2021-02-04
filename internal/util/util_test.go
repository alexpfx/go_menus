package util

import (
	"reflect"
	"testing"
)

func TestAppendIf(t *testing.T) {
	original := []string{"a"}

	type args struct {
		slice    []string
		argName  string
		argValue interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				slice:    original,
				argName:  "-b",
				argValue: true,
			},
			want: []string{"a", "-b"},
		},
		{
			name: "t2",
			args: args{
				slice:    original,
				argName:  "-c",
				argValue: "string",
			},
			want: []string{"a", "-c", "string"},
		},
		{
			name: "t3",
			args: args{
				slice:    original,
				argName:  "-d",
				argValue: 1,
			},
			want: []string{"a", "-d", "1"},
		},
		{
			name: "t4",
			args: args{
				slice:    original,
				argName:  "-e",
				argValue: 1,
			},
			want: []string{"a", "-e", "1"},
		},
		{
			name: "t5",
			args: args{
				slice:   original,
				argName: "-f",
				argValue: []string{
					"multi", "param",
				},
			},
			want: []string{"a", "-f", "multi", "param"},
		},
		{
			name: "t6",
			args: args{
				slice:   original,
				argName: "-g",
				argValue: map[string]string{
					"not": "supported",
				},
			},
			want: []string{"a"},
		},
		{
			name: "t7",
			args: args{
				slice:   original,
				argName: "-h",
				argValue: []map[string]string{
					{
						"not": "supported",
					},
					{
						"not": "supported",
					},
				},
			},
			want: []string{"a"},
		},
		{
			name: "t8",
			args: args{
				slice:    original,
				argName:  "-i",
				argValue: 0,
			},
			want: []string{"a", "-i", "0"},
		},
		{
			name: "t9",
			args: args{
				slice:    original,
				argName:  "-j",
				argValue: "",
			},
			want: []string{"a", "-j", ""},
		},
		{
			name: "t10",
			args: args{
				slice:    original,
				argName:  "",
				argValue: "",
			},
			want: []string{"a", ""},
		},
		{
			name: "t11",
			args: args{
				slice:    original,
				argName:  "",
				argValue: []string{"abc", "def"},
			},
			want: []string{"a", "abc", "def"},
		},
		{
			name: "t12",
			args: args{
				slice:    original,
				argName:  "abc",
				argValue: nil,
			},
			want: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendIf(tt.args.slice, tt.args.argName, tt.args.argValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIf() = %v, want %v", got, tt.want)
			}
		})
	}
}
