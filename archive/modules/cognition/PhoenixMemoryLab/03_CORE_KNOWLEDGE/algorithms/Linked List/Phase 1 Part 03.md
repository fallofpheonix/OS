# Phase 1 Part 03


Result: 10 -> 20 -> 30 -> null
```

##### Time Complexity

- O(n)

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(20);
list.insertAtTail(30);

list.traverse();      // Output: 10 -> 10 -> 20 -> 20 -> 30 -> null
list.deleteDuplicates();
list.traverse();      // Output: 10 -> 20 -> 30 -> null
```

---

### **11. PRINT REVERSE (Without Reversing List)**

##### Theory

```
Use recursion: process recursively, then print while returning
OR use reverse traversal with stack (advanced)
```

##### Code

java

```java
void printReverse(Node head) {
    if(head == null) return;
    
    printReverse(head.next);  // Recursive call first
    System.out.print(head.data + " -> ");  // Print while returning
}

// Call it like:
void printReverse() {
    printReverse(head);
    System.out.println("null");
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> null

```
Call: printReverse(Node(10))
  |
  +-> Call: printReverse(Node(20))
        |
        +-> Call: printReverse(Node(30))
              |
              +-> Call: printReverse(null)
                    return (base case)
                    
              Print: 30 -> 
              Return to Node(20)
        
        Print: 20 -> 
        Return to Node(10)
  
  Print: 10 -> 
  Return

Output: 30 -> 20 -> 10 -> null
```

##### Time Complexity

- O(n) time
- O(n) space (recursion stack)

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);

list.traverse();      // Output: 10 -> 20 -> 30 -> null
list.printReverse();  // Output: 30 -> 20 -> 10 -> null
```

---

### **12. COMMON MISTAKES & HOW TO AVOID**

#### Mistake 1: Losing Forward Link

❌ **WRONG:**

java

```java
curr.next = prev;
curr = curr.next;  // ❌ curr.next is now pointing backward!
```

✅ **RIGHT:**

java

```java
Node next = curr.next;  // Save forward first
curr.next = prev;
curr = next;  // Use saved reference
```

---

#### Mistake 2: Null Pointer Exception

❌ **WRONG:**

java

```java
while(temp != null) {
    System.out.println(temp.next.next);  // ❌ Crashes if temp.next is null
}
```

✅ **RIGHT:**

java

```java
while(temp != null && temp.next != null) {
    System.out.println(temp.next.next);  // ✓ Safe
}
```

---

#### Mistake 3: Infinite Loop

❌ **WRONG:**

java

```java
while(temp != null) {
    System.out.println(temp.data);
    // Missing: temp = temp.next;
}  // ❌ Infinite loop!
```

✅ **RIGHT:**

java

```java
while(temp != null) {
    System.out.println(temp.data);
    temp = temp.next;  // ✓ Move forward
}
```

---

#### Mistake 4: Losing Head

❌ **WRONG:**

java

```java
void someOperation() {
    Node temp = head;
    head = head.next;  // ❌ Lost first node
    // ... rest of operations
}
```

✅ **RIGHT:**

java

```java
void someOperation() {
    Node temp = head;
    // Don't modify head unless intentional!
}
```

---

#### Mistake 5: Not Handling Empty List

❌ **WRONG:**

java

```java
void deleteHead() {
    head = head.next;  // ❌ Crashes if head is null
}
```

✅ **RIGHT:**

java

```java
void deleteHead() {
    if(head == null) return;  // ✓ Check first
    head = head.next;
}
```

---

### **COMPLEXITY SUMMARY TABLE**

|Operation|Time|Space|Notes|
|---|---|---|---|
|Traverse|O(n)|O(1)|Visit all nodes|
|Create empty|O(1)|O(1)|-|
|Insert at head|O(1)|O(1)|No traversal|
|Insert at tail|O(n)|O(1)|Must find tail|
|Insert at position|O(n)|O(1)|Must find position|
|Delete head|O(1)|O(1)|No traversal|
|Delete tail|O(n)|O(1)|Must find 2nd-to-last|
|Delete at position|O(n)|O(1)|Must find position|
|Search|O(n)|O(1)|Worst case entire list|
|Find length|O(n)|O(1)|Count all nodes|
|Find middle|O(n)|O(1)|Half traversal|
|Count occurrences|O(n)|O(1)|Check all nodes|

---

