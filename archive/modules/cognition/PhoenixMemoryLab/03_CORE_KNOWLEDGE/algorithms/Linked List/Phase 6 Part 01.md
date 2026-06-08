# Phase 6 Part 01

PHASE 6: LINKED LISTS AS A FOUNDATION FOR ADVANCED DATA STRUCTURES — With Questions, Examples & Explanations

This phase is about zooming out. Instead of just solving linked list problems, you now understand _why_ linked lists exist everywhere in computer science — and how mastering them unlocks trees, graphs, caches, and real systems.

---

### 1. THE BIG REALIZATION: LINKED LISTS ARE POINTER NETWORKS

Every advanced data structure is really a network of nodes connected by pointers. Mastering linked lists means you've already mastered the core skill — pointer reasoning. The rest is just applying it in different shapes.

---

### 2. LINKED LIST → STACK

**Concept:** A stack is simply a linked list where you can only touch the head.

```
TOP
 ↓
10 -> 20 -> 30 -> null
```

Push = insertAtHead, Pop = deleteHead — both O(1).

**Question: Implement Stack Using Linked List**

java

```java
class Stack {
    Node top;

    void push(int data) {
        Node newNode = new Node(data);
        newNode.next = top;
        top = newNode;
    }

    int pop() {
        if (top == null) throw new RuntimeException("Stack empty");
        int val = top.data;
        top = top.next;
        return val;
    }

    int peek() {
        if (top == null) throw new RuntimeException("Stack empty");
        return top.data;
    }
}
```

**Dry Run:**

```
push(10): top → [10|null]
push(20): top → [20|*] → [10|null]
push(30): top → [30|*] → [20|*] → [10|null]
pop():    returns 30, top → [20|*] → [10|null]
peek():   returns 20
```

**Why LL beats Array for Stack:** No fixed size. You never need to resize.

---

### 3. LINKED LIST → QUEUE

**Concept:** A queue needs O(1) enqueue AND O(1) dequeue. With a plain array or singly LL with only a head pointer, one of those is O(n). The trick: maintain both head AND tail pointers.

```
front → [10] → [20] → [30] ← rear
```

Enqueue = insert at tail (using tail pointer = O(1)), Dequeue = delete at head = O(1).

**Question: Implement Queue Using Linked List**

java

```java
class Queue {
    Node front, rear;

    void enqueue(int data) {
        Node newNode = new Node(data);
        if (rear != null) rear.next = newNode;
        rear = newNode;
        if (front == null) front = rear;
    }

    int dequeue() {
        if (front == null) throw new RuntimeException("Queue empty");
        int val = front.data;
        front = front.next;
        if (front == null) rear = null;  // Queue is now empty
        return val;
    }
}
```

**Dry Run:**

```
enqueue(10): front=rear=[10]
enqueue(20): front=[10]→[20]=rear
enqueue(30): front=[10]→[20]→[30]=rear
dequeue():   returns 10, front=[20]→[30]=rear
dequeue():   returns 20, front=rear=[30]
```

**The key insight:** You need the tail pointer. Without it, every enqueue is O(n). This is a pattern you'll see in BFS — queues everywhere.

---

### 4. LINKED LIST → HASHMAP (Separate Chaining)

**Concept:** When two keys hash to the same index (collision), you need somewhere to put both. The answer is a linked list at that bucket.

```
index 0: null
index 1: [key=1, val="A"] → [key=11, val="K"] → null
index 2: [key=2, val="B"] → null
index 3: null
```

**Question: Implement a Simple HashMap Using Linked Lists**

java

