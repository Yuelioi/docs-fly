---
title: eulertoquaternion
order: 5
---
[Unit quaternions are used to represent rotations](http://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation). This VEX function creates a unit quaternion that represents a rotation.

`vector4  eulertoquaternion(vector rotations, int order)`

Creates a vector4 representing a unit quaternion from a vector representing Euler rotations in X, Y, and Z.

The angles are in radians. Use the `radians()` function to convert degrees into radians.

Show/hide arguments

`order`

One of the rotation order constants listed below, which can be imported from `$HFS/houdini/vex/include/math.h`.

| Constant name | Rotation Order |
| --- | --- |
| XFORM_XYZ | Rotate order X, Y, Z |
| XFORM_XZY | Rotate order X, Z, Y |
| XFORM_YXZ | Rotate order Y, X, Z |
| XFORM_YZX | Rotate order Y, Z, X |
| XFORM_ZXY | Rotate order Z, X, Y |
| XFORM_ZYX | Rotate order Z, Y, X |
