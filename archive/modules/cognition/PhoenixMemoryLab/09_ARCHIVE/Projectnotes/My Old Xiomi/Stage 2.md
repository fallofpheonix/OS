# STAGE 2: MEMORY FUNDAMENTALS (VECTOR)

---

## 1. Objective

Build a **resizable contiguous memory structure** to understand:

- heap allocation
- pointer manipulation
- dynamic resizing
- memory safety

---

## 2. Conceptual Model

A vector is:

> a dynamically resizing array backed by heap memory

Memory layout:

[data pointer] → [x][x][x][x]...[unused space]  
                  ↑  
                elements

---

## 3. Data Structure Definition

typedef struct {  
    int *data;  
    int size;  
    int capacity;  
} Vector;

---

### Field Semantics

#### `data`

- pointer to heap memory
- start of contiguous block

#### `size`

- number of valid elements

#### `capacity`

- total allocated slots

---

## 4. Core Invariants (must always hold)

0 ≤ size ≤ capacity  
data != NULL (after init)

Violation → undefined behavior.

---

## 5. Memory Lifecycle

---

### (A) Allocation

v->data = malloc(sizeof(int) * capacity);

- reserves memory on heap
- uninitialized values

---

### (B) Growth

Triggered when:

size == capacity

Resize logic:

capacity *= 2;  
data = realloc(data, capacity * sizeof(int));

---

### Why doubling matters

Without doubling:

- resizing every insert → O(n²)

With doubling:

- amortized O(1)

---

### (C) Deallocation

free(v->data);

- releases heap memory
- prevents leaks

---

## 6. Operation Breakdown

---

### init

Initializes structure:

- sets size = 0
- sets capacity = small base (2)
- allocates memory

---

### push_back

Steps:

1. Check if resize needed
2. Resize (if required)
3. Insert at `data[size]`
4. Increment `size`

---

### get

- validates index
- returns value

---

### set

- validates index
- modifies element

---

### pop_back

- decreases size
- does NOT shrink memory

---

### free_vector

- frees heap
- resets fields

---

## 7. Complexity Analysis

|Operation|Time|
|---|---|
|push_back|amortized O(1)|
|get|O(1)|
|set|O(1)|
|resize|O(n)|

---

## 8. Memory Behavior

### Contiguous allocation

Advantages:

- fast access
- cache-friendly

Disadvantages:

- expensive resizing
- memory copying

---

### Pointer stability

After `realloc`:

- pointer may change
- old pointer becomes invalid

Correct usage:

new_ptr = realloc(...)  
if (new_ptr != NULL)  
    data = new_ptr;

---

## 9. Error Handling Model

Current design:

- prints error
- returns sentinel (`INT_MIN`)

Problem:

- ambiguous return values
- no propagation

Better design:

- return status
- separate output parameter

---

## 10. Testing Strategy (important)

You implemented:

### Functional tests

- insertion
- retrieval
- modification

---

### Structural tests

- resizing behavior
- capacity correctness

---

### Edge cases

- invalid index
- empty pop

---

### Stress test

- large insertions

---

## 11. Failure Modes (critical learning)

---

### Segmentation fault

Cause:

- accessing out-of-bounds
- invalid pointer

---

### Memory leak

Cause:

- missing `free`

---

### Realloc failure

Cause:

- memory exhaustion

---

### Off-by-one errors

Cause:

- incorrect size updates

---

## 12. Design Tradeoffs

---

### Why vector?

- fast random access
- simple structure

---

### Limitations

- inefficient insert/remove in middle
- requires resizing
- cannot shrink automatically

---

## 13. Conceptual Insight

You built:

> an abstraction layer over raw heap memory

This is the same underlying idea used in:

- dynamic arrays
- buffers
- system-level memory pools

---

## 14. What you did NOT build (important)

- generic type support
- memory shrink strategy
- iterator system
- thread safety

---

## 15. Key Takeaways

- memory is not automatic
- resizing has cost
- pointer correctness is critical
- invariants define correctness

---

## 16. Transition to next stage

Vector taught:

> contiguous memory + resizing

Next:

> non-contiguous memory + pointer chaining

That is:

→ Linked List

---

## Final evaluation

Stage 2 is complete only if you understand:

- why doubling works
- how realloc behaves
- why bounds checking matters
- what happens in memory during operations
