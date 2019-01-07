package main

func (j *job) less(k *job) bool {
	if j.cost > k.cost {
		return true
	} else if j.cost == k.cost {
		return j.weight >= k.weight
	}
	return false
}

type job struct {
	weight int
	length int
	cost   float64
}
