# Phase 2 Part 04


---

### **11. REMOVE NTH NODE FROM END**

#### Question

**Remove the nth node from the end of the list.**

**Examples:**

```
1 -> 2 -> 3 -> 4 -> 5, n=2
Remove 4: 1 -> 2 -> 3 -> 5

1 -> 2, n=1
Remove 2: 1

1 -> 2, n=2
Remove 1: 2
```

#### Theory

Use fast-slow pointers with gap:

- Move fast n steps ahead
- Then move both until fast reaches end
- Delete the node at slow

#### Code

java

```java
Node removeNthFromEnd(Node head, int n) {
    Node dummy = new Node(0);
    dummy.next = head;
    
    Node slow = dummy;
    Node fast = dummy;
    
    // Move fast n+1 steps ahead
    for(int i = 0; i <= n; i++) {
        if(fast == null) return head;
        fast = fast.next;
    }
    
    // Move both until fast reaches end
    while(fast != null) {
        slow = slow.next;
        fast = fast.next;
    }
    
    // Delete the node
    slow.next = slow.next.next;
    
    return dummy.next;
}
```

#### Why Use Dummy Node?

**Without dummy:**

```
1 -> 2, n=2 (remove 1)
If you need to delete head, you lose reference!
```

**With dummy:**

```
dummy -> 1 -> 2, n=2
Delete at position, then return dummy.next
Safely handles head deletion
```

#### Detailed Dry Run

**List:** 1 -> 2 -> 3 -> 4 -> 5, n=2 (remove 4)

```
Initial:
dummy -> 1 -> 2 -> 3 -> 4 -> 5
slow = dummy, fast = dummy

Step 1: Move fast n+1 = 3 steps
fast = dummy, i=0: fast = 1
fast = 1, i=1: fast = 2
fast = 2, i=2: fast = 3

State: slow = dummy, fast = 3

Step 2: Move both until fast.next == null
Iteration 1:
slow = 1, fast = 4

Iteration 2:
slow = 2, fast = 5

Iteration 3:
fast.next = null, exit loop

State: slow = 2 (node before 4), fast = 5

Step 3: Delete
slow.next = slow.next.next
2.next = 4.next = 5

Result: 1 -> 2 -> 3 -> 5
```

#### Example Code

java

```java
class LinkedList {
    Node head;
    
    void insertAtTail(int data) { /*...*/ }
    
    void traverse() {
        Node temp = head;
        while(temp != null) {
            System.out.print(temp.data + " -> ");
            temp = temp.next;
        }
        System.out.println("null");
    }
    
    Node removeNthFromEnd(int n) {
        Node dummy = new Node(0);
        dummy.next = head;
        
        Node slow = dummy;
        Node fast = dummy;
        
        // Move fast n+1 steps
        for(int i = 0; i <= n; i++) {
            if(fast == null) return head;
            fast = fast.next;
        }
        
        // Move both until fast reaches end
        while(fast != null) {
            slow = slow.next;
            fast = fast.next;
        }
        
        // Delete
        slow.next = slow.next.next;
        
        head = dummy.next;
        return head;
    }
}

// TEST
public class Main {
    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        for(int i = 1; i <= 5; i++) {
            list.insertAtTail(i);
        }
        
        System.out.println("Original:");
        list.traverse();  // 1 -> 2 -> 3 -> 4 -> 5 -> null
        
        list.removeNthFromEnd(2);
        System.out.println("After removing 2nd from end:");
        list.traverse();  // 1 -> 2 -> 3 -> 5 -> null
    }
}
```

#### Output:

```
Original:
1 -> 2 -> 3 -> 4 -> 5 -> null
After removing 2nd from end:
1 -> 2 -> 3 -> 5 -> null
```

#### Time & Space Complexity

- **Time:** O(n)
- **Space:** O(1)

---

### **12. PALINDROME LINKED LIST**

#### Question

**Check if a linked list is a palindrome.**

**Examples:**

