# Phase 3 Part 01

Here is Phase 3 fully explained with examples, dry runs, and code.

---

## PHASE 3: MERGE & STRUCTURAL MANIPULATION

**Goal:** Master list combination, splitting/reconnecting, and the dummy node technique.

This is where linked list problems become truly interview-level.

---

### 1. THE DUMMY NODE TECHNIQUE (Most Important Tool of Phase 3)

Before jumping into problems, you must understand this technique deeply because almost every problem in this phase uses it.

**The problem without a dummy node:**

When building a new list (like merging two lists), you need to track the head. But the very first node you attach IS the head, creating a special case that makes your code messy.

**The dummy node solution:**

Create a fake "placeholder" node before you start. Build your result list starting from this fake node. At the end, return `dummy.next` which skips the fake node and gives you the real answer.

java

```java
Node dummy = new Node(-1);  // Fake head
Node temp = dummy;          // Pointer to build list

// ... build list using temp ...

return dummy.next;          // Skip dummy, return real head
```

**Why this is genius:** You never have to write special code for "attach the first node." Every node, including the first real one, is attached the same way: `temp.next = someNode; temp = temp.next;`

---

### 2. MERGE TWO SORTED LISTS

**Problem:** Given two sorted linked lists, merge them into one sorted linked list.

```
Input:
List 1:  1 -> 3 -> 5 -> null
List 2:  2 -> 4 -> 6 -> null

Output:
         1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
```

**Core idea:** At every step, compare the front of both lists. Pick the smaller one, attach it to your result, and advance that list's pointer.

**Code:**

java

```java
Node mergeTwoLists(Node l1, Node l2) {
    Node dummy = new Node(-1);
    Node temp = dummy;

    while (l1 != null && l2 != null) {
        if (l1.data <= l2.data) {
            temp.next = l1;
            l1 = l1.next;
        } else {
            temp.next = l2;
            l2 = l2.next;
        }
        temp = temp.next;  // CRITICAL: always move temp forward
    }

    // Attach remaining nodes (one list is exhausted)
    if (l1 != null) temp.next = l1;
    if (l2 != null) temp.next = l2;

    return dummy.next;
}
```

**IMPORTANT:** We are NOT creating new nodes. We are rewiring existing nodes. That's why it's O(1) space.

**Detailed Dry Run:**

```
l1: 1 -> 3 -> 5
l2: 2 -> 4 -> 6
dummy -> null, temp = dummy

--- Iteration 1 ---
l1.data(1) <= l2.data(2)? YES
temp.next = l1 (Node 1)
l1 = l1.next = Node 3
temp = temp.next = Node 1
State: dummy -> 1,  l1=3, l2=2

--- Iteration 2 ---
l1.data(3) <= l2.data(2)? NO
temp.next = l2 (Node 2)
l2 = l2.next = Node 4
temp = temp.next = Node 2
State: dummy -> 1 -> 2,  l1=3, l2=4

--- Iteration 3 ---
l1.data(3) <= l2.data(4)? YES
temp.next = Node 3
l1 = Node 5
temp = Node 3
State: dummy -> 1 -> 2 -> 3,  l1=5, l2=4

--- Iteration 4 ---
l1.data(5) <= l2.data(4)? NO
temp.next = Node 4
l2 = Node 6
temp = Node 4

--- Iteration 5 ---
l1.data(5) <= l2.data(6)? YES
temp.next = Node 5
l1 = null
temp = Node 5

--- Loop exits (l1 is null) ---
l2 still has: 6 -> null
temp.next = l2

Final: dummy -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
return dummy.next = Node 1
```

Time: O(n + m) | Space: O(1)

**Common Bug 1 — Forgetting to move temp:**

java

```java
// ❌ WRONG
temp.next = l1;
l1 = l1.next;
// temp never moves! You keep overwriting the same connection

// ✅ RIGHT
temp.next = l1;
l1 = l1.next;
temp = temp.next;  // MUST move temp forward
```

**Common Bug 2 — Forgetting to attach the remaining list:**

