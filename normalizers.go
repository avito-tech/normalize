package normalize

import (
	"regexp"
	"strings"
	"unicode"
)

type Option func(string) string

var specialCharsPattern = regexp.MustCompile(`(?i:[^äöüa-zа-яё0-9])`)

// WithRemoveSpecialChars any char except latin/cyrillic letters, German umlauts (`ä`, `ö`, `ü`) and digits are removed
func WithRemoveSpecialChars() Option {
	return func(str string) string {
		return specialCharsPattern.ReplaceAllString(str, "")
	}
}

var rareCyrillicChars = withUpperPairs(map[rune]rune{
	'ё': 'е',
	'й': 'и',
})

// WithFixRareCyrillicChars rare cyrillic letters `ё` and `й` are replaced with  common equivalents `е` and `и`
func WithFixRareCyrillicChars() Option {
	return WithRuneMapping(rareCyrillicChars)
}

var cyrillicTolatinsLookAlike = withUpperPairs(map[rune]rune{
	'а': 'a',
	'е': 'e',
	'т': 't',
	'у': 'y',
	'о': 'o',
	'р': 'p',
	'н': 'h',
	'к': 'k',
	'х': 'x',
	'с': 'c',
	'б': 'b',
	'м': 'm',
	'д': 'd',
	'л': 'l',
	'в': 'b',
	'г': 'g',
	'ф': 'f',
})

// WithCyrillicToLatinLookAlike Latin/cyrillic look-alike pairs are normalized to latin letters so `В (в)` becomes `B (b)`, etc.
func WithCyrillicToLatinLookAlike() Option {
	return WithRuneMapping(cyrillicTolatinsLookAlike)
}

var umlautsToLatin = withUpperPairs(map[rune]rune{
	'ä': 'a',
	'ö': 'o',
	'ü': 'u',
})

// WithUmlautToLatinLookAlike german umlauts `ä`, `ö`, `ü` get converted to latin `a`, `o`, `u`
func WithUmlautToLatinLookAlike() Option {
	return WithRuneMapping(umlautsToLatin)
}

// WithRuneMapping configures arbitrary rune mapping, case sensitive
func WithRuneMapping(mapping map[rune]rune) Option {
	return func(str string) string {
		return strings.Map(func(letter rune) rune {
			if newLetter, ok := mapping[letter]; ok {
				return newLetter
			}
			return letter
		}, str)
	}
}

func withUpperPairs(m map[rune]rune) map[rune]rune {
	for from, to := range m {
		m[unicode.ToUpper(from)] = unicode.ToUpper(to)
	}
	return m
}

// WithLowerCase converts string to lowercase
func WithLowerCase() Option {
	return func(str string) string {
		return strings.ToLower(str)
	}
}

// WithNoNormalization applies no changes to string
func WithNoNormalization() Option {
	return func(str string) string {
		return str
	}
}
