---
title: abs
order: 4
---
`int  abs(int n)`

`float  abs(float n)`

`<vector> abs(<vector>v)`

Returns the absolute (positive) equivalent of the number. For vectors, this is done per-component.

Examples

## examples

Scalar example

```vex
if (abs(n) > 1) {
    // n is greater than 1 or less than -1
}

```

Vector example

```vex
vector v = {1.0, -0.5, 1.1}
if (abs(v) > 1.0) {
    // vector is greater than unit scale
}

```
