# Phase 4 Part 01

## PHASE 4: ADVANCED STRUCTURAL MANIPULATION

**Goal:** Develop true interview-level mastery. Problems here are not just traversal — they are **pointer architecture problems**.

The mindset shift: Earlier you moved _through_ lists. Now you _restructure_ them safely using segments, boundaries, and reconnect operations.

---

### 1. REVERSE NODES IN K-GROUP

**Problem:** Reverse every k consecutive nodes. If remaining nodes are less than k, leave them as is.

```
Input:  1 -> 2 -> 3 -> 4 -> 5,  k = 2
Output: 2 -> 1 -> 4 -> 3 -> 5

Input:  1 -> 2 -> 3 -> 4 -> 5,  k = 3
Output: 3 -> 2 -> 1 -> 4 -> 5
```

**What makes this hard:** It's not just reversal. You must:

1. Find the kth node (boundary of group)
2. Detach the group
3. Reverse it
4. Reconnect it to previous group AND next group
5. Repeat for the rest of the list

**The 4 pointers you need:**

|Pointer|Purpose|
|---|---|
|`groupPrev`|Node just before current group (to reconnect after reversal)|
|`kth`|Last node of current group (group end boundary)|
|`groupNext`|Node just after current group (to reconnect after reversal)|
|`prev/curr`|Standard reversal pointers inside the group|

**Visual before each group reversal:**

```
groupPrev -> [1 -> 2 -> 3] -> groupNext
              ↑            ↑
             curr          kth
```

After reversal:

```
groupPrev -> [3 -> 2 -> 1] -> groupNext
```

**Code:**

java

```java
// Helper: find the kth node from curr
Node getKth(Node curr, int k) {
    while (curr != null && k > 0) {
        curr = curr.next;
        k--;
    }
    return curr;
}

Node reverseKGroup(Node head, int k) {
    Node dummy = new Node(0);
    dummy.next = head;
    Node groupPrev = dummy;

    while (true) {
        // Step 1: Find kth node
        Node kth = getKth(groupPrev, k);
        if (kth == null) break;  // Less than k nodes remaining, stop

        Node groupNext = kth.next;  // Save what comes after group

        // Step 2: Reverse the group
        Node prev = groupNext;  // prev starts at groupNext (for reconnection!)
        Node curr = groupPrev.next;

        while (curr != groupNext) {
            Node next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }

        // Step 3: Reconnect groupPrev to new group head
        Node tmp = groupPrev.next;  // This is old head of group (now group tail after reversal)
        groupPrev.next = kth;       // groupPrev now points to new group head
        groupPrev = tmp;            // Move groupPrev to end of reversed group (for next iteration)
    }

    return dummy.next;
}
```

**Why `prev = groupNext` at the start of reversal?**

Because after reversing the group, its last node must point to `groupNext` (what comes after the group). Starting `prev` at `groupNext` ensures the reversal automatically connects the group tail to the next group. Genius trick.

**Detailed Dry Run (k=2):**

```
List: 1 -> 2 -> 3 -> 4 -> 5
dummy -> 1 -> 2 -> 3 -> 4 -> 5
groupPrev = dummy

--- GROUP 1 ---
getKth(dummy, 2):  dummy -> 1 -> 2  → kth = Node(2)
groupNext = Node(3)

Reverse group (from groupPrev.next=1 to groupNext=3):
  prev = groupNext = Node(3)
  curr = Node(1)

  Iteration 1: curr=1 != groupNext(3)
    next = 2
    1.next = prev(3) → 1 -> 3
    prev = 1
    curr = 2

  Iteration 2: curr=2 != groupNext(3)
    next = 3
    2.next = prev(1) → 2 -> 1
    prev = 2
    curr = 3

  curr=3 == groupNext → STOP
  Group reversed: 2 -> 1 -> 3 -> 4 -> 5

Reconnect:
  tmp = groupPrev.next = Node(1) (old head, now tail of group)
  groupPrev.next = kth = Node(2) → dummy -> 2 -> 1 -> 3 -> 4 -> 5
  groupPrev = tmp = Node(1)

--- GROUP 2 ---
getKth(Node(1), 2): 1 -> 3 -> 4 → kth = Node(4)
groupNext = Node(5)

Reverse (from groupPrev.next=3 to groupNext=5):
  prev = Node(5)
  curr = Node(3)

  3.next = 5, prev = 3, curr = 4
  4.next = 3, prev = 4, curr = 5
  curr = groupNext → STOP

Reconnect:
  tmp = Node(3) (was groupPrev.next, now tail)
  groupPrev.next = Node(4) → 1 -> 4 -> 3 -> 5
  groupPrev = Node(3)

--- GROUP 3 ---
getKth(Node(3), 2): 3 -> 5 -> null → kth = null
kth == null → BREAK

Final: dummy -> 2 -> 1 -> 4 -> 3 -> 5 ✓
```

Time: O(n) | Space: O(1)

**Edge case:** If list length is not a multiple of k, the last group is left as-is because `getKth` returns null.

---

### 2. ODD-EVEN LINKED LIST

**Problem:** Group all odd-positioned nodes first, then all even-positioned nodes. Position 1, 3, 5... are odd; position 2, 4, 6... are even.

```
Input:  1 -> 2 -> 3 -> 4 -> 5
Output: 1 -> 3 -> 5 -> 2 -> 4

Input:  2 -> 1 -> 3 -> 5 -> 6 -> 4 -> 7
Output: 2 -> 3 -> 6 -> 7 -> 1 -> 5 -> 4
```

**Critical reminder:** This is about positions (1st, 2nd, 3rd...), NOT the values being odd or even!

**Core idea:** Maintain two chains simultaneously — odd-position chain and even-position chain. At the end, connect the odd chain's tail to the even chain's head.

```
Original:    1 -> 2 -> 3 -> 4 -> 5
Odd chain:   1 -> 3 -> 5
Even chain:  2 -> 4
Connect:     1 -> 3 -> 5 -> 2 -> 4
```

**Code:**

java

```java
Node oddEvenList(Node head) {
    if (head == null || head.next == null) return head;

    Node odd = head;           // Points to current odd node
    Node even = head.next;     // Points to current even node
    Node evenHead = even;      // Save even head to connect at end

    while (even != null && even.next != null) {
        odd.next = even.next;   // Skip even, connect odd to next odd
        odd = odd.next;          // Advance odd pointer

        even.next = odd.next;    // Skip odd, connect even to next even
        even = even.next;        // Advance even pointer
    }

    odd.next = evenHead;  // Connect odd tail to even head
    return head;
}
```

**Detailed Dry Run:**

```
List: 1 -> 2 -> 3 -> 4 -> 5
odd = 1, even = 2, evenHead = 2

--- Iteration 1 ---
even(2) != null && even.next(3) != null ✓
odd.next = even.next → 1.next = 3
odd = odd.next = 3
even.next = odd.next → 2.next = 4
even = even.next = 4

State: 1 -> 3 -> 4 -> 5  (odd chain: 1->3)
       2 -> 4 -> 5        (even chain: 2->4)

--- Iteration 2 ---
even(4) != null && even.next(5) != null ✓
odd.next = even.next → 3.next = 5
odd = odd.next = 5
even.next = odd.next → 4.next = null
even = even.next = null

State: odd chain: 1 -> 3 -> 5
       even chain: 2 -> 4 -> null

--- Loop check ---
even = null → EXIT

odd.next = evenHead → 5.next = 2

Final: 1 -> 3 -> 5 -> 2 -> 4 -> null ✓
```

Time: O(n) | Space: O(1)

**Key insight:** Save `evenHead` before the loop. Once you start rewiring, you'll lose access to where the even chain starts.

---

### 3. ROTATE LIST

**Problem:** Rotate list to the right by k places.

```
Input:  1 -> 2 -> 3 -> 4 -> 5,  k = 2
Output: 4 -> 5 -> 1 -> 2 -> 3
```

**Understanding rotation:** Rotating right by 2 means the last 2 nodes come to the front.

**Strategy:**

1. Find length and make the list circular (tail.next = head)
2. Find the new tail: it's at position `length - (k % length) - 1`
3. The new head is `newTail.next`
4. Break the circle: `newTail.next = null`

**Why `k % length`?** Rotating by k=5 on a list of length 5 brings you back to the same list. So `k % length` gives the effective rotation.

**Code:**

java

