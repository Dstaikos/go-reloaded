# Architecture Comparison: Pipeline vs FSM

## Pipeline Architecture

The Pipeline architecture divides text processing into sequential stages, where each stage takes the result of the previous one and modifies it.

### Core Functions
- Detecting and converting `(hex)` and `(bin)` numbers
- Applying text transformations: `(up)`, `(low)`, `(cap)`
- Correcting punctuation marks
- Replacing "a" with "an" where necessary

### Processing Flow
```
Initial text → Number conversion → Word formatting → Punctuation correction → Final text
```

#### ✅ Advantages
- **Clear separation of concerns** - Each stage has a single responsibility
- **Easy testing** - Each stage can be tested independently
- **High readability** - Simple, auditable code structure
- **Maintainable** - Easy to add or modify individual stages

#### ❌ Disadvantages
- **Context dependency** - Some stages need access to previous words (e.g., `(up, 3)`)
- **Multiple passes** - Text may be processed multiple times

## FSM (Finite State Machine) Architecture

The FSM processes text token by token, changing states based on encountered patterns.

### State Examples
- **Normal**: Processing regular words
- **Marker**: Handling `(up)`, `(low)`, or `(cap)` commands
- **Quote**: Processing text within single quotes
- **Number**: Converting numbers before `(hex)` or `(bin)`

#### ✅ Advantages
- **Single pass efficiency** - Processes text in one iteration
- **Precise context control** - Easy to apply transformations like `(up, 3)`
- **Optimal parsing** - Natural fit for sequential text processing

#### ❌ Disadvantages
- **Complex implementation** - More intricate state management
- **Testing challenges** - Interconnected components harder to isolate
- **Higher cognitive load** - Requires understanding state transitions and interactions

## Final Architecture Decision

**Chosen: Pipeline Architecture**

### Rationale
For this text processing application, **modularity and maintainability are prioritized over performance optimization**. The Pipeline approach provides:

- **Modular design** - Each transformation can be developed and tested independently
- **Clear data flow** - Easy to trace how text is transformed at each stage
- **Extensibility** - New transformation rules can be added without affecting existing stages
- **Debugging simplicity** - Issues can be isolated to specific pipeline stages

While an FSM offers better performance through single-pass processing, the Pipeline architecture better aligns with the project's emphasis on code maintainability and feature extensibility.