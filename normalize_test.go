package normalize_test

import (
	"testing"

	"github.com/avito-tech/normalize"
)

func Test_Normalize(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		normalizers []normalize.Option
		want        string
	}{
		{
			name: "default_options_string_gets_lowercased",
			str:  "PART123",
			want: "part123",
		},
		{
			name: "default_options_lowercase_string_is_unchanged",
			str:  "part123",
			want: "part123",
		},
		{
			name: "default_options_spaces_are_trimmed",
			str:  " part123 ",
			want: "part123",
		},
		{
			name: "default_options_special_characters_and_spaces_are_removed",
			str:  "part - #123_50",
			want: "part12350",
		},
		{
			name: "default_options_cyrillic_are_not_removed",
			str:  "Часть-123",
			want: "чactь123",
		},
		{
			name: "default_options_cyrillic_in_string_is_converted_to_latin",
			str:  "Wewtаренбсхдумтокдлвгф",
			want: "wewtapehbcxdymtokdlbgf",
		},
		{
			name: "default_options_ё_replaced_with_e_latin",
			str:  "ёжикЁжик",
			want: "eжиkeжиk",
		},
		{
			name: "default_options_й_replaced_with_и",
			str:  "йодЙод",
			want: "иodиod",
		},
		{
			name: "default_options_umlauts_converted_to_latin",
			str:  "äöüÄÖÜ",
			want: "aouaou",
		},
		{
			name:        "no_normalizers_options",
			str:         "Wewtаренбсхдумтокдлвгф",
			normalizers: []normalize.Option{normalize.WithNoNormalization()},
			want:        "Wewtаренбсхдумтокдлвгф",
		},
		{
			name:        "only_some_options",
			str:         "АВ-тест",
			normalizers: []normalize.Option{normalize.WithLowerCase(), normalize.WithCyrillicToLatinLookAlike()},
			want:        "ab-tect",
		},
		{
			name:        "only_cyr2lat_cyrillic_letters_are_fixed_with_keeping_case",
			str:         "АВС",
			normalizers: []normalize.Option{normalize.WithCyrillicToLatinLookAlike()},
			want:        "ABC",
		},
		{
			name:        "only_fix_rare_cyrillic_rare_cyrillic_letters_are_fixed_with_keeping_case",
			str:         "ЙЁ",
			normalizers: []normalize.Option{normalize.WithFixRareCyrillicChars()},
			want:        "ИЕ",
		},
		{
			name:        "only_fix_umlaut2lat_umlauts_converted_to_latin_keeping_case",
			str:         "äöüÄÖÜ",
			normalizers: []normalize.Option{normalize.WithUmlautToLatinLookAlike()},
			want:        "aouAOU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize.Normalize(tt.str, tt.normalizers...)
			if tt.want != got {
				t.Errorf("want = %s, got = %s", tt.want, got)
			}
		})
	}
}

func Test_NormalizeMany(t *testing.T) {
	tests := []struct {
		name    string
		manyStr []string
		want    []string
	}{
		{
			name:    "normalize_several_strings",
			manyStr: []string{"Part-1", "Часть 2", "WewtаренБсхдумтОк"},
			want:    []string{"part1", "чactь2", "wewtapehbcxdymtok"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize.Many(tt.manyStr)
			if len(tt.want) != len(got) {
				t.Errorf("got slices of different length, want = %s, got = %s", tt.want, got)
			}
			for i, wantStr := range tt.want {
				if wantStr != got[i] {
					t.Errorf("at pos %d want = %s, got = %s", i, tt.want, got)
				}
			}
		})
	}
}
