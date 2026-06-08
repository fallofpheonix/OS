# Phase 7 Part 02

    state[node] = 2;  // Mark as done
    return false;
}
```

**Dry Run:**

```
courses=2, prerequisites=[[1,0]]  (0 → 1)

graph: 0→[1], 1→[]

DFS from 0: state[0]=1, visit 1: state[1]=1, no neighbors, state[1]=2. state[0]=2.
No cycle → return true (can finish) ✓

prerequisites=[[1,0],[0,1]]  (0→1 and 1→0 = cycle)

DFS from 0: state[0]=1, visit 1: state[1]=1, visit 0: state[0]==1 → CYCLE!
return false ✓
```

**State = 3 colors trick:** Unvisited / Currently-on-path / Done. This exact 3-state pattern appears in: alien dictionary, parallel courses, build order — any dependency problem.

---

### 9. CONSTRAINT ANALYSIS — KNOWING WHAT SOLUTION IS ALLOWED

In a real interview, the constraints tell you what complexity is acceptable. Reading constraints is a skill.

|Input Size (n)|Maximum Complexity You Can Use|
|---|---|
|n ≤ 20|O(2ⁿ) — backtracking, bitmask DP|
|n ≤ 1,000|O(n²) — nested loops, interval DP|
|n ≤ 100,000|O(n log n) — sorting, heap, merge sort|
|n ≤ 1,000,000|O(n) — single pass, two pointers|
|n ≤ 10^9|O(log n) — binary search, math|

**Question: Can you solve this in O(n log n)?**

"Given n tasks with deadlines, find maximum tasks completable."

n ≤ 100,000 → O(n log n) is expected → Sort by deadline, use a greedy approach with a heap.

java

```java
int maxTasks(int[][] tasks) {
    Arrays.sort(tasks, (a, b) -> a[1] - b[1]);  // Sort by deadline
    PriorityQueue<Integer> heap = new PriorityQueue<>();  // Min-heap of durations

    int time = 0;

    for (int[] task : tasks) {
        heap.offer(task[0]);  // Add task duration
        time += task[0];

        if (time > task[1]) {  // Missed deadline
            time -= heap.poll();  // Remove longest task
        }
    }

    return heap.size();
}
```

Seeing n ≤ 100,000 immediately tells you: O(n²) won't pass, you need a heap or sort-based approach. Constraints guide your solution choice before you even think about the algorithm.

---

### 10. TRADEOFF PROBLEMS — TIME VS SPACE

Many problems have two valid approaches: fast but memory-heavy, or slow but memory-lean. Know both.

**Question: Find Duplicate in Array**

Three approaches with different tradeoffs:

```
Approach 1: Sort + scan         → O(n log n) time, O(1) space (modifies input)
Approach 2: HashSet             → O(n) time, O(n) space
Approach 3: Floyd on array      → O(n) time, O(1) space (Phase 5 pattern!)
```

When interviewer says "O(1) space, don't modify input" → only Floyd works. Knowing all three and their tradeoffs is what separates good candidates from great ones.

---

### 11. THE FINAL LEVEL: THINKING IN STATES AND INVARIANTS

The highest level of algorithmic thinking is when the data structure disappears from your mind. You no longer think "I need a linked list here." You think:

- **What states exist?** (nodes in the abstract graph of the problem)
- **What transitions are valid?** (edges)
- **What constraint must always hold?** (invariant)
- **What do I want to minimize/maximize?** (objective)

**Example — LRU Cache at the highest level:**

- State: which key was used at what time
- Invariant: most-recently-used stays, least-recently-used goes when full
- Objective: O(1) access and O(1) eviction
- Implication: need O(1) lookup (HashMap) + O(1) ordered removal (DLL)

You derive the implementation from the constraints. You don't memorize "LRU = HashMap + DLL" — you derive it every time from first principles. That's real algorithmic thinking.

---

### 12. COMPLETE PHASE 7 PRACTICE QUESTIONS

**Pattern Recognition:**

- Given a problem, identify the pattern before writing any code (practice with 20 random LeetCode mediums)

**Optimization Practice:**

- Two Sum → HashMap optimization
- Best Time to Buy/Sell Stock → single-pass optimization
- Sliding Window Maximum → monotonic deque

**DP Practice (in order):**

- Climbing Stairs (linear DP)
- House Robber (linear DP with skip)
- Coin Change (unbounded knapsack)
- Longest Common Subsequence (2D DP)
- Word Break (string DP)

**Graph Practice:**

- Number of Islands (DFS + state)
- Course Schedule (cycle detection = topological sort)
- Shortest Path in Binary Matrix (BFS)
- Network Delay Time (Dijkstra)

**Hybrid Problems (combine 2+ patterns):**

- Reorder List (fast-slow + reversal + merge) — already solved
- Sliding Window Maximum (window + monotonic deque)
- Word Search II (trie + DFS + backtracking)
- Merge K Lists (heap + linked lists) — already solved

**Must-complete list (Blind 75 core):**

✅ Two Sum, Three Sum (array + hashmap) ✅ Climbing Stairs, House Robber (linear DP) ✅ Coin Change (knapsack DP) ✅ Number of Islands (graph DFS) ✅ Course Schedule (topological sort) ✅ Merge Intervals (sorting + greedy) ✅ Valid Parentheses (stack) ✅ LRU Cache (DLL + HashMap) ✅ Sliding Window Maximum (deque) ✅ Word Break (string DP)

---

### 13. WHAT TO DO NEXT — YOUR EXACT ROADMAP

**Month 1:** DP Deep Dive

- Linear DP → 2D DP → Knapsack → String DP → Interval DP

**Month 2:** Graph Algorithms

- BFS/DFS → Topological Sort → Union Find → Dijkstra → Bellman-Ford

**Month 3:** Advanced DS

- Segment Trees → Tries → Advanced Heap (Lazy propagation)

**Interview Grind (parallel):**

- NeetCode 150 in order (it's pattern-organized, not random)
- 2 problems per day, always write brute force first
- Review your solution after 3 days (spaced repetition)

---

### FINAL REALIZATION

After 7 phases, here is what you actually learned:

Phases 1-2 taught you **pointer safety** — how to manipulate connected memory without losing nodes.

Phases 3-4 taught you **pattern templates** — the 5-6 core LL patterns that appear in every interview.

Phase 5 taught you **invariant thinking** — maintaining guarantees throughout execution.

Phase 6 taught you **structural generalization** — linked lists are just the simplest pointer network; the skills transfer everywhere.

Phase 7 teaches you **problem decomposition** — any new problem is just a combination of known patterns.

The goal was never linked lists. It was teaching you to think in structures, transitions, and invariants. You now have that foundation. Everything else — trees, graphs, DP — is building on it. 🎯

---

### 14. CODE REVIEW SUMMARY

This phase is strong as an educational sequence, but the code examples should be treated as teaching snippets rather than drop-in production code.

**What is working well**

- The examples map each problem to one clear pattern, which makes the learning path easy to follow.
- Complexity reasoning is explicit and usually paired with a dry run, which is exactly what a reviewable teaching note should do.
- The linked list material correctly emphasizes pointer safety, invariants, and structural transformations.

**Review findings**

- The Java snippets are intentionally partial in a few places, so a reader may need to cross-reference adjacent sections to see the full method body.
- Several snippets omit class wrappers and imports, which is fine for a note but should be called out clearly as pseudocode-style examples.
- In the graph and DP sections, the key learning is the transition logic, so the surrounding explanation is more important than the literal syntax.
- The `Course Schedule` example is structurally correct, but the note should keep the completion step explicit so readers remember that DFS coloring must end with a terminal state.

**Code review guidance for this phase**

1. Read the problem statement first and identify the pattern before the code.
2. Validate invariants before optimizing the implementation.
3. Trace one dry run by hand for every recursive, graph, or DP example.
4. Document the complexity, base cases, and state transitions immediately after the snippet.
5. Treat mutation-heavy examples, like `Number of Islands`, as in-place strategies that may need a non-mutating alternative in real systems.

---

### 15. DOCUMENTATION NOTES

If you are explaining this phase to GPT or Claude, use this structure:

**Phase purpose**
- Phase 7 is about pattern synthesis, not memorizing isolated problems.

**What to explain for each example**
- What structure the problem has.
- What state is being tracked.
- What invariant must stay true.
- Why the chosen data structure is the right one.
- What the bottleneck was and how it was removed.
- What the time and space complexity become.

**Best reusable documentation pattern**

```markdown
## Problem
## Pattern
## State
## Transition
## Base Case
## Invariant
## Dry Run
## Complexity
## Common Mistakes
## Why This Works
```

**Teaching note for AI assistants**
- This phase is best summarized as: "recognize the pattern, justify the state, prove the transition, and verify the invariant."

**Completion marker**
- Phase 7 documentation is complete when a reader can take a new problem, identify the pattern, and explain the solution without looking at the code first.
