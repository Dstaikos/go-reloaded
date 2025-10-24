## 🧠 TASK-005 — Hexadecimal Conversion

- **ID**: TASK-005  
- **Owner**: Backend Developer  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-004  
- **Soft Dependencies**: TASK-006  
- **Related Blueprint Pillars**: Numeric Conversion

### Mission Profile
Detect words followed by `(hex)` and convert them to decimal.

### Steps
1. **Write Tests:** `"1E (hex)"` → `"30"`; `"G1 (hex)"` → ignored.  
2. **Implement:** `HexConv(s string)` using `strconv.ParseInt(num, 16, 64)`.  
3. **Validate:** Run `go test ./pkg`.