---
title: prerotate
order: 17
---
| Since | 17.5 |
| --- | --- |

`void  prerotate(matrix3 &m, float amount, vector axis)`

`void  prerotate(matrix &m, float amount, vector axis)`

`void  prerotate(matrix3 &m, vector angles, int xyz)`

`void  prerotate(matrix &m, vector angles, int xyz)`

`void  prerotate(matrix3 &m, float angle, int axis)`

`void  prerotate(matrix &m, float angle, int axis)`

Applies a prerotation to the given matrix. The angles must be given in
radians and the axis must be normalized. The xyz argument is the rotate order.
The axis can also be given as an integer where XAXIS=1, YAXIS=2 and ZAXIS=4.
