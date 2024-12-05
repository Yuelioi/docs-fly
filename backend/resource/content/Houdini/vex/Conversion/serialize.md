---
title: serialize
order: 12
---
`float [] serialize(<vector>vectors[])`

`float [] serialize(<matrix>matrices[])`

These functions will serialize the arrays of tuple values.
That is, the values of the tuples are extracted one by one into a
flat floating point array.
Examples

## examples

```vex
vector v[] = { {1,2,3}, {7,8,9} }; // A vector[] of length 2
float  f[];

f = serialize(v);
// Now f[] has a length of 6 and equals { 1,2,3,7,8,9 }

```
