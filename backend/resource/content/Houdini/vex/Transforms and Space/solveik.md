---
title: solveik
order: 29
---
| Since | 17.5 |
| --- | --- |

`vector [] solveik(float lengths[], vector targetpos, vector twistpos, float twist, int twistflag, float dampen, int resiststraight, float trackingthres, matrix relmat, vector constraints[])`

`matrix3 [] solveik(float lengths[], vector targetpos, vector twistpos, float twist, int twistflag, float dampen, int resiststraight, float trackingthres, matrix relmat, vector constraints[])`

Returns a array of local bone rotations in degrees.

Show/hide arguments

`lengths`

The lengths of all the bones to solve.

`targetpos`

A target position in worldspace.

`twistpos`

A twist affect position in worldspace.

`twist`

A twist angle in degrees.

`twistflag`

Apply twist using the twist affector or not.

`dampen`

Dampening factor for the whole chain.

`resiststraight`

Resist straightening.

`trackingthres`

Tracking threshold.

`relmat`

A relative matrix used to transform the target and twist positions relative to the origin.
This is normally the invert matrix of the root of the chain.

`constraints`

This is a vector array used to define the per bone restangles, dampening, minangles, maxangles, min damp, max damp, and rolloff.
If the array is empty, same default values present in the bone objects are used.
If the array size is equal to the number of input bones, the rest angles are defined.
If the array size is equal to 2 times the number of input bones, the rest angles and the dampening are defined.
If the array size is equal to 3 times the number of input bones, the rest angles, the dampening, the min/max angles are defined. Min/Max angles share the same values.
If the array size is equal to 4 times the number of input bones, the rest angles, the dampening, the min/max angles are defined. Min/Max angles have different values.
If the array size is equal to 5 times the number of input bones, the rest angles, the dampening, the min/max angles and damp angles are defined.
If the array size is equal to 6 times the number of input bones, the rest angles, the dampening, the min/max angles, min/max damp angles are defined.
If the array size is equal to 7 times the number of input bones, the rest angles, the dampening, the min/max angles, min/max damp angles, and rolloff are defined.
