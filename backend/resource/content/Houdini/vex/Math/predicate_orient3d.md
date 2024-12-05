---
title: predicate_orient3d
order: 57
---
`float  predicate_orient3d(vector a, vector b, vector c, vector d)`

Given 3 points `a`, `b` and `c` in space, return a negative value if `d` is behind the
plane defined by the triangle `abc` (with right hand rule winding order), a positive value if its in front of
the plane, and zero if points `a`, `b`, `c` and `d` are coplanar.

More precisely, this function computes the determinant of the matrix:

```vex
[a_x a_y a_z 1; b_x b_y b_z 1; c_x c_y c_z 1; d_x d_y d_z 1]
```

…with a guaranteed correct sign.
