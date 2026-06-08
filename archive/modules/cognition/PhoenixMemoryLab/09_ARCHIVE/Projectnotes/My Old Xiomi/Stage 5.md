# Stage 5: Linked List (Detailed Notes)

---

## 1. What you are building

> **Singly Linked List using heap-allocated nodes**

Unlike vector:

- NOT contiguous memory
- each element is a separate allocation
- connected via pointers

---

## 2. Core structure

typedef struct Node {  
    int data;  
    struct Node *next;  
} Node;

---

## 3. Conceptual model

Vector:

[data][data][data][data]

Linked List:

[data | *] -> [data | *] -> [data | NULL]

Each node:

- stores value
- points to next node

---

## 4. Memory model (critical difference)

### Vector

- single `malloc`
- resizing via `realloc`

### Linked List

- multiple `malloc` calls
- no resizing
- nodes scattered in memory

---

## 5. Core operations breakdown

---

### (A) create_node

Node* create_node(int value);

**Purpose:**

- allocate memory for one node
- initialize fields

---

### Key concept

Node *n = malloc(sizeof(Node));

Each node = independent heap allocation.

---

### (B) push_front

void push_front(Node **head, int value);

---

### Why `Node **head`?

Because:

- you may modify the head pointer itself
- need pointer to pointer

---

### Logic

new_node -> next = current head  
head = new_node

---

### (C) push_back

void push_back(Node **head, int value);

---

### Logic

1. if list empty → new node = head
2. else traverse until `next == NULL`
3. attach node

---

### Key cost

- O(n) traversal

---

### (D) get

int get(Node *head, int index);

---

### Logic

- traverse node by node
- stop at index

---

### Key difference from vector

|Operation|Vector|Linked List|
|---|---|---|
|access|O(1)|O(n)|

---

### (E) delete_value

void delete_value(Node **head, int value);

---

### Logic cases

1. delete head
2. delete middle
3. delete last

---

### Critical concept

You must track:

previous node  
current node

---

### Memory handling

free(node)

Missing this → memory leak.

---

### (F) free_list

void free_list(Node *head);

---

### Logic

Traverse and free:

while (head != NULL):  
    temp = head  
    head = head->next  
    free(temp)

---

## 6. Core invariants

Must always hold:

head == NULL OR valid node  
last node -> next == NULL  
no cycles (in this version)

---

## 7. Common failure points (you WILL hit these)

---

### (A) Losing head pointer

Incorrect:

head = head->next;

If done wrong → list becomes unreachable.

---

### (B) Memory leak

Forgetting:

free(node);

---

### (C) Dangling pointer

Using node after `free`.

---

### (D) Incorrect traversal

Infinite loop if:

next pointer misassigned

---

## 8. Complexity comparison (important insight)

|Operation|Vector|Linked List|
|---|---|---|
|insert end|amortized O(1)|O(n)|
|insert front|O(n)|O(1)|
|access|O(1)|O(n)|
|memory|compact|fragmented|

---

## 9. Why this matters

This teaches:

- pointer chaining
- non-contiguous memory handling
- dynamic structure growth without resizing

These are used in:

- OS kernel structures
- schedulers
- memory allocators

---

## 10. Design weaknesses (you must avoid)

- silent failures
- no null checks
- modifying head incorrectly
- mixing traversal + modification logic

---

## 11. Testing requirements

You must write tests like vector:

### Required tests

- insert sequence
- insert at front
- delete head
- delete middle
- delete non-existent value
- get invalid index
- free full list

---

## 12. Key mental shift

Vector thinking:

> index-based access

Linked list thinking:

> pointer traversal

If you still think in indices → you will fail.

---

## 13. What you will learn (core outcomes)

- pointer-to-pointer usage
- memory ownership
- traversal correctness
- manual lifecycle management

---

## 14. Relationship to next stages

After this:

- trees = linked structures
- graphs = generalized linked structures
- OS = heavy pointer-based systems

---

## 15. Final directive

Do NOT start coding yet if:

- you don’t understand `Node **head`
- you can’t trace pointer changes step by step

---
