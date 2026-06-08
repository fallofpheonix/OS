# Phase 2 Part 01

## PHASE 2: POINTER MASTERY - COMPLETE WITH EXAMPLES & EXPLANATIONS

---

### **PART 1: REVERSAL PATTERN**

#### **What is Reversal?**

**Original List:**

```
10 -> 20 -> 30 -> null
```

**Reversed List:**

```
null <- 10 <- 20 <- 30
```

Or visually:

```
30 -> 20 -> 10 -> null
```

**Goal:** Change the direction of all pointers (links) in the list.

---

### **1. WHY REVERSAL IS HARD**

#### The Problem

Students think: "Just change the arrows!"

But the real issue:

- **If you change one pointer without saving the next node, you lose all remaining nodes!**

#### Example of Losing Nodes

java

```java
// ❌ WRONG APPROACH
curr.next = prev;           // Change pointer
curr = curr.next;           // Try to move forward
// ERROR: curr.next is now pointing backward!
// You lost the forward link!
```

**Visual:**

```
Before:  curr -> [10 | *] -> [20 | next]

Step 1:  curr.next = prev
         curr -> [10 | prev]  (backward link)

Step 2:  curr = curr.next
         curr -> points to prev (not forward!)
         The [20 | next] is now LOST!
```

---

### **2. THE CORE IDEA: THREE POINTERS**

You need **3 pointers** to safely reverse:

|Pointer|Purpose|Location|
|---|---|---|
|`prev`|Points to reversed part|Behind current|
|`curr`|Current node being reversed|Current position|
|`next`|Save the future node|Ahead of current|

#### Visual Concept

```
During Reversal:

REVERSED PART    CURRENT    UNVISITED PART
    ↓               ↓              ↓
null <- 10 <- 20   30  ->  40 -> 50 -> null
        ↑           ↑        ↑
       prev        curr     next
```

---

### **3. COMPLETE ITERATIVE REVERSAL**

#### Code

java

```java
Node reverse(Node head) {
    Node prev = null;      // Start: no previous
    Node curr = head;      // Start: at head
    
    while(curr != null) {
        // Step 1: Save the future
        Node next = curr.next;
        
        // Step 2: Reverse the link
        curr.next = prev;
        
        // Step 3: Move prev forward
        prev = curr;
        
        // Step 4: Move curr forward
        curr = next;
    }
    
    return prev;  // New head!
}
```

#### Critical Understanding

```
Why return prev, not curr?

After loop:
- curr = null (reached end)
- prev = last node (which is new head)
```

---

### **4. DETAILED DRY RUN - STEP BY STEP**

**Original List:** 10 -> 20 -> 30 -> null

#### **Initial State**

```
prev = null
curr = Node(10)

List: [10 | *] -> [20 | *] -> [30 | null]
```

---

#### **ITERATION 1**

**Step 1: Save next**

```
next = curr.next = Node(20)
```

**Step 2: Reverse link**

```
curr.next = prev
[10 | null]  (now points back to null)
```

**Step 3: Move prev**

```
prev = curr = Node(10)
```

**Step 4: Move curr**

```
curr = next = Node(20)
```

**After Iteration 1:**

```
prev = [10 | null]  (REVERSED PART)
curr = [20 | *]     (CURRENT NODE)
next = [30 | null]  (FUTURE)

Visual:
null <- [10]    [20] -> [30] -> null
         ↑       ↑
        prev    curr
```

---

#### **ITERATION 2**

**Step 1: Save next**

```
next = curr.next = Node(30)
```

**Step 2: Reverse link**

```
curr.next = prev
[20 | prev] where prev = Node(10)
```

**Step 3: Move prev**

```
prev = curr = Node(20)
```

**Step 4: Move curr**

```
curr = next = Node(30)
```

**After Iteration 2:**

```
prev = [20 | *]     (REVERSED PART)
         ↓
       [10 | null]

curr = [30 | *]     (CURRENT NODE)
next = null         (FUTURE)

Visual:
[20 | *] -> [10 | null]    [30] -> null
 ↑                         ↑
prev                      curr
```

