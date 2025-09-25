package unpack_test

import (
	"testing"
	"wbtasksl2/9/unpack"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "basic unpacking with digits",
			input:   "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "string without digits",
			input:   "abcd",
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "only digits - invalid",
			input:   "45",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "single character",
			input:   "a",
			want:    "a",
			wantErr: false,
		},
		{
			name:    "character with repetition",
			input:   "a5",
			want:    "aaaaa",
			wantErr: false,
		},
		{
			name:    "multiple digits after character",
			input:   "a10",
			want:    "aaaaaaaaaa",
			wantErr: false,
		},
		{
			name:    "starts with digit - invalid",
			input:   "3a",
			want:    "",
			wantErr: true,
		},
		{
			name:    "consecutive digits",
			input:   "a11b",
			want:    "aaaaaaaaaaab",
			wantErr: false,
		},
		{
			name:    "special characters",
			input:   "!2@3#1",
			want:    "!!@@@#",
			wantErr: false,
		},
		{
			name:    "unicode characters",
			input:   "я3ß2",
			want:    "яяяßß",
			wantErr: false,
		},
		{
			name:    "escape sequences",
			input:   `\`,
			want:    ``,
			wantErr: false,
		},
		{
			name:    "digit as character with escape",
			input:   `\42`,
			want:    `44`,
			wantErr: false,
		},
		{
			name:    "escape backslash",
			input:   `\\`,
			want:    `\`,
			wantErr: false,
		},
		{
			name:    "escaped backslash with repetition",
			input:   `\\3`,
			want:    `\\\`,
			wantErr: false,
		},
		{
			name:    "escaped digit",
			input:   `\4`,
			want:    `4`,
			wantErr: false,
		},
		{
			name:    "escaped digit with repetition",
			input:   `\42`,
			want:    `44`,
			wantErr: false,
		},
		{
			name:    "mixed escaped and normal characters",
			input:   `a\3b2\\c`,
			want:    `a3bb\c`,
			wantErr: false,
		},
		{
			name:    "multiple escaped digits",
			input:   `\4\5\6`,
			want:    `456`,
			wantErr: false,
		},
		{
			name:    "escaped digit followed by normal digit",
			input:   `a\23`,
			want:    `a222`,
			wantErr: false,
		},
		{
			name:    "complex escape sequence",
			input:   `\\2\33a\52`,
			want:    `\\333a55`,
			wantErr: false,
		},
		{
			name:    "escape at end of string",
			input:   `abc\`,
			want:    `abc`,
			wantErr: false,
		},
		{
			name:    "only escaped characters",
			input:   `\4\5\6`,
			want:    `456`,
			wantErr: false,
		},
		{
			name:    "escaped character with large repetition",
			input:   `\510`,
			want:    `5555555555`,
			wantErr: false,
		},
		{
			name:    "invalid escape sequence - ends with backslash",
			input:   `a\`,
			want:    `a`,
			wantErr: false,
		},
		{
			name:    "double escape",
			input:   `\\\\`,
			want:    `\\`,
			wantErr: false,
		},
		{
			name:    "double escape with repetition",
			input:   `\\\\2`,
			want:    `\\\`,
			wantErr: false,
		},
		{
			name:    "escaped letter (should be treated as normal)",
			input:   `\a2`,
			want:    `aa`,
			wantErr: false,
		},
		{
			name:    "mixed escapes and special characters",
			input:   `\@3\!2\n`,
			want:    `@@@!!n`,
			wantErr: false,
		},
		{
			name:    "escape before EOF",
			input:   `test\`,
			want:    `test`,
			wantErr: false,
		},
		{
			name:    "consecutive escapes",
			input:   `\\\3\42`,
			want:    `\344`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := unpack.UnpackString(tt.input)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UnpackString() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnpackString() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("UnpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
