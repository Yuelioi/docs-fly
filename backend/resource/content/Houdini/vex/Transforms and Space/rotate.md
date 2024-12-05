---
title: rotate
order: 21
---
`void  rotate(matrix2 &m, float amount)`

`void  rotate(matrix3 &m, float amount, vector axis)`

`void  rotate(matrix &m, float amount, vector axis)`

`void  rotate(matrix3 &m, vector angles, int xyz)`

`void  rotate(matrix &m, vector angles, int xyz)`

`void  rotate(matrix3 &m, float angle, int axis)`

`void  rotate(matrix &m, float angle, int axis)`

Applies a rotation to the given matrix. The angles must be given in
radians and the axis must be normalized. The xyz argument is the rotate order.
The axis can also be given as an integer where XAXIS=1, YAXIS=2 and ZAXIS=4.