```
1 -> 2 -> 2 -> 1  → Palindrome: true
1 -> 2 -> 3 -> 2 -> 1  → Palindrome: true
1 -> 2 -> 3  → Palindrome: false
```

#### Theory

Strategy:

1. Find middle using fast-slow
2. Reverse second half
3. Compare first half with reversed second half

#### Code

java

```java
boolean isPalindrome(Node head) {
    if(head == null || head.next == null) {
        return true;
    }
    
    // Step 1: Find middle
    Node slow = head;
    Node fast = head;
    
    while(fast != null && fast.next != null) {
        slow = slow.next;
        fast = fast.next.next;
    }
    
    // Step 2: Reverse second half
    Node secondHalf = slow;
    secondHalf = reverse(secondHalf);
    
    // Step 3: Compare
    Node first = head;
    Node second = secondHalf;
    
    while(second != null) {  // Second half might be shorter
        if(first.data != second.data) {
            return false;
        }
        first = first.next;
        second = second.next;
    }
    
    return true;
}

Node reverse(Node head) {
    Node prev = null;
    Node curr = head;
    
    while(curr != null) {
        Node next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    
    return prev;
}
```

#### Detailed Dry Run

**List:** 1 -> 2 -> 3 -> 2 -> 1

**Step 1: Find Middle**

```
slow: 1 -> 2 -> 3
fast: 1 -> 3 -> 1

Middle: 3
```

**Step 2: Reverse Second Half**

```
Before: 3 -> 2 -> 1
After:  1 -> 2 -> 3
```

**Step 3: Compare**

```
First:  1 -> 2
Second: 1 -> 2

Match! Palindrome = true
```

#### Example Code

java

```java
class LinkedList {
    Node head;
    
    void insertAtTail(int data) { /*...*/ }
    
    Node reverse(Node head) {
        Node prev = null;
        Node curr = head;
        
        while(curr != null) {
            Node next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        
        return prev;
    }
    
    boolean isPalindrome() {
        if(head == null || head.next == null) {
            return true;
        }
        
        // Find middle
        Node slow = head;
        Node fast = head;
        
        while(fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        
        // Reverse second half
        Node secondHalf = reverse(slow);
        
        // Compare
        Node first = head;
        Node second = secondHalf;
        boolean result = true;
        
        while(second != null) {
            System.out.println("Compare: " + first.data + " vs " + second.data);
            if(first.data != second.data) {
                result = false;
                break;
            }
            first = first.next;
            second = second.next;
        }
        
        return result;
    }
}

// TEST
public class Main {
    public static void main(String[] args) {
        // Test 1: Palindrome
        LinkedList list1 = new LinkedList();
        list1.insertAtTail(1);
        list1.insertAtTail(2);
        list1.insertAtTail(3);
        list1.insertAtTail(2);
        list1.insertAtTail(1);
        
        System.out.println("List 1 is palindrome: " + list1.isPalindrome());  // true
        
        // Test 2: Not palindrome
        LinkedList list2 = new LinkedList();
        list2.insertAtTail(1);
        list2.insertAtTail(2);
        list2.insertAtTail(3);
        
        System.out.println("List 2 is palindrome: " + list2.isPalindrome());  // false
    }
}
```

#### Output:

```
Compare: 1 vs 1
Compare: 2 vs 2
List 1 is palindrome: true
Compare: 1 vs 3
List 2 is palindrome: false
```

#### Time & Space Complexity

- **Time:** O(n)
- **Space:** O(1) if reversal is in-place

---

### **13. EDGE CASES TO ALWAYS TEST**

#### 1. Empty List

java

```java
head = null
```

#### 2. Single Node

java

```java
1 -> null
```

#### 3. Two Nodes

java

```java
1 -> 2 -> null
```

#### 4. All Identical

java

```java
1 -> 1 -> 1 -> null
```

#### 5. Cycle at Head

java

```java
1 -> 2 -> 1 (cycle at 1)
```

---

### **14. COMMON MISTAKES IN PHASE 2**

#### Mistake 1: Losing Forward Link (Reversal)

❌ **WRONG:**

java