### **COMPLETE WORKING CODE**

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
    
    // 1. TRAVERSE
    void traverse() {
        Node temp = head;
        while(temp != null) {
            System.out.print(temp.data + " -> ");
            temp = temp.next;
        }
        System.out.println("null");
    }
    
    // 2. INSERT AT HEAD
    void insertAtHead(int data) {
        Node newNode = new Node(data);
        newNode.next = head;
        head = newNode;
    }
    
    // 3. INSERT AT TAIL
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
    
    // 4. INSERT AT POSITION
    void insertAtPosition(int data, int position) {
        if(position == 0) {
            insertAtHead(data);
            return;
        }
        Node newNode = new Node(data);
        Node temp = head;
        for(int i = 0; i < position - 1; i++) {
            if(temp == null) return;
            temp = temp.next;
        }
        if(temp == null) return;
        newNode.next = temp.next;
        temp.next = newNode;
    }
    
    // 5. DELETE HEAD
    void deleteHead() {
        if(head == null) return;
        head = head.next;
    }
    
    // 6. DELETE TAIL
    void deleteTail() {
        if(head == null) return;
        if(head.next == null) {
            head = null;
            return;
        }
        Node temp = head;
        while(temp.next.next != null) {
            temp = temp.next;
        }
        temp.next = null;
    }
    
    // 7. DELETE AT POSITION
    void deleteAtPosition(int position) {
        if(position == 0) {
            deleteHead();
            return;
        }
        Node temp = head;
        for(int i = 0; i < position - 1; i++) {
            if(temp == null) return;
            temp = temp.next;
        }
        if(temp == null || temp.next == null) return;
        temp.next = temp.next.next;
    }
    
    // 8. SEARCH
    boolean search(int key) {
        Node temp = head;
        while(temp != null) {
            if(temp.data == key) return true;
            temp = temp.next;
        }
        return false;
    }
    
    // 9. FIND LENGTH
    int findLength() {
        int count = 0;
        Node temp = head;
        while(temp != null) {
            count++;
            temp = temp.next;
        }
        return count;
    }
    
    // 10. PRINT MIDDLE
    void printMiddle() {
        if(head == null) return;
        int length = findLength();
        int middlePos = length / 2;
        Node temp = head;
        for(int i = 0; i < middlePos; i++) {
            temp = temp.next;
        }
        System.out.println("Middle: " + temp.data);
    }
    
    // 11. COUNT OCCURRENCES
    int countOccurrences(int key) {
        int count = 0;
        Node temp = head;
        while(temp != null) {
            if(temp.data == key) count++;
            temp = temp.next;
        }
        return count;
    }
    
    // 12. DELETE DUPLICATES
    void deleteDuplicates() {
        if(head == null) return;
        Node temp = head;
        while(temp != null && temp.next != null) {
            if(temp.data == temp.next.data) {
                temp.next = temp.next.next;
            } else {
                temp = temp.next;
            }
        }
    }
    
    // 13. PRINT REVERSE
    void printReverse(Node head) {
        if(head == null) return;
        printReverse(head.next);
        System.out.print(head.data + " -> ");
    }
    
    void printReverse() {
        printReverse(head);
        System.out.println("null");
    }
}

// MAIN - Test all operations
public class Main {
    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        
        System.out.println("=== INSERT OPERATIONS ===");
        list.insertAtTail(10);
        list.insertAtTail(20);
        list.insertAtTail(40);
        list.traverse();  // 10 -> 20 -> 40 -> null
        
        list.insertAtPosition(30, 2);
        list.traverse();  // 10 -> 20 -> 30 -> 40 -> null
        
        list.insertAtHead(5);
        list.traverse();  // 5 -> 10 -> 20 -> 30 -> 40 -> null
        
        System.out.println("\n=== SEARCH & LENGTH ===");
        System.out.println("Search 30: " + list.search(30));  // true
        System.out.println("Search 99: " + list.search(99));  // false
        System.out.println("Length: " + list.findLength());  // 5
        
        System.out.println("\n=== MIDDLE & COUNT ===");
        list.printMiddle();  // Middle: 20
        System.out.println("Count 20: " + list.countOccurrences(20));  // 1
        
        System.out.println("\n=== DELETION ===");
        list.deleteHead();
        list.traverse();  // 10 -> 20 -> 30 -> 40 -> null
        
        list.deleteTail();