```java
class HashMap {
    static final int SIZE = 10;
    Node[] buckets = new Node[SIZE];

    static class Node {
        int key;
        String val;
        Node next;
        Node(int k, String v) { key = k; val = v; }
    }

    int hash(int key) {
        return key % SIZE;
    }

    void put(int key, String value) {
        int idx = hash(key);
        Node curr = buckets[idx];

        // Update if key exists
        while (curr != null) {
            if (curr.key == key) { curr.val = value; return; }
            curr = curr.next;
        }

        // Insert at head of bucket
        Node newNode = new Node(key, value);
        newNode.next = buckets[idx];
        buckets[idx] = newNode;
    }

    String get(int key) {
        int idx = hash(key);
        Node curr = buckets[idx];
        while (curr != null) {
            if (curr.key == key) return curr.val;
            curr = curr.next;
        }
        return null;
    }
}
```

**Dry Run:**

```
put(1, "A"):  hash=1, bucket[1] → [1,"A"]
put(11, "K"): hash=1, bucket[1] → [11,"K"] → [1,"A"]
get(11):      hash=1, scan bucket → found "K" ✓
get(1):       hash=1, scan bucket → found "A" ✓
```

**Why this matters:** Java's actual HashMap uses this exact pattern. When a bucket's chain grows too long (>8 nodes), Java converts it to a Red-Black Tree. You now understand that internal detail.

---

### 5. LINKED LIST → LRU CACHE (DLL + HashMap)

You already implemented this in Phase 4. Here the focus is understanding _why_ it's a doubly linked list, not a singly linked list, and what that means for system design.

**Question: Why does LRU Cache need a DLL and not a singly LL?**

To move a node to "most recently used" position, you need to remove it from its current position. Removal from a singly LL requires the _previous_ node, which means O(n) traversal to find it. A DLL gives you the prev pointer for free, making removal O(1).

```
Singly LL:  To remove Node X, traverse to find X.prev → O(n)
Doubly LL:  Node X has X.prev built in → O(1)
```

This is the entire reason LRU Cache uses DLL. One pointer (prev) changes the whole complexity.

**Real system analogy:**

```
CPU Cache:       L1 → L2 → L3 → RAM  (LRU eviction at each level)
Browser:         Cache evicts least recently visited pages
Redis:           Uses LRU eviction policy for memory management
```

---

### 6. LINKED LIST → TREE

**Concept:** A tree node is just a linked list node with two next pointers instead of one.

java

```java
// Linked list node
class LLNode {
    int data;
    LLNode next;       // one pointer
}

// Tree node  
class TreeNode {
    int data;
    TreeNode left;     // two pointers
    TreeNode right;
}
```

**Question: Convert Sorted Linked List to Binary Search Tree**

This bridges both worlds directly.

java

```java
Node middleNode(Node head) {
    Node slow = head, fast = head, prev = null;
    while (fast != null && fast.next != null) {
        prev = slow;
        slow = slow.next;
        fast = fast.next.next;
    }
    if (prev != null) prev.next = null;  // Split list
    return slow;
}

TreeNode sortedListToBST(Node head) {
    if (head == null) return null;
    if (head.next == null) return new TreeNode(head.data);

    Node mid = middleNode(head);

    TreeNode root = new TreeNode(mid.data);
    root.left = sortedListToBST(head);     // Left half (head to mid-1)
    root.right = sortedListToBST(mid.next); // Right half
    return root;
}
```

**Dry Run:**

```
List: 1 → 2 → 3 → 4 → 5

middleNode finds 3, splits into [1→2] and [4→5]

BST:        3
           / \
          2   4
         /     \
        1       5
```

You used fast-slow (Phase 2 skill) + split (Phase 3 skill) to solve a tree problem. That's the cross-topic synthesis Phase 6 is about.

---

### 7. LINKED LIST → GRAPH (Adjacency List)

**Concept:** A graph's adjacency list is literally an array of linked lists. Each index is a node, and its linked list stores the neighbors.

```
Graph: 0—1—2
           |
           3

Adjacency List:
0: [1]
1: [0, 2, 3]
2: [1]
3: [1]
```

**Question: Build and Traverse a Graph Using Linked List Adjacency List**

java

