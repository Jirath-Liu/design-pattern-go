package simple

import (
	"testing"
)

func TestFactory_Build(t *testing.T) {
	type args struct {
		name     string
		typeName string
	}
	tests := []struct {
		name   string
		args   args
		want1  bool
	}{
		{
			name:  "test",
			args:  args{
				name:     "one",
				typeName: "one",
			},
			want1: true,
		},
		{
			name:  "test2",
			args:  args{
				name:     "three",
				typeName: "three",
			},
			want1: true,
		},
		{
			name:  "test3",
			args:  args{
				name:     "two",
				typeName: "two",
			},
			want1: true,
		},
		{
			name:  "testFalse",
			args:  args{
				name:     "ne",
				typeName: "ne",
			},
			want1: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory()
			got, got1 := f.Build(tt.args.name, tt.args.typeName)
			if got1 != tt.want1 {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got1 {
				t.Logf("values:%s", got.GetName())
			}

		})
	}
}
