## TASK-002 — Uppercase Transformation

- **ID**: TASK-002  
- **Owner**: Backend Developer  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-001  
- **Soft Dependencies**: TASK-003  
- **Related Blueprint Pillars**: Text Transformation

### Mission Profile
Implement `(up)` and `(up, x)` functionality that capitalizes the previous one or `x` words in the text, modifying the full string without tokenization.

### Deliverables
- `/pkg/Upper.go` (already started).  
- Unit tests for all valid and invalid `(up)` formats.

### Steps (TDD Flow)
1. **Write Tests:** `"this is nice (up)"` → `"this is NICE"`; `"wow cool text (up, 2)"` → `"wow COOL TEXT"`.  
2. **Implement:** Complete `Upper()` using `unicode` and helper functions.  
3. **Validate:** Run `go test ./pkg` — all uppercase cases pass.