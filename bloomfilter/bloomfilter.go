package bloomfilter

import "encoding/json"

type BloomFilter struct {
	hashCount int   //defalut value 9
	bitSize   int64 //default value 36919
	table     []bool
}

func NewBloomFilter(bitSize int, hashCount int) *BloomFilter {
	if bitSize <= 0 {
		bitSize = 36919
	}
	if hashCount <= 0 || hashCount > 17 {
		hashCount = 3
	}
	return &BloomFilter{
		hashCount: hashCount,
		bitSize:   int64(bitSize),
		table:     make([]bool, bitSize, bitSize),
	}
}
func (this *BloomFilter) InsertElement(v interface{}) {
	var bitIndex int
	buff, _ := json.Marshal(v)
	for i := 0; i < this.hashCount; i++ {
		bitIndex = bKDRHash(string(buff), int64(i), this.bitSize)
		this.table[bitIndex] = true
	}
}
func (this *BloomFilter) ExistsElement(v interface{}) bool {
	var bitIndex int
	buff, _ := json.Marshal(v)
	for i := 0; i < this.hashCount; i++ {
		bitIndex = bKDRHash(string(buff), int64(i), this.bitSize)
		if this.table[bitIndex] == false {
			return false
		}
	}
	return true
}