```java
class Graph {
    int vertices;
    Node[] adj;  // Array of linked lists

    Graph(int v) {
        vertices = v;
        adj = new Node[v];
    }

    void addEdge(int u, int v) {
        // Add v to u's list
        Node newNode = new Node(v);
        newNode.next = adj[u];
        adj[u] = newNode;

        // Add u to v's list (undirected)
        newNode = new Node(u);
        newNode.next = adj[v];
        adj[v] = newNode;
    }

    void bfs(int start) {
        boolean[] visited = new boolean[vertices];
        Queue<Integer> queue = new LinkedList<>();

        visited[start] = true;
        queue.add(start);

        while (!queue.isEmpty()) {
            int node = queue.poll();
            System.out.print(node + " ");

            Node curr = adj[node];
            while (curr != null) {
                if (!visited[curr.data]) {
                    visited[curr.data] = true;
                    queue.add(curr.data);
                }
                curr = curr.next;
            }
        }
    }
}
```

**Dry Run (BFS from 0):**

```
Visit 0. Queue: [1]
Visit 1. Queue: [0,2,3] → 0 already visited → Queue: [2,3]
Visit 2. Queue: [3,1] → 1 visited → Queue: [3]
Visit 3. Queue: [1] → 1 visited → Queue: []

Output: 0 1 2 3
```

**The linked list skills you used:** traversal of the adjacency list chain, visited check (like cycle detection), queue (which you just built from LL).

---

### 8. LINKED LIST → UNDO/REDO SYSTEM

**Question: Design an Undo/Redo System**

java

```java
class UndoRedo {
    class Node {
        String state;
        Node prev, next;
        Node(String s) { state = s; }
    }

    Node current;

    UndoRedo(String initialState) {
        current = new Node(initialState);
    }

    void execute(String newState) {
        Node newNode = new Node(newState);
        newNode.prev = current;
        current.next = newNode;
        // Future redo history is erased
        current = newNode;
    }

    String undo() {
        if (current.prev == null) return current.state;
        current = current.prev;
        return current.state;
    }

    String redo() {
        if (current.next == null) return current.state;
        current = current.next;
        return current.state;
    }
}
```

**Dry Run:**

```
execute("type A"): A
execute("type B"): A ↔ B (current)
execute("type C"): A ↔ B ↔ C (current)
undo():            A ↔ B (current) ↔ C → return "type B"
undo():            A (current) ↔ B ↔ C → return "type A"
redo():            A ↔ B (current) ↔ C → return "type B"
execute("type X"): A ↔ B ↔ X (current)  — C is gone!
```

Notice `execute()` deliberately doesn't connect `newNode.next`. When you take a new action after undoing, the redo history is erased. Exactly like browser forward history disappearing when you visit a new page.

---

### 9. LINKED LIST → MEMORY ALLOCATOR (Free List)

**Concept:** Operating systems track available memory blocks using a linked list called a "free list." Allocating memory = removing a node. Freeing memory = inserting a node back.

```
freeList: [block at 0x1000, size=100] → [block at 0x2500, size=200] → null

malloc(50):  Remove [0x1000, size=100], split → [0x1000, size=50] used + [0x1050, size=50] back to free list
free(0x1000): Insert [0x1000, size=50] back into free list
```

You won't implement this in interviews, but understanding it means you now know why `new Node()` in Java is fast (heap allocator uses a free list internally).

---

### 10. COMPLETE TRANSITION MAP

|Linked List Skill|Where It Appears|
|---|---|
|Traversal|Graph BFS/DFS, tree traversal|
|Reversal|Tree inversion, expression parsing|
|Fast-Slow|Cycle detection in any pointer structure|
|Dummy node|Complex structural edits in trees|
|Merge|Merge sort, heap-based merging|
|DLL|LRU cache, undo/redo, browser history|
|Circular LL|OS schedulers, playlist systems|
|Random pointer|Graph cloning, deep copying structures|
|Split/reconnect|BST construction, balanced tree building|

