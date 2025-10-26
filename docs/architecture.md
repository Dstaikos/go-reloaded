# Architecture — go-reloaded

## Overview

This document describes the high-level architecture of the `go-reloaded` repository, the responsibilities of its components, the main dataflow, contracts, edge cases, build/run instructions, and recommended next steps.

## Repository layout

- cmd/
  - `main.go` — CLI entry point. Reads an input text file, applies transformations from `pkg`, and prints the result.
- pkg/
  - Library code that implements the text transformation rules (package `piscine`). Key files:
    - `Upper.go`, `Lower.go`, `UpLowCap.go` — detect and apply `(up)`, `(low)`, `(cap)` commands and variants like `(up, 2)`.
    - `RemoveSpces.go` — helper `removeTrailingSpace` that trims trailing spaces from a rune slice.
    - `hexconv.go` — placeholder for hex conversion logic (not yet implemented).
- docs/ — documentation and analysis files (this file added here).
- tasks/ — task definitions and TODOs.

## Components & responsibilities

- CLI (`cmd/main.go`)
  - Reads file bytes from disk.
  - Converts bytes to a `string` and feeds it to the `piscine` package functions.
  - Prints processed text to stdout.

- `piscine` package (`pkg/`)
  - Responsible for parsing inline transformation directives (e.g., `(up)`, `(low, 2)`, `(hex)`) and punctuation/spacing rules.
  - Works at rune level to correctly handle Unicode.
  - Contains helper routines for trimming spaces and applying casing transformations to the previous word(s).

## Main data flow

1. Input file -> bytes -> string
2. `piscine.Upper(text)` (as called in `main.go`) — applies `(up)` transformations present in text
3. `piscine.Lower(text)` — applies `(low)` transformations present in text
   - Note: The current `main.go` calls `Upper` then `Lower`. Order matters; review if this is the intended pipeline.
4. Result printed to stdout

Internally, each transformation function converts the string into a `[]rune`, iterates runes, appends to a working `newRunes` slice, and when it sees a directive (like `(up)`), it removes any trailing spaces from `newRunes` then applies the appropriate `applyX` function to the previous N words.

## Contracts (inputs/outputs/error modes)

- Public API (package `piscine`): functions accept and return `string`.
  - Input: arbitrary UTF-8 text that may contain directives and punctuation.
  - Output: transformed text following the spec in `docs/analysis.md`.
- Error modes:
  - CLI `main.go` currently assumes `os.Args[1]` exists; missing args will cause a runtime panic. Add a guard to return a friendly error.
  - Parsing directives is strict about exact spacing and format (e.g., `(up, 2)`); malformed directives are treated as plain text.

## Important implementation details

- Unicode-aware: operations use `[]rune` and `unicode` package functions (`IsSpace`, `IsLetter`, `ToUpper`, `ToLower`) to handle multi-byte characters correctly.
- Helper `removeTrailingSpace(newRunes *[]rune)` mutates the working rune slice to remove trailing whitespace before applying transformations. This function is unexported and used by multiple files.
- Some files (e.g., `UpLowCap.go`) call `removeTrailingSpaces` (plural) while your actual helper is named `removeTrailingSpace` (singular). Keep function names consistent to avoid "undefined" errors.

## Edge cases and constraints

- Directive parsing is strict: extra or missing spaces inside parentheses will cause directives to be ignored.
- Punctuation rules in `analysis.md` (attach punctuation to previous word and ensure one space after) need careful handling when punctuation sequences like `...` or `!?` appear.
- Apostrophes: the spec requires paired single quotes with no surrounding spaces — parsing and enforcement need careful tokenization.
- Single-letter article rule (`a`/`A` -> `an`/`An` before vowels or `h`) must consider punctuation and beginning-of-line cases.
- When applying transformations to multiple previous words (e.g., `(up, 3)`), ensure there are at least that many words — code currently stops early if there aren't enough words.

## Build and run (PowerShell)

From repository root:

```powershell
# Build everything
go build ./...

# Run main with an input file (provide a path to your .txt file)
go run cmd/main.go .\cmd\your-input.txt
```

Notes:
- `cmd/main.go` currently reads `os.Args[1]` without validating args. Use `go run cmd/main.go path\to\file.txt` or add a guard in `main.go`.

## Recommendations & next steps

1. Fix name inconsistencies: ensure all callers use the exact helper name `removeTrailingSpace` (or rename the helper to `removeTrailingSpaces` and update callers) to avoid compile errors.
2. Implement the `hexconv` and other missing converters in `hexconv.go`.
3. Add unit tests for each transformation function (happy paths and edge cases). Place tests under `pkg` using Go's testing framework.
4. Add argument validation and a small CLI usage message in `cmd/main.go`.
5. Consider exporting a stable API from `pkg` (e.g. `ProcessText(s string) string`) that runs all pipeline stages in a defined order.
6. Add CI (GitHub Actions) to run `go build` and `go test` on commits.

## Quick diagram (text)

[Input file] -> cmd/main.go (read) -> piscine.Process pipeline (Upper -> Lower -> other funcs) -> stdout


---
Generated from `docs/analysis.md` and source layout. If you want, I can also:
- Add a sequence diagram or ASCII art flowchart.
- Implement the name-fix (make helper name consistent) and run `go build` to verify.
- Add a minimal test harness for one transformation.
