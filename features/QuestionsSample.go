package features

type QuestionsSampleStruct struct {
	MapQuestionsSample map[string]string
}


var QuestionsSample = QuestionsSampleStruct {
	MapQuestionsSample: map[string]string{
		"operator":`Sample:
CalculateWorkingCarsPerHour(1547, 90)
expect:
// => 1392.3`,
		"operator 1":`Sample:
CalculateWorkingCarsPerHour(1547, 90)
expect:
// => 1392.3`,
	},
}