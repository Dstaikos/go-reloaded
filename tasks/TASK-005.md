## TASK-005 — "a" / "an" Article Rule

- **ID**: TASK-005
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-004
- **Soft Dependencies**: TASK-006
- **Related Blueprint Pillars**: Grammar Correction

### Mission Profile
Replace "a"/"A" with "an"/"An" when followed by a vowel or `h`.

### Deliverables
- `/pkg/anconv.go` with `AnConvert()` function.  
- Unit tests for article conversion rules.

### Steps
1. **Write Tests:** `"A amazing idea"` → `"An amazing idea"`; `"a hour"` → `"an hour"`.  
2. **Implement:** `AnConvert()` scanning for patterns using regex or manual parsing.  
3. **Validate:** Run `go test ./pkg` — all article conversion tests pass.
