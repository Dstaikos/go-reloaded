## 🧠 TASK-006 — Binary Conversion

- **ID**: TASK-006  
- **Owner**: Backend Developer  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-005  
- **Soft Dependencies**: TASK-007  
- **Related Blueprint Pillars**: Numeric Conversion

### Mission Profile
Detect `(bin)` and convert the preceding binary number to decimal.

### Steps
1. **Write Tests:** `"101 (bin)"` → `"5"`; `"abc (bin)"` → ignored.  
2. **Implement:** `BinConv(s string)` using `strconv.ParseInt(num, 2, 64)`.  
3. **Validate:** Unit tests pass and conversions correct.