package serv

import "math"

func Search(key string) (int, error) {
	ok := len(TagData)
	ng := -1
	for math.Abs(float64(ok-ng)) > 1 {
		mid := (ok + ng) / 2
		if TagData[mid].Name < key {
			ng = mid
		} else {
			ok = mid
		}
	}
	return ok, nil
}
