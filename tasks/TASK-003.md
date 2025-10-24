## TASK-003 — Lowercase Transformation

- **ID**: TASK-003  
- **Owner**: Backend Developer  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: TASK-004  
- **Related Blueprint Pillars**: Text Transformation

### Mission Profile
Implement `(low)` and `(low, x)` modifiers to lowercase the previous one or `x` words.

### Steps
1. **Write Tests:** `"I AM HERE (low, 2)"` → `"I am here"`.  
2. **Implement:** `Lower(s string)` with logic similar to `Upper()`.  
3. **Validate:** Unit tests confirm correctness and skip invalid forms.
