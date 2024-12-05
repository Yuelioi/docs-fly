---
title: svddecomp
order: 83
---
| Since | 18.0 |
| --- | --- |

`void  svddecomp(matrix2 input_M, matrix2 &output_U, vector2 &output_S, matrix2 &output_V)`

`void  svddecomp(matrix3 input_M, matrix3 &output_U, vector &output_S, matrix3 &output_V)`

`void  svddecomp(matrix input_M, matrix &output_U, vector4 &output_S, matrix &output_V)`

Computes the [singular value decomposition](http://en.wikipedia.org/wiki/Singular_value_decomposition) of a
given matrix. More precisely, computes `U`, `S`, `V` such that
`M = U*T*transpose(V)`, where `T` is the diagonal matrix constructed from `S`,
the vector of singular values.

`vector2  svddecomp(matrix2 input_M)`

`vector  svddecomp(matrix3 input_M)`

`vector4  svddecomp(matrix input_M)`

The second form of this function simply returns the vector of singular values.
