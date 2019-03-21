package bloomfilter

import (
	"math/rand"
	"testing"
	)

func TestBKDRHashSeed(t *testing.T) {
	testData := [][]int64{
		{0, 13},
		{1, 131},
		{2, 1313},
		{17, 1313131313131313131},
	}
	for _, value := range testData {
		if value[1] != BKDRHashSeed(value[0]) {
			t.Errorf("Test data %v-%v failed", value[0], value[1])
			t.Fatal("TestBKDRHashSeed error.")
		}
	}
}

func TestBKDRHash(t *testing.T) {
	testData := []string{}
	for i := 0; i < 100; i++ {
		testData = append(testData, string(StrRand(rand.Intn(1000))))
	}
	for _, value := range testData {
		BKDRHash(value, rand.Int63n(9))
	}
}