java

```java
// ❌ WRONG - after loop, one list still has nodes, they get lost
return dummy.next;

// ✅ RIGHT
if (l1 != null) temp.next = l1;
if (l2 != null) temp.next = l2;
return dummy.next;
```

---

### 3. SORT LINKED LIST (Merge Sort)

**Problem:** Sort a linked list in ascending order.

```
Input:  4 -> 2 -> 1 -> 3 -> null
Output: 1 -> 2 -> 3 -> 4 -> null
```

**Why Merge Sort and NOT Quick Sort for linked lists?**

Quick sort needs random access to pick a pivot and partition efficiently. Linked lists don't have random access — reaching index i takes O(n). Merge sort only needs to split in half and then merge, both of which work perfectly with pointers.

**The 4 steps:**

1. Base case: if 0 or 1 node, already sorted — return it
2. Find the middle using fast-slow pointer
3. Split into two halves by setting `slow.next = null`
4. Recursively sort both halves, then merge

**Code:**

java

```java
Node sortList(Node head) {
    // Base case
    if (head == null || head.next == null) {
        return head;
    }

    // Step 1: Find middle
    Node slow = head;
    Node fast = head.next;  // fast starts at head.next for better split

    while (fast != null && fast.next != null) {
        slow = slow.next;
        fast = fast.next.next;
    }

    // Step 2: Split
    Node second = slow.next;
    slow.next = null;  // PHYSICALLY disconnect the two halves

    // Step 3: Recursively sort
    Node left = sortList(head);
    Node right = sortList(second);

    // Step 4: Merge
    return mergeTwoLists(left, right);
}
```

**Why `fast = head.next` instead of `fast = head`?**

If you start both at head, for a 2-node list `1 -> 2`:

- slow would end at Node 2 (last node)
- `slow.next = null` means second = null
- You'd recursively call `sortList(1->2)` and `sortList(null)` — infinite recursion!

Starting fast at head.next ensures slow stops at the first half's last node.

**Detailed Dry Run:**

```
sortList(4 -> 2 -> 1 -> 3)

Find middle:
slow = 4, fast = 2
  slow = 2, fast = 3   (fast.next.next skips to 3.next = null — wait)

Let me trace carefully:
Iteration 1: fast=2 (not null), fast.next=1 (not null) → slow=2, fast=1.next.next? 
  fast = head.next = Node(2)
  fast.next = Node(1), fast.next.next = Node(3)
  So: slow = Node(2), fast = Node(3)
Iteration 2: fast=3 (not null), fast.next=null → STOP
  slow is at Node(2)

Split:
  second = slow.next = Node(1)
  slow.next = null
  Left:  4 -> 2 -> null
  Right: 1 -> 3 -> null

Recursively sort left (4 -> 2):
  Find middle: slow=4, fast=2
    fast.next = null, STOP, slow = Node(4)
  second = Node(2), slow.next = null
  Left: 4, Right: 2
  sortList(4) → returns 4
  sortList(2) → returns 2
  merge(4, 2) → 2 -> 4

Recursively sort right (1 -> 3):
  Find middle: slow=1, fast=3
    fast.next = null, STOP, slow = Node(1)
  second = Node(3), slow.next = null
  sortList(1) → returns 1
  sortList(3) → returns 3
  merge(1, 3) → 1 -> 3

Final merge(2->4, 1->3):
  Compare 2 vs 1: pick 1
  Compare 2 vs 3: pick 2
  Compare 4 vs 3: pick 3
  Remaining: 4
  Result: 1 -> 2 -> 3 -> 4
```

Time: O(n log n) | Space: O(log n) recursion stack

---

### 4. REORDER LIST

**Problem:** Given list `1 -> 2 -> 3 -> 4 -> 5`, reorder it to `1 -> 5 -> 2 -> 4 -> 3`.

The pattern is: first, last, second, second-last, third, third-last...

**This is Phase 3's boss problem.** It combines three patterns you've already learned:

