package normalize_test

import (
	"testing"

	"github.com/avito-tech/normalize"
)

func Test_areStringsSimilar(t *testing.T) {
	tests := []struct {
		name        string
		one         string
		other       string
		threshold   float64
		normalizers []normalize.Option
		want        bool
	}{
		{
			name:  "same_strings_must_be_similar",
			one:   "hello",
			other: "hello",
			want:  true,
		},
		{
			name:  "non_normalized_strings_distance",
			one:   "hella",
			other: "hello",
			want:  false,
		},
		{
			name:      "non_normalized_strings_distance_with_threshold",
			one:       "hella",
			other:     "hello",
			threshold: 0.25,
			want:      true,
		},
		{
			name:      "non_normalized_strings_distance_with_threshold",
			one:       "hela",
			other:     "hello",
			threshold: 0.25,
			want:      false,
		},
		{
			name:      "non_normalized_stringds_of_different_length",
			one:       "hell",
			other:     "hello",
			threshold: 0.34,
			want:      true,
		},
		{
			name:      "non_normalized_stringds_of_different_length_flipped",
			one:       "hello",
			other:     "hell",
			threshold: 0.34,
			want:      true,
		},
		{
			name:      "non_normalized_strings_of_different_length_flipped",
			one:       "hello",
			other:     "hell",
			threshold: 0.34,
			want:      true,
		},
		{
			name:  "normalized_strings",
			one:   "A b",
			other: "АВ", // all cyrillic
			want:  true,
		},
		{
			name:      "normalized_strings_with_threshold",
			one:       "AB-test",
			other:     "АВ тест", // all cyrillic
			threshold: 0.17,
			want:      true,
		},
		{
			name:        "normalized_strings_with_custom_options",
			one:         "AB",
			other:       "АВ",                                          // all cyrillic
			normalizers: []normalize.Option{normalize.WithLowerCase()}, // no cyr2lat
			want:        false,
		},
		{
			name:        "normalized_strings_with_custom_options",
			one:         "AB",
			other:       "АВ", // all cyrillic
			normalizers: []normalize.Option{normalize.WithLowerCase(), normalize.WithCyrillicToLatinLookAlike()},
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize.AreStringsSimilar(tt.one, tt.other, tt.threshold, tt.normalizers...); got != tt.want {
				t.Errorf("AreStringsSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
