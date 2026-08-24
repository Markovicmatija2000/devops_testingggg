package domain

import "testing"

func TestStatus_String(t *testing.T) {
	cases := []struct {
		name   string
		status Status
		want   string
	}{
		{"pending", Pending, "Pending"},
		{"working", Working, "Working"},
		{"done", Done, "Done"},
		{"unknown value", Status(99), "Unknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.status.String(); got != c.want {
				t.Errorf("Status(%d).String() = %q, want %q", c.status, got, c.want)
			}
		})
	}
}

func TestParseTaskStatus(t *testing.T) {
	cases := []struct {
		name    string
		input   int
		want    Status
		wantErr bool
	}{
		{"pending", 0, Pending, false},
		{"working", 1, Working, false},
		{"done", 2, Done, false},
		{"out of range", 42, Pending, true},
		{"negative", -1, Pending, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseTaskStatus(c.input)
			if (err != nil) != c.wantErr {
				t.Fatalf("ParseTaskStatus(%d) error = %v, wantErr %v", c.input, err, c.wantErr)
			}
			if got != c.want {
				t.Errorf("ParseTaskStatus(%d) = %v, want %v", c.input, got, c.want)
			}
		})
	}
}

func TestParseTaskStatus2(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    Status
		wantErr bool
	}{
		{"pending", "Pending", Pending, false},
		{"working", "Working", Working, false},
		{"done", "Done", Done, false},
		{"unrecognized string", "InProgress", Pending, true},
		{"empty string", "", Pending, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseTaskStatus2(c.input)
			if (err != nil) != c.wantErr {
				t.Fatalf("ParseTaskStatus2(%q) error = %v, wantErr %v", c.input, err, c.wantErr)
			}
			if got != c.want {
				t.Errorf("ParseTaskStatus2(%q) = %v, want %v", c.input, got, c.want)
			}
			if c.wantErr && err != nil {
				// The error message should contain the offending value, not a
				// mangled fmt verb (this used to be "%d" for a string argument).
				wantMsg := "invalid status value: " + c.input
				if err.Error() != wantMsg {
					t.Errorf("ParseTaskStatus2(%q) error message = %q, want %q", c.input, err.Error(), wantMsg)
				}
			}
		})
	}
}
