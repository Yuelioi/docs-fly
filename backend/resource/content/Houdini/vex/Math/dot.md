---
title: dot
order: 22
---
`float  dot(vector2 a, vector2 b)`

`float  dot(vector a, vector b)`

`float  dot(vector4 a, vector4 b)`

`float  dot(vector a, vector4 b)`

`float  dot(vector4 a, vector b)`

`float  dot(matrix2 a, matrix2 b)`

`float  dot(matrix3 a, matrix3 b)`

`float  dot(matrix a, matrix b)`

Returns the [dot product](http://en.wikipedia.org/wiki/Dot_product) of the arguments.

When dotting a `vector` with a `vector4`, only the first three
components are used.

`float  dot(<type>a[], <type>b[])`

`int  dot(int a[], int b[])`

Returns the sum of dot products i.e. `dot(a, b) = dot(a[0], b[0]) + ... + dot(a[n-1], b[n-1])` where `n = min(len(a), len(b))`.
