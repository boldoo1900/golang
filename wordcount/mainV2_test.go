package main

import (
	"fmt"
	"testing"
)

// func byteToInt(chara byte) uint64 {
// 	if 'A' <= chara && chara <= 'Z' {
// 		return uint64(chara-'A') + 1
// 	} else if 'a' <= chara && chara <= 'z' {
// 		return uint64(chara-'a') + 1
// 	} else {
// 		return 0
// 	}
// }

func TestByteToInt(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want uint64
	}{
		{name: "case1", b: 65, want: 1},
		{name: "case2", b: 63, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteToInt(tt.b); got != tt.want {
				fmt.Println(tt.b, got)
				t.Errorf("byteToInt(%v) = %v, want %v", tt.b, got, tt.want)
			}
		})
	}
}
