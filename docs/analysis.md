# Go-Reloaded: Text Processing Tool Analysis

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Command Specifications](#command-specifications)
- [Word Definition](#word-definition)

## Installation

### Clone the Repository
```bash
git clone https://platform.zone01.gr/git/dstaikos/go-reloaded.git
```

### Navigate to Project
```bash
cd go-reloaded/cmd
```

## Usage

Place your input text file in the `cmd` folder and run:

```bash
go run main.go <input.txt> <output.txt>
```

## Features

This tool processes and corrects text files based on the following transformation rules:

### Number Conversions

#### Hexadecimal to Decimal
- **Trigger:** `(hex)` after a hexadecimal number
- **Example:** `"1E (hex) files were added"` → `"30 files were added"`

#### Binary to Decimal
- **Trigger:** `(bin)` after a binary number
- **Example:** `"It has been 10 (bin) years"` → `"It has been 2 years"`

### Text Transformations

#### Uppercase
- **Trigger:** `(up)` after a word
- **Example:** `"Ready, set, go (up)!"` → `"Ready, set, GO!"`

#### Lowercase
- **Trigger:** `(low)` after a word
- **Example:** `"I should stop SHOUTING (low)"` → `"I should stop shouting"`

#### Capitalize
- **Trigger:** `(cap)` after a word
- **Example:** `"Welcome to the Brooklyn bridge (cap)"` → `"Welcome to the Brooklyn Bridge"`

#### Multiple Word Transformations
- **Trigger:** `(up, x)`, `(low, x)`, or `(cap, x)` where x is the number of previous words
- **Example:** `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`

### Punctuation Correction

Punctuation marks (`. , ! ? : ;`) are automatically corrected to:
- **Attach directly** to the previous word (no spaces before)
- **Have exactly one space** after them

**Examples:**
- `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`
- `"I was thinking ... You were right"` → `"I was thinking... You were right"`

### Quote Formatting

Single quotes (`'`) are automatically paired and formatted:
- **No spaces** between quotes and enclosed text
- **Proper pairing** of opening and closing quotes

**Examples:**
- `"I am exactly how they describe me: ' awesome '"` → `"I am exactly how they describe me: 'awesome'"`
- `"As Elton John said: ' I am the most well-known homosexual in the world '"` → `"As Elton John said: 'I am the most well-known homosexual in the world'"`

### Article Correction

The article "a" (or "A") is automatically changed to "an" (or "An") when followed by:
- **Vowels:** a, e, i, o, u
- **The letter h**

**Example:** `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`

## Command Specifications

### Format Requirements

⚠️ **Important:** Commands must follow exact formatting:

- **Valid:** `(up, 2)`
- **Invalid:** `( up,2 )`, `(up , 2)` - treated as normal text

### Number Constraints

- Numbers must be **non-negative integers** (≥ 0)
- **Negative numbers** or **non-integers** are treated as normal text
- If requested word count exceeds available words, transformation applies to all available words

### Processing Behavior

When word count exceeds available words, the modifier applies to every available word and stops.

## Word Definition

For transformation purposes, a **word** is defined as any sequence of alphanumeric characters separated by whitespace:

- **Individual letters:** `a`, `Z`
- **Individual numbers:** `5`, `42`
- **Mixed sequences:** `abc123`, `test2`
- **Alphabetic words:** `hello`, `WORLD`
- **Numeric sequences:** `12345`

### Modifier Consumption

When applying counted modifiers (e.g., `(up, 3)`):
- Consumes exactly the specified number of words **backwards** from modifier position
- Words with no transformable letters are counted but remain unchanged
- Only letters are transformed; numbers and special characters remain intact

**Example:** 
- Input: `"i love you 2 (up, 3)"`
- Consumed words: `"2"`, `"you"`, `"love"`
- Output: `"i LOVE YOU 2"`