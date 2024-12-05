---
title: quaternion
order: 65
---
[Unit quaternions are used to represent rotations](http://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation). This VEX function creates a unit quaternion that represents a rotation.

`vector4  quaternion(matrix3 rotations)`

Creates a vector4 representing a unit quaternion from a 3×3 rotational matrix.

`vector4  quaternion(float angle, vector axis)`

Creates a vector4 representing a unit quaternion from an angle and a normalized axis. The angle is specified in radians.

`vector4  quaternion(vector angleaxis)`

Creates a vector4 representing a unit quaternion from a combined angle/axis. This is the normalized rotation axis multiplied by the rotation angle in radians.

There used to be a fourth form that took a rotation vector. It has been renamed to `eulertoquaternion` and now takes radians.

For more information, see [Data types](../lang.html#data-types) and [Dot operator](../lang.html#dot-operator).
