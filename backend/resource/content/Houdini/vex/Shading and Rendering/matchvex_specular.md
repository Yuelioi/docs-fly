---
title: matchvex_specular
order: 53
---
`bsdf  matchvex_specular(float exponent, ...)`

`bsdf  matchvex_specular(vector nml, float exponent, ...)`

![](../../images/rendering/matchvex_specular.png)
The BSDF produced by [specular](specular.html "Returns a specular BSDF or computes specular shading.") is not the same as the traditional VEX `specular()` output. Use this function to produce a closer approximate match to the traditional VEX `specular()`.
