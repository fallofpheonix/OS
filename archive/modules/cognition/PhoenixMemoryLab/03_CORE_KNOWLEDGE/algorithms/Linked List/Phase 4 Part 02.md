# Phase 4 Part 02

    // Pass 1: Create all copy nodes
    Node curr = head;
    while (curr != null) {
        map.put(curr, new Node(curr.val));
        curr = curr.next;
    }

    // Pass 2: Assign next and random using map
    curr = head;
    while (curr != null) {
        map.get(curr).next = map.get(curr.next);
        map.get(curr).random = map.get(curr.random);
        curr = curr.next;
    }

    return map.get(head);
}
```

Time: O(n) | Space: O(n)

#### Approach 2: Interleaving (Optimal, O(1) space)

Transform the original list by inserting copy nodes right after each original node:

```
Original:    A -> B -> C
After step1: A -> A' -> B -> B' -> C -> C'
```

Now `A.next = A'`, so `A.next.random = A'.random`. We can assign random pointers as: `curr.next.random = curr.random.next` (the copy of curr.random is right after curr.random).

Then separate the two lists.

java

```java
Node copyRandomList(Node head) {
    if (head == null) return null;

    // Step 1: Interleave copy nodes
    Node curr = head;
    while (curr != null) {
        Node copy = new Node(curr.val);
        copy.next = curr.next;
        curr.next = copy;
        curr = copy.next;  // Move to original next
    }

    // Step 2: Assign random pointers to copies
    curr = head;
    while (curr != null) {
        if (curr.random != null) {
            curr.next.random = curr.random.next;  // copy's random = copy of original's random
        }
        curr = curr.next.next;  // Jump to next original node
    }

    // Step 3: Separate the two lists
    Node dummy = new Node(0);
    Node copyCurr = dummy;
    curr = head;

    while (curr != null) {
        copyCurr.next = curr.next;    // Extract copy node
        curr.next = curr.next.next;   // Restore original list
        copyCurr = copyCurr.next;
        curr = curr.next;
    }

    return dummy.next;
}
```

**Dry Run of Interleaving:**

```
Original: 1 -> 2 -> 3
1.random = 3, 2.random = 1, 3.random = null

Step 1: Interleave
  curr=1: insert 1' after 1 → 1 -> 1' -> 2 -> 3
  curr=2: insert 2' after 2 → 1 -> 1' -> 2 -> 2' -> 3
  curr=3: insert 3' after 3 → 1 -> 1' -> 2 -> 2' -> 3 -> 3'

Step 2: Assign random to copies
  curr=1: 1.random=3, so 1'.random = 3.next = 3'  ✓
  curr=2: 2.random=1, so 2'.random = 1.next = 1'  ✓
  curr=3: 3.random=null, skip

Step 3: Separate
  Extract 1', restore 1.next=2
  Extract 2', restore 2.next=3
  Extract 3', restore 3.next=null

Original restored: 1 -> 2 -> 3
Copy:             1' -> 2' -> 3' (with correct randoms)
```

Time: O(n) | Space: O(1)

---

### 6. LRU CACHE

**Problem:** Design a data structure with O(1) `get(key)` and `put(key, value)`. When capacity is exceeded, evict the Least Recently Used item.

**Why a Doubly Linked List?**

- You need to move any node to "most recently used" position (front) in O(1)
- To remove a node from the middle of a list in O(1), you need its `prev` pointer
- Singly linked list would require O(n) to find the previous node
- DLL gives O(1) removal when you have the node directly

**Why also a HashMap?**

- `get(key)` must be O(1) — you can't traverse the DLL to find the key
- HashMap maps key → node directly for O(1) access

**Structure:**

```
HashMap: key → node

DLL:   dummy_head <-> (MRU node) <-> ... <-> (LRU node) <-> dummy_tail

Most Recently Used = near head
Least Recently Used = near tail (gets evicted)
```

Using two dummy nodes (head and tail) eliminates edge cases when adding/removing from ends.

**Code:**

java

```java
class LRUCache {
    class Node {
        int key, val;
        Node prev, next;
        Node(int k, int v) { key = k; val = v; }
    }

    int capacity;
    HashMap<Integer, Node> map;
    Node head, tail;  // Dummy head and tail

    LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>();

        // Dummy nodes eliminate edge cases
        head = new Node(0, 0);
        tail = new Node(0, 0);
        head.next = tail;
        tail.prev = head;
    }

    // Remove a node from its current position in DLL
    private void remove(Node node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    // Insert node right after head (most recently used position)
    private void insertAfterHead(Node node) {
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
    }

    public int get(int key) {
        if (!map.containsKey(key)) return -1;

        Node node = map.get(key);
        remove(node);           // Remove from current position
        insertAfterHead(node);  // Move to front (most recently used)
        return node.val;
    }

    public void put(int key, int value) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.val = value;     // Update value
            remove(node);
            insertAfterHead(node);
        } else {
            if (map.size() == capacity) {
                // Evict LRU (node just before dummy tail)
                Node lru = tail.prev;
                remove(lru);
                map.remove(lru.key);
            }
            Node newNode = new Node(key, value);
            map.put(key, newNode);
            insertAfterHead(newNode);
        }
    }
}
```

**Dry Run:**

```
LRUCache(2)  capacity=2
DLL: head <-> tail

put(1, 1):
  New node (1,1). Insert after head.
  DLL: head <-> (1,1) <-> tail
  map: {1: Node(1,1)}

put(2, 2):
  New node (2,2). Insert after head.
  DLL: head <-> (2,2) <-> (1,1) <-> tail
  map: {1: Node(1,1), 2: Node(2,2)}

get(1):
  Found key=1. Remove from DLL. Insert after head.
  DLL: head <-> (1,1) <-> (2,2) <-> tail
  return 1

put(3, 3):
  map.size(2) == capacity(2) → evict LRU
  LRU = tail.prev = Node(2,2)
  Remove Node(2,2). map.remove(2).
  DLL: head <-> (1,1) <-> tail
  Insert new Node(3,3) after head.
  DLL: head <-> (3,3) <-> (1,1) <-> tail
  map: {1: Node(1,1), 3: Node(3,3)}

get(2):
  Key 2 not in map → return -1
```

Time: O(1) for both get and put | Space: O(capacity)

**The 4 things that must stay in sync — always do both:**

java

```java
// When accessing a key: update DLL AND it's already in map (no map change)
// When adding a key: update BOTH map AND DLL
// When evicting: remove from BOTH map AND DLL
```

---

### 7. COMMON PHASE 4 MISTAKES

**Mistake 1: Forgetting to reconnect after group reversal (K-Group)**

java

```java
// ❌ WRONG: After reversing group, group is floating
// The previous group's tail still points to the old group head!

// ✅ RIGHT: Always reconnect groupPrev.next = kth (new group head)
```

**Mistake 2: Not setting `large.next = null` in Partition**

java

```java
// ❌ WRONG: large chain's last node still points somewhere in original list
// Creates corrupted list or infinite loop!

// ✅ RIGHT: large.next = null before connecting chains
```

**Mistake 3: Using wrong k in Rotate List**

java

```java
// ❌ WRONG: k can be larger than list length
// Rotating 7 times a list of 5 = rotating 2 times
// Must: k = k % length

// ✅ RIGHT: Always normalize k first
```

**Mistake 4: HashMap and DLL out of sync in LRU**

java

```java
// ❌ WRONG: Remove from DLL but forget to remove from map (or vice versa)
// Next operation will find a stale/ghost entry in map

// ✅ RIGHT: Every DLL operation must have a corresponding map operation
```

**Mistake 5: Comparing node values instead of references**

java

```java
// ❌ WRONG
if (slow.val == fast.val)

// ✅ RIGHT
if (slow == fast)
```

---

### 8. COMPLEXITY SUMMARY

|Problem|Time|Space|Key Technique|
|---|---|---|---|
|Reverse K-Group|O(n)|O(1)|Segmentation + reversal + reconnection|
|Odd-Even List|O(n)|O(1)|Two chains + connect|
|Rotate List|O(n)|O(1)|Circular + break at new tail|
|Partition List|O(n)|O(1)|Two dummy chains + connect|
|Copy Random Pointer (HashMap)|O(n)|O(n)|Map old→new|
|Copy Random Pointer (Interleave)|O(n)|O(1)|Interleave + assign + separate|
|LRU Cache|O(1) per op|O(n)|HashMap + DLL|

---

### 9. COMPLETE WORKING CODE

java

```java
public class Phase4 {
    static class Node {
        int data;
        Node next;
        Node(int d) { data = d; next = null; }
    }

