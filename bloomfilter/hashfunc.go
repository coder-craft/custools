package bloomfilter

import (
	"math/rand"
	"time"
)

//1313131313131313131
func BKDRHashSeed(index int64) int64 {
	var seed = int64(13)
	for i := int64(0); i < index; i++ {
		if i%2 == 0 {
			seed = seed*10 + 1
		} else {
			seed = seed*10 + 3
		}
	}
	return seed
}
func BKDRHash(data string, index int64) int {
	var hash = int64(1)
	var seed = int64(BKDRHashSeed(index))
	for i := 0; i < len(data); i++ {
		hash = hash*seed + int64(data[i])
	}
	return int(hash % 36919)
}
func StrRand(size int) []byte {
	kinds, result := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		ikind := rand.Intn(3)
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
