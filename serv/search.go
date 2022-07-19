package serv

import "math"

func Serch(key string) (RetTag, error){
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
	ans, err := TagToRetTag(TagData[ok])
	if err != nil {
		return RetTag{}, nil
	}
	return ans, nil 
}
