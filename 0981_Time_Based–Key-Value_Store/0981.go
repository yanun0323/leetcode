package main

type TimeMap struct {
	hash map[string][]Time
}

func Constructor() TimeMap {
	return TimeMap{
		hash: make(map[string][]Time, 0),
	}
}

type Time struct {
	value     string
	timestamp int
}

func NewTime(value string, timestamp int) Time {
	return Time{
		value:     value,
		timestamp: timestamp,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.hash[key] == nil {
		this.hash[key] = make([]Time, 0, 10)
	}
	this.hash[key] = append(this.hash[key], NewTime(value, timestamp))
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.hash[key]) == 0 {
		return ""
	}
	if len(this.hash[key]) == 1 {
		return this.hash[key][0].value
	}

	lp := 0
	rp := len(this.hash[key]) - 1
	for lp < rp {
		mid := (lp + rp + 1) / 2
		if this.hash[key][mid].timestamp > timestamp {
			rp = mid - 1
		} else {
			lp = mid
		}
	}

	if this.hash[key][lp].timestamp <= timestamp {
		return this.hash[key][lp].value
	} else {
		return ""
	}

}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
