---
title: solvecurve
order: 27
---
| Since | 17.5 |
| --- | --- |

`vector [] solvecurve(float lengths[], int closed, int orienttonormal, vector tangent, vector points[], vector normals[])`

`matrix3 [] solvecurve(float lengths[], int closed, int orienttonormal, vector tangent, vector points[], vector normals[])`

`vector [] solvecurve(float &outlength, vector &outpos, float lengths[], int closed, int orienttonormal, int normalmode, vector tangent, vector points[], vector normals[])`

`matrix3 [] solvecurve(float &outlength, vector &outpos, float lengths[], int closed, int orienttonormal, int normalmode, vector tangent, vector points[], vector normals[])`

`vector [] solvecurve(float &outlength, vector &outpos, float lengths[], int closed, int orienttonormal, int normalmode, vector tangent, vector points[], vector normals[], float twists[], float initialtwists[], int fmt, int order, float lod)`

`matrix3 [] solvecurve(float &outlength, vector &outpos, float lengths[], int closed, int orienttonormal, int normalmode, vector tangent, vector points[], vector normals[], float twists[], float initialtwists[], int fmt, int order, float lod)`

`vector [] solvecurve(string op, float lengths[], int closed, int orienttonormal, vector tangent, int normalcalcmethod, matrix relmat)`

`matrix3 [] solvecurve(string op, float lengths[], int closed, int orienttonormal, vector tangent, int normalcalcmethod, matrix relmat)`

`vector [] solvecurve(string op, float lengths[], int closed, int orienttonormal, vector tangent, int normalcalcmethod, matrix relmat, int primnum, float lod)`

`matrix3 [] solvecurve(string op, float lengths[], int closed, int orienttonormal, vector tangent, int normalcalcmethod, matrix relmat, int primnum, float lod)`

Returns a array of vectors or matrix3 representing local bone rotations. Angles are in degrees.

Show/hide arguments

`op`

The SOP path to a curve to evaluate.

`outlength`

Return the length of the curve where the solution ends. This is different than the sum of all lengths array.

`outpos`

Return the computed position of the last joint.

`lengths`

The lengths of all the bones to solve.

`closed`

Close the curve.

`orienttonormal`

Use the normals from the curve to orient the bones.

`normalmode`

Define how the normals/twists are computed from the control points.

Use the constants defined in `$HH/vex/include/math.h`.

`tangent`

A tangent vector to orient the end tip of the curve.

`points`

An array of vectors to use as points to define the curve.

`normals`

An array of vectors to use as normals to define the curve.

`twists`

An optional array of floats to use as twist angles in degree to define the curve.

`initialtwists`

An optional array of floats to use as initial twist angles in degree define the curve.

`normalcalcmethod`

A normal calculation method when evaluating using a SOP. (0 default, 1 none, 2 interpolate with quaternions, 3 interpolate with twist angles in 0,180 range, 4 interpolate with twist angles in any range.)

`relmat`

A relative matrix used to transform the points, normals and tangent relative to the origin.
This is normally the invert matrix of the root of the chain.

`fmt`

The curve type to create.
Use constants defined in `$HH/vex/include/math.h`, or 0 to create a polygon curve, 1 to create a bezier curve or 2 to create a NURBS curve.

`order`

The curve order for NURBS or Bezier curves. This is ignored for polygon curves.
