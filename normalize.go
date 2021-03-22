package normalize

var defaultNormalizers = []Option{
	WithRemoveSpecialChars(),
	WithFixRareCyrillicChars(),
	WithCyrillicToLatinLookAlike(),
	WithUmlautToLatinLookAlike(),
	WithLowerCase(),
}

// Normalize returns normalized string.
// If not normalizers specified default set of normalizers is used.
func Normalize(str string, normalizers ...Option) string {
	if len(normalizers) == 0 {
		normalizers = defaultNormalizers
	}
	result := str
	for _, normalizer := range normalizers {
		result = normalizer(result)
	}
	return result
}

// Many normalizes slice of strings returning new slice with normalized elements.
func Many(strings []string, normalizers ...Option) []string {
	normalizedStrings := make([]string, 0, len(strings))
	for _, str := range strings {
		normalizedStrings = append(normalizedStrings, Normalize(str, normalizers...))
	}
	return normalizedStrings
}
