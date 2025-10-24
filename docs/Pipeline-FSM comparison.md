<h2>Comparison Between Pipeline and FSM Architectures</h2>

<h3>Pipeline</h3>

The Pipeline architecture divides text processing into sequential stages.
Each stage takes the result of the previous one and modifies it.
This allows having separate functions such as:

• Detecting and converting (hex) and (bin)

• Applying (up), (low), (cap)

• Correcting punctuation marks

• Replacing a with an where necessary

Example flow:

Initial text → Number conversion → Word formatting → Punctuation correction → Final text

<h4>Advantages:</h4>

• Clear separation of processes

• Easy testing, since each stage can be checked independently

• Readable and simple code (ideal for auditing)

<h4>Disadvantages:</h4>

• Some stages require access to previous words (like (up, 3)), so some extra logic is needed

• The code may pass over the same text multiple times

<h3>FSM (Finite State Machine)</h3>

The FSM works differently: the program reads the text step by step (token by token) and changes its state depending on what it encounters.

Example states:

• Normal: reading regular words

• Marker: when encountering “(up)”, “(low)”, or “(cap)”

• Quote: when inside ' '

• Number: when encountering a number before (hex) or (bin)

The program has one main loop that checks the current state and decides what to do with the next word.

<h4>Advantages:</h4>

• Everything happens in a single pass through the text (efficient)

• Very precise context control (e.g., easily apply (up, 3) to previous words)

• Good logic for parsing

<h4>Disadvantages:</h4>

• More complex code

• Harder to test since everything is interconnected

• Lower readability for auditors

<h3>Architecture Choice</h3>

I chose to use the Pipeline approach.

The program will be reviewed by other auditors, so clarity and simplicity of the code are more important than optimization.

With the Pipeline, I can write small, clear functions that each apply one specific rule and explain the program flow step by step.

An FSM would be more efficient but harder to understand or audit.