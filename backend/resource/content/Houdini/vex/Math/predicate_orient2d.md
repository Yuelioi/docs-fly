---
title: predicate_orient2d
order: 56
---
`float  predicate_orient2d(vector2 a, vector2 b, vector2 c)`

Given 2 points `a` and `b` in the plane, return a positive value if `c` is on the left
of the segment `(a,b)`, a negative value if it is on the right of the segment, and
zero if `a`, `b` and `c` are colinear.

More precisely, this function computes the determinant of the matrix:

```vex
[a_x a_y 1; b_x b_y 1; c_x c_y 1]
```

…with a guaranteed correct sign.
