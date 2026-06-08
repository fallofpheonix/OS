# Phase 5 Part 01

## PHASE 5: LINKED LIST MASTERY & INTERVIEW HARD PROBLEMS

**Goal:** Stop solving known problems from memory. Start **deriving solutions** for unseen problems by understanding the underlying structure.

This phase is about upgrading your thinking from template-matching to pattern synthesis.

---

### 1. THINKING IN INVARIANTS (Most Important Concept of Phase 5)

**What is an invariant?** A condition that stays true throughout execution of your algorithm. This is how experts think about code — not "what does this line do" but "what is always true at this point."

**Example: Reversal invariant**

At every point during reversal, this is always true:

- `prev` always points to the already-reversed portion
- `curr` always points to the not-yet-reversed portion

```
State after 2 iterations on: 1 -> 2 -> 3 -> 4 -> 5

REVERSED PART     REMAINING PART
null <- 1 <- 2    3 -> 4 -> 5
              ↑   ↑
            prev  curr

Invariant holds: prev = reversed, curr = remaining ✓
```

**Why this matters for hard problems:**

When you hit a bug, ask: "Which invariant got violated?" That tells you exactly where the problem is. This is far faster than reading line by line.

**Example: Fast-Slow invariant**

At every iteration:

- `fast` is always exactly 2× steps ahead of `slow`
- In a cycle: the gap between them shrinks by 1 each iteration (since fast gains 1 step per round)
- They will always meet — this is the mathematical guarantee

**Example: Merge invariant**

At every step of merging two sorted lists:

- `temp` always points to the last node of the result list built so far
- Everything before `temp` is already correctly sorted

Internalize these invariants. When you write code, you're not writing steps — you're maintaining invariants.

---

### 2. SEGMENT-BASED THINKING

Hard LL problems stop being about individual nodes. They're about **segments** (bounded regions of nodes).

**Always mentally split your list into regions:**

```
[processed] [current segment] [unprocessed]
     ↑              ↑               ↑
  done already   working here   handle later
```

**Examples of segment thinking:**

K-Group Reversal:

```
[groupPrev] | [node1 -> node2 -> node3] | [groupNext -> ...]
              ↑ this segment gets reversed ↑
```

Reorder List:

```
[first half: 1->2->3] | [second half reversed: 5->4]
  merge these two alternately
```

Palindrome Check:

```
[first half: 1->2->3] | [second half reversed: 3->2->1]
  compare these two
```

**Rule:** Never try to think about the entire list at once. Identify the boundaries of the segment you're working on, manipulate that segment, then move to the next.

---

### 3. THE FLOYD CYCLE MATH — DEEPLY EXPLAINED

You've used fast-slow pointers. Now understand WHY they work mathematically, so you can apply the logic to new problems.

**Setup:**

```
Non-cycle part length = x
Cycle length = c
Meeting point is y steps into the cycle from cycle start
```

```
head ---x--- cycleStart ---y--- meetingPoint
                  ↑                    |
                  |__________c-y_______|
```

**When they meet:**

- Slow has traveled: `x + y`
- Fast has traveled: `x + y + n*c` (fast did extra full loops, n ≥ 1)
- Since fast = 2 × slow: `2(x + y) = x + y + n*c`
- Simplify: `x + y = n*c`
- Therefore: `x = n*c - y`

**What does `x = n*c - y` mean?**

If you now move one pointer from `head` and one pointer from the `meetingPoint`, both moving 1 step:

- Head pointer travels `x` steps to reach cycle start
- Meeting point pointer travels `n*c - y` steps, which is also exactly the cycle start (it goes around the cycle and lands exactly at cycle start)

This is WHY resetting one pointer to head and moving both at speed 1 always finds the cycle start. It's not magic — it's distance equivalence.

**Applying this thinking to new problems:**

Any problem that says "find the position/node at distance k from end" or "find where two paths converge" — think about whether fast-slow distance math applies.

---

### 4. DEEP COPY VS SHALLOW COPY — INTERVIEW TRAP

This is a concept where interviewers catch candidates who don't truly understand references.

**Shallow copy:**

java

```java
Node b = a;  // b and a point to the SAME node
b.data = 99; // a.data is also 99 now!
```

**Deep copy:**

java

