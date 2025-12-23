## TASK-004 — Hex & Binary Conversion

- **ID**: TASK-004
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-002
- **Soft Dependencies**: TASK-003
- **Related Blueprint Pillars**: Numeric Conversion

### Mission Profile
Implement `(hex)` and `(bin)` conversions in a single function that detects and converts hexadecimal and binary numbers to decimal.

### Deliverables
- `/pkg/NumberConvert.go` with `NumberConvert()` function handling both conversions.
- Helper functions: `applyHex()`, `applyBin()`, `trimTokenBounds()`, `removeTrailingSpaces()`.
- Unit tests for valid/invalid hex and binary formats.

### Steps (TDD Flow)
1. **Write Tests:** `"1E (hex)"` → `"30"`; `"101 (bin)"` → `"5"`; invalid formats ignored.
2. **Implement:** Single `NumberConvert()` function using `strconv.ParseInt()` with base 16 and 2, plus token boundary trimming.
3. **Validate:** Run `go test ./pkg` — all conversion tests pass.
