package utils

import (
	"strconv"
	"sync"
)

// RedisMap ...
type RedisMap struct {
	sync.Mutex
	mapObj map[string]string
}

// NewRedisMap ...
func NewRedisMap() *RedisMap {
	return &RedisMap{
		mapObj: make(map[string]string),
	}
}

// Get ...
func (rm *RedisMap) Get(key string) (value string, ok bool) {
	rm.Lock()
	defer rm.Unlock()
	result, ok := rm.mapObj[key]

	return result, ok
}

// Set ...
func (rm *RedisMap) Set(key, value string) {
	rm.Lock()
	defer rm.Unlock()
	rm.mapObj[key] = value
}

// Delete ...
func (rm *RedisMap) Delete(key string) {
	rm.Lock()
	defer rm.Unlock()

	delete(rm.mapObj, key)
}

// Increment ...
func (rm *RedisMap) Increment(key string, inc string) (string, bool) {
	rm.Lock()
	defer rm.Unlock()

	value, hasVal := rm.mapObj[key]
	if hasVal {
		oldVal, err := strconv.Atoi(value)
		if err != nil {
			return "value is not an integer or out of range", false
		}

		incVal, err := strconv.Atoi(inc)
		if err != nil {
			return "value is not an integer or out of range", false
		}

		rm.mapObj[key] = strconv.Itoa(oldVal + incVal)
		return rm.mapObj[key], true
	}

	_, err := strconv.Atoi(inc)
	if err != nil {
		return "value is not an integer or out of range", false
	}

	rm.mapObj[key] = inc
	return inc, true
}
