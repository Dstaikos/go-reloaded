## TASK-006 — "a" / "an" Article Rule

- **ID**: TASK-006
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-002, TASK-003, TASK-004, TASK-005
- **Soft Dependencies**: None
- **Related Blueprint Pillars**: Grammar Correction

### Mission Profile
Replace "a"/"A" with "an"/"An" when followed by a vowel or `h`.

### Deliverables
- `/pkg/ArticleConvert.go` with `ArticleConvert()` function.
- Helper functions: `isVowel()`.
- Unit tests for article conversion rules.

### Steps (TDD Flow)
1. **Write Tests:** `"A amazing idea"` → `"An amazing idea"`; `"a hour"` → `"an hour"`.
2. **Implement:** `ArticleConvert()` scanning for patterns using regex or manual parsing.
3. **Validate:** Run `go test ./pkg` — all article conversion tests pass.
