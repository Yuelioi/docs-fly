---
title: wireblinn
order: 80
---
`bsdf  wireblinn(vector tangent, float exponent, ...)`

![](../../images/rendering/wireblinn.png)
Blinn function defined around a tangent vector. You can use this to produce the average specular illumination for thin wire-like primitives such as hair.

- `tangent` – tangent vector along the hair.
- `exponent` – blinn exponent.
