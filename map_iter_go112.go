//go:build go1.12
// +build go1.12

package luar

import lua "github.com/epikur-io/gopher-lua"

func mapCall(L *lua.LState) int {
	ref, _ := check(L, 1)

	iter := ref.MapRange()
	exhausted := false
	fn := func(L *lua.LState) int {
		if exhausted || !iter.Next() {
			exhausted = true
			return 0
		}
		L.Push(New(L, iter.Key().Interface()))
		L.Push(New(L, iter.Value().Interface()))
		return 2
	}
	L.Push(L.NewFunction(fn))
	return 1
}
