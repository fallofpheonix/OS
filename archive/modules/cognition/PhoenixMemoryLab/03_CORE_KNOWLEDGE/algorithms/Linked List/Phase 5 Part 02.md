# Phase 5 Part 02


curr=3: no child
curr=4: no child
Done!

Final: 1 <-> 2 <-> 7 <-> 8 <-> 3 <-> 4 ✓
```

Time: O(n) | Space: O(1)

**Why this works in one pass:** After inserting the child list, `curr` moves to `curr.next` which is the child's head. If that child also had children, we handle it when we reach it — the while loop naturally processes recursively-inserted levels.

---

### 8. DESIGN PROBLEMS — BROWSER HISTORY

**Problem:** Design browser history navigation.

java

```java
BrowserHistory(String homepage)  // Start here
void visit(String url)            // Visit new page (clears forward history)
String back(int steps)            // Go back up to 'steps' pages
String forward(int steps)         // Go forward up to 'steps' pages
```

**Why DLL?** You need to go both forward and backward in O(1) per step. A DLL with a `current` pointer is perfect.

java

```java
class BrowserHistory {
    class Node {
        String url;
        Node prev, next;
        Node(String url) { this.url = url; }
    }

    Node current;

    BrowserHistory(String homepage) {
        current = new Node(homepage);
    }

    void visit(String url) {
        Node newPage = new Node(url);
        current.next = newPage;   // New page after current
        newPage.prev = current;
        // Don't connect newPage.next — forward history is erased
        current = newPage;
    }

    String back(int steps) {
        while (steps > 0 && current.prev != null) {
            current = current.prev;
            steps--;
        }
        return current.url;
    }

    String forward(int steps) {
        while (steps > 0 && current.next != null) {
            current = current.next;
            steps--;
        }
        return current.url;
    }
}
```

**Dry Run:**

```
BrowserHistory("google.com")
current: [google.com]

visit("youtube.com"):
current: google.com <-> youtube.com
                              ↑ current

visit("gmail.com"):
current: google.com <-> youtube.com <-> gmail.com
                                              ↑ current

back(1):
current moves to youtube.com. return "youtube.com"

visit("maps.com"):
youtube.com <-> maps.com
Forward history (gmail) is gone! (maps.next = null)
current: youtube.com <-> maps.com
                               ↑ current

forward(2):
current.next = null (no forward history) → stay at maps.com
return "maps.com"

back(2):
maps → youtube → google. return "google.com"
```

---

### 9. POINTER SAFETY FRAMEWORK (How to Never Lose a Node Again)

This is the systematic checklist experts run mentally before writing any pointer manipulation:

**Before modifying any pointer, ask these 3 questions:**

1. **"What am I about to lose?"** — If I change `X.next`, I lose the node `X.next` was pointing to. Have I saved it?
2. **"Who owns this node after the change?"** — After I rewire, is this node reachable from some other pointer? Or is it now floating in memory with no reference?
3. **"What region am I affecting?"** — Am I only touching the current segment? Or am I accidentally disconnecting something from a segment I haven't processed yet?

**Practical checklist before any pointer rewiring:**

```
□ Save all nodes I'm about to lose access to (save curr.next before changing curr.next)
□ Identify the boundaries of my current segment
□ Check that all nodes outside my segment are still properly connected
□ After rewiring, verify invariants still hold
□ Check the loop termination condition hasn't been corrupted
```

**The most common node-loss pattern:**

java

```java
// ❌ Classic mistake
curr.next = prev;
curr = curr.next;  // curr.next is now prev (backwards!), not forward

// ✅ Always save forward link FIRST
Node next = curr.next;  // Save before touching
curr.next = prev;       // Now safe to rewire
curr = next;            // Use saved reference
```

---

### 10. RECURSION VS ITERATION — WHEN TO USE WHICH

This comes up in interviews. Know both and know the trade-offs.

||Iteration|Recursion|
|---|---|---|
|Space|O(1)|O(n) call stack|
|Readability|More verbose|Often cleaner|
|When to prefer|When space matters|When structure naturally decomposes|
|Risk|More pointer bugs|Stack overflow on long lists|

**Key interview point:** Recursive linked list solutions are NOT O(1) space. The call stack uses O(n) memory. If an interviewer asks for O(1) space, you must use iteration.

**Example — Reverse in groups of k:**

Recursively, every call handles one group and delegates the rest. Clean to read, but O(n/k) stack depth.

Iteratively, you use the `groupPrev` pointer loop. O(1) space.

**When recursion genuinely helps:**

Flatten Multilevel DLL — the recursive version is very natural because the problem itself is recursive (child lists can have their own children). Iteration requires manually simulating the stack.

---

### 11. COMPLETE PATTERN MAP WITH PROBLEM EXAMPLES

|Pattern|Core Mechanism|Example Problems|
|---|---|---|
|Reversal|3-pointer (prev, curr, next)|Reverse LL, Reverse K-Group, Palindrome|
|Fast-Slow|2-speed pointers|Cycle detection, Middle node, Find Duplicate|
|Dummy Node|Fake head to avoid edge cases|Merge sorted, Partition, Delete nth|
|Merge|Compare heads, attach smaller|Merge Two Lists, Merge K Lists, Sort List|
|Split|`slow.next = null` to disconnect|Sort List, Reorder List, Palindrome|
|Two Chains|Build two lists, connect at end|Odd-Even, Partition, Separate components|
|Segment Manipulation|groupPrev/kth/groupNext|Reverse K-Group, Rotate List|
|Deep Copy|HashMap or interleave|Copy Random Pointer, Clone Graph|
|DLL + HashMap|O(1) access + O(1) ordering|LRU Cache, LFU Cache, Browser History|

---

### 12. THE 6-STEP INTERVIEW SOLVING FRAMEWORK

When you see a new linked list problem in an interview, apply this in order:

**Step 1 — Identify the pattern.** Does it involve reversing? Detecting cycles? Merging? Structural rearrangement? Match it to your pattern map.

**Step 2 — Identify segment boundaries.** What regions am I working on? What comes before, during, and after my manipulation?

**Step 3 — Name your invariants.** What should be true at the start of every loop iteration? Write it in a comment if needed.

**Step 4 — Identify what you must save before modifying.** List every pointer that will be overwritten and make sure you have a saved reference before touching it.

**Step 5 — Handle edge cases.** Empty list? Single node? Two nodes? k > length? Cycle already exists?

**Step 6 — Trace through a small example mentally.** Especially at boundaries — first iteration, last iteration, and the iteration before loop exit.

---

### 13. ADVANCED QUESTION: ALL O(1) DATA STRUCTURE

**Problem:** Design a data structure supporting insert, delete, getRandom — all O(1).

java

```java
class RandomizedSet {
    ArrayList<Integer> list;     // For O(1) getRandom (random index)
    HashMap<Integer, Integer> map; // val → index in list (for O(1) delete)

