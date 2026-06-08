# Question To Do Part 01

## COMPLETE LINKED LIST QUESTION LIST WITH DIFFICULTY & PATTERNS

---

### **A. BASICS (Foundation - Start Here!)**

#### Level: Easy

1. **Traverse List**
    - Pattern: Basic traversal
    - Concepts: Pointer movement, while loop
    - Time: O(n) | Space: O(1)
2. **Insert Node**
    - Pattern: Basic insertion
    - Concepts: Pointer rewiring
    - Variations: At head, at tail, at position
    - Time: O(1) head, O(n) tail | Space: O(1)
3. **Delete Node**
    - Pattern: Basic deletion
    - Concepts: Pointer reconnection
    - Variations: Delete head, delete tail, delete by value
    - Time: O(1) head, O(n) tail | Space: O(1)
4. **Search Node**
    - Pattern: Linear search
    - Concepts: Conditional check while traversing
    - Time: O(n) | Space: O(1)
5. **Find Length**
    - Pattern: Count while traversing
    - Concepts: Counter variable
    - Time: O(n) | Space: O(1)

**✓ Complete these without any help before moving forward!**

---

### **B. REVERSAL PROBLEMS (Most Important Pattern)**

#### Level: Medium

6. **Reverse Linked List** ⭐ (MUST MASTER)
    - Pattern: Iterative reversal with 3 pointers
    - Concepts: prev, curr, next pointer management
    - Edge Cases: Empty list, single node, cycle
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Very High
    - **Dry Run Essential!**
7. **Reverse Linked List Recursive**
    - Pattern: Recursive reversal
    - Concepts: Recursion, backtracking pointer
    - Key Line: `head.next.next = head`
    - Time: O(n) | Space: O(n) - recursion stack
    - **Interview Frequency:** Medium
8. **Reverse Nodes in K-Group** ⭐ (VERY HARD)
    - Pattern: Segmentation + reversal + reconnection
    - Concepts: Multi-pointer (groupPrev, kth, groupNext, prev, curr, next)
    - Edge Cases: List not multiple of k
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Very High
    - **Difficulty:** Hard
    - **Prerequisites:** Master basic reversal first!
9. **Reverse Linked List II** (MEDIUM-HARD)
    - Pattern: Partial reversal
    - Concepts: Find boundaries, reverse segment, reconnect
    - Edge Cases: Reverse from head, reverse to tail
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** High
10. **Reverse Alternate K Nodes** (MEDIUM-HARD)
    - Pattern: Selective reversal
    - Concepts: Reverse k nodes, skip k nodes, repeat
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium

**✓ After mastering reversal, you unlock 30% of LL problems!**

---

### **C. FAST-SLOW POINTER (Second Most Important Pattern)**

#### Level: Medium

11. **Middle of Linked List** ⭐ (MUST MASTER)
    - Pattern: Fast-slow pointer
    - Concepts: fast = 2 steps, slow = 1 step
    - Edge Cases: Even/odd length lists
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Very High
    - **Used In:** Merge sort, reorder list, palindrome check
12. **Linked List Cycle** ⭐ (MUST MASTER)
    - Pattern: Floyd Cycle Detection
    - Concepts: Meeting point detection
    - Edge Cases: No cycle, single node cycle
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Very High
    - **Key:** Check `slow == fast` (same node), not values
13. **Linked List Cycle II** (MEDIUM-HARD)
    - Pattern: Floyd + distance reasoning
    - Concepts: Find cycle start point
    - Edge Cases: Cycle at head, cycle later
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** High
    - **Key Insight:** Mathematical distance property
14. **Remove Nth Node From End** ⭐
    - Pattern: Fast-slow with gap
    - Concepts: Maintain k-node distance
    - Edge Cases: Remove head, single node
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Very High
15. **Find Kth Node From End**
    - Pattern: Fast-slow gap
    - Concepts: Move fast k steps first, then both move together
    - Edge Cases: k > length
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium
16. **Happy Number** (VARIANT)
    - Pattern: Cycle detection in numbers (not LL directly)
    - Concepts: Fast-slow on function iterations
    - Time: O(1) on average | Space: O(1)
    - **Interview Frequency:** Medium
17. **Palindrome Linked List** (MEDIUM)
    - Pattern: Combination of fast-slow + reversal
    - Concepts: Find middle, reverse second half, compare
    - Edge Cases: Single node, two nodes
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** High
    - **Tests:** Pattern combination ability

**✓ Fast-slow is the most powerful LL technique. Master thoroughly!**

---

### **D. MERGE PROBLEMS (Ordered Combining)**

#### Level: Medium

18. **Merge Two Sorted Lists** ⭐ (MUST MASTER)
    - Pattern: Two pointer merge with dummy node
    - Concepts: Compare and attach nodes
    - Edge Cases: Empty lists, single nodes
    - Time: O(n + m) | Space: O(1)
    - **Interview Frequency:** Very High
    - **Key:** Don't create new nodes, rewire existing ones
