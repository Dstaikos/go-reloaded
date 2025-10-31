## TASK-001 — Project Initialization & File Handling

- **ID**: TASK-001
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: None
- **Soft Dependencies**: None
- **Related Blueprint Pillars**: File I/O, CLI UX

### Mission Profile
Set up the project structure and CLI entry point to read from an input file, apply string transformations, and display or write results.

### Deliverables
- `/cmd/main.go` reading `os.Args[1]` (input) and optionally writing output.
- Graceful error messages for missing files or args.
- Unit test verifying file reading/writing and CLI argument validation.

### Steps (TDD Flow)
1. **Write Tests:** Verify file read returns correct content; missing file triggers error message.
2. **Implement:** Use `os.ReadFile` and optional `os.WriteFile`.
3. **Validate:** Run `go run ./cmd/main.go input.txt output.txt` — text printed to console.
