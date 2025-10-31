## TASK-002 — Helper Functions & Utilities

- **ID**: TASK-002-helper
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-001
- **Soft Dependencies**: None
- **Related Blueprint Pillars**: Code Utilities

### Mission Profile
Implement core helper functions that will be used across multiple transformation modules, particularly space handling utilities.

### Deliverables
- `/pkg/RemoveSpces.go` with `removeTrailingSpaces()` function.
- Unit tests for space removal edge cases.

### Steps (TDD Flow)
1. **Write Tests:** Verify trailing space removal from rune slices; empty slice handling; mixed whitespace characters.
2. **Implement:** `removeTrailingSpaces(newRunes *[]rune)` using `unicode.IsSpace()` to clean up text before applying transformations.
3. **Validate:** Run `go test ./pkg` — all utility function tests pass.
