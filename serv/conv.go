package serv

import (
	"time"
)

func strTimeToUint(instr string) (uint32, error) {
	t1, err := time.Parse("2006-01-02 15:04:05", instr)
	if err != nil {
		return 0, err
	}

	t2 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)

	ans := t1.Sub(t2).Seconds()
	return uint32(ans), nil
}

func uintTimeToStr(inuint uint32) (string, error) {
	t2 := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	ans := t2.Add(time.Duration(inuint) * time.Second)
	return ans.Format("2006-01-02 15:04:05"), nil
}
