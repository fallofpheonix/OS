# Bit Manipulation

## Pattern Cheatsheet


- **XOR Pattern**: Use XOR properties (a^a=0, a^0=a) to cancel or isolate values.
- **Bit Masking**: Use bits to represent states or toggle/check specific positions.

---

### Usage

- **Subset via Bits**: Generate all subsets by iterating over bitmasks from 0 to 2ⁿ−1.
- **Bit Checks**: Test, set, or clear bits using bitwise operations (&, |, ^, ~).
- **Prefix XOR**: Store cumulative XOR to answer subarray XOR queries efficiently.
---

## XOR Cancellation Pattern


**Bit Manipulation** is a technique where operations are performed directly on binary representations of numbers using operators like `&`, `|`, `^`, `~`, `<<`, `>>`.

---

## 2. XOR Operator (Most Important)

### Key Properties

1. `a ^ a = 0`
2. `a ^ 0 = a`
3. XOR is **commutative** → `a ^ b = b ^ a`
4. XOR is **associative** → `(a ^ b) ^ c = a ^ (b ^ c)`

From notes: XOR cancels identical values

---

## 3. Core Intuition

- Same values cancel out
- Only **unique/unpaired element remains**

---

## 4. When to Use XOR

Apply XOR when:

- One element is **extra or unique**
- All other elements appear **twice**
- Order does not matter
- Need **O(1) space**

---

## 5. Standard Template

### (A) Find Unique Element

```
int result = 0;for(int num : arr){    result ^= num;}return result;
```

---

### (B) Compare Two Strings (Your Problem)

```
char result = 0;for(char c : s.toCharArray()) result ^= c;for(char c : t.toCharArray()) result ^= c;return result;
```

---

## 6. Example

s = "abcd"  
t = "abcde"

```
(a ^ b ^ c ^ d) ^ (a ^ b ^ c ^ d ^ e)= e
```

---

## 7. Common XOR Problems

- Find single non-duplicate element
- Find missing number
- Find extra character (this problem)
- Swap two numbers without temp
- Subarray XOR problems

---

## 8. Complexity

- Time: **O(n)**
- Space: **O(1)**

---

## 9. Advantages

- No extra memory
- Very fast (bit-level)
- Clean and minimal code

---

## 10. Common Mistakes

- Confusing XOR with addition
- Forgetting cancellation property
- Using wrong data type (overflow in sum method)
- Applying XOR when duplicates are not guaranteed

---

## 11. Limitations

- Works only when pattern fits (pairs + one unique)
- Not useful for general frequency problems

---

## 12. Quick Revision Sheet

- XOR = cancellation tool
- `a ^ a = 0`
- `a ^ 0 = a`
- Use when:
    - one extra / unique element
- No sorting, no hashmap needed

---

## 13. Conclusion

XOR-based bit manipulation is a **pattern-driven optimization technique** used to eliminate duplicates and isolate unique elements efficiently in linear time and constant space.
## Related
- [[Two-Pointer Technique]]
- [[Linked List]]

#concept #dsa/bit-manipulation
