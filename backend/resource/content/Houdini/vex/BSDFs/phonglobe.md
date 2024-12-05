---
title: phonglobe
order: 18
---
`bsdf  phonglobe(vector dir, float exponent, ...)`

`bsdf  phonglobe(vector nml, vector dir, float exponent, ...)`

`bsdf  phonglobe(vector dir, float exponentx, float exponenty, vector framex, vector framey, ...)`

`bsdf  phonglobe(vector nml, vector dir, float exponentx, float exponenty, vector framex, vector framey, ...)`

![](../../images/rendering/phonglobe.png)
A phong (blurred) reflection along a given direction vector. This will produce the same result as `phong()` when the direction vector is the reflection vector, but with this function you can also gather illumination from other directions (such as transmission).

It is possible to create anisotropic phong lobes by providing x and y exponents and tangent vectors.

Show/hide arguments

`dir`

the direction of specularity.

`nml`

optional normal to specify the hemisphere for reflection directions.

`exponent`

phong exponent.

`exponentx`

phong exponent along the `framex` vector.

`exponenty`

phong exponent along the `framey` vector.

`framex`

highlight X direction

`framey`

highlight Y direction
