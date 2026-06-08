# Phase 7 Part 01

PHASE 7: ADVANCED PROBLEM SOLVING & PATTERN SYNTHESIS — With Questions, Examples & Explanations

This is the final phase. The goal shifts from "solve this problem" to "recognize the structure, identify the pattern, derive the solution." Every section here includes a real question, full explanation, and dry run.

---

### 1. THE CORE REALIZATION: ~20 PATTERNS, NOT 10,000 PROBLEMS

Every LeetCode problem you will ever see in an interview is either one pattern or a combination of patterns. Once you internalize this, interviews stop being scary — they become pattern recognition exercises.

The complete map:

|Domain|Pattern|Core Idea|
|---|---|---|
|Array|Prefix Sum|precompute cumulative values|
|Array|Sliding Window|dynamic range that grows/shrinks|
|Array|Two Pointers|relative movement from ends or same direction|
|Array|Binary Search|exploit monotonicity|
|Array|Kadane|local vs global optimization|
|Linked List|Reversal|3-pointer edge flipping|
|Linked List|Fast-Slow|relative speed detection|
|Linked List|Merge|ordered combining|
|Linked List|Split/Reconnect|structural transformation|
|Stack|Monotonic Stack|nearest greater/smaller element|
|Stack|Expression Parsing|precedence management|
|Tree|DFS|recursive decomposition|
|Tree|BFS|level-by-level processing|
|Tree|Divide & Conquer|subtree independence|
|Graph|BFS|shortest unweighted path|
|Graph|DFS|traversal + state|
|Graph|Topological Sort|dependency ordering|
|Graph|Union Find|connectivity grouping|
|Graph|Dijkstra|weighted shortest path|
|DP|Linear DP|sequential dependency|
|DP|Knapsack|choose/skip decisions|
|DP|Grid DP|movement transitions|
|DP|Interval DP|subrange optimization|

You don't need to memorize 1000 solutions. You need to deeply understand these ~20 patterns.

---

### 2. THE 4-QUESTION DECOMPOSITION FRAMEWORK

When you see any new problem, immediately ask these 4 questions before writing a single line of code:

**Q1: What structure exists in the data?**

- Numbers in a line → array patterns
- Nodes with connections → graph/tree patterns
- Ordered sequence → two pointer / binary search
- Nested decisions → DP / backtracking

**Q2: What operation repeats?**

- "Find nearest..." → monotonic stack
- "Best path from A to B..." → BFS/Dijkstra
- "All combinations/subsets..." → backtracking
- "Min/Max over range..." → prefix sum / segment tree

**Q3: Is state reusable?**

- Same subproblem appears multiple times → DP candidate
- If brute force has overlapping subproblems → definitely DP

**Q4: Is there a monotonic property?**

- "If X works, X+1 also works" → binary search on answer
- "As window grows, result only increases" → sliding window

Practice applying these 4 questions to every problem you see. The pattern will reveal itself within 2 minutes.

---

### 3. OPTIMIZATION THINKING — THE MOST IMPORTANT INTERVIEW SKILL

Interviewers don't just want a correct answer. They want to see you _think about reducing complexity_.

**The standard optimization routes:**

|Bottleneck|Solution|Why|
|---|---|---|
|O(n²) nested search|HashMap|O(1) lookup replaces O(n) scan|
|Recomputing same range|Prefix Sum|precompute once, answer in O(1)|
|Recomputing same subproblem|Memoization/DP|cache previous results|
|Sorting k times|Heap|maintain sorted order dynamically|
|O(n) search in sorted data|Binary Search|halve search space each step|

**Question: Two Sum (Classic optimization example)**

Brute force: For every element, scan rest of array for complement. O(n²).

Optimization question: "What's slow?" The inner scan. Can we look up a complement in O(1)? Yes — HashMap.

java

```java
int[] twoSum(int[] nums, int target) {
    HashMap<Integer, Integer> map = new HashMap<>();  // value → index

    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];

        if (map.containsKey(complement)) {
            return new int[]{map.get(complement), i};
        }

        map.put(nums[i], i);
    }

    return new int[]{};
}
```

**Dry Run:**

```
nums = [2, 7, 11, 15], target = 9

i=0: complement = 7. Map has 7? No. Put {2:0}.
i=1: complement = 2. Map has 2? YES → return [0, 1] ✓
```

**The thinking process:** Brute → bottleneck = inner search → replace with HashMap → O(n). This exact thought process is what interviewers are watching for.

---

### 4. BRUTE FORCE FIRST — NON-NEGOTIABLE RULE

Never jump to optimal. The correct interview flow is always:

1. State brute force + its complexity
2. Identify the bottleneck ("the slow part is...")
3. Optimize that bottleneck with a data structure or insight
4. State new complexity

**Question: Sliding Window Maximum**

