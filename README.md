# Maximum XOR With an Element From Array

> Striver SDE Sheet 191 · Day 27 Trie — the offline trick: sort queries by their bound, feed elements into the XOR trie as they qualify, answer in sorted order, return in original order. Go, stdlib only, tested.

## Problem

Each query (x, m) asks for the maximum of x XOR e over array elements e ≤ m, or -1 if nothing qualifies — the bounded version of #190, and the sheet's 191st and final problem.

## Approach

The bound m is the twist: the trie from #190 can't un-insert large elements per query. Going **offline** dissolves the problem: sort the elements, sort the queries by m, and sweep — before each query, insert every element that just became admissible (each element inserted exactly once across all queries), then run the standard opposite-bit descent. An index-sort lets the answers return in the original query order despite the processing order, which the tests pin. Offline processing — reordering queries so a cheap incremental structure suffices — is the final tool the sheet teaches, and a fitting one to end on.

**Complexity:** O((n + q) log(n + q) + (n + q)·32) time.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
