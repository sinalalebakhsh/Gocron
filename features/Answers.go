package features

type AllAnswers struct {
	SingleAnws map[string]string
}

var Answers = AllAnswers{
	SingleAnws: map[string]string{
		"operator 1": "Hints and Tips Calculate the number of working cars produced per hour The percentage (passed as an argument) is a number between 0-100. To make this percentage a bit easier to work with, start by dividing it by 100. To compute the number of cars produced successfully, multiply the percentage (divided by 100) by the number of cars produced. When multiplying two numbers together, they both need to be of the same type. Use type conversions if needed.",
	},
}