19. **Merge K Sorted Lists** (HARD)
    - Pattern: Min heap + merge
    - Concepts: Heap maintains smallest node
    - Edge Cases: Empty lists in array
    - Time: O(n log k) | Space: O(k)
    - **Interview Frequency:** Very High
    - **Difficulty:** Hard
    - **Tests:** Multi-DS integration
20. **Sort Linked List** (MEDIUM-HARD)
    - Pattern: Merge sort (not quick sort!)
    - Concepts: Find middle, split, recursively sort, merge
    - Edge Cases: Single node, two nodes
    - Time: O(n log n) | Space: O(log n) - recursion stack
    - **Interview Frequency:** High
    - **Key:** Why merge sort? No random access, efficient merging
    - **Tests:** Can you combine fast-slow + recursion + merge?

**✓ Merge pattern unlocks sorting and advanced problems!**

---

### **E. STRUCTURAL REARRANGEMENT (Complex Rewiring)**

#### Level: Medium to Hard

21. **Reorder List** (MEDIUM) ⭐
    - Pattern: Combination of fast-slow + reversal + merge
    - Concepts: Find middle, reverse second half, merge alternately
    - Edge Cases: Single node, two nodes
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** High
    - **Example:** 1->2->3->4->5 becomes 1->5->2->4->3
    - **Tests:** Can you combine 3 patterns?
22. **Odd-Even Linked List** (MEDIUM)
    - Pattern: Maintain two chains and connect
    - Concepts: Separate odd-position and even-position nodes
    - Edge Cases: Empty, single node, two nodes
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium-High
    - **Example:** 1->2->3->4->5 becomes 1->3->5->2->4
    - **Key:** Odd-even positions (1st, 3rd, 5th...), not odd-even values!
23. **Rotate List** (MEDIUM-HARD)
    - Pattern: Circular connection + split
    - Concepts: Calculate effective rotation, create circle, split
    - Edge Cases: k > length, rotation by 0
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium
    - **Example:** 1->2->3->4->5, k=2 becomes 4->5->1->2->3
24. **Partition List** (MEDIUM)
    - Pattern: Build two lists and connect
    - Concepts: Maintain smaller and larger chains
    - Edge Cases: All smaller, all larger
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium
    - **Example:** Partition by x=3: 1->4->2->5->2 becomes 1->2->2->4->5
25. **Swap Nodes in Pairs** (MEDIUM)
    - Pattern: Pairwise reversal/swapping
    - Concepts: Swap adjacent nodes without creating new nodes
    - Edge Cases: Odd-length list
    - Time: O(n) | Space: O(1) iterative, O(n) recursive
    - **Interview Frequency:** Medium-High

**✓ These test your ability to segment and manipulate structures!**

---

### **F. DELETION VARIANTS (Node Removal Patterns)**

#### Level: Easy to Medium

