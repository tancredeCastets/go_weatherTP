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
