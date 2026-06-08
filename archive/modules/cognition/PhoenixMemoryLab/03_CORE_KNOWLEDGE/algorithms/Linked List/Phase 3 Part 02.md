# Phase 3 Part 02

    }

    return p1;  // Either intersection node, or null (no intersection)
}
```

**Dry Run:**

```
A: 1 -> 2 -> 8 -> 9 -> null
B:      5 -> 8 -> 9 -> null

p1=1, p2=5  (not equal)
p1=2, p2=8  (not equal)
p1=8, p2=9  (not equal)
p1=9, p2=null → p2 switches to a: p2=1
p1=null → p1 switches to b: p1=5, p2=2
p1=8, p2=8  ← EQUAL! Return Node(8) ✓

No intersection case:
Both pointers eventually reach null at the same time.
p1 = null, p2 = null → loop exits, return null
```

Time: O(n + m) | Space: O(1)

---

### 6. MERGE K SORTED LISTS

**Problem:** Merge k sorted linked lists into one sorted list.

```
Input:
1 -> 4 -> 5
1 -> 3 -> 4
2 -> 6

Output:
1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
```

**Brute force:** Merge them one by one. First merge list 1 and 2, then merge result with list 3, etc. This is O(k×n) — slow.

**Optimal:** Use a Min Heap. Always pick the globally smallest node next.

**Logic:**

1. Insert the head of every list into a min heap (ordered by node value)
2. Extract the minimum node, add to result
3. If that node had a next, insert next into the heap
4. Repeat until heap is empty

java

```java
Node mergeKLists(Node[] lists) {
    // Min heap ordered by node's data value
    PriorityQueue<Node> heap = new PriorityQueue<>((a, b) -> a.data - b.data);

    // Add all heads to heap
    for (Node list : lists) {
        if (list != null) heap.offer(list);
    }

    Node dummy = new Node(-1);
    Node temp = dummy;

    while (!heap.isEmpty()) {
        Node smallest = heap.poll();   // Extract minimum
        temp.next = smallest;
        temp = temp.next;

        if (smallest.next != null) {   // Add that list's next node
            heap.offer(smallest.next);
        }
    }

    return dummy.next;
}
```

**Dry Run:**

```
Lists: [1->4->5], [1->3->4], [2->6]

Heap initially: {1(L1), 1(L2), 2(L3)}

Step 1: Poll 1(L1). result: 1. Add 4(L1) to heap.
Heap: {1(L2), 2(L3), 4(L1)}

Step 2: Poll 1(L2). result: 1->1. Add 3(L2) to heap.
Heap: {2(L3), 3(L2), 4(L1)}

Step 3: Poll 2(L3). result: 1->1->2. Add 6(L3).
Heap: {3(L2), 4(L1), 6(L3)}

Step 4: Poll 3(L2). result: 1->1->2->3. Add 4(L2).
Heap: {4(L1), 4(L2), 6(L3)}

Step 5: Poll 4(L1). result: ...->4. Add 5(L1).
Heap: {4(L2), 5(L1), 6(L3)}

Step 6: Poll 4(L2). result: ...->4->4. (L2 exhausted)
Step 7: Poll 5(L1). result: ...->5. (L1 exhausted)
Step 8: Poll 6(L3). result: ...->6. (L3 exhausted)

Final: 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6 ✓
```

Time: O(n log k) where n = total nodes, k = number of lists | Space: O(k) for heap

---

### 7. COMMON MISTAKES IN PHASE 3

**Mistake 1: Forgetting `slow.next = null` when splitting**

java

```java
// ❌ WRONG
Node second = slow.next;
// Forgot: slow.next = null;
Node left = sortList(head);    // head is still connected to second!
Node right = sortList(second); // Both halves share nodes = chaos
```

java

```java
// ✅ RIGHT
Node second = slow.next;
slow.next = null;  // MUST physically disconnect
Node left = sortList(head);
Node right = sortList(second);
```

**Mistake 2: Forgetting `temp = temp.next` in merge**

java

```java
// ❌ WRONG - overwrites the same connection every time
while (l1 != null && l2 != null) {
    temp.next = l1;
    l1 = l1.next;
    // Missing: temp = temp.next;
}
// Result: only the last picked node gets attached!
```

**Mistake 3: Comparing values instead of references (intersection)**

java

```java
// ❌ WRONG - two nodes can have the same value but not be the same node
if (p1.data == p2.data) return p1;

// ✅ RIGHT - check if they're the exact same object in memory
if (p1 == p2) return p1;
```

**Mistake 4: Not handling null in reorder's interleave step**

java

```java
// ❌ WRONG - if second becomes null, accessing second.next crashes
while (first != null && second != null) {
    Node tmp1 = first.next;
    Node tmp2 = second.next;  // Crashes when second = null
    ...
}

