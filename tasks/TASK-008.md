## 🧠 TASK-008 — Apostrophe Formatting

- **ID**: TASK-008  
- **Owner**: Backend Developer  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-007  
- **Soft Dependencies**: TASK-009  
- **Related Blueprint Pillars**: Grammar & Readability

### Mission Profile
Ensure single quotes are correctly paired and without inner spacing.

### Steps
1. **Write Tests:** `"I am ' cool '"` → `"I am 'cool'"`.  
2. **Implement:** `FixApostrophes(s string)` using regex or manual state tracking.  
3. **Validate:** Unit tests ensure multiple quoted segments handled.
