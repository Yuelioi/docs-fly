---
title: opdigits
order: 21
---
`int  opdigits(string str)`

`int  opdigits()`

Returns the integer value of the last sequence of digits in the input string.

If no argument is passed, the code is equivalent to

```vex
string dir, name;
splitpath(opfullpath("."), dir, name);
return opdigits(name);

```

Examples

## examples

- `opdigits("/obj/geo34/box21")` - returns 21
- `opdigits("/obj/geo34/box")` - returns 34
- `opdigits("/obj/geo34/box2.1")` - returns 1 (“.” is not a digit)
