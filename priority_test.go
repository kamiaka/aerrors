package aerrors

import "testing"

func TestErrorPriority_HigherThan(t *testing.T) {
	cases := []struct {
		a    ErrorPriority
		b    ErrorPriority
		want bool
	}{
		{
			a:    Emergency,
			b:    Alert,
			want: true,
		},
		{
			a:    Emergency,
			b:    Info,
			want: true,
		},
		{
			a:    Emergency,
			b:    Emergency,
			want: false,
		},
		{
			a:    Info,
			b:    Emergency,
			want: false,
		},
	}
	for i, tc := range cases {
		got := tc.a.HigherThan(tc.b)
		if got != tc.want {
			t.Errorf("#%d: ErrorPriority(%d).HigherThan(ErrorPriority(%d)) == %v, want %v", i, tc.a, tc.b, got, tc.want)
		}
	}
}

func TestErrorPriority_String(t *testing.T) {
	cases := []struct {
		priority ErrorPriority
		want     string
	}{
		{
			priority: Error,
			want:     "Error",
		},
		{
			priority: ErrorPriority(999),
			want:     "ErrorPriority(999)",
		},
	}

	for i, tc := range cases {
		got := tc.priority.String()
		if got != tc.want {
			t.Errorf("#%d: (ErrorPriority(%d)).String() == %#v, want %#v", i, tc.priority, got, tc.want)
		}
	}
}
