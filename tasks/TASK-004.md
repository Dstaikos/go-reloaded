## TASK-004 — Capitalization Transformation

- **ID**: TASK-004  
- **Owner**: Backend Developer  
- **Size**: M  
- **Confidence**: Medium  
- **Hard Dependencies**: TASK-003  
- **Soft Dependencies**: TASK-005  
- **Related Blueprint Pillars**: Text Transformation

### Mission Profile
Add `(cap)` and `(cap, x)` modifiers that capitalize only the first letter of the previous one or `x` words.

### Steps
1. **Write Tests:** `"hello world (cap, 2)"` → `"Hello World"`.  
2. **Implement:** `Capitalize(s string)` reusing `unicode.ToUpper` for first rune per word.  
3. **Validate:** Test both single and multi-word use.