Given array `nums` and window size `k`, find the maximum in every window of size k.

**Brute force:** For every window, scan all k elements. O(n×k).

```
nums = [1, 3, -1, -3, 5, 3, 6, 7], k = 3
Windows: [1,3,-1]=3, [3,-1,-3]=3, [-1,-3,5]=5, ...
```

**Bottleneck:** The inner O(k) scan per window.

**Optimization insight:** We need to efficiently track the maximum as the window slides. When element enters, it might be the new max. When it leaves, we need the next max. A monotonic deque (decreasing order) handles this: it always stores the maximum at the front.

java

```java
int[] maxSlidingWindow(int[] nums, int k) {
    int n = nums.length;
    int[] result = new int[n - k + 1];
    Deque<Integer> deque = new ArrayDeque<>();  // stores indices

    for (int i = 0; i < n; i++) {
        // Remove elements outside the window
        if (!deque.isEmpty() && deque.peekFirst() < i - k + 1) {
            deque.pollFirst();
        }

        // Remove elements smaller than current (they can never be the max)
        while (!deque.isEmpty() && nums[deque.peekLast()] < nums[i]) {
            deque.pollLast();
        }

        deque.offerLast(i);

        // Window is fully formed
        if (i >= k - 1) {
            result[i - k + 1] = nums[deque.peekFirst()];
        }
    }

    return result;
}
```

**Dry Run:**

```
nums = [1, 3, -1, -3, 5, 3, 6, 7], k = 3

i=0: deque=[0(val=1)]
i=1: 3 > 1, remove 0. deque=[1(val=3)]
i=2: -1 < 3, keep. deque=[1,2]. Window ready: result[0]=nums[1]=3
i=3: -3 < -1, keep. deque=[1,2,3]. result[1]=nums[1]=3
i=4: 5 > -3,-1,3 → clear. deque=[4]. result[2]=5
i=5: 3 < 5. deque=[4,5]. result[3]=5
i=6: 6 > 3,5 → clear. deque=[6]. result[4]=6
i=7: 7 > 6. deque=[7]. result[5]=7

Output: [3, 3, 5, 5, 6, 7] ✓
```

**Pattern used:** Sliding Window + Monotonic Deque. Two patterns combined.

---

### 5. RECURSION TREE THINKING

Every recursive algorithm creates a tree of calls. Visualizing this tree tells you the time and space complexity immediately.

**Question: Generate All Subsets**

java

```java
void subsets(int[] nums, int start, List<Integer> current, List<List<Integer>> result) {
    result.add(new ArrayList<>(current));  // Add current state

    for (int i = start; i < nums.length; i++) {
        current.add(nums[i]);              // Include
        subsets(nums, i + 1, current, result);
        current.remove(current.size() - 1); // Exclude (backtrack)
    }
}
```

**Recursion Tree for [1, 2, 3]:**

```
                    []
          /          |          \
        [1]         [2]         [3]
       /   \          \
    [1,2] [1,3]      [2,3]
      |
   [1,2,3]

Results: [], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]  → 2³ = 8 subsets
```

**What the tree tells you:**

- Depth = n (one decision per element) → O(n) recursion stack space
- Total nodes = 2ⁿ → O(2ⁿ) time and result space
- Each node = one include/exclude choice

Whenever you see "all combinations/subsets/permutations" → draw a binary recursion tree. The depth gives stack space, the number of leaves gives time complexity.

---

### 6. STATE MANAGEMENT — THE UNIFYING CONCEPT

The deepest insight of Phase 7: almost every hard problem is a state management problem. The data structure you choose is just a way to manage state efficiently.

|Problem Type|State|How Managed|
|---|---|---|
|DFS/Graph|visited nodes|HashSet|
|Backtracking|current path|List (add/remove)|
|DP|subproblem answers|dp[] array or HashMap|
|BFS|current level nodes|Queue|
|LRU Cache|access recency|DLL + HashMap|
|Sliding Window|current window|two index pointers|

**Question: Number of Islands (DFS + state management)**

java

```java
int numIslands(char[][] grid) {
    int count = 0;
    int rows = grid.length, cols = grid[0].length;

    for (int r = 0; r < rows; r++) {
        for (int c = 0; c < cols; c++) {
            if (grid[r][c] == '1') {
                count++;
                dfs(grid, r, c);  // Sink the entire island
            }
        }
    }
    return count;
}

void dfs(char[][] grid, int r, int c) {
    if (r < 0 || c < 0 || r >= grid.length || c >= grid[0].length || grid[r][c] != '1') return;

    grid[r][c] = '0';  // Mark visited by sinking (state change)

    dfs(grid, r + 1, c);
    dfs(grid, r - 1, c);
    dfs(grid, r, c + 1);
    dfs(grid, r, c - 1);
}
```

**Dry Run:**

