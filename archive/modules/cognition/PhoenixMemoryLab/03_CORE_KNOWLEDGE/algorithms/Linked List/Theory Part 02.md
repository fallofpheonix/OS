# Theory Part 02


**Concept:** Use two pointers moving at different speeds

- **Slow:** moves 1 step per iteration
- **Fast:** moves 2 steps per iteration

**Why it works:** Fast pointer eventually meets slow in a cycle or reaches the end first

**Used for:**

- Finding middle node
- Detecting cycles
- Finding kth node from end

---

#### Pattern 3: DUMMY NODE

**Purpose:** Simplify handling of special cases

java

```java
Node dummy = new Node(-1);  // Fake node
Node temp = dummy;
// Build solution through temp
return dummy.next;  // Skip dummy, return actual head
```

**Advantages:**

- Avoids special case for head deletion/insertion
- Makes code cleaner
- Handles null head gracefully

---

#### Pattern 4: MERGE PATTERN

**Purpose:** Combine two sorted linked lists

```
1 -> 3 -> 5
2 -> 4 -> 6

Result:
1 -> 2 -> 3 -> 4 -> 5 -> 6
```

**Key Insight:** At each step, choose the smaller node

---

#### Pattern 5: SPLIT AND RECONNECT

**Purpose:** Break list into parts and rejoin differently

Used in:

- Merge sort (split list in half)
- Reorder list (split, reverse, merge)
- Odd-even rearrangement

**Critical operation:**

java

```java
Node second = slow.next;
slow.next = null;  // Physically disconnect
```

---

### 11. CRITICAL POINTER CONCEPTS

#### Concept 1: POINTER REWIRING (Most Important)

**Wrong way:**

java

```java
curr.next = prev;
curr = curr.next;  // ❌ curr.next is now pointing backward!
```

**Problem:** After changing `curr.next`, you lose the forward direction!

**Correct way:**

java

```java
Node next = curr.next;  // Save future first
curr.next = prev;       // Rewire
prev = curr;            // Move prev forward
curr = next;            // Move curr forward (using saved next)
```

---

#### Concept 2: LOSING NODES

**DANGER:** If you don't save pointers before modifying, nodes become unreachable!

java

```java
// WRONG - node at position 3 is lost!
temp.next = newNode;
// temp.next.next is now inaccessible

// RIGHT - save before modifying
Node savedNext = temp.next;
temp.next = newNode;
newNode.next = savedNext;
```

---

#### Concept 3: NULL POINTER EXCEPTION

**Always check before accessing next:**

java

```java
// ❌ WRONG - crashes if temp.next is null
temp.next.next;

// ✓ RIGHT - check first
while(temp != null && temp.next != null) {
    // Safe to access temp.next.next
}
```

---

#### Concept 4: NODE EQUALITY VS VALUE EQUALITY

java

```java
// ❌ WRONG - compares values, not nodes
if(slow.data == fast.data)

// ✓ RIGHT - compares if they're the same node object
if(slow == fast)
```

Two different nodes can have the same value, but we want to detect if they're literally the same node.

---

### 12. IMPORTANT EDGE CASES (Always Test!)

#### 1. Empty List

java

```java
if(head == null) {
    // Handle empty case
}
```

#### 2. Single Node

```
10 -> null
```

#### 3. Two Nodes (Critical for reversal bugs!)

```
10 -> 20 -> null
```

#### 4. Cycle Exists

```
1 -> 2 -> 3
     ^    |
     |____|
```

Without proper checking, you'll get infinite loop!

#### 5. Head Modification

If head changes during operation, return the new head!

#### 6. Tail Modification

Make sure you don't lose the tail reference.

#### 7. Duplicate Values

Multiple nodes with same data.

---

### 13. COMMON BEGINNER MISTAKES

#### Mistake 1: Losing the List

java

```java
// ❌ WRONG
curr.next = prev;
curr = curr.next;  // Lost forward link!
```

#### Mistake 2: Infinite Loop

java

```java
// ❌ WRONG - no pointer movement
while(temp != null){
    // Process
    // temp is never updated
}
```

#### Mistake 3: Not Handling Null

java

```java
// ❌ WRONG - crashes on last node
temp.next.next;

// ✓ RIGHT
if(temp.next != null) {
    temp.next.next;
}
```

#### Mistake 4: Confusing References

java

```java
// ❌ Thinks copying variable copies node
Node b = a;  // b and a point to same node!

// ✓ To actually copy a node
Node copy = new Node(original.data);
```

---

### 14. MOST CRITICAL INTERVIEW RULES

#### Rule 1: Always Save Next Before Modifying

java

```java
Node next = curr.next;
curr.next = prev;  // Now safe to modify
```

#### Rule 2: Never Lose Head

Track the new head after operations!

#### Rule 3: Test with Edge Cases

- Empty list
- Single node
- Two nodes
- Cycle exists

#### Rule 4: Use Dry Run

**Physically write out:**

- Every pointer
- Every step
- Every change

This catches bugs instantly!

#### Rule 5: Use Dummy Node When Possible

Simplifies head handling dramatically.

---

### 15. CORE LINKED LIST OPERATIONS SUMMARY

|Operation|Complexity|Notes|
|---|---|---|
|Traverse|O(n)|Must visit every node|
|Insert at head|O(1)|Fastest insertion|
|Insert at tail|O(n)|Need to find tail|
|Insert at position|O(n)|Need to find position|
|Delete head|O(1)|Fastest deletion|
|Delete at position|O(n)|Need to find position|
|Search|O(n)|Linear search|
|Access by index|O(n)|No random access|

---

### 16. WHEN TO USE WHAT

#### Use Singly Linked List When:

- Simple traversal needed
- Only forward direction required
- Memory is constrained
- Most common choice

#### Use Doubly Linked List When:

- Need backward traversal
- O(1) deletion when node is known
- LRU cache, browser history
- More complex operations

#### Use Circular Linked List When:

- Continuous cycle needed
- Round-robin scheduling
- Cyclic processes

---

This is the **complete theoretical foundation** for linked lists. Master these concepts before moving to actual coding problems! 🎯