    boolean insert(int val) {
        if (map.containsKey(val)) return false;
        list.add(val);
        map.put(val, list.size() - 1);
        return true;
    }

    boolean remove(int val) {
        if (!map.containsKey(val)) return false;

        // Swap with last element (so we can remove from end in O(1))
        int idx = map.get(val);
        int last = list.get(list.size() - 1);
        list.set(idx, last);
        map.put(last, idx);

        // Remove last
        list.remove(list.size() - 1);
        map.remove(val);
        return true;
    }

    int getRandom() {
        Random rand = new Random();
        return list.get(rand.nextInt(list.size()));
    }
}
```

**Why this works:** You can't delete from the middle of an array in O(1). But you can swap the target element with the last element, then remove the last element — both O(1). Update the map to reflect the new index of the swapped element.

This problem shows how linked list thinking (maintaining references/pointers to elements for O(1) manipulation) applies to array-based structures too.

---

### 14. EDGE CASE MASTERY — THE COMPLETE LIST

For every linked list problem, automatically test these mentally before submitting:

|Edge Case|What to check|
|---|---|
|Empty list|`head == null` — return immediately|
|Single node|`head.next == null` — often base case|
|Two nodes|Most reversal and fast-slow bugs surface here|
|All same values|Doesn't break pointer logic|
|Already sorted|Merge and sort should handle gracefully|
|Already reversed|Reverse should still work|
|k > length|Modulo: `k % length` before using k|
|Cycle at head|Cycle start detection must handle this|
|k = 0|Rotation by 0 = no change|
|Duplicate values|Intersection uses `==` (reference), not `.data`|

---

### 15. FINAL PRACTICE QUESTION SET

#### Easy (Warm up — should be instant):

- Reverse Linked List
- Find Middle Node
- Detect Cycle

#### Medium (Core interview level — 15-20 min each):

- Reorder List
- Palindrome Linked List
- Partition List
- Odd-Even List
- Sort List (Merge Sort)
- Swap Nodes in Pairs

#### Hard (Senior/FAANG level — 30-45 min each):

- Reverse Nodes in K-Group
- Merge K Sorted Lists
- Copy List with Random Pointer (O(1) space approach)
- Flatten Multilevel DLL
- LRU Cache
- LFU Cache
- Find Duplicate Number (Floyd on array)

#### Design-Heavy (System design cross-over):

- Browser History
- Undo/Redo System (same idea, different domain)
- All O(1) Data Structure

---

### PHASE 5 CHECKPOINT

You have completed the full linked list journey when you can do all of these:

✅ Explain what invariant is maintained at every step of your reversal code

✅ Derive Floyd's cycle math from scratch and explain why resetting one pointer to head works

✅ Solve Find Duplicate Number using fast-slow on an array (recognize the hidden LL pattern)

✅ Implement LFU Cache with frequency buckets — explain why minFreq resets to 1 on every `put`

✅ Implement Flatten Multilevel DLL and explain why one pass handles all nesting levels

✅ Implement Browser History and explain why `visit` must NOT connect `newPage.next`

✅ For any new linked list problem, identify the pattern within 2 minutes

✅ For any bug in a linked list, immediately draw the node-pointer diagram rather than reading code

---

### THE FINAL REALIZATION

Linked Lists are NOT about storing data. They are about **managing relationships between nodes safely**.

Every hard linked list problem reduces to: "How do I safely rewire relationships between nodes without losing anything, while maintaining structural invariants?"

Once you internalize this, you stop memorizing solutions. You start deriving them. That is the difference between passing and failing a FAANG interview. 🎯