```java
Node b = new Node(a.data);  // b is a completely NEW node
b.data = 99;                 // a.data is still original value
```

**Why this matters:**

In Copy List with Random Pointer, if you do shallow copy, your "new" list and original list share nodes — modifying one corrupts the other. A true deep copy means every single node is a brand new object in memory.

**The mental test:** After your copy operation, if I delete every node in the original list, does the copy still work perfectly? If yes, it's a deep copy. If no, it's shallow.

---

### 5. TORTOISE-HARE BEYOND LINKED LISTS

The fast-slow pattern is not just for linked lists. Once you understand the underlying logic, you can apply it anywhere a "next pointer" concept exists.

**Find Duplicate Number (LeetCode 287):**

Given array where values are in range [1, n] and one value appears twice:

```
nums = [1, 3, 4, 2, 2]
index:  0  1  2  3  4
```

Treat `nums[i]` as a "next pointer": from index i, go to index `nums[i]`. This creates a linked list structure with a cycle (the duplicate creates the cycle, because two different indices point to the same next index).

java

```java
int findDuplicate(int[] nums) {
    // Phase 1: Find meeting point (Floyd detection)
    int slow = nums[0];
    int fast = nums[0];

    do {
        slow = nums[slow];        // 1 step
        fast = nums[nums[fast]];  // 2 steps
    } while (slow != fast);

    // Phase 2: Find cycle start (duplicate number)
    slow = nums[0];
    while (slow != fast) {
        slow = nums[slow];
        fast = nums[fast];
    }

    return slow;
}
```

**Dry Run:**

```
nums = [1, 3, 4, 2, 2]
Step trace (slow/fast as values, not indices):
  Start: slow=1, fast=1
  Step1: slow=nums[1]=3, fast=nums[nums[1]]=nums[3]=2
  Step2: slow=nums[3]=2, fast=nums[nums[2]]=nums[4]=2
  slow==fast==2 → meeting point

  Reset slow=nums[0]=1
  Step1: slow=nums[1]=3, fast=nums[2]=4
  Step2: slow=nums[3]=2, fast=nums[4]=2
  slow==fast==2 → duplicate = 2 ✓
```

**Key insight:** You recognized a problem that looks nothing like a linked list, but applied the same cycle detection logic. That's pattern abstraction.

---

### 6. LFU CACHE — THE HARDEST CACHE DESIGN

**Problem:** LFU (Least Frequently Used) Cache — when capacity is exceeded, evict the item that was accessed the fewest times. If there's a tie, evict the least recently used among those.

**Why harder than LRU:** LRU only cares about recency. LFU cares about frequency AND recency (as tiebreaker).

**Data structures needed:**

- `keyMap`: key → node (for O(1) key lookup)
- `freqMap`: frequency → DLL of nodes at that frequency (for O(1) eviction)
- `minFreq`: tracks current minimum frequency (so we know which DLL to evict from)

**Structure visualization:**

```
keyMap:  {1: Node(1), 2: Node(2), 3: Node(3)}

freqMap: {
  1: DLL [Node(3)]         ← accessed 1 time
  2: DLL [Node(1)]         ← accessed 2 times (MRU at head)
  3: DLL [Node(2)]         ← accessed 3 times
}

minFreq = 1
```

When we access Node(3): its frequency goes from 1 to 2. Remove it from freqMap[1], add to freqMap[2]. If freqMap[1] is now empty AND minFreq==1, increment minFreq to 2.

**Code:**

java

