// slice包主要包含一些切片的常用切官方库没有的操作
package slice

import "sort"

// 去除字符串切片重复和空元素。
func DistinctStr(in []string) (out []string) {
	sort.Strings(in)
	for i := 0; i < len(in); i++ {
		if (i > 0 && in[i-1] == in[i]) || len(in[i]) == 0 {
			continue
		}
		out = append(out, in[i])
	}
	return
}
