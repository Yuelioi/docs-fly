---
title: translucent
order: 78
---
`bsdf  translucent(vector nml, vector ng, ...)`

Returns a diffuse BSDF for the transmission direction. This can
be used as a cheap alternative to subsurface scattering for thin surfaces,
to allow illumination to pass from one side of the object to the other
while also broadly diffusing the illumination like diffuse.
