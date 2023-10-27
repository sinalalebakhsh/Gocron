package features

type AnswersSampleStruct struct {
	MapAnswersSample map[string]string
}

var AnswersSample = AnswersSampleStruct {
	MapAnswersSample: map[string]string {
		"operator 1": `// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	last := float64(productionRate) * successRate / 100
	return last
}`,
	},
}