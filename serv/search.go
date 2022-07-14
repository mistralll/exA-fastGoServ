package serv

import (
	"fmt"
	"strconv"
)

func Serch(key string) []Image {
	fmt.Println(key)
	ans := []Image{}
	for _, row := range TagData {
		fmt.Print(row.Name + " ")
		if row.Name == key {
			ans = row.Imgs
			break
		}
	}

	fmt.Println(strconv.Itoa(len(ans)))

	return ans
}
