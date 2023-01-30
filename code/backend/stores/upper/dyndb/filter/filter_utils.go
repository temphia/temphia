package filter

import (
	"strconv"
	"strings"
)

type LData struct {
	lat      float32
	long     float32
	distance float32
}

func locationData(rawstr string) LData {
	points := strings.Split(rawstr, " ")

	return LData{
		lat:      getFloat(points[0]),
		long:     getFloat(points[1]),
		distance: getFloat(points[2]),
	}
}

func getFloat(fstr string) float32 {
	f, err := strconv.ParseFloat(fstr, 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}
