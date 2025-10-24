## 🧠 TASK-007 — Punctuation Normalization

- **ID**: TASK-007  
- **Owner**: Backend Developer  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-006  
- **Soft Dependencies**: TASK-008  
- **Related Blueprint Pillars**: Grammar & Formatting

### Mission Profile
Fix spacing before and after punctuation like `. , ! ? : ;` across the full string.

### Steps
1. **Write Tests:** `"Hello , world !"` → `"Hello, world!"`; `"Wow!!  Great"` → `"Wow!! Great"`.  
2. **Implement:** `FixPunctuation(s string)` using regex replacements.  
3. **Validate:** Run `go test ./pkg`.