```java
class LFUCache {
    class Node {
        int key, val, freq;
        Node prev, next;
        Node(int k, int v) { key = k; val = v; freq = 1; }
    }

    class DLL {
        Node head, tail;
        int size;

        DLL() {
            head = new Node(0, 0);
            tail = new Node(0, 0);
            head.next = tail;
            tail.prev = head;
            size = 0;
        }

        void addToFront(Node node) {
            node.next = head.next;
            node.prev = head;
            head.next.prev = node;
            head.next = node;
            size++;
        }

        void remove(Node node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
            size--;
        }

        Node removeLast() {  // Remove LRU from this frequency bucket
            if (size == 0) return null;
            Node last = tail.prev;
            remove(last);
            return last;
        }
    }

    int capacity, minFreq;
    HashMap<Integer, Node> keyMap;      // key → node
    HashMap<Integer, DLL> freqMap;      // freq → DLL of nodes

    LFUCache(int capacity) {
        this.capacity = capacity;
        minFreq = 0;
        keyMap = new HashMap<>();
        freqMap = new HashMap<>();
    }

    // Update frequency: move node from freq bucket to freq+1 bucket
    private void updateFreq(Node node) {
        int freq = node.freq;
        freqMap.get(freq).remove(node);

        // If this bucket is now empty AND it was the minimum, increment minFreq
        if (freqMap.get(freq).size == 0 && freq == minFreq) {
            minFreq++;
        }

        node.freq++;
        freqMap.putIfAbsent(node.freq, new DLL());
        freqMap.get(node.freq).addToFront(node);
    }

    public int get(int key) {
        if (!keyMap.containsKey(key)) return -1;
        Node node = keyMap.get(key);
        updateFreq(node);   // Every access increases frequency
        return node.val;
    }

    public void put(int key, int value) {
        if (capacity == 0) return;

        if (keyMap.containsKey(key)) {
            Node node = keyMap.get(key);
            node.val = value;
            updateFreq(node);
        } else {
            if (keyMap.size() == capacity) {
                // Evict from the minimum frequency bucket (LRU among those)
                DLL minDLL = freqMap.get(minFreq);
                Node evicted = minDLL.removeLast();
                keyMap.remove(evicted.key);
            }

            Node newNode = new Node(key, value);
            keyMap.put(key, newNode);
            freqMap.putIfAbsent(1, new DLL());
            freqMap.get(1).addToFront(newNode);
            minFreq = 1;  // New node always has freq = 1
        }
    }
}
```

**Dry Run:**

```
LFUCache(2)

put(1, 1):  keyMap={1:Node(1,freq=1)}, freqMap={1:[1]}, minFreq=1
put(2, 2):  keyMap={1,2}, freqMap={1:[2,1]}, minFreq=1

get(1):     Node(1) freq 1→2. freqMap={1:[2], 2:[1]}, minFreq=1
            return 1

put(3, 3):  capacity full. Evict from minFreq=1 bucket.
            freqMap[1] = [2], remove last = Node(2). keyMap removes 2.
            Add Node(3,freq=1). freqMap={1:[3], 2:[1]}, minFreq=1

get(2):     Not in keyMap → return -1

get(3):     Node(3) freq 1→2. freqMap={2:[3,1]}, minFreq=2
            return 3
```

---

### 7. FLATTEN MULTILEVEL DOUBLY LINKED LIST

**Problem:** A doubly linked list where some nodes have a `child` pointer pointing to another DLL. Flatten it into a single-level DLL.

```
Input:
1 <-> 2 <-> 3 <-> 4
      |
      7 <-> 8 <-> 9
                  |
                  10 <-> 11

Output:
1 <-> 2 <-> 7 <-> 8 <-> 9 <-> 10 <-> 11 <-> 3 <-> 4
```

**Core idea:** Whenever you encounter a node with a child, insert the entire child list between the current node and its next. Think of it like DFS — go deep first, then continue right.

java

```java
class Node {
    int val;
    Node prev, next, child;
}

Node flatten(Node head) {
    Node curr = head;

    while (curr != null) {
        if (curr.child != null) {
            Node child = curr.child;
            Node next = curr.next;

            // Connect curr to child list head
            curr.next = child;
            child.prev = curr;
            curr.child = null;  // Clear child pointer!

            // Find tail of child list
            Node tail = child;
            while (tail.next != null) {
                tail = tail.next;
            }

            // Connect child list tail to next
            tail.next = next;
            if (next != null) next.prev = tail;
        }
        curr = curr.next;
    }

    return head;
}
```

**Dry Run:**

```
Start: 1 <-> 2 <-> 3 <-> 4
              |
              7 <-> 8

curr=1: no child, move on
curr=2: HAS CHILD!
  child = Node(7)
  next = Node(3)

  2.next = 7, 7.prev = 2, 2.child = null
  Find child tail: 7 -> 8 -> null, tail = Node(8)
  8.next = 3, 3.prev = 8

  List now: 1 <-> 2 <-> 7 <-> 8 <-> 3 <-> 4
