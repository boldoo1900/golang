package test

import (
	"server/utils"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"ping"}, want: "+PONG"},
		{name: "case2", arr: []string{"pinG"}, want: "+PONG"},
		{name: "case3", arr: []string{"ping", "test"}, want: "+test"},
		{name: "case4", arr: []string{"ping", "hello world"}, want: "+hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CommandPing(&tt.arr); got != tt.want {
				t.Errorf("CommandPing(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"get", "nonexisting"}, want: "$-1"}, // special value (nil)
		{name: "case2", arr: []string{"get", "mykey"}, want: "hello"},
		{name: "case3", arr: []string{"GeT", "mykey"}, want: "hello"},
	}

	mapObj := utils.NewRedisMap()
	mapObj.Set("mykey", "hello")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.name == "case2" || tt.name == "case3" {
				got = utils.CommandGet(&tt.arr, mapObj)[4:]
			} else {
				got = utils.CommandGet(&tt.arr, mapObj)
			}

			if got != tt.want {
				t.Errorf("CommandGet(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"set", "mykey", "hello"}, want: "+OK"},
		{name: "case3", arr: []string{"set", "mykey1", "hello1", "nx"}, want: "+OK"},     // set only if not exists
		{name: "case4", arr: []string{"set", "mykeyExist", "hello1", "xx"}, want: "+OK"}, // set only existing element
	}
	for _, tt := range tests {
		mapObj := utils.NewRedisMap()
		mapObj.Set("mykeyExist", "123")
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CommandSet(&tt.arr, mapObj); got != tt.want {
				t.Errorf("CommandSet(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})

	}
}

func TestDel(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"del", "test123"}, want: ":0"}, // (integer) 0
		{name: "case2", arr: []string{"DeL", "test123"}, want: ":0"},
		{name: "case3", arr: []string{"del", "key1", "key2", "key3"}, want: ":2"}, // only deleted value count (note: ":" redis integer symbol)
	}

	mapObj := utils.NewRedisMap()
	mapObj.Set("key1", "123")
	mapObj.Set("key2", "123")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CommandDel(&tt.arr, mapObj); got != tt.want {
				t.Errorf("CommandDel(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

func TestIncrby(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"incrby", "mykey", "5"}, want: ":15"}, //	(note: ":" redis integer)
		{name: "case2", arr: []string{"INCRBY", "mykey", "11"}, want: ":26"},
		{name: "case2", arr: []string{"INCRBY", "test", "1"}, want: ":1"},
	}

	mapObj := utils.NewRedisMap()
	mapObj.Set("mykey", "10")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CommandIncrby(&tt.arr, mapObj); got != tt.want {
				t.Errorf("CommandIncrby(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
