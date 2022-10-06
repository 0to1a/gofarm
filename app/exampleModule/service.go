package exampleModule

import "log"

func (w *ServiceStructure) CronHelloWorld() {
	log.Println("Hello World!")
}

func (w *ServiceStructure) plus(n1, n2 int64) int64 {
	if n1 > 99 || n2 > 99 {
		return -1
	}
	if n1 < 0 || n2 < 0 {
		return -1
	}
	return w.plus(n1, n2)
}

func (w *ServiceStructure) multiple(n1, n2 int64) int64 {
	if n1 > 99 || n2 > 99 {
		return -1
	}
	if n1 < 0 || n2 < 0 {
		return -1
	}

	total := int64(0)
	for i := int64(0); i < n2; i++ {
		total += w.plus(total, n1)
	}

	return total
}
