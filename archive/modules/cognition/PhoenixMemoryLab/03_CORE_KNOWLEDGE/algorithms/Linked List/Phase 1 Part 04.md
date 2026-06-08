# Phase 1 Part 04

        list.traverse();  // 10 -> 20 -> 30 -> null
        
        list.deleteAtPosition(1);
        list.traverse();  // 10 -> 30 -> null
        
        System.out.println("\n=== PRINT REVERSE ===");
        list.printReverse();  // 30 -> 10 -> null
    }
}
```

#### Output:

```
=== INSERT OPERATIONS ===
10 -> 20 -> 40 -> null
10 -> 20 -> 30 -> 40 -> null
5 -> 10 -> 20 -> 30 -> 40 -> null

=== SEARCH & LENGTH ===
Search 30: true
Search 99: false
Length: 5

=== MIDDLE & COUNT ===
Middle: 20
Count 20: 1

=== DELETION ===
10 -> 20 -> 30 -> 40 -> null
10 -> 20 -> 30 -> null
10 -> 30 -> null

=== PRINT REVERSE ===
30 -> 10 -> null
```

---

### **YOUR PRACTICE PLAN**

#### **Day 1-2: Master Traversal & Insertion**

- [ ]  Implement Node class
- [ ]  Implement LinkedList class with head
- [ ]  Implement traverse()
- [ ]  Implement insertAtHead()
- [ ]  Implement insertAtTail()
- [ ]  Implement insertAtPosition()
- [ ]  Test by creating list: 10 -> 20 -> 30 -> null

#### **Day 3: Master Deletion**

- [ ]  Implement deleteHead()
- [ ]  Implement deleteTail()
- [ ]  Implement deleteAtPosition()
- [ ]  Test all deletion operations

#### **Day 4: Master Helper Operations**

- [ ]  Implement search()
- [ ]  Implement findLength()
- [ ]  Implement printMiddle()
- [ ]  Implement countOccurrences()

#### **Day 5: Master Advanced Operations**

- [ ]  Implement deleteDuplicates()
- [ ]  Implement printReverse()
- [ ]  Test edge cases (empty list, single node, two nodes)

#### **Day 6-7: Dry Run Practice**

- [ ]  Manually dry run each operation on paper
- [ ]  Draw pointers for each step
- [ ]  Identify what happens to head and pointers

#### **Day 8: Consolidation**

- [ ]  Solve all 14 questions without help
- [ ]  Explain each operation to someone
- [ ]  Run all code examples

---

### **CHECKPOINT: Are You Ready for Phase 2?**

You can move to Phase 2 (Reversal) only if you can:

✅ Implement singly linked list from scratch ✅ Explain pointer movement for each operation ✅ Handle all edge cases (empty, single, two nodes) ✅ Solve insertion/deletion without looking at code ✅ Do dry runs manually on paper ✅ Identify common mistakes immediately

**If you can do all 6 things confidently → Ready for Phase 2! 🎯**
