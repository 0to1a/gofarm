package calculateModule

import "framework/app/structure"

func (w *serviceStructure) plus(n1, n2 int64) int64 {
	if n1 > 99 || n2 > 99 {
		return -1
	}
	if n1 < 0 || n2 < 0 {
		return -1
	}
	return calculatePlus(n1, n2)
}

func (w *serviceStructure) multiple(n1, n2 int64) int64 {
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

	total = int64(float64(total) * structure.SystemConf.SampleMultiplier)

	return total
}
