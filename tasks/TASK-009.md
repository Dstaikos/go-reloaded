## TASK-009 — “a” / “an” Article Rule

- **ID**: TASK-009  
- **Owner**: Backend Developer  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-008  
- **Soft Dependencies**: TASK-010  
- **Related Blueprint Pillars**: Grammar Correction

### Mission Profile
Replace “a”/“A” with “an”/“An” when followed by a vowel or `h`.

### Steps
1. **Write Tests:** `"A amazing idea"` → `"An amazing idea"`; `"a hour"` → `"an hour"`.  
2. **Implement:** `FixArticles(s string)` scanning for patterns using regex.  
3. **Validate:** Run `go test ./pkg`.
