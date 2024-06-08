package answer1

func WeightedStrings(s string, queries []int) []string {
	// Peta untuk menyimpan bobot substring
	weights := make(map[int]struct{})
	n := len(s)
	i := 0

	// Memproses string untuk menghitung bobot karakter dan substring
	for i < n {
		weight := int(s[i] - 'a' + 1)
		weights[weight] = struct{}{}
		// Menghitung substring dari karakter yang berulang
		for i+1 < n && s[i] == s[i+1] {
			weight += int(s[i] - 'a' + 1)
			weights[weight] = struct{}{}
			i++
		}
		i++
	}

	// Hasil untuk setiap query
	results := make([]string, len(queries))

	for j, query := range queries {
		if _, exists := weights[query]; exists {
			results[j] = "Yes"
		} else {
			results[j] = "No"
		}
	}

	return results
}
