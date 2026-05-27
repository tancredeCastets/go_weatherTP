package weather

func FilterByCountry(stations []Station, iso string) []Station {
	result := make([]Station, 0)
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

func AvgTemperature(s Station) float64 {
	total := 0.0
	if len(s.Observation) == 0 {
		return 0
	}
	for _, o := range s.Observation {
		total += o.Temperature
	}

	total = total / float64(len(s.Observation))
	return total
}

func MaxWindGust(stations []Station) (Station, float64) {
	maxstation := Station{}
	maxspeedwind := 0.0
	for _, s := range stations {
		for _, o := range s.Observation {
			if maxspeedwind < o.Wind.Speed {
				maxspeedwind = o.Wind.Speed
				maxstation = s
			}
		}
	}
	return maxstation, maxspeedwind
}

func CountByCountry(stations []Station) map[string]int {
	result := make(map[string]int)
	for _, s := range stations {
		result[s.Country]++
	}
	return result
}
