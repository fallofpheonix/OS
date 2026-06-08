# Phase 1 Part 02

```

##### Code

java

```java
void deleteTail() {
    if(head == null) return;  // Empty
    if(head.next == null) {   // Single node
        head = null;
        return;
    }
    
    Node temp = head;
    while(temp.next.next != null) {  // Find second-to-last
        temp = temp.next;
    }
    
    temp.next = null;  // Disconnect last node
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> null

```
Step 1: Find second-to-last node
        temp = Node(10)
        temp.next.next = Node(30) (not null, continue)
        temp = temp.next = Node(20)
        temp.next.next = null (STOP - found!)

Step 2: temp.next = null
        Node(20).next = null

Before: [10 | *] -> [20 | *] -> [30 | null]
After:  [10 | *] -> [20 | null]

The [30 | null] node is now unreachable
```

##### Time Complexity

- Must traverse to second-to-last: **O(n)**

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);

list.traverse();       // Output: 10 -> 20 -> 30 -> null
list.deleteTail();
list.traverse();       // Output: 10 -> 20 -> null
```

---

#### C. DELETE SPECIFIC NODE (By Position)

##### Theory

```
Before: 10 -> 20 -> 30 -> 40
Delete node at position 2 (value 30)
After:  10 -> 20 -> 40

Steps:
1. Find node before target
2. Skip the target node
```

##### Code

java

```java
void deleteAtPosition(int position) {
    if(position == 0) {
        deleteHead();  // Special case
        return;
    }
    
    Node temp = head;
    
    // Find node before position
    for(int i = 0; i < position - 1; i++) {
        if(temp == null) return;
        temp = temp.next;
    }
    
    if(temp == null || temp.next == null) return;
    
    temp.next = temp.next.next;  // Bypass the target node
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> 40 -> null **Delete at position:** 2 (delete 30)

```
Step 1: Find position-1 = position 1
        temp = head = Node(10)
        i = 0: temp = temp.next = Node(20)
        Loop ends

Step 2: temp.next = temp.next.next
        
Before: Node(20).next = Node(30)
        Node(30).next = Node(40)

After:  Node(20).next = Node(40)

Structure:
[10 | *] -> [20 | *] ---------> [40 | null]
                      ↑ skipped
                    [30 | *]

Result: 10 -> 20 -> 40 -> null
```

##### Time Complexity

- Traverse to position: **O(n)**

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);
list.insertAtTail(40);

list.traverse();          // Output: 10 -> 20 -> 30 -> 40 -> null
list.deleteAtPosition(2); // Delete 30
list.traverse();          // Output: 10 -> 20 -> 40 -> null
```

---

### **6. SEARCHING**

##### Theory

```
Traverse list and check each node's value
Return true if found, false if reached null
```

##### Code

java

```java
boolean search(int key) {
    Node temp = head;
    
    while(temp != null) {
        if(temp.data == key) {
            return true;  // Found!
        }
        temp = temp.next;
    }
    
    return false;  // Not found
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> null **Search for:** 20

```
Step 1: temp = head = Node(10)
        10 == 20? No
        temp = temp.next = Node(20)

Step 2: temp = Node(20)
        20 == 20? YES!
        return true
```

##### Time Complexity

- O(n) - worst case traverse entire list

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);

System.out.println(list.search(20));  // Output: true
System.out.println(list.search(40));  // Output: false
```

---

### **7. FIND LENGTH**

##### Theory

```
Count nodes while traversing
```

##### Code

java

```java
int findLength() {
    int count = 0;
    Node temp = head;
    
    while(temp != null) {
        count++;
        temp = temp.next;
    }
    
    return count;
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> null

```
temp = Node(10),  count = 1
temp = Node(20),  count = 2
temp = Node(30),  count = 3
temp = null,      STOP
return 3
```

##### Time Complexity

- O(n)

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(30);

System.out.println(list.findLength());  // Output: 3
```

---

### **8. PRINT MIDDLE NODE**

##### Theory

```
Find length, then go to (length/2)th node
OR use two pointer technique (will learn in Phase 2)
```

##### Code

java

```java
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
```

##### Dry Run Example

**List:** 10 -> 20 -> 30 -> 40 -> 50 -> null

```
length = 5
middlePos = 5 / 2 = 2

temp = Node(10)
i = 0: temp = Node(20)
i = 1: temp = Node(30)
Loop ends

Output: Middle: 30
```

##### Time Complexity

- O(n) - must find length first

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
for(int i = 10; i <= 50; i += 10) {
    list.insertAtTail(i);
}
list.printMiddle();  // Output: Middle: 30
```

---

### **9. COUNT OCCURRENCES**

##### Theory

```
Traverse and count matching nodes
```

##### Code

java

```java
int countOccurrences(int key) {
    int count = 0;
    Node temp = head;
    
    while(temp != null) {
        if(temp.data == key) {
            count++;
        }
        temp = temp.next;
    }
    
    return count;
}
```

##### Dry Run Example

**List:** 10 -> 20 -> 10 -> 30 -> 10 -> null **Count:** 10

```
temp = Node(10), 10 == 10? YES, count = 1
temp = Node(20), 20 == 10? NO
temp = Node(10), 10 == 10? YES, count = 2
temp = Node(30), 30 == 10? NO
temp = Node(10), 10 == 10? YES, count = 3
temp = null, STOP
return 3
```

##### Time Complexity

- O(n)

##### Example Code Run

java

```java
LinkedList list = new LinkedList();
list.insertAtTail(10);
list.insertAtTail(20);
list.insertAtTail(10);
list.insertAtTail(30);
list.insertAtTail(10);

System.out.println(list.countOccurrences(10));  // Output: 3
System.out.println(list.countOccurrences(20));  // Output: 1
```

---

### **10. DELETE DUPLICATES (Sorted List)**

##### Theory

```
For sorted list: 10 -> 10 -> 20 -> 20 -> 30
If current == next, skip next
Result: 10 -> 20 -> 30
```

##### Code

java

```java
void deleteDuplicates() {
    if(head == null) return;
    
    Node temp = head;
    while(temp != null && temp.next != null) {
        if(temp.data == temp.next.data) {
            temp.next = temp.next.next;  // Skip duplicate
        } else {
            temp = temp.next;
        }
    }
}
```

##### Dry Run Example

**List:** 10 -> 10 -> 20 -> 20 -> 30 -> null

```
Step 1: temp = Node(10)
        10 == 10? YES
        temp.next = Node(20)
        (Don't move temp, check again)

Step 2: temp = Node(10)
        10 == 20? NO
        temp = temp.next = Node(20)

Step 3: temp = Node(20)
        20 == 20? YES
        temp.next = Node(30)
        (Don't move temp)

Step 4: temp = Node(20)
        20 == 30? NO
        temp = temp.next = Node(30)

Step 5: temp = Node(30)
        30 == null? (no next, exit loop)
