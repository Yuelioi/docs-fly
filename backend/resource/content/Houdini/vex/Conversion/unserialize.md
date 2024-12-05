---
title: unserialize
order: 13
---
`<vector>[] unserialize(float values[])`

`<matrix>[] unserialize(float values[])`

The inverse operation to [serialize](serialize.html "Flattens an array of vector or matrix types into an array of floats."). This operation takes an array of float values
and creates a new array of vectors or floats by taking each float and assigning it to the
next component of the vector or matrix in the output array. For example:

Examples

## examples

```vex
vector v[]
float  f[] = { 1, 2, 3,  7, 8, 9 };

v = vector(unserialize(f));
// Now v has a length of 2 and contains { {1,2,3}, {7,8,9} }

```
