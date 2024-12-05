---
title: eigenvalues
order: 23
---
`void  eigenvalues(int &nroot, matrix3 mat, vector &real, vector &imaginary)`

Computes the [eigenvalues](http://en.wikipedia.org/wiki/Eigenvalues_and_eigenvectors) of a 3×3 matrix.

Show/hide arguments

`nroot`

The function overwrites this variable with the number of real roots.

`mat`

The matrix to compute the eigenvalues for.

`real`, `imaginary`

The components of these two vectors are overwritten with corresponding pairs of real and imaginary parts of each eigenvalue.
For example, `real[0]` and `imaginary[0]` contain the real and imaginary parts of the first eigenvalue.
