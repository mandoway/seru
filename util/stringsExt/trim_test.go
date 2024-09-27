package stringsExt

import "testing"

func TestTrimAfterLast(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "selector",
			args: args{
				s:      "foo.bar",
				substr: ".",
			},
			want: "foo",
		},
		{
			name: "selector full",
			args: args{
				s:      "foo.bar",
				substr: ".bar",
			},
			want: "foo",
		},
		{
			name: "email",
			args: args{
				s:      "example@mail.com",
				substr: "@",
			},
			want: "example",
		},
		{
			name: "not found",
			args: args{
				s:      "example@mail.com",
				substr: "?",
			},
			want: "example@mail.com",
		},
		{
			name: "trim to empty",
			args: args{
				s:      "example@mail.com",
				substr: "example",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimAfterLast(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("TrimAfterLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
