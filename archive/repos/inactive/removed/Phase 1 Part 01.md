# Phase 1 Part 01

## PHASE 1: LINKED LIST FOUNDATIONS - WITH EXAMPLES & EXPLANATIONS

---

### **1. UNDERSTAND THE NODE**

#### Basic Node Structure

java

```java
class Node {
    int data;      // The actual value stored
    Node next;     // Reference to next node
    
    Node(int data) {
        this.data = data;
        this.next = null;  // Initially points to nothing
    }
}
```

#### Visualizing Nodes

```
Single Node:
[10 | null]

Chain of 3 nodes:
[10 | *] -> [20 | *] -> [30 | null]
 ↑           ↑           ↑
data        data        data

Where * means "address of next node"
```

#### Example: Creating Nodes

java

```java
// Create 3 separate nodes
Node node1 = new Node(10);
Node node2 = new Node(20);
Node node3 = new Node(30);

// Connect them
node1.next = node2;  // node1 points to node2
node2.next = node3;  // node2 points to node3
node3.next = null;   // node3 points to nothing (end)

// Now we have: 10 -> 20 -> 30 -> null
```

#### Understanding References

java

```java
Node a = new Node(10);
Node b = a;  // b now points to SAME node as a

a.data = 15;  // Change through a
System.out.println(b.data);  // Prints 15! (same object)

// a and b are two references to the SAME object in memory
```

**Key Insight:** Variables store references (memory addresses), not objects themselves!

---

### **2. HEAD POINTER - The Entry Point**

java

```java
public class LinkedList {
    Node head;  // Entry point to entire list
    
    LinkedList() {
        head = null;  // Empty list initially
    }
}
```

#### Visualization

```
Empty List:
head = null

After adding first node:
head → [10 | null]

After adding more nodes:
head → [10 | *] → [20 | *] → [30 | null]
```

**Critical Rule:** If you lose `head`, the entire list is lost!

---

### **3. TRAVERSAL - Visit Every Node**

#### Theory

```
Start at head
While current node is not null:
    Process current node
    Move to next node
Stop when you reach null
```

#### Code Example

java

```java
void traverse() {
    Node temp = head;
    
    while(temp != null) {
        System.out.print(temp.data + " -> ");
        temp = temp.next;
    }
    System.out.println("null");
}
```

#### Example Execution

**List:** 10 -> 20 -> 30 -> null

```
Iteration 1:  temp = Node(10)  →  Print "10"  →  temp = temp.next = Node(20)
Iteration 2:  temp = Node(20)  →  Print "20"  →  temp = temp.next = Node(30)
Iteration 3:  temp = Node(30)  →  Print "30"  →  temp = temp.next = null
Iteration 4:  temp = null      →  STOP
```

**Output:** 10 -> 20 -> 30 -> null

#### Time Complexity

- Must visit all n nodes: **O(n)**
- No extra space: **O(1)**

---

### **4. INSERTION OPERATIONS**

#### A. INSERT AT HEAD (Beginning)

##### Theory

```
Before: 10 -> 20
After:  5 -> 10 -> 20

Steps:
1. Create new node
2. Point new node to current head
3. Make new node the head
```

##### Code

java

```java
void insertAtHead(int data) {
    Node newNode = new Node(data);
    newNode.next = head;  // New node points to old head
    head = newNode;       // New node becomes head
}
```

##### Dry Run Example

**Initial List:** 10 -> 20 -> 30 -> null **Insert:** 5

```
Step 1: Create newNode = Node(5)
        newNode.next = null (initial)

Step 2: newNode.next = head
        newNode → [5 | *]
                    ↓
                  [10 | *] -> [20 | *] -> [30 | null]

Step 3: head = newNode
        head is now pointing to [5 | *]

Result: 5 -> 10 -> 20 -> 30 -> null
```

##### Time Complexity

- O(1) - No traversal needed!

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtHead(30);  // List: 30
list.insertAtHead(20);  // List: 20 -> 30
list.insertAtHead(10);  // List: 10 -> 20 -> 30
list.traverse();        // Output: 10 -> 20 -> 30 -> null
```

---

#### B. INSERT AT END (Tail)

##### Theory

```
Before: 10 -> 20
After:  10 -> 20 -> 30

