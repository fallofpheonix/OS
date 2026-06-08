# Theory Part 01

### 1. WHAT IS A LINKED LIST?

A Linked List is a linear data structure where:

- Elements are stored in separate memory locations
- Each node stores:
    - **Data** (the actual value)
    - **Reference/address of next node** (pointer)

#### Structure of a Node

```
[data | next]
```

**Example:**

```
10 -> 20 -> 30 -> null
```

Each node points to the next node in sequence.

---

### 2. WHY LINKED LISTS EXIST?

#### Array Limitations

- Fixed size (cannot grow dynamically)
- Expensive insertion/deletion (requires shifting)
- Requires contiguous memory

#### Linked List Advantages

- Dynamic size (grows as needed)
- Efficient insertion/deletion at known positions: O(1)
- Memory can be scattered (non-contiguous)

#### Linked List Disadvantages

- No random access (must traverse linearly): O(n)
- Extra memory for pointers (overhead per node)
- Cannot perform binary search

---

### 3. TYPES OF LINKED LISTS

#### A. Singly Linked List

Each node points to **only the next node**.

```
10 -> 20 -> 30 -> null
```

**Node Structure:**

java

```java
class Node {
    int data;
    Node next;
    
    Node(int data){
        this.data = data;
        this.next = null;
    }
}
```

**Properties:**

- One-directional traversal only
- Simplest form
- Most common in interviews

---

#### B. Doubly Linked List

Each node has **both next and previous pointers**.

```
null <- 10 <-> 20 <-> 30 -> null
```

**Node Structure:**

java

```java
class Node {
    int data;
    Node prev;
    Node next;
}
```

**Advantages:**

- Can traverse both directions (forward & backward)
- Deletion is O(1) if node is known (no need for previous node pointer)

**Disadvantages:**

- Extra memory for previous pointer
- More complex pointer manipulation

---

#### C. Circular Linked List

Last node points **back to head** (creates a circle).

```
10 -> 20 -> 30
^            |
|____________|
```

**Types:**

- Circular singly (one direction loop)
- Circular doubly (both direction loop)

**Uses:**

- Round-robin scheduling
- Cyclic buffer systems
- Continuous playback (music players)

---

### 4. MEMORY REPRESENTATION

#### Array

```
[10][20][30]  ← Contiguous memory locations
```

#### Linked List

```
[10|1000] -> [20|2050] -> [30|null]
  ↑ data              ↑ address
  
Nodes can be anywhere in memory (non-contiguous)
```

**Key Difference:** Linked lists don't need continuous memory blocks.

---

### 5. IMPORTANT TERMINOLOGY

|Term|Meaning|
|---|---|
|**Head**|Pointer to the first node|
|**Tail**|The last node (next = null)|
|**Node**|Single unit containing data + pointer|
|**Next**|Reference/pointer to next node|
|**Prev**|Pointer to previous node (DLL only)|
|**Null**|Marks the end of the list|
|**Reference**|Memory address of a node|

---

### 6. REFERENCE VS OBJECT (CRITICAL)

#### Understanding Pointers

java

```java
Node a = new Node(10);
```

- `a` is a **reference** (stores memory address)
- The actual **object** is in heap memory

#### Aliasing Problem

java

```java
Node a = new Node(10);
Node b = a;
```

- Both `a` and `b` point to **the same object**
- Changing through `a` affects what `b` sees
- **They reference the same node in memory**

This is fundamental to linked list manipulation!

---

### 7. HEAD POINTER

The `head` pointer is your **entry point** to the entire linked list.

```
head
 ↓
[10 | *] -> [20 | *] -> [30 | null]
```

**If head is lost:**

- Entire list becomes inaccessible
- Memory leak (nodes cannot be freed)
- Data is lost forever

**Empty List:**

java

```java
head = null;  // No nodes exist
```

---

### 8. BASIC OPERATIONS

#### A. TRAVERSAL

Visit every node sequentially from head to tail.

**Logic:**

java

```java
Node temp = head;
while(temp != null){
    System.out.println(temp.data);  // Process node
    temp = temp.next;               // Move to next
}
```

**Time Complexity:** O(n) - must visit all n nodes

---

#### B. INSERTION

##### 1. Insert at Beginning (Head)

**Before:**

```
10 -> 20
```

**After inserting 5:**

```
5 -> 10 -> 20
```

**Steps:**

1. Create new node: `Node newNode = new Node(5);`
2. Point new node to current head: `newNode.next = head;`
3. Move head pointer: `head = newNode;`

**Time Complexity:** O(1) ✓ Very efficient!

---

##### 2. Insert at End (Tail)

**Before:**

```
10 -> 20
```

**After inserting 30:**

```
10 -> 20 -> 30
```

**Logic:**

- Traverse until `temp.next == null` (find tail)
- Attach new node: `temp.next = newNode;`

**Time Complexity:** O(n) - must traverse to find tail

---

##### 3. Insert at Specific Position

**Before:**

```
10 -> 20 -> 40
```

**After inserting 30 at position 2:**

```
10 -> 20 -> 30 -> 40
```

**Core Operation:**

java

```java
newNode.next = prev.next;  // Step 1: Save old link
prev.next = newNode;       // Step 2: Insert new node
```

**Important:** Order matters! Must save the old link first!

---

#### C. DELETION

##### 1. Delete Head

**Before:**

```
10 -> 20 -> 30
```

**After:**

```
20 -> 30
```

**Code:**

java

```java
head = head.next;
```

**Time Complexity:** O(1) ✓ Very efficient!

---

##### 2. Delete Tail

Must find the **second-to-last node** because we need to break its link.

**Logic:**

- Traverse until `temp.next.next == null`
- Break link: `temp.next = null;`

**Time Complexity:** O(n)

---

##### 3. Delete Specific Node

**Before:**

```
10 -> 20 -> 30 -> 40
```

**After deleting 30:**

```
10 -> 20 ---------> 40
```

**Operation:**

java

```java
prev.next = curr.next;  // Bypass current node
```

**Time Complexity:** O(1) if you have the previous node pointer

---

#### D. SEARCHING

Find if a value exists in the list.

java

```java
boolean search(int key){
    Node temp = head;
    while(temp != null){
        if(temp.data == key){
            return true;
        }
        temp = temp.next;
    }
    return false;
}
```

**Time Complexity:** O(n) - must check all nodes in worst case

---

### 9. TIME COMPLEXITY TABLE

|Operation|Time|Space|
|---|---|---|
|Traverse all nodes|O(n)|O(1)|
|Access by index|O(n)|O(1)|
|Search for value|O(n)|O(1)|
|Insert at head|O(1)|O(1)|
|Insert at tail|O(n)|O(1)|
|Insert at position|O(n)|O(1)|
|Delete head|O(1)|O(1)|
|Delete tail|O(n)|O(1)|
|Delete by value|O(n)|O(1)|

---

### 10. CORE LINKED LIST PATTERNS

#### Pattern 1: REVERSAL

**Goal:** Reverse the direction of all links

Original: `10 -> 20 -> 30 -> null` Reversed: `null <- 10 <- 20 <- 30`

**Why it's hard:** Changing one pointer can disconnect the list

**Key Insight:** Need 3 pointers to reverse safely:

- `prev` - points to reversed part
- `curr` - current node being reversed
- `next` - save future node before modifying

---

#### Pattern 2: FAST-SLOW POINTER
