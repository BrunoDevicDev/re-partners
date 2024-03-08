package algorithm

type val struct {
	numBoxes int
	M        map[int]int
}

// Solve is the main function to solve the algorithm
func Solve(target int, packSizes []int) (int, val) {
	arr := make([]val, target+findMax(packSizes))

	for _, b := range packSizes {
		arr[b] = val{
			numBoxes: 1,
			M: map[int]int{
				b: 1,
			},
		}

		for i, v := range arr {
			if v.numBoxes == 0 {
				continue
			}
			i2 := i + b

			m2 := CopyMap(v.M)
			m2[b]++
			v2 := val{
				numBoxes: v.numBoxes + 1,
				M:        m2,
			}
			if i2 >= len(arr) {
				continue
			}
			if arr[i2].numBoxes == 0 || v2.numBoxes < arr[i2].numBoxes {
				arr[i2] = v2
			}
		}
	}

	for i := target; i < len(arr); i++ {
		if arr[i].numBoxes > 0 {
			return target, arr[i]
		}
	}

	return 0, val{}
}

func CopyMap[K comparable, V any](original map[K]V) map[K]V {
	m := make(map[K]V, len(original))
	for k, v := range original {
		m[k] = v
	}
	return m
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