// ✅ RIGHT - check second before accessing second.next
while (second != null) {
    ...
}
```

---

### 8. COMPLETE WORKING CODE — ALL PHASE 3 OPERATIONS

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
        if (head == null) { head = newNode; return; }
        Node temp = head;
        while (temp.next != null) temp = temp.next;
        temp.next = newNode;
    }

    void traverse() {
        Node temp = head;
        while (temp != null) { System.out.print(temp.data + " -> "); temp = temp.next; }
        System.out.println("null");
    }

    // ⭐ 1. MERGE TWO SORTED LISTS
    Node mergeTwoLists(Node l1, Node l2) {
        Node dummy = new Node(-1);
        Node temp = dummy;
        while (l1 != null && l2 != null) {
            if (l1.data <= l2.data) { temp.next = l1; l1 = l1.next; }
            else { temp.next = l2; l2 = l2.next; }
            temp = temp.next;
        }
        if (l1 != null) temp.next = l1;
        if (l2 != null) temp.next = l2;
        return dummy.next;
    }

    // ⭐ 2. SORT LIST (Merge Sort)
    Node sortList(Node head) {
        if (head == null || head.next == null) return head;
        Node slow = head, fast = head.next;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        Node second = slow.next;
        slow.next = null;
        Node left = sortList(head);
        Node right = sortList(second);
        return mergeTwoLists(left, right);
    }

    // ⭐ 3. REORDER LIST
    void reorderList(Node head) {
        if (head == null || head.next == null) return;
        Node slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next; fast = fast.next.next;
        }
        Node second = slow.next;
        slow.next = null;
        second = reverse(second);
        Node first = head;
        while (second != null) {
            Node tmp1 = first.next, tmp2 = second.next;
            first.next = second;
            second.next = tmp1;
            first = tmp1;
            second = tmp2;
        }
    }

    // ⭐ 4. INTERSECTION
    Node getIntersection(Node a, Node b) {
        Node p1 = a, p2 = b;
        while (p1 != p2) {
            p1 = (p1 == null) ? b : p1.next;
            p2 = (p2 == null) ? a : p2.next;
        }
        return p1;
    }

    private Node reverse(Node head) {
        Node prev = null, curr = head;
        while (curr != null) {
            Node next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        return prev;
    }
}

public class Main {
    public static void main(String[] args) {
        LinkedList helper = new LinkedList();

        // Test 1: Merge Two Sorted Lists
        System.out.println("=== MERGE TWO SORTED LISTS ===");
        Node l1 = new Node(1); l1.next = new Node(3); l1.next.next = new Node(5);
        Node l2 = new Node(2); l2.next = new Node(4); l2.next.next = new Node(6);
        Node merged = helper.mergeTwoLists(l1, l2);
        Node t = merged;
        while (t != null) { System.out.print(t.data + " -> "); t = t.next; }
        System.out.println("null");  // 1->2->3->4->5->6

        // Test 2: Sort List
        System.out.println("\n=== SORT LIST ===");
        Node unsorted = new Node(4);
        unsorted.next = new Node(2);
        unsorted.next.next = new Node(1);
        unsorted.next.next.next = new Node(3);
        Node sorted = helper.sortList(unsorted);
        t = sorted;
        while (t != null) { System.out.print(t.data + " -> "); t = t.next; }
        System.out.println("null");  // 1->2->3->4

        // Test 3: Reorder List
        System.out.println("\n=== REORDER LIST ===");
        LinkedList list3 = new LinkedList();
        for (int i = 1; i <= 5; i++) list3.insertAtTail(i);
        list3.traverse();         // 1->2->3->4->5
        helper.reorderList(list3.head);
        list3.traverse();         // 1->5->2->4->3
    }
}
```

**Output:**

```
=== MERGE TWO SORTED LISTS ===
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null

=== SORT LIST ===
1 -> 2 -> 3 -> 4 -> null

=== REORDER LIST ===
1 -> 2 -> 3 -> 4 -> 5 -> null
1 -> 5 -> 2 -> 4 -> 3 -> null
```

---

### 9. PHASE 3 COMPLEXITY SUMMARY

|Problem|Time|Space|Key Technique|
|---|---|---|---|
|Merge Two Sorted Lists|O(n+m)|O(1)|Dummy node + two pointers|
|Sort List|O(n log n)|O(log n)|Fast-slow + merge sort|
|Reorder List|O(n)|O(1)|Fast-slow + reversal + merge|
|Intersection|O(n+m)|O(1)|Two pointer switch|
|Merge K Lists|O(n log k)|O(k)|Min heap|

### 10. PRACTICE ORDER FOR PHASE 3

**Day 1-2:** Merge Two Sorted Lists — practice until you can write it without thinking

**Day 3-4:** Sort List — combines fast-slow from Phase 2 with merge; this is a big leap

**Day 5-6:** Reorder List — hardest of the phase; combine all three patterns

**Day 7:** Intersection of Two Lists — elegant two-pointer problem

**Day 8+:** Merge K Sorted Lists — bring in the heap

---

### PHASE 3 CHECKPOINT: Are You Ready for Phase 4?

Move on only when you can do all of these without help:

✅ Implement merge two sorted lists with dummy node from scratch

✅ Explain why `slow.next = null` is mandatory before recursion in sort list

✅ Implement reorder list by combining fast-slow + reversal + interleave

✅ Explain the intersection two-pointer trick mathematically

✅ Dry run all problems on paper with pointer diagrams

✅ Handle edge cases: empty list, single node, two nodes

If all yes → **Ready for Phase 4!** 🎯

  
