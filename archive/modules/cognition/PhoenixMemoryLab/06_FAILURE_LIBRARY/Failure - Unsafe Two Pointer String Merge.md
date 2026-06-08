# Failure: Unsafe Two Pointer String Merge

## Environment
Java / DSA practice

## Symptom
String merge implementation produced incorrect output or index errors when input lengths differed.

## Root Cause
Implementation mixed immutable `String` APIs with mutable-builder assumptions, used unsafe loop bounds, and failed to process remaining characters after the shared traversal ended.

## Fix
Use `StringBuilder`, traverse while both indexes are valid, then drain each remaining suffix.

```java
StringBuilder out = new StringBuilder();
int i = 0, j = 0;
while (i < a.length() && j < b.length()) {
    out.append(a.charAt(i++));
    out.append(b.charAt(j++));
}
while (i < a.length()) out.append(a.charAt(i++));
while (j < b.length()) out.append(b.charAt(j++));
return out.toString();
```

## Prevention
Before coding any two-pointer merge, write the main loop condition and both drain loops first. Use `StringBuilder` for Java string accumulation.

## Related Concepts
[[Two-Pointer Technique]], [[Bit Manipulation]], [[Linked List]]

## Related Failures
None

## Original Notes


### (A) Language-Level Mistakes

1. **Using `append()` on String**
    - `String` is immutable in Java
    - `append()` works only with `StringBuilder`
2. **Wrong Character Combination**
    - `char + char` → integer (ASCII addition), not string
    - Must use `.append(char)` or convert to string

---

### (B) Scope and Variable Errors

3. **Variable declared inside loop**
    - Goes out of scope immediately
    - Cannot return it later
4. **No accumulation**
    - Reassigning instead of building result

---

### (C) Logic Errors

5. **Incorrect loop condition**
    - Looping till longer string but accessing both  
        → causes `IndexOutOfBoundsException`
6. **Wrong use of if-else on length**
    - Problem is symmetric
    - Should not branch logic based on length

---

### (D) Missing Edge Case Handling

7. **Remaining characters not handled**
    - If lengths differ, leftover part ignored

---

## 2. Core Concept You Missed

This is a **two-pointer string merging problem**  
Pattern: sequential traversal of two arrays/strings

---

## 3. Checklist for Next Questions

### (1) Before Coding

- Identify pattern:
    - Two pointers?
    - Sliding window?
    - Prefix?
- Check:
    - Equal or unequal sizes?
    - Need for merging or comparison?

---

### (2) While Coding

- Use correct data structure:
    - `StringBuilder` for string building
- Loop safely:
    
    ```
    while(i < n && j < m)
    ```
    
- Never assume equal lengths

---

### (3) After Main Loop

Always ask:

- “Did I process remaining elements?”

---

### (4) Index Safety Rule

If accessing two arrays/strings:

```
i < len1 && i < len2
```

---

### (5) Java-Specific Rules

- String → immutable
- Use:
    - `charAt(i)`
    - `StringBuilder.append()`
- Avoid:
    - `+` in loops (inefficient)

---

## 4. Common Pattern Template (Memorize)

```
while(i < n && j < m){    // process both}while(i < n){    // remaining}while(j < m){    // remaining}
```

---

## 5. Typical Exam Mistake Traps

- Using `+` instead of `append`
- Ignoring leftover elements
- Index out of bounds
- Overcomplicating with if-else
- Not identifying pattern

---

## 6. Conclusion

Your errors are not conceptual difficulty but **implementation discipline failures**:

- incorrect API usage
- unsafe indexing
- incomplete traversal logic

Focus on patterns + boundary handling.
#failure/debugging #tech/java #dsa/two-pointer
