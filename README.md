# Fidgetlang

This is mostly useful if you want to learn how to implement
a simple state-based Brainfuck interpreter.

Otherwise I advise to not use this

## Language Reference

Fidgetlang has the full brainfuck instruction set, it's simply
mapped to symbol sequences.

Like in brainfuck, unknown sequences are ignored.

Fidgetlang uses 256 32bit wide registers and supports up to 
4 billion instructions. (If your memory holds it)

| Instruction       | Brainfuck | Fidgetlang |
|-------------------|:---------:|:----------:|
| Increment Pointer |     >     |      ߷     |
| Decrement Pointer |     <     |     ߷߷     |
| Increment Value   |     +     |     ߷߷߷    |
| Decrement Value   |     -     |    ߷߷߷߷    |
| Print Value       |     .     |    ߷߷߷߷߷   |
| Read Value        |     ,     |   ߷߷߷߷߷߷   |
| Jump Ahead        |     [     |   ߷߷߷߷߷߷߷  |
| Jump Behind       |     ]     |  ߷߷߷߷߷߷߷߷߷ |

## When to Use it

Never

## License

Mozilla Public License 2.0