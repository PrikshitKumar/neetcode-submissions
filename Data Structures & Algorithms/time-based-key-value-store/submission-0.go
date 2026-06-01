type Entry struct {
	timestamp int
	value string
}

type TimeMap struct {
	data map[string][]Entry
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	this.data[key] = append(this.data[key], Entry {
		timestamp: timestamp, 
		value: value,
	})

}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, exists := this.data[key]
	if !exists {
		return ""
	}

	left, right := 0, len(entries) - 1

	result := ""

	for left <= right {
		mid := left + (right - left) / 2
		if entries[mid].timestamp <= timestamp {
			result = entries[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
