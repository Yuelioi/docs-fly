---
title: diagonalizesymmetric
order: 18
---
| Since | 17.0 |
| --- | --- |

`matrix2  diagonalizesymmetric(matrix2 symmat, vector2 &diag)`

`matrix3  diagonalizesymmetric(matrix3 symmat, vector &diag)`

`matrix  diagonalizesymmetric(matrix symmat, vector4 &diag)`

[Diagonalize](http://en.wikipedia.org/wiki/Diagonalizable_matrix) a [symmetric matrix](http://en.wikipedia.org/wiki/Symmetric_matrix).

Returns the orthogonal matrix which, combined with the diagonal matrix
implicit in the second argument, will form the original symmetric matrix.

This can be useful for analyzing the result of summing a series of
outerproduct updates.

Show/hide arguments

`symmat`

The symmetric matrix to diagonalize.

`diag`

The diagonal elements of the diagonal matrix.
