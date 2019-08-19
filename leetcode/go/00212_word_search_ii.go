package main

type TrieNode struct {
	next [26]*TrieNode
	word string
}

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	if m == 0 {
		return []string{}
	}
	n := len(board[0])
	if n == 0 {
		return []string{}
	}
	results := []string{}

	trie := buildTrie(words)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, trie, &results)
		}
	}
	return results
}

func dfs(board [][]byte, i, j int, trie *TrieNode, results *[]string) {
	cell := board[i][j]
	if cell == '#' || trie.next[int(cell-'a')] == nil {
		return
	}

	trie = trie.next[int(cell-'a')]
	if len(trie.word) > 0 {
		*results = append(*results, trie.word)
		trie.word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, trie, results)
	}

	if i < len(board)-1 {
		dfs(board, i+1, j, trie, results)
	}

	if j > 0 {
		dfs(board, i, j-1, trie, results)
	}

	if j < len(board[0])-1 {
		dfs(board, i, j+1, trie, results)
	}

	board[i][j] = cell
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{}
	for _, word := range words {
		// copy of reference
		cur := root
		for _, char := range word {
			cIdx := int(char - 'a') // int a-z = 97-123
			if cur.next[cIdx] == nil {
				cur.next[cIdx] = &TrieNode{}
			}
			cur = cur.next[cIdx]
		}
		cur.word = word
	}
	return root
}

// Algorithm:
// buildTrie
// loop through matrix
// - for each cell, dfs(board, i, j, trie, results) // results = *[]string
//      - dfs