```
Grid:
1 1 0
1 0 0
0 0 1

(0,0)=1: count=1, DFS sinks entire top-left island → all 1s become 0s
(2,2)=1: count=2, DFS sinks it

Output: 2 ✓
```

**State here:** visited cells. Managed by modifying grid in-place (no extra HashSet needed). That's O(1) extra space vs O(n×m) for a visited array — an optimization that comes from recognizing you can modify state in the input itself.

---

### 7. DYNAMIC PROGRAMMING TRANSITION

DP is just memoized recursion with a systematic framework. You're now ready for it because you've internalized recursion trees.

**The 4 DP questions for any problem:**

1. What defines the state? (what changes between subproblems)
2. What are the transitions? (how do we get from smaller to larger)
3. What is the recurrence relation?
4. What are the base cases?

**Question: Climbing Stairs**

You can climb 1 or 2 steps. How many ways to reach step n?

**Step 1: State.** dp[i] = number of ways to reach step i.

**Step 2: Transitions.** To reach step i, you came from step i-1 (took 1 step) or step i-2 (took 2 steps).

**Step 3: Recurrence.** dp[i] = dp[i-1] + dp[i-2]

**Step 4: Base cases.** dp[1] = 1, dp[2] = 2

java

```java
int climbStairs(int n) {
    if (n <= 2) return n;

    int[] dp = new int[n + 1];
    dp[1] = 1;
    dp[2] = 2;

    for (int i = 3; i <= n; i++) {
        dp[i] = dp[i-1] + dp[i-2];
    }

    return dp[n];
}
```

**Dry Run (n=5):**

```
dp[1]=1, dp[2]=2
dp[3] = dp[2]+dp[1] = 3
dp[4] = dp[3]+dp[2] = 5
dp[5] = dp[4]+dp[3] = 8

Output: 8 ✓
```

This is just Fibonacci with renamed variables. The pattern: "ways to reach current = ways from each predecessor" appears in coin change, house robber, decode ways — all linear DP.

**Question: 0/1 Knapsack (Choose/Skip pattern)**

Given weights and values, pick items to maximize value within capacity W.

java

```java
int knapsack(int[] weights, int[] values, int W) {
    int n = weights.length;
    int[][] dp = new int[n + 1][W + 1];

    for (int i = 1; i <= n; i++) {
        for (int w = 0; w <= W; w++) {
            // Skip item i
            dp[i][w] = dp[i-1][w];

            // Include item i (if it fits)
            if (weights[i-1] <= w) {
                dp[i][w] = Math.max(dp[i][w], values[i-1] + dp[i-1][w - weights[i-1]]);
            }
        }
    }

    return dp[n][W];
}
```

**Dry Run:**

```
weights=[1,3,4,5], values=[1,4,5,7], W=7

dp[1][w]: item1(w=1,v=1)
  dp[1][0]=0, dp[1][1]=1, dp[1][2]=1, ... dp[1][7]=1

dp[2][w]: item2(w=3,v=4)
  dp[2][3] = max(dp[1][3], 4+dp[1][0]) = max(1,4) = 4
  dp[2][4] = max(1, 4+dp[1][1]) = max(1,5) = 5
  dp[2][7] = max(1, 4+dp[1][4]) = max(1,5) = 5

... continuing ...

dp[4][7] = 9 (items 2+4: weight 3+5=8>7, or items 2+3: 4+5=7, value 4+5=9) ✓
```

The "choose/skip" decision structure in Knapsack appears in: coin change, partition equal subset sum, target sum — recognize the pattern, apply the same template.

---

### 8. GRAPH THINKING — PROBLEMS DISGUISED AS GRAPHS

Many problems that look nothing like graphs are actually graph problems. The key is recognizing that states are nodes and transitions are edges.

**Question: Course Schedule (Cycle Detection = Topological Sort)**

Can you finish all courses given prerequisites? This is: "does this directed graph have a cycle?"

java

```java
boolean canFinish(int numCourses, int[][] prerequisites) {
    List<List<Integer>> graph = new ArrayList<>();
    for (int i = 0; i < numCourses; i++) graph.add(new ArrayList<>());

    for (int[] pre : prerequisites) {
        graph.get(pre[1]).add(pre[0]);  // pre[1] must come before pre[0]
    }

    int[] state = new int[numCourses];  // 0=unvisited, 1=visiting, 2=done

    for (int i = 0; i < numCourses; i++) {
        if (hasCycle(graph, state, i)) return false;
    }
    return true;
}

boolean hasCycle(List<List<Integer>> graph, int[] state, int node) {
    if (state[node] == 1) return true;   // Currently visiting = cycle!
    if (state[node] == 2) return false;  // Already done = safe

    state[node] = 1;  // Mark as visiting

    for (int neighbor : graph.get(node)) {
        if (hasCycle(graph, state, neighbor)) return true;
    }

