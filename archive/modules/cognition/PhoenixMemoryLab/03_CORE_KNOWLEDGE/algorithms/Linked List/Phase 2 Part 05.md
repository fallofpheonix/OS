# Phase 2 Part 05

```java
curr.next = prev;
curr = curr.next;  // curr.next is now backward!
```

✅ **RIGHT:**

java

```java
Node next = curr.next;
curr.next = prev;
curr = next;
```

---

#### Mistake 2: Wrong Loop Condition (Fast-Slow)

❌ **WRONG:**

java

```java
while(fast != null) {  // Doesn't check fast.next!
    fast = fast.next.next;  // NullPointerException
}
```

✅ **RIGHT:**

java

```java
while(fast != null && fast.next != null) {
    fast = fast.next.next;
}
```

---

#### Mistake 3: Comparing Values Instead of References (Cycle)

❌ **WRONG:**

java

```java
if(slow.data == fast.data) {  // Could match by coincidence!
    // Cycle detected
}
```

✅ **RIGHT:**

java

```java
if(slow == fast) {  // Same object in memory
    // Cycle detected
}
```

---

#### Mistake 4: Returning Wrong Pointer (Reversal)

❌ **WRONG:**

java

```java
while(curr != null) {
    // ... reversal logic
}
return curr;  // curr is null!
```

✅ **RIGHT:**

java

```java
while(curr != null) {
    // ... reversal logic
}
return prev;  // prev is new head!
```

---

#### Mistake 5: Not Saving Next (Recursive Reversal)

❌ **WRONG:**

java

```java
head.next.next = head;
head.next = null;  // Lost the cycle!
```

✅ **RIGHT:** Save newHead from recursive call, then reconnect.

---

### **15. FINAL PRACTICE CHECKLIST**

#### Problems to Solve

✅ **Reversal:**

1. Reverse Linked List (iterative)
2. Reverse Linked List (recursive)
3. Reverse Linked List II
4. Reverse Nodes in K-Group

✅ **Fast-Slow:**

1. Find Middle of Linked List
2. Linked List Cycle
3. Linked List Cycle II
4. Remove Nth Node From End

✅ **Combination:**

1. Palindrome Linked List

#### Mastery Checklist

- [ ]  Can implement iterative reversal without help
- [ ]  Can implement recursive reversal and explain critical line
- [ ]  Can find middle using fast-slow in one pass
- [ ]  Can detect cycle using Floyd's algorithm
- [ ]  Can find cycle start using distance reasoning
- [ ]  Can remove nth node from end safely
- [ ]  Can check palindrome by combining patterns
- [ ]  Can dry run all operations on paper
- [ ]  Can identify common mistakes immediately
- [ ]  Can handle all edge cases

---

