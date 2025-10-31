## TASK-002 — Case Transformation (up/low/cap)

- **ID**: TASK-002  
- **Owner**: Backend Developer  
- **Size**: L  
- **Confidence**: High  
- **Hard Dependencies**: TASK-001  
- **Soft Dependencies**: TASK-003  
- **Related Blueprint Pillars**: Text Transformation

### Mission Profile
Implement all case transformations in a single function: `(up)`, `(up, x)`, `(low)`, `(low, x)`, `(cap)`, and `(cap, x)` that modify the previous one or `x` words.

### Deliverables
- `/pkg/UpLowCap.go` with `UpLowCap()` function handling all three transformations.  
- Helper functions: `applyUp()`, `applyLow()`, `applyCap()`, `removeTrailingSpaces()`.  
- Unit tests for all valid and invalid formats.

### Steps (TDD Flow)
1. **Write Tests:** `"this is nice (up)"` → `"this is NICE"`; `"I AM HERE (low, 2)"` → `"I am here"`; `"hello world (cap, 2)"` → `"Hello World"`.  
2. **Implement:** Single `UpLowCap()` function with pattern detection for all three modifiers using rune-based parsing.  
3. **Validate:** Run `go test ./pkg` — all case transformation tests pass.