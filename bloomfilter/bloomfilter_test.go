package bloomfilter

import "testing"

func TestNewBloomFilter(t *testing.T) {
	testData := [][]string{
		{"Student", "Teacher", "Math", "Logic"},
		{"U.S.A", "China", "Jap", "Russia", "England", "French", "Indin", "Canada"},
		{"laocuan", "laoxi", "laori", "laopu", "laomei", "shabi", "shadiao", "erhuo"},
	}
	for _, arr := range testData {
		bt := NewBloomFilter(0, 0)
		for i := 0; i < len(arr); i++ {
			bt.InsertElement(arr[i])
		}
		for i := 0; i < len(arr); i++ {
			if bt.ExistsElement(arr[i]) == false {
				t.Error("Test data %v not in %v", arr[i], arr)
				t.Fatal("TestNewBloomFilter fatal!")
			}
		}
		if bt.ExistsElement("bloomfilter") == true {
			//very low probability error that can be ignored,or check the code
			t.Fatal("TestNewBloomFilter error!")
		}
	}
}
