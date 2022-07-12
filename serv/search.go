package serv

import (
	"fmt"
	"math"
	"strconv"
)

func Serch(key string) ImageList {
	fmt.Println("serch: start to setch '" + key + "'")
	// find tag mathed
	matchedTags := []Tag{}
	for _, row := range tags {
		if row.name == key {
			matchedTags = append(matchedTags, row)
		}
	}
	fmt.Println("serch: math id(s) count is " + strconv.Itoa(len(matchedTags)))

	// serch image data
	ans := ImageList{}
	for _, row := range matchedTags {
		ok := len(imgs)
		ng := -1
		for math.Abs(float64(ok-ng)) > 1 {
			mid := (ok + ng) / 2
			if imgs[mid].id < row.id {
				ng = mid
			} else {
				ok = mid
			}
		}

		if 0 <= ok && ok < len(imgs) {
			ans.images = append(ans.images, imgs[ok])
		}
			
	}

	fmt.Println("serch: serch is done. count is " + strconv.Itoa(len(ans.images)))

	return ans
}
