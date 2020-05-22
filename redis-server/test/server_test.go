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
		{name: "case2", arr: []string{"set", "mykey", "hello"}, want: "+OK"},
		{name: "case3", arr: []string{"get", "mykey"}, want: "hello"},
	}
	for _, tt := range tests {
		if tt.name == "case2" {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandSet(&tt.arr); got != tt.want {
					t.Errorf("CommandSet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		} else {
			t.Run(tt.name, func(t *testing.T) {
				var got string
				if tt.name == "case3" {
					got = utils.CommandGet(&tt.arr)[4:]
				} else {
					got = utils.CommandGet(&tt.arr)
				}

				if got != tt.want {
					t.Errorf("CommandGet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"set", "mykey", "hello"}, want: "+OK"},
		{name: "case2", arr: []string{"get", "mykey"}, want: "hello"},
		{name: "case3", arr: []string{"set", "mykey1", "hello1", "nx"}, want: "+OK"}, // set only if not exists
		{name: "case4", arr: []string{"set", "mykey", "hello1", "xx"}, want: "+OK"},  // set only existing value
		{name: "case5", arr: []string{"get", "mykey"}, want: "hello1"},
	}
	for _, tt := range tests {
		if tt.name == "case2" || tt.name == "case5" {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandGet(&tt.arr); got[4:] != tt.want {
					t.Errorf("CommandGet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		} else {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandSet(&tt.arr); got != tt.want {
					t.Errorf("CommandSet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		}
	}
}

func TestDel(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"set", "key1", "hello"}, want: "+OK"},
		{name: "case2", arr: []string{"set", "key2", "world"}, want: "+OK"},
		{name: "case3", arr: []string{"set", "del", "key1", "key2", "key3"}, want: ":2"}, // only deleted value count (note: ":" redis integer symbol)
	}
	for _, tt := range tests {
		if tt.name == "case1" || tt.name == "case2" {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandSet(&tt.arr); got != tt.want {
					t.Errorf("CommandSet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		} else {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandDel(&tt.arr); got != tt.want {
					t.Errorf("CommandDel(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		}
	}
}

func TestIncrby(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{name: "case1", arr: []string{"set", "mykey", "10"}, want: "+OK"},
		{name: "case2", arr: []string{"incrby", "mykey", "5"}, want: ":15"}, //	(note: ":" redis integer symbol)
		{name: "case3", arr: []string{"get", "mykey"}, want: "15"},
	}
	for _, tt := range tests {
		if tt.name == "case1" {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandSet(&tt.arr); got != tt.want {
					t.Errorf("CommandSet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		} else if tt.name == "case3" {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandGet(&tt.arr); got[4:] != tt.want {
					t.Errorf("CommandGet(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		} else {
			t.Run(tt.name, func(t *testing.T) {
				if got := utils.CommandIncrby(&tt.arr); got != tt.want {
					t.Errorf("CommandIncrby(%v) = %v, want %v", tt.arr, got, tt.want)
				}
			})
		}
	}
}
