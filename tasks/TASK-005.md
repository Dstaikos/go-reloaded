## TASK-005 — Punctuation Normalization

- **ID**: TASK-005
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-002, TASK-003, TASK-004
- **Soft Dependencies**: None
- **Related Blueprint Pillars**: Grammar & Formatting

### Mission Profile
Fix spacing before and after punctuation like `. , ! ? : ;` and handle apostrophe formatting across the full string.

### Deliverables
- `/pkg/Punctuation.go` with `FixPunctuation()` function.
- Unit tests for punctuation spacing and apostrophe handling.

### Steps (TDD Flow)
1. **Write Tests:** `"Hello , world !"` → `"Hello, world!"`; `"I am ' cool '"` → `"I am 'cool'"`.
2. **Implement:** `FixPunctuation()` using regex replacements or manual parsing.
3. **Validate:** Run `go test ./pkg` — all punctuation tests pass.
