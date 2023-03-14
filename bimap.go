package bimap

import "fmt"

// A BidirectionalMap is a map-like data structure that can be
// searched either by keys or by values (bidirectional).
type BidirectionalMap struct {
	m1 map[string]string
	m2 map[string]string
}

// insert inserts a key/value pair into the BidirectionalMap.
func (m *BidirectionalMap) insert(key, value string) {
	m.m1[key], m.m2[value] = value, key
}

// findByKey searches for a string in a BidirectionalMap based on map keys.
//
// findByKey returns the value if the key was found, or an empty
// string if it wasn't.
func (m BidirectionalMap) findByKey(key string) string {
	if val, found := m.m1[key]; found {
		return val
	}
	return ""
}

// findByValue searches for a string in a BidirectionalMap based on map values.
//
// findByValue returns the key if the value was found, or an empty
// string if it wasn't.
func (m BidirectionalMap) findByValue(value string) string {
	if key, found := m.m2[value]; found {
		return key
	}
	return ""
}

// find checks whether a string exists either as a key or a value inside of a map.
// It first checks if a key exists and if not, then searches for a value.
//
// find returns either the key or value if the string was found,
// or an empty string if neither were found.
func (m BidirectionalMap) find(str string) string {
	if val, found := m.m1[str]; found {
		return val
	}

	if key, found := m.m2[str]; found {
		return key
	}

	return ""
}

// newBDMap creates a new BidirectionalMap with optionally pre-populated data from a map.
func newBDMap(from map[string]string) *BidirectionalMap {
	// Create a new BidirectionalMap and initialize maps m1 and m2.
	bm := &BidirectionalMap{
		m1: make(map[string]string),
		m2: make(map[string]string),
	}

	// If data was passed in, iterate over the map and insert it's data into bm.
	if from != nil {
		for key, val := range from {
			bm.insert(key, val)
		}
	}

	// return BidirectionalMap bm.
	return bm
}

func main() {
	// Create some data to store in a BidirectionalMap.
	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}

	// Create a new BidirectionalMap bm with some key/value pair data from map m.
	bm := newBDMap(m)

	// Could also have created bm using the following:
	// m := &BidirectionalMap{
	// 	m1: map[string]string{
	// 		"one":   "1",
	// 		"two":   "2",
	// 		"three": "3",
	// 	},
	// 	m2: map[string]string{
	// 		"1": "one",
	// 		"2": "two",
	// 		"3": "three",
	// 	},
	// }

	// Insert some more key/value pairs into bm.
	bm.insert("four", "4")
	bm.insert("five", "5")

	// Find in bm by key.
	fmt.Println(bm.findByKey("one"))
	fmt.Println(bm.findByKey("ten"))

	// Find in bm by value.
	fmt.Println(bm.findByValue("1"))
	fmt.Println(bm.findByValue("10"))

	// Generic find in bm. Searches using keys and values.
	fmt.Println(bm.find("five"))
	fmt.Println(bm.find("5"))
	fmt.Println(bm.find("twenty"))
	fmt.Println(bm.find("20"))
}
