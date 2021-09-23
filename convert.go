package utils

import (
	"fmt"
	"math"
)

var units = []string{"","Ki","Mi","Gi","Ti","Pi","Ei","Zi"}

// SizeOf from stack overflow
func SizeOf(size int64) string {
	s64 := float64(size)
	for _, unit := range units {
		if math.Abs(s64) < 1024 {
			return fmt.Sprintf("%3.1f%sB",s64, unit)
		}
		s64/=1024
	}
	return fmt.Sprintf("%.1fYiB",s64)
}
