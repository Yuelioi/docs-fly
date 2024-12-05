---
title: length
order: 37
---
To get the length of a string, or number of items in an array, use [len](len.html "Returns the length of an array.").

`float  length(float f)`

Simply returns the given number.

`float  length(vector2 v)`

`float  length(vector v)`

`float  length(vector4 v)`

Returns the distance of the vector or vector4 from the origin.

If you want the squared length, using [length2](length2.html "Returns the squared distance of the vector or vector4.") is faster than squaring the result of this function.

Examples

## examples

```vex
length({1.0, 0, 0}) == 1.0;
length({1.0, 1.0, 0}) == 1.41421;

```
