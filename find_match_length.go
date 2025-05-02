package brotli

import (
	"encoding/binary"
	"math/bits"

	"github.com/asvedr/brotli/matchfinder"
)

/* Copyright 2010 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

func findMatchLengthWithLimit32(s1 []byte, s2 []byte, limit uint) uint {
	var matched uint
	for matched+4 <= limit {
		w1 := binary.LittleEndian.Uint32(s1[matched:])
		w2 := binary.LittleEndian.Uint32(s2[matched:])
		if w1 != w2 {
			return matched + uint(bits.TrailingZeros32(w1^w2)>>3)
		}
		matched += 4
	}
	return matched
}

func findMatchLengthWithLimit64(s1 []byte, s2 []byte, limit uint) uint {
	var matched uint
	for matched+8 <= limit {
		w1 := binary.LittleEndian.Uint64(s1[matched:])
		w2 := binary.LittleEndian.Uint64(s2[matched:])
		if w1 != w2 {
			return matched + uint(bits.TrailingZeros64(w1^w2)>>3)
		}
		matched += 8
	}
	return matched
}

var findMatchLengthWithLimitPrecheck = func() func([]byte, []byte, uint) uint {
	is32, is64 := matchfinder.CheckArch()
	switch {
	case is64:
		return findMatchLengthWithLimit64
	case is32:
		return findMatchLengthWithLimit32
	default:
		return func([]byte, []byte, uint) uint { return 0 }
	}
}()

/* Function to find maximal matching prefixes of strings. */
func findMatchLengthWithLimit(s1 []byte, s2 []byte, limit uint) uint {
	_, _ = s1[limit-1], s2[limit-1] // bounds check
	matched := findMatchLengthWithLimitPrecheck(s1, s2, limit)
	for matched < limit && s1[matched] == s2[matched] {
		matched++
	}
	return matched
}