Steps:
1. Create new node
2. Traverse to tail (find node where next is null)
3. Make tail point to new node
```

##### Code

java

```java
void insertAtTail(int data) {
    Node newNode = new Node(data);
    
    if(head == null) {
        head = newNode;  // Empty list case
        return;
    }
    
    Node temp = head;
    while(temp.next != null) {  // Find tail
        temp = temp.next;
    }
    
    temp.next = newNode;  // Attach new node
}
```

##### Dry Run Example

**List:** 10 -> 20 -> null **Insert:** 30

```
Step 1: newNode = Node(30)

Step 2: Find tail
        temp = head = Node(10)
        temp.next = Node(20) (not null, continue)
        temp = temp.next = Node(20)
        temp.next = null (STOP - found tail)

Step 3: temp.next = newNode
        Node(20).next = Node(30)

Result: 10 -> 20 -> 30 -> null
```

##### Time Complexity

- Must traverse to tail: **O(n)**
- No extra space: **O(1)**

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);  // List: 10
list.insertAtTail(20);  // List: 10 -> 20
list.insertAtTail(30);  // List: 10 -> 20 -> 30
list.traverse();        // Output: 10 -> 20 -> 30 -> null
```

---

#### C. INSERT AT SPECIFIC POSITION

##### Theory

```
Before: 10 -> 20 -> 40
Insert 30 at position 2 (0-indexed)
After:  10 -> 20 -> 30 -> 40

Steps:
1. Find the node BEFORE insertion point (position 1)
2. Save its next pointer
3. Insert new node in between
4. Connect new node to saved pointer
```

##### Code

java

```java
void insertAtPosition(int data, int position) {
    if(position == 0) {
        insertAtHead(data);  // Special case
        return;
    }
    
    Node newNode = new Node(data);
    Node temp = head;
    
    // Find node at position-1
    for(int i = 0; i < position - 1; i++) {
        if(temp == null) return;  // Invalid position
        temp = temp.next;
    }
    
    if(temp == null) return;  // Invalid position
    
    // Insert new node
    newNode.next = temp.next;  // NEW node points to OLD next
    temp.next = newNode;        // OLD node points to NEW node
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 40 -> null **Insert:** 30 at position 2

```
Step 1: Find position-1 = position 1
        temp = head = Node(10)
        i = 0: temp = temp.next = Node(20)
        Loop ends (i < 1)

Step 2: newNode = Node(30)

Step 3: newNode.next = temp.next
        newNode.next = Node(40)
        
        temp.next = newNode
        Node(20).next = Node(30)

Step 4: Structure
        [10 | *] -> [20 | *] -> [30 | *] -> [40 | null]
                        ↑            ↑          ↑
                       temp      newNode    saved link

Result: 10 -> 20 -> 30 -> 40 -> null
```

##### Time Complexity

- Traverse to position: **O(n)**
- No extra space: **O(1)**

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(40);
list.insertAtPosition(30, 2);  // Insert 30 at position 2
list.traverse();               // Output: 10 -> 20 -> 30 -> 40 -> null
```

---

### **5. DELETION OPERATIONS**

#### A. DELETE HEAD

##### Theory

```
Before: 10 -> 20 -> 30
After:  20 -> 30

Just move head to next node!
```

##### Code

java

```java
void deleteHead() {
    if(head == null) return;  // Empty list
    head = head.next;         // Skip first node
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> null

```
Before: head → [10 | *]
                   ↓
               [20 | *] -> [30 | null]

After:  head → [20 | *] -> [30 | null]

The [10 | *] node is now unreachable and can be garbage collected
```

##### Time Complexity

- O(1) - Just change one pointer!

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);

list.traverse();       // Output: 10 -> 20 -> 30 -> null
list.deleteHead();
list.traverse();       // Output: 20 -> 30 -> null
```

---

#### B. DELETE TAIL

##### Theory

```
Before: 10 -> 20 -> 30
After:  10 -> 20

Steps:
1. Find second-to-last node (where next.next == null)
2. Set its next to null