### **COMPLETE WORKING CODE - ALL PHASE 2 OPERATIONS**

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
    
    void traverse() {
        Node temp = head;
        System.out.print("List: ");
        while(temp != null) {
            System.out.print(temp.data + " -> ");
            temp = temp.next;
        }
        System.out.println("null");
    }
    
    // ⭐ 1. REVERSE ITERATIVE
    void reverseIterative() {
        Node prev = null;
        Node curr = head;
        
        while(curr != null) {
            Node next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        
        head = prev;
    }
    
    // ⭐ 2. REVERSE RECURSIVE
    void reverseRecursive() {
        head = reverseHelper(head);
    }
    
    private Node reverseHelper(Node head) {
        if(head == null || head.next == null) {
            return head;
        }
        
        Node newHead = reverseHelper(head.next);
        head.next.next = head;
        head.next = null;
        
        return newHead;
    }
    
    // ⭐ 3. FIND MIDDLE
    Node findMiddle() {
        if(head == null || head.next == null) {
            return head;
        }
        
        Node slow = head;
        Node fast = head;
        
        while(fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        
        return slow;
    }
    
    // ⭐ 4. DETECT CYCLE
    boolean hasCycle() {
        if(head == null || head.next == null) {
            return false;
        }
        
        Node slow = head;
        Node fast = head;
        
        while(fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            
            if(slow == fast) {
                return true;
            }
        }
        
        return false;
    }
    
    // ⭐ 5. FIND CYCLE START
    Node findCycleStart() {
        if(head == null || head.next == null) {
            return null;
        }
        
        Node slow = head;
        Node fast = head;
        
        // Detect cycle
        while(fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            
            if(slow == fast) {
                // Find cycle start
                slow = head;
                while(slow != fast) {
                    slow = slow.next;
                    fast = fast.next;
                }
                return slow;
            }
        }
        
        return null;
    }
    
    // ⭐ 6. REMOVE NTH FROM END
    void removeNthFromEnd(int n) {
        Node dummy = new Node(0);
        dummy.next = head;
        
        Node slow = dummy;
        Node fast = dummy;
        
        // Move fast n+1 steps
        for(int i = 0; i <= n; i++) {
            if(fast == null) return;
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
    }
    
    // ⭐ 7. PALINDROME CHECK
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
        
        while(second != null) {
            if(first.data != second.data) {
                return false;
            }
            first = first.next;
            second = second.next;
        }
        
        return true;
    }
    
    private Node reverse(Node head) {
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
}

// TEST ALL OPERATIONS
public class Main {
    public static void main(String[] args) {
        System.out.println("=== PHASE 2: POINTER MASTERY ===\n");
        
        // Test 1: Reverse
        System.out.println("1. REVERSE ITERATIVE:");
        LinkedList list1 = new LinkedList();
        for(int i = 1; i <= 5; i++) {
            list1.insertAtTail(i);
        }
        list1.traverse();
        list1.reverseIterative();
        list1.traverse();
        
        // Test 2: Find Middle
        System.out.println("\n2. FIND MIDDLE:");
        LinkedList list2 = new LinkedList();
        for(int i = 1; i <= 5; i++) {
            list2.insertAtTail(i);
        }
        list2.traverse();
        Node middle = list2.findMiddle();
        System.out.println("Middle: " + middle.data);
        
        // Test 3: Remove Nth from End
        System.out.println("\n3. REMOVE NTH FROM END:");
        LinkedList list3 = new LinkedList();
        for(int i = 1; i <= 5; i++) {
            list3.insertAtTail(i);
        }
        list3.traverse();
        list3.removeNthFromEnd(2);
        list3.traverse();
        
        // Test 4: Palindrome
        System.out.println("\n4. PALINDROME CHECK:");
        LinkedList list4 = new LinkedList();
        list4.insertAtTail(1);
        list4.insertAtTail(2);
        list4.insertAtTail(3);
        list4.insertAtTail(2);
        list4.insertAtTail(1);
        list4.traverse();
        System.out.println("Is palindrome: " + list4.isPalindrome());
    }
}
```

#### Output:

```
=== PHASE 2: POINTER MASTERY ===

1. REVERSE ITERATIVE:
List: 1 -> 2 -> 3 -> 4 -> 5 -> null
List: 5 -> 4 -> 3 -> 2 -> 1 -> null

2. FIND MIDDLE:
List: 1 -> 2 -> 3 -> 4 -> 5 -> null
Middle: 3

3. REMOVE NTH FROM END:
List: 1 -> 2 -> 3 -> 4 -> 5 -> null
List: 1 -> 2 -> 3 -> 5 -> null

4. PALINDROME CHECK:
List: 1 -> 2 -> 3 -> 2 -> 1 -> null
Is palindrome: true
```

---

### **PHASE 2 MASTERY CHECKPOINT**

#### Can You?

✅ Implement iterative reversal with 3 pointers ✅ Implement recursive reversal and explain `head.next.next = head` ✅ Explain why fast-slow works mathematically ✅ Use fast-slow to find middle in one pass ✅ Detect cycle using Floyd's algorithm ✅ Find cycle start using distance reasoning ✅ Remove Nth node from end safely ✅ Check palindrome by combining patterns ✅ Handle all edge cases (empty, single, two nodes) ✅ Avoid all 5 common mistakes

**If YES to all → Ready for Phase 3! 🎯**

---

### **NEXT STEPS**

**After mastering Phase 2, you will:**

- Unlock 60% of linked list interview problems
- Be able to combine patterns confidently
- Have deep pointer manipulation intuition
- Solve most medium-level LL problems
- Be ready for advanced structural manipulation (Phase 3)

**Time Commitment:** 1-2 weeks of daily practice **Questions to Solve:** 15+ problems (with variations) **Key to Success:** Dry running every operation multiple times!
