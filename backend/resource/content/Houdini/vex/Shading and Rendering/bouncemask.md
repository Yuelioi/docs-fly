---
title: bouncemask
order: 4
---
`int  bouncemask(string labels)`

Show/hide arguments

`labels`

A label or space-separated list of labels.

Returns

A bitmask that matches any of the labels.

Mantra tags different types of rays using shading component *labels*, such as “diffuse”, “reflect”, “refract”, “volume”, and “sss”. A custom BSDF can also specify its own labels in addition to existing ones (see [cvex_bsdf](cvex_bsdf.html "Creates a bsdf object from two CVEX shader strings.") for more information).

Some VEX functions take or return a *component bitmask*, which specifies a combination of one or more of these labels using the bits of an integer.

To get the bit value associated with a label, use [bouncemask](bouncemask.html), for example `bouncemask("diffuse")`. To get a mask that matches multiple labels, use a space-separated list:

```vex
reflect_or_refract = bouncemask("reflect refract")

```

To construct a bitmask that matches all labels, use `bouncemask("all")`. To match no labels, use `0`.

When you get a bitmask as a return value, you can check if it matches a certain label using `&`. For example:

```vex
mask = getbounces(mybsdf)
if (mask & bouncemask("reflect")) {
    ...
}

```

(As an alternative to basic uses of `bouncemask()`, you can `#import "pbr.h"` and work with the constants `PBR_DIFFUSE_MASK`, `PBR_REFLECT_MASK`, `PBR_REFRACT_MASK`, `PBR_VOLUME_MASK`, `PBR_SSS_MASK`, as well as `PBR_ALL_MASK` and `PBR_NO_MASK`. You can combine the constants using `|`, for example `reflect_or_refract = PBR_REFLECT_MASK | PBR_REFRACT_MASK`.)