---

#### **ITERATION 3**

**Step 1: Save next**

```
next = curr.next = null
```

**Step 2: Reverse link**

```
curr.next = prev
[30 | prev] where prev = Node(20)
```

**Step 3: Move prev**

```
prev = curr = Node(30)
```

**Step 4: Move curr**

```
curr = next = null
```

**After Iteration 3:**

```
prev = [30 | *]  (NEW HEAD!)
         ↓
       [20 | *]
         ↓
       [10 | null]

curr = null  (LOOP EXITS)

Visual:
[30 | *] -> [20 | *] -> [10 | null]
 ↑
prev (this becomes the new head)
```

---

#### **Loop Condition: curr != null?**

```
curr = null → FALSE
Loop exits!
```

---

#### **Return:**

```
return prev = Node(30)
```

**Final Reversed List:**

```
30 -> 20 -> 10 -> null
```

---

### **5. COMPLETE WORKING CODE WITH EXAMPLE**

java

```java
class Node {
    int data;
    Node next;
    
    Node(int data) {
        this.data = data;
        this.next = null;
    }
}

class LinkedList {
    Node head;
    
    void traverse() {
        Node temp = head;
        while(temp != null) {
            System.out.print(temp.data + " -> ");
            temp = temp.next;
        }
        System.out.println("null");
    }
    
    void insertAtTail(int data) {
        Node newNode = new Node(data);
        if(head == null) {
            head = newNode;
            return;
        }
        Node temp = head;
        while(temp.next != null) {
            temp = temp.next;
        }
        temp.next = newNode;
    }
    
    // ⭐ REVERSE ITERATIVE
    void reverseIterative() {
        Node prev = null;
        Node curr = head;
        
        System.out.println("\n=== REVERSAL PROCESS ===");
        int step = 1;
        
        while(curr != null) {
            // Save next
            Node next = curr.next;
            System.out.println("Step " + step + ": curr=" + curr.data + 
                             ", next=" + (next != null ? next.data : "null") + 
                             ", prev=" + (prev != null ? prev.data : "null"));
            
            // Reverse
            curr.next = prev;
            
            // Move pointers
            prev = curr;
            curr = next;
            step++;
        }
        
        head = prev;  // New head
        System.out.println("Reversed! New head: " + head.data);
    }
    
    // ⭐ REVERSE RECURSIVE
    Node reverseRecursive(Node head) {
        // Base case: empty or single node
        if(head == null || head.next == null) {
            return head;
        }
        
        // Recursive case
        Node newHead = reverseRecursive(head.next);
        
        // Create backward connection
        head.next.next = head;  // CRITICAL LINE
        
        // Break forward connection
        head.next = null;
        
        return newHead;
    }
}

// TEST
public class Main {
    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        
        // Create list: 10 -> 20 -> 30 -> 40 -> null
        list.insertAtTail(10);
        list.insertAtTail(20);
        list.insertAtTail(30);
        list.insertAtTail(40);
        
        System.out.println("Original List:");
        list.traverse();  // Output: 10 -> 20 -> 30 -> 40 -> null
        
        list.reverseIterative();
        
        System.out.println("\nReversed List:");
        list.traverse();  // Output: 40 -> 30 -> 20 -> 10 -> null
    }
}
```

#### Output:

```
Original List:
10 -> 20 -> 30 -> 40 -> null

=== REVERSAL PROCESS ===
Step 1: curr=10, next=20, prev=null
Step 2: curr=20, next=30, prev=10
Step 3: curr=30, next=40, prev=20
Step 4: curr=40, next=null, prev=30
Reversed! New head: 40

Reversed List:
40 -> 30 -> 20 -> 10 -> null
```

---

### **6. RECURSIVE REVERSAL**

#### Code

java

```java
Node reverseRecursive(Node head) {
    // Base case
    if(head == null || head.next == null) {
        return head;
    }
    
    // Recursive case
    Node newHead = reverseRecursive(head.next);
    
    // Critical step: create backward link
    head.next.next = head;
