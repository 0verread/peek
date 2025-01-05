package cout

import (
	"bytes"
	"os"
	"testing"
)

func TestStats(t *testing.T) {
	// Redirect stdout to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test case
	Stats(200, 150)

	// Restore stdout
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if len(output) == 0 {
		t.Error("Expected stats output, got empty string")
	}
}

func TestHeader(t *testing.T) {
	// Redirect stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test case
	Header("http://example.com", "GET")

	// Restore stdout
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if len(output) == 0 {
		t.Error("Expected header output, got empty string")
	}
}

func TestPrettyPrint(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid json object",
			input:   []byte(`{"key": "value"}`),
			wantErr: false,
		},
		{
			name:    "valid json array",
			input:   []byte(`[{"key": "value"}]`),
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   []byte(`invalid json`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrettyPrint(tt.input)

			// Restore stdout
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			if tt.wantErr && len(output) > 0 {
				t.Errorf("PrettyPrint() expected error, got output: %v", output)
			}
			if !tt.wantErr && len(output) == 0 {
				t.Error("PrettyPrint() expected output, got empty string")
			}
		})
	}
}
