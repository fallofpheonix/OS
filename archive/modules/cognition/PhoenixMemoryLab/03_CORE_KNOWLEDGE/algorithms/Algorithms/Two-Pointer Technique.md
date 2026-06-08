## 1. Definition

The **two-pointer technique** is an algorithmic approach where two indices traverse a data structure to solve problems in linear time by avoiding redundant computation.

---

## 2. Intuition

Brute force often compares every pair → **O(n²)**  
Two-pointer reduces this by:

- Using order (sorted data) or structure
- Moving pointers based on conditions

Result: **O(n)** traversal

---

## 3. When to Use

Apply two-pointer when:

- Array/string is **sorted** or sequential
- Need to find **pairs, subarrays, or merges**
- Problem involves **continuous segments**
- Comparing elements from **both ends or two sources**

---

## 4. Types of Two-Pointer Patterns

### (A) Opposite Direction (Left–Right)

**Structure:**

```
int left = 0, right = n - 1;while(left < right){    // condition check    left++;    right--;}
```

**Use cases:**

- Palindrome check
- Two sum (sorted array)

---

### (B) Same Direction (Fast–Slow)

**Structure:**

```
int slow = 0;for(int fast = 0; fast < n; fast++){    if(condition){        // process        slow++;    }}
```

**Use cases:**

- Remove duplicates
- Partitioning
- Cycle detection

---

### (C) Two Arrays / Merge Pattern

**Structure:**

```
int i = 0, j = 0;while(i < n && j < m){    // process both    i++;    j++;}
```

**Use case:**

- Merge arrays/strings

---

## 5. Standard Template (Important for Exams)

```
int i = 0, j = 0;while(i < n && j < m){    // main logic    i++;    j++;}// leftoverwhile(i < n) i++;while(j < m) j++;
```

---

## 6. Example

### Merge Alternately

Input:  
word1 = "ab", word2 = "pqrs"

Steps:

- a + p
- b + q
- remaining: r, s

Output: "apbqrs"

---

## 7. Complexity

- Time Complexity: **O(n)**
- Space Complexity:
    - O(1) if in-place
    - O(n) if building result

---

## 8. Advantages

- Eliminates nested loops
- Efficient linear traversal
- Works well with sorted data
- Easy to implement once pattern is recognized

---

## 9. Common Mistakes

- Incorrect loop condition (`i < n` instead of `i < n && j < m`)
- Index out of bounds
- Not handling leftover elements
- Wrong pointer movement
- Misidentifying pattern

---

## 10. Revision Notes (Quick Recall)

- Two-pointer = **two indices, one pass**
- Patterns:
    - Opposite ends
    - Fast–slow
    - Merge traversal
- Always:
    - Check bounds
    - Handle leftovers
- Use when:
    - Sorted / sequential data
    - Pair or segment problems

---

## 11. Conclusion

Two-pointer is a core optimization technique that transforms quadratic solutions into linear ones.  
Correct pointer movement and boundary handling determine correctness.