- Fast-slow pointer (find middle)
- Reversal (reverse second half)
- Merge (interleave two halves)

**The 3 steps:**

```
Original:        1 -> 2 -> 3 -> 4 -> 5

Step 1 (split):  1 -> 2 -> 3    and    4 -> 5

Step 2 (reverse second half):   5 -> 4

Step 3 (interleave):  1 -> 5 -> 2 -> 4 -> 3
```

**Code:**

java

```java
void reorderList(Node head) {
    if (head == null || head.next == null) return;

    // Step 1: Find middle
    Node slow = head;
    Node fast = head;
    while (fast != null && fast.next != null) {
        slow = slow.next;
        fast = fast.next.next;
    }

    // Step 2: Reverse second half
    Node second = slow.next;
    slow.next = null;          // Disconnect first half from second
    second = reverse(second);  // Reverse second half

    // Step 3: Merge alternately
    Node first = head;
    while (second != null) {
        Node tmp1 = first.next;   // Save next of first half
        Node tmp2 = second.next;  // Save next of second half

        first.next = second;      // first -> second
        second.next = tmp1;       // second -> next of first

        first = tmp1;             // Advance first pointer
        second = tmp2;            // Advance second pointer
    }
}

Node reverse(Node head) {
    Node prev = null;
    Node curr = head;
    while (curr != null) {
        Node next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}
```

**Detailed Dry Run:**

```
List: 1 -> 2 -> 3 -> 4 -> 5

Step 1: Find middle
slow: 1 -> 2 -> 3
fast: 1 -> 3 -> 5
slow stops at Node(3)

Split:
First half:  1 -> 2 -> 3 -> null
Second half: 4 -> 5 -> null

Step 2: Reverse second half
reverse(4 -> 5):
  prev=null, curr=4
  next=5, 4.next=null, prev=4, curr=5
  next=null, 5.next=4, prev=5, curr=null
  return 5
Second half reversed: 5 -> 4 -> null

Step 3: Interleave
first=1, second=5

--- Round 1 ---
tmp1 = first.next = Node(2)
tmp2 = second.next = Node(4)
first.next = second → 1 -> 5
second.next = tmp1  → 5 -> 2
first = tmp1 = Node(2)
second = tmp2 = Node(4)

--- Round 2 ---
tmp1 = Node(3)
tmp2 = null
first.next = second → 2 -> 4
second.next = tmp1  → 4 -> 3
first = Node(3)
second = null

--- Loop exits (second is null) ---

Final: 1 -> 5 -> 2 -> 4 -> 3 -> null ✓
```

Time: O(n) | Space: O(1)

---

### 5. INTERSECTION OF TWO LINKED LISTS

**Problem:** Find the node where two linked lists intersect (share the same node object, not just the same value).

```
A: 1 -> 2 -> 8 -> 9 -> null
                ↑ (same node)
B:      5 -> 8 -> 9 -> null

Answer: Node(8)
```

**Key point:** Intersection means the actual same node in memory. After the intersection, both lists share the exact same nodes.

**Brute force:** For every node in A, scan all of B. O(n²) — too slow.

**Optimal approach (Two Pointer trick):**

The insight is that both pointers need to travel the same total distance to meet at the intersection. So let each pointer traverse both lists:

- p1 goes through A then switches to B
- p2 goes through B then switches to A

They'll both have traveled `lengthA + lengthB` distance, meeting at the intersection.

```
A length = 4 (1->2->8->9)
B length = 3 (5->8->9)

p1 path: 1 -> 2 -> 8 -> 9 -> null -> 5 -> 8 (meets here!)
p2 path: 5 -> 8 -> 9 -> null -> 1 -> 2 -> 8 (meets here!)

Both traveled 7 nodes total, meeting at Node(8) ✓
```

**Code:**

java

```java
Node getIntersectionNode(Node a, Node b) {
    Node p1 = a;
    Node p2 = b;

    while (p1 != p2) {
        p1 = (p1 == null) ? b : p1.next;
        p2 = (p2 == null) ? a : p2.next;
