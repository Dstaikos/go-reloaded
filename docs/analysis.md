We will create a tool that processes and “corrects” a user’s text file based on specific rules:

• When **(hex)** appears, it converts the hexadecimal number that comes *before* it into its decimal form.
Example: `"1E (hex) files were added"` → `"30 files were added"`

• When **(bin)** appears, it converts the binary number that comes *before* it into its decimal form.
Example: `"It has been 10 (bin) years"` → `"It has been 2 years"`

• When **(up)** appears, it changes the word that comes *before* it so that **all its letters become uppercase**.
Example: `"Ready, set, go (up)!"` → `"Ready, set, GO!"`

• When **(low)** appears, it changes the word that comes *before* it so that **all its letters become lowercase**.
Example: `"I should stop SHOUTING (low)"` → `"I should stop shouting"`

• When **(cap)** appears, it changes the word that comes *before* it so that **only its first letter is capitalized**.
Example: `"Welcome to the Brooklyn bridge (cap)"` → `"Welcome to the Brooklyn Bridge"`

• If **(up)**, **(low)**, or **(cap)** are followed by a comma and a number *x* inside the parentheses, the transformation applies to the **x previous words**.
Example: `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`

• When punctuation marks like **. , ! ? : ;** appear, they must:

* Be **attached** directly to the previous word (remove any spaces before them).

* Have **exactly one space** after them (add one if missing, or remove extra ones).
  Example: `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`

* If multiple punctuation marks appear together (e.g. `...` or `!?`), they should be treated as **a single unit**, and the same spacing rules apply.
  Example: `"I was thinking ... You were right"` → `"I was thinking... You were right"`

• Every **'** (apostrophe) must appear in **pairs** and be placed **directly around** the word(s) they enclose, with **no spaces** between the quotes and the words.
Example: `"I am exactly how they describe me: ' awesome '"` → `"I am exactly how they describe me: 'awesome'"`

* If there is more than one word inside the quotes, there should still be **no spaces** between the quotes and the inner text.
  Example: `"As Elton John said: ' I am the most well-known homosexual in the world '"` → `"As Elton John said: 'I am the most well-known homosexual in the world'"`

• When the single letter **“a” (or “A”)** appears by itself and is followed by a vowel (**a, e, i, o, u**) or **h**, it should change to **“an” (or “An”)**.
Example: `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`