26. **Remove Nth Node From End** ⭐
    - (Already covered in Fast-Slow section - #14)
27. **Delete Duplicates from Sorted List** (EASY)
    - Pattern: Linear traversal with condition
    - Concepts: Skip duplicate nodes
    - Edge Cases: All duplicates, no duplicates
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium
    - **Variation:** What if list is unsorted? (Use HashMap)
28. **Delete Middle Node** (EASY-MEDIUM)
    - Pattern: Find middle and delete
    - Concepts: Use fast-slow or count
    - Edge Cases: Single node, two nodes
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Low-Medium
29. **Remove Elements** (EASY)
    - Pattern: Delete all nodes with specific value
    - Concepts: Handle head deletion specially
    - Edge Cases: All nodes match, no match, head is match
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Low

**✓ Deletion practice solidifies pointer safety!**

---

### **G. PALINDROME PROBLEMS**

#### Level: Medium

30. **Palindrome Linked List** ⭐
    - (Already covered in Fast-Slow section - #17)
    - Pattern: Fast-slow + reversal + comparison
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** High

---

### **H. RANDOM POINTER PROBLEMS (Deep Copy)**

#### Level: Hard

31. **Copy List with Random Pointer** (HARD) ⭐
    - Pattern: Deep copy with pointer preservation
    - Node has: `int val, Node next, Node random`
    - Approach 1: HashMap (easy, O(n) space)
    - Approach 2: Interleaving (optimal, O(1) space)
    - Edge Cases: Null list, single node, circular structure
    - Time: O(n) | Space: O(1) interleaving, O(n) hashmap
    - **Interview Frequency:** Very High
    - **Difficulty:** Hard
    - **Key Concept:** Deep copy vs shallow copy
    - **Interleaving Method:**
        - A -> A' -> B -> B' -> C -> C'
        - Assign random pointers
        - Separate the two lists

**✓ This tests true pointer understanding!**

---

### **I. DESIGN PROBLEMS (System-Level)**

#### Level: Hard

32. **LRU Cache** (HARD) ⭐⭐
    - Pattern: HashMap + Doubly Linked List
    - Concepts: O(1) get and put operations
    - Data Structures: HashMap (fast lookup) + DLL (ordering)
    - Edge Cases: Single capacity, repeated access
    - Time: O(1) get/put | Space: O(capacity)
    - **Interview Frequency:** Very High
    - **Difficulty:** Hard
    - **Company Favorites:** Google, Amazon, Meta
    - **Key Operations:**
        - get(): move to head
        - put(): add to head or update
        - evict(): remove from tail
    - **Critical Code:**

java

```java
      node.prev.next = node.next;
      node.next.prev = node.prev;
```

33. **LFU Cache** (VERY HARD)
    - Pattern: HashMap + Frequency buckets + DLL
    - Concepts: Frequency tracking
    - Data Structures: HashMap + HashMap of frequencies + DLL
    - Time: O(1) get/put | Space: O(capacity)
    - **Interview Frequency:** Very High
    - **Difficulty:** Very Hard
    - **Company Favorites:** Google, Meta
    - **Key:** Maintain frequency and evict least frequently used
34. **Browser History** (MEDIUM-HARD)
    - Pattern: Doubly Linked List + current pointer
    - Concepts: Forward/backward navigation
    - Operations: visit, back, forward
    - Time: O(1) all operations | Space: O(n)
    - **Interview Frequency:** Medium
    - **Variation:** Text editor cursor
35. **Music Playlist** (MEDIUM)
    - Pattern: Circular doubly linked list
    - Concepts: Continuous playback, prev/next song
    - Operations: next, prev, play
    - Time: O(1) | Space: O(n)
    - **Interview Frequency:** Low

**✓ Design problems test your systems thinking!**

---

### **J. CIRCULAR LINKED LIST**

#### Level: Medium to Hard

36. **Detect Circularity** (MEDIUM)
    - Pattern: Floyd cycle detection
    - Concepts: Same as Cycle Detection, but might be circular
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Medium
37. **Josephus Problem** (HARD)
    - Pattern: Circular elimination
    - Concepts: Simulate elimination in circular list
    - Time: O(n*k) or O(n) with math | Space: O(1)
    - **Interview Frequency:** Low-Medium
    - **Use Case:** Round-robin elimination
38. **Split Circular Linked List** (HARD)
    - Pattern: Break circle and split
    - Concepts: Find break point, reconnect
    - Time: O(n) | Space: O(1)
    - **Interview Frequency:** Low

**✓ Circular LL is less common but good for completeness!**

---

### **K. DOUBLY LINKED LIST**

#### Level: Medium

39. **Insert/Delete in Doubly Linked List** (MEDIUM)
    - Pattern: Maintain both prev and next
    - Concepts: Bidirectional pointer updates
    - Time: O(1) if position known, O(n) otherwise
    - **Interview Frequency:** Medium
    - **Complexity:** More pointer updates than singly LL
40. **LRU Cache Implementation** (MEDIUM-HARD)
    - (Same as #32 LRU Cache - uses DLL!)
    - Key difference: Uses DLL instead of singly LL
    - Why DLL? O(1) deletion when node is known

---

### **RECOMMENDED PRACTICE ORDER**

#### **Phase 1: Foundations** (Days 1-2)

```
1. Traverse list
2. Insert node
3. Delete node
4. Search node
5. Find length
```

**Goal:** Comfortable with basic pointer movement

---

#### **Phase 2: Reversal Mastery** (Days 3-5)

```
6. Reverse Linked List (iterative)
7. Reverse Linked List (recursive)
8. Reverse Linked List II
```

**Do NOT do k-group yet!**

**Goal:** Deep understanding of 3-pointer reversal

---

#### **Phase 3: Fast-Slow Pattern** (Days 6-8)

```
11. Middle of Linked List
12. Linked List Cycle
13. Linked List Cycle II
14. Remove Nth Node From End
15. Find Kth Node From End
```

**Goal:** Master fast-slow completely

---

#### **Phase 4: Pattern Combination** (Days 9-10)

```
17. Palindrome Linked List (combines fast-slow + reversal)
18. Sort List (combines fast-slow + merge)
19. Reorder List (combines fast-slow + reversal + merge)
```

**Goal:** Combine patterns confidently

---

#### **Phase 5: Advanced Reversal** (Days 11-12)

```
8. Reverse Nodes in K-Group ⭐⭐
9. Reverse Alternate K Nodes
```

**Goal:** Master complex segmentation + reversal

---

#### **Phase 6: Merge & Rearrangement** (Days 13-15)

```
18. Merge Two Sorted Lists
19. Merge K Sorted Lists
20. Odd-Even Linked List
21. Rotate List
22. Partition List
```

**Goal:** Comfortable with structural manipulation

---

#### **Phase 7: Complex Problems** (Days 16-18)

