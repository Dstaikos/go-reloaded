## TASK-006 — Full Transformation Pipeline

- **ID**: TASK-006
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-005
- **Soft Dependencies**: TASK-007
- **Related Blueprint Pillars**: Integration

### Mission Profile
Chain all text processing functions in the main pipeline that processes the entire input string sequentially.

### Deliverables
- Updated `/cmd/main.go` with complete transformation pipeline.
- Integration tests with multiple transformations.

### Steps
1. **Write Tests:** Combined examples with multiple transformations.
2. **Implement:**
   ```go
   text = pkg.HexBin(text)
   text = pkg.UpLowCap(text)
   text = pkg.FixPunctuation(text)
   text = pkg.AnConvert(text)
   ```
3. **Validate:** Integration test matches expected final output.
