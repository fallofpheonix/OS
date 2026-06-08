# Percentages Simple And Compound Interest Part 01


## Definition
Percentage, simple interest, and compound interest problems are arithmetic models for relative change over fixed bases and time.

## How It Works
Use explicit bases, multiplicative factors for successive percentage change, and separate linear simple-interest growth from compounding growth.

## Complexity / Costs
Time: O(1) per formula-based problem. Space: O(1). Main cost is choosing the correct base and time segment.

## Common Mistakes
- Adding successive percentage changes directly.
- Ignoring base changes.
- Treating compound interest as linear simple interest.
- Guessing reverse percentage problems instead of forming equations.

## Examples And Problem Patterns


#Percentages
## 1. Reverse Percentage Trap

**Question:**  
A number is increased by 25% and then decreased by 20%. What is the net percentage change?

**Solution:**

**Concept:** Successive percentage change  
Final multiplier = (1 + 25/100)(1 − 20/100)  
= 1.25 × 0.8 = 1.0

**Result:** No change (0%)

**Key Point:**  
+25% and −20% do not cancel algebraically; they cancel multiplicatively here.

---

## 2. Hidden Base Change

**Question:**  
The price of a product increases by 20%, and then the quantity demanded decreases by 10%. What is the net effect on total revenue?

**Solution:**

Revenue = Price × Quantity  
New Revenue = (1.2P)(0.9Q) = 1.08PQ

**Net Effect:** 8% increase

**Trap:** Students assume increase and decrease cancel.

---

## 3. Successive Increase Comparison

**Question:**  
Which is greater:  
A: Increase by 30% then decrease by 30%  
B: Decrease by 30% then increase by 30%

**Solution:**

A = 1.3 × 0.7 = 0.91  
B = 0.7 × 1.3 = 0.91

**Conclusion:** Both equal (9% decrease)

**Insight:** Order does not matter in multiplication.

---

## 4. Conditional Percentage Logic

**Question:**  
If A is 20% more than B, and B is 25% less than C, compare A and C.

**Solution:**

A = 1.2B  
B = 0.75C

So,  
A = 1.2 × 0.75C = 0.9C

**Conclusion:** A is 10% less than C

---

## 5. Population Growth Trap

**Question:**  
A town’s population increases by 10% annually for 2 years. What is the total increase?

**Solution:**

Final multiplier = (1.1)² = 1.21

**Increase = 21%**

**Trap:** Not 20%

---

## 6. Weighted Percentage Trick

**Question:**  
In a class, boys increased by 20% and girls increased by 30%. Total class strength increased by 25%. Find ratio of boys to girls initially.

**Solution:**

Let boys = B, girls = G

Total increase:  
(0.2B + 0.3G) / (B + G) = 0.25

0.2B + 0.3G = 0.25B + 0.25G

⇒ 0.05G = 0.05B  
⇒ B = G

**Answer:** Ratio = 1:1

---

## 7. Replacement Concept

**Question:**  
A container has milk. 20% is removed and replaced with water. This is repeated twice. What % of milk remains?

**Solution:**

Remaining milk = (1 − 0.2)² = (0.8)² = 0.64

**Answer:** 64%

---

## 8. Base Shift Trick

**Question:**  
Income increases by 20% and expenditure increases by 25%. What happens to savings?

**Solution:**

Let income = 100, expenditure = 80 → savings = 20

New income = 120  
New expenditure = 100

New savings = 20

**Conclusion:** No change

**Trap:** Percentages mislead unless base is tracked.

---

## 9. Ratio to Percentage Trap

**Question:**  
If A:B = 3:5, by what % is B greater than A?

**Solution:**

Difference = 2  
Base = A = 3

% = (2/3) × 100 = 66.67%

---

## 10. Advanced Reverse Engineering

**Question:**  
After a 20% discount, the selling price is ₹800. What was the marked price?

**Solution:**

Let MP = x

0.8x = 800  
x = 1000

---

## 11. Multi-step Logical Trap

**Question:**  
A salary increases by 50% in year 1, then decreases by 40% in year 2. Compare final salary with original.

**Solution:**

1.5 × 0.6 = 0.9

**Result:** 10% decrease

---

## 12. Extreme Case Reasoning

**Question:**  
A number increases by 100% and then decreases by 50%. Net effect?

**Solution:**

2 × 0.5 = 1

**Answer:** No change

---

## 13. Percentage of Percentage Trap

**Question:**  
30% of A is 20% of B. Find A:B.

