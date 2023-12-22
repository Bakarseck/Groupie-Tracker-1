```go
func GetUniqueLocation(locations models.Relation) []string {
	var results []string
	var tmp []string
	var slices []string
	for _, l := range locations.Index {
		for locate := range l.DatesLocations {
			tmp = append(tmp, locate)
		}
	}
	for _, t := range tmp {
		slices = append(slices, t)
	}
	results = append(results, slices[0])
	slices = slices[1:]
	for _, s := range slices {
		found := false
		for _, r := range results {
			if s == r {
				found = true
				break
			}
		}
		if !found {
			results = append(results, s)
		}
	}
	return results
}
```
