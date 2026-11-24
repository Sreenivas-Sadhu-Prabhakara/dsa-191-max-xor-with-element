package main

import (
	"fmt"
	"sort"
)

const topBit = 31

type bitTrie struct {
	children [2]*bitTrie
}

func (t *bitTrie) insert(v int) {
	n := t
	for b := topBit; b >= 0; b-- {
		bit := (v >> b) & 1
		if n.children[bit] == nil {
			n.children[bit] = &bitTrie{}
		}
		n = n.children[bit]
	}
}

func (t *bitTrie) bestXor(v int) int {
	n := t
	result := 0
	for b := topBit; b >= 0; b-- {
		bit := (v >> b) & 1
		if n.children[1-bit] != nil {
			result |= 1 << b
			n = n.children[1-bit]
		} else {
			n = n.children[bit]
		}
	}
	return result
}

// Query asks: max of x XOR e over elements e ≤ m (-1 when none qualifies).
type Query struct {
	X, M int
}

// MaxXorQueries answers all queries offline: sort elements and queries by
// the bound, insert elements as they become admissible, then trie-query.
func MaxXorQueries(nums []int, queries []Query) []int {
	sorted := append([]int{}, nums...)
	sort.Ints(sorted)
	idx := make([]int, len(queries))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return queries[idx[a]].M < queries[idx[b]].M })

	out := make([]int, len(queries))
	root := &bitTrie{}
	inserted := 0
	for _, qi := range idx {
		q := queries[qi]
		for inserted < len(sorted) && sorted[inserted] <= q.M {
			root.insert(sorted[inserted])
			inserted++
		}
		if inserted == 0 {
			out[qi] = -1
		} else {
			out[qi] = root.bestXor(q.X)
		}
	}
	return out
}

func main() {
	fmt.Println(MaxXorQueries([]int{0, 1, 2, 3, 4}, []Query{{3, 1}, {1, 3}, {5, 6}}))
}
