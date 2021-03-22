package normalize

import "github.com/agnivade/levenshtein"

// AreStringsSimilar returns true if relative distance between 2 strings after normalization is lower than provided threshold.
func AreStringsSimilar(one string, other string, threshold float64, normalizers ...Option) bool {
	one, other = Normalize(one, normalizers...), Normalize(other, normalizers...)
	greatestLen := greatest(len(one), len(other))
	distance := levenshtein.ComputeDistance(one, other)
	return float64(distance)/float64(greatestLen) <= threshold
}

func greatest(a, b int) int {
	if a > b {
		return a
	}
	return b
}