**Solution:**

0.3A = 0.2B  
A/B = 2/3

---

## 14. Real Exam-Level Twist

**Question:**  
A shopkeeper marks goods 50% above cost and gives 20% discount. Find profit %.

**Solution:**

CP = 100  
MP = 150  
SP = 150 × 0.8 = 120

Profit = 20

**Profit % = 20%**

---

## 15. High Difficulty Conceptual

**Question:**  
If x% of y equals y% of x, what can be concluded?

**Solution:**

(x/100)y = (y/100)x

⇒ xy = xy

**Conclusion:** Always true

---

## Final Observations

- Most errors come from ignoring **multiplicative nature** of percentages.
- Always convert to **multipliers** instead of additive thinking.
- Track **base values explicitly** in multi-step problems.
- Reverse problems require algebraic setup, not intuition.


#SimpleInterest

# 1. Time-Split Interest Trap

**Question:**  
A sum is invested at 10% p.a. for first 2 years and 15% p.a. for next 3 years. Total simple interest earned is ₹6500. Find the principal.

**Solution:**

**Concept:** Piecewise SI  
SI = P×2×10/100 + P×3×15/100

= 0.2P + 0.45P = 0.65P

0.65P = 6500  
P = 10000

---

# 2. Changing Principal Midway

**Question:**  
A person borrows ₹20,000 at 10% SI. After 1 year, he repays ₹8000. Find total amount due at end of 2nd year.

**Solution:**

Year 1 interest = 20000 × 10% = 2000  
Amount = 22000

After payment → remaining = 14000

Year 2 interest = 14000 × 10% = 1400

Final amount = 15400

---

# 3. Ratio-Based Interest Comparison

**Question:**  
Two sums are in ratio 3:5 and are invested at same rate for same time. Ratio of interests is 9:10. Find ratio of time.

**Solution:**

SI ∝ P × T

So,  
(3T₁)/(5T₂) = 9/10

⇒ 30T₁ = 45T₂  
⇒ T₁/T₂ = 3/2

---

# 4. Reverse Engineering Amount

**Question:**  
A sum becomes ₹1800 in 2 years and ₹2100 in 5 years at SI. Find principal.

**Solution:**

Difference = 300 over 3 years  
Yearly SI = 100

SI for 2 years = 200

P = 1800 − 200 = 1600

---

# 5. Hidden Rate Trap

**Question:**  
A sum triples in 20 years at SI. Find rate.

**Solution:**

Final amount = 3P  
SI = 2P

Using SI formula:  
2P = P × R × 20 / 100

R = 10%

---

# 6. Effective Rate Trick

**Question:**  
₹1000 is lent at 10% for first year and at unknown rate for second year. Total SI for 2 years = ₹250. Find second year rate.

**Solution:**

Year 1 SI = 100  
Remaining SI = 150

150 = 1000 × R /100

R = 15%

---

# 7. Multi-Loan Comparison

**Question:**  
A lends ₹5000 to B at 10% SI and ₹8000 to C at 15% SI. After how many years will total interest from both be ₹9500?

**Solution:**

SI = (5000×10T)/100 + (8000×15T)/100  
= 500T + 1200T = 1700T

1700T = 9500  
T ≈ 5.59 years

---

# 8. Principal Split Strategy

**Question:**  
₹10,000 is divided into two parts such that SI at 10% and 20% for 1 year are equal. Find ratio.

**Solution:**

Let parts = x and (10000 − x)

10% of x = 20% of (10000 − x)

0.1x = 0.2(10000 − x)

0.1x = 2000 − 0.2x

0.3x = 2000  
x = 6666.67

Ratio ≈ 2:1

---

# 9. Advanced Logical Trap

**Question:**  
A sum earns ₹400 interest in 4 years and ₹600 in 6 years. Find rate.

**Solution:**

Difference = ₹200 for 2 years → ₹100/year

So yearly SI = 100

Let P = unknown

100 = P×R/100

PR = 10000

Using SI for 4 years:  
400 = 4×100 (consistent)

Insufficient data for individual P and R, but product known.

---

# 10. Time-Reversal Trick

**Question:**  
At SI, a sum becomes ₹1200 in 2 years and ₹1500 in 5 years. In how many years will it become ₹1800?

**Solution:**

Difference: 300 in 3 years → 100/year

To reach 1800:  
Increase from 1200 = 600

Time = 600/100 = 6 years

Total = 2 + 6 = 8 years

---

# 11. Fractional Rate Trap

**Question:**  
