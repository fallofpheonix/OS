### Sliding Window

- **Fixed Size**: Maintain a window of size k and update results in O(1) per step.
- **Variable Size (Expand–Shrink)**: Dynamically adjust window boundaries to satisfy a constraint.
- **Monotonic Window**: Use a deque to maintain max/min in a moving window in O(n).

---

### Two Pointer

- **Opposite Ends**: Move two pointers inward to exploit sorted order or symmetry.
- **Same Direction (Fast–Slow)**: Use two speeds to detect structure or compress data in-place.
- **Partition / Dutch Flag**: Rearrange elements into groups using multiple boundary pointers.

---

### Prefix Based

- **Prefix Sum**: Store cumulative sums to answer range queries in O(1).
- **Prefix XOR**: Use cumulative XOR to transform subarray XOR queries into lookups.
- **2D Prefix**: Extend prefix sums to matrices for constant-time submatrix queries.

---

### Kadane / Subarray

- **Max Subarray Sum (Kadane’s)**: Track best running sum by discarding negative prefixes.
- **Max Product Subarray**: Track both max and min due to sign flips from negatives.
- **Subarray with Given Sum/XOR**: Use prefix + hashmap to detect matching subarrays.

---

### Binary Search

- **On Index**: Search a sorted structure by halving the search space each step.
- **On Answer**: Binary search over solution space using a monotonic feasibility check.

## Related
- [[Dsa Patterns]]
