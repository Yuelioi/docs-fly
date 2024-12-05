---
title: blackbody
order: 1
---
`vector  blackbody(float temperature, float luminance)`

Given a temperature, in Kelvin, and a luminance value, computes the
color of an incandescent black body as CIE XYZ tristimulus values.

The computation uses a fast approximation, which is valid for
temperature values between 1666K and 25000K. Values outside of this
range are clamped to the nearest valid in-range value.

The returned value can be converted to linear sRGB values using
the [xyztorgb](xyztorgb.html "Convert CIE XYZ tristimulus values to a linear sRGB triplet.") function.
