---
title: polardecomp
order: 16
---
`matrix3  polardecomp(matrix3 transform)`

Computes the stretch matrix (S) and the orthogonal matrix (Q) such that `M = S*Q`.
This is very useful for shape matching or blending of transforms.

Show/hide arguments

`transform`

The matrix (M) to undergo polar decomposition.

Returns

'Q', the orthogonal matrix that best matches the given transform.

`void  polardecomp(matrix3 transform, matrix3 &rot, matrix3 &stretch, int check_determinant=1)`

Show/hide arguments

`&rot`

Returns the orthogonal matrix of the polar decomposition.

`&stretch`

Returns the stretch matrix of the polar decomposition.

`check_determinant`

Whether or not to check if there is a negative determinant (scale). If there is and this is not set to 0, the orthogonal and scale matrices will be negated.
