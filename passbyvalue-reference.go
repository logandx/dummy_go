package main

type Example struct {
	count int
}

func createExample() Example {
	return Example{
		count: 1,
	}
}

func (e *Example) changeCount(countVal int) int {
	(*e).count = countVal
	countVal = 2
	return countVal
}
