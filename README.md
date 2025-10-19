Text Completion & Auto-Correction Tool

📖 Description

This project is a text completion and correction tool written in Go.

It takes an input text file, applies a series of automatic corrections and transformations, and writes the modified text into an output file.

The purpose is to simulate a small-scale text editor capable of fixing grammar, formatting, and capitalization — automatically!


How it works

You run your program like this:  go run . input.txt output.txt

🗂️ The program will read the contents of input.txt, apply all transformation rules, and write the corrected version to output.txt.


Features & Rules


🔢 Number Conversions

(hex) ➡️ Converts the previous hexadecimal number to decimal

"1E (hex) files were added" → "30 files were added"

(bin) ➡️ Converts the previous binary number to decimal

"It has been 10 (bin) years" → "It has been 2 years"

🔤 Text Case Modifiers

(up) ➡️ Turns the previous word UPPERCASE

"Ready, set, go (up)!" → "Ready, set, GO!"

(low) ➡️ Turns the previous word lowercase

"STOP SHOUTING (low)" → "STOP shouting"

(cap) ➡️ Capitalizes the previous word

"welcome to the bridge (cap)" → "Welcome to the Bridge"



📦 You can also apply them to multiple words using a number:

"This is so exciting (up, 2)" → "This is SO EXCITING"



✍️ Punctuation Rules

Punctuation marks . , ! ? : ; should:

Stick to the previous word (no space before)

Be followed by one space (no double or missing spaces)

Example:

"I was sitting over there ,and then BAMM !!" → "I was sitting over there, and then BAMM!!"



💡 For grouped punctuations (like ... or !?), treat them as one unit:

"I was thinking ... You were right" → "I was thinking... You were right"



🗣️ Quotation Formatting

Single quotes ' ' should hug the inner words tightly, without spaces.

"I am ' awesome '" → "I am 'awesome'"

If there are multiple words inside, the quotes still stay tight:

"As Elton John said: ' I am the most well-known homosexual in the world '" → "As Elton John said: 'I am the most well-known homosexual in the world'"



🅰️ Article Correction

When "a" or "A" is followed by a vowel (a, e, i, o, u) or h, it becomes "an" or "An" respectively.

"A amazing rock" → "An amazing rock"

"a honest man" → "an honest man"