```java
Node rotateRight(Node head, int k) {
    if (head == null || head.next == null || k == 0) return head;

    // Step 1: Find length and tail
    int length = 1;
    Node tail = head;
    while (tail.next != null) {
        tail = tail.next;
        length++;
    }

    // Step 2: Effective rotation
    k = k % length;
    if (k == 0) return head;  // No rotation needed

    // Step 3: Make circular
    tail.next = head;

    // Step 4: Find new tail (length - k steps from original head)
    int stepsToNewTail = length - k;
    Node newTail = head;
    for (int i = 1; i < stepsToNewTail; i++) {
        newTail = newTail.next;
    }

    // Step 5: Set new head and break circle
    Node newHead = newTail.next;
    newTail.next = null;

    return newHead;
}
```

**Detailed Dry Run:**

```
List: 1 -> 2 -> 3 -> 4 -> 5,  k = 2

Step 1: length = 5, tail = Node(5)

Step 2: k = 2 % 5 = 2

Step 3: tail.next = head → circular: 1->2->3->4->5->1->...

Step 4: stepsToNewTail = 5 - 2 = 3
  Start at Node(1)
  i=1: newTail = Node(2)
  i=2: newTail = Node(3)
  Loop ends. newTail = Node(3)

Step 5:
  newHead = newTail.next = Node(4)
  newTail.next = null → breaks circle at 3

Final: 4 -> 5 -> 1 -> 2 -> 3 -> null ✓
```

Time: O(n) | Space: O(1)

---

### 4. PARTITION LIST

**Problem:** Partition list around value x such that all nodes less than x come before nodes greater than or equal to x. **Preserve original relative order.**

```
Input:  1 -> 4 -> 3 -> 2 -> 5 -> 2,  x = 3
Output: 1 -> 2 -> 2 -> 4 -> 3 -> 5
```

**Core idea:** Build two separate lists simultaneously — one for nodes less than x, one for nodes greater than or equal to x. Then connect them.

**Code:**

java

```java
Node partition(Node head, int x) {
    // Two dummy heads for two chains
    Node smallDummy = new Node(0);
    Node largeDummy = new Node(0);

    Node small = smallDummy;
    Node large = largeDummy;

    Node curr = head;

    while (curr != null) {
        if (curr.data < x) {
            small.next = curr;
            small = small.next;
        } else {
            large.next = curr;
            large = large.next;
        }
        curr = curr.next;
    }

    // CRITICAL: break the large chain's tail link
    large.next = null;

    // Connect small chain to large chain
    small.next = largeDummy.next;

    return smallDummy.next;
}
```

**Detailed Dry Run:**

```
List: 1 -> 4 -> 3 -> 2 -> 5 -> 2,  x = 3
smallDummy -> null,  largeDummy -> null

curr = 1: 1 < 3 → small chain: smallDummy -> 1
curr = 4: 4 >= 3 → large chain: largeDummy -> 4
curr = 3: 3 >= 3 → large chain: largeDummy -> 4 -> 3
curr = 2: 2 < 3 → small chain: smallDummy -> 1 -> 2
curr = 5: 5 >= 3 → large chain: largeDummy -> 4 -> 3 -> 5
curr = 2: 2 < 3 → small chain: smallDummy -> 1 -> 2 -> 2

large.next = null → breaks any dangling link from Node(5)

small.next = largeDummy.next = Node(4)
→ 1 -> 2 -> 2 -> 4 -> 3 -> 5 -> null ✓
```

**Why `large.next = null`?** The last node added to the large chain (Node 5 here) still has its original `next` pointer (pointing to Node 2 from the original list). If you don't set it to null, the large chain has a dangling connection creating a corrupted or circular list.

Time: O(n) | Space: O(1)

---

### 5. COPY LIST WITH RANDOM POINTER

**Problem:** Deep copy a linked list where each node has a `next` pointer AND a `random` pointer that can point to any node in the list (or null).

java

```java
class Node {
    int val;
    Node next;
    Node random;
}
```

```
Input:  1 -> 2 -> 3
        1.random = 3
        2.random = 1
        3.random = null

Output: New copied list with exact same structure
```

**Why this is hard:** When you create a copy of node 1, you need to set copy1.random to point to the copy of node 3 — not the original node 3. But how do you find the copy of node 3 while building the copies?

#### Approach 1: HashMap (Easier, O(n) space)

Map every original node to its copy. Then assign pointers using the map.

java

```java
Node copyRandomList(Node head) {
    if (head == null) return null;

    HashMap<Node, Node> map = new HashMap<>();

