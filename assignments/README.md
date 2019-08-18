# Programming assignments for practice

- [ ] [JSON Parser](#JSON-Parser)
- [x] [LRU Cache](#LRU-cache)
- [ ] [Differ](#Differ)

## JSON Parser

Things learned:
- JSON
- Lexer
- Parser
- Tokenizer
- Building a library

## LRU Cache

Least Recently Used Cache. Provide fast and efficient way of retrieving data

##### API for LRUCache:

`Constructor(capacity)`
`Get(key)`
`Put(key, val)`

##### Strategy:
- HashMap: keys, vals for O(1) lookup and deletion (put, get operations)
- DoublyLinkedList: nodes can remove themselves without reference to others

**Time complexity** : O(1) both for `put` and `get`.

**Space complexity** : O(_capacity_) since the space is used only for a hashmap and double linked list with at most `capacity + 1` elements

**Requirements:**
- Fixed Size: bounds to limit memory usages
- Fast Access: Cache Insert and lookup operation should be fast , preferably O(1) time.
- Replacement of Entry in case , Memory Limit is reached: A cache should have efficient algorithm to evict the entry when memory is full.

**Things learned:**
- Caching technique
- Queue (Deque)
- Doubly Linked List
    - Adding and removing nodes from DLL
- Hashmap

## Differ

Diff two files

LCS