    // ===== 1. ODD EVEN LIST =====
    static Node oddEvenList(Node head) {
        if (head == null || head.next == null) return head;
        Node odd = head, even = head.next, evenHead = even;
        while (even != null && even.next != null) {
            odd.next = even.next;  odd = odd.next;
            even.next = odd.next;  even = even.next;
        }
        odd.next = evenHead;
        return head;
    }

    // ===== 2. ROTATE LIST =====
    static Node rotateRight(Node head, int k) {
        if (head == null || head.next == null || k == 0) return head;
        int length = 1;
        Node tail = head;
        while (tail.next != null) { tail = tail.next; length++; }
        k = k % length;
        if (k == 0) return head;
        tail.next = head;
        int stepsToNewTail = length - k;
        Node newTail = head;
        for (int i = 1; i < stepsToNewTail; i++) newTail = newTail.next;
        Node newHead = newTail.next;
        newTail.next = null;
        return newHead;
    }

    // ===== 3. PARTITION LIST =====
    static Node partition(Node head, int x) {
        Node sd = new Node(0), ld = new Node(0);
        Node small = sd, large = ld, curr = head;
        while (curr != null) {
            if (curr.data < x) { small.next = curr; small = small.next; }
            else               { large.next = curr; large = large.next; }
            curr = curr.next;
        }
        large.next = null;
        small.next = ld.next;
        return sd.next;
    }

    // ===== 4. REVERSE K GROUP =====
    static Node getKth(Node curr, int k) {
        while (curr != null && k-- > 0) curr = curr.next;
        return curr;
    }

    static Node reverseKGroup(Node head, int k) {
        Node dummy = new Node(0); dummy.next = head;
        Node groupPrev = dummy;
        while (true) {
            Node kth = getKth(groupPrev, k);
            if (kth == null) break;
            Node groupNext = kth.next;
            Node prev = groupNext, curr = groupPrev.next;
            while (curr != groupNext) {
                Node next = curr.next;
                curr.next = prev; prev = curr; curr = next;
            }
            Node tmp = groupPrev.next;
            groupPrev.next = kth;
            groupPrev = tmp;
        }
        return dummy.next;
    }

    // ===== UTILITY =====
    static void print(Node head) {
        while (head != null) { System.out.print(head.data + " -> "); head = head.next; }
        System.out.println("null");
    }

    static Node build(int... vals) {
        Node dummy = new Node(0), t = dummy;
        for (int v : vals) { t.next = new Node(v); t = t.next; }
        return dummy.next;
    }

    public static void main(String[] args) {
        System.out.println("=== ODD EVEN LIST ===");
        print(oddEvenList(build(1,2,3,4,5)));  // 1->3->5->2->4

        System.out.println("=== ROTATE LIST k=2 ===");
        print(rotateRight(build(1,2,3,4,5), 2));  // 4->5->1->2->3

        System.out.println("=== PARTITION x=3 ===");
        print(partition(build(1,4,3,2,5,2), 3));  // 1->2->2->4->3->5

        System.out.println("=== REVERSE K GROUP k=2 ===");
        print(reverseKGroup(build(1,2,3,4,5), 2));  // 2->1->4->3->5

        System.out.println("=== REVERSE K GROUP k=3 ===");
        print(reverseKGroup(build(1,2,3,4,5), 3));  // 3->2->1->4->5
    }
}
```

**Output:**

```
=== ODD EVEN LIST ===
1 -> 3 -> 5 -> 2 -> 4 -> null
=== ROTATE LIST k=2 ===
4 -> 5 -> 1 -> 2 -> 3 -> null
=== PARTITION x=3 ===
1 -> 2 -> 2 -> 4 -> 3 -> 5 -> null
=== REVERSE K GROUP k=2 ===
2 -> 1 -> 4 -> 3 -> 5 -> null
=== REVERSE K GROUP k=3 ===
3 -> 2 -> 1 -> 4 -> 5 -> null
