---
title: ctransform
order: 2
---
`vector  ctransform(string fromspace, string tospace, vector clr)`

`vector  ctransform(string tospace, vector clr)`

If you don’t give a fromspace, assumes `"cspace:rgb"`.

Transforms color tuple clr from one color space to another.

The possible arguments for fromspace and tospace are
`"cspace:rgb"`, `"cspace:hsl"`, `"cspace:hsv"`, `"cspace:XYZ"`,
`"cspace:Lab"`, and `"cspace:tmi"`.
Notes

## notes

- The hue-based systems are normalized with the hue going from `0` to `1`. LAB and TMI are not normalized.
- For `"cspace:rgb"`, the primaries are assumed to be in linear NTSC space (gamma 1.0), using C reference white.
- C reference white is used for conversion when converting from XYZ to LAB space (and vice versa).
