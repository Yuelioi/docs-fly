---
title: ocio_transform
order: 7
---
`vector  ocio_transform(string dest, vector clr)`

`vector  ocio_transform(string src, string dest, vector clr)`

`vector  ocio_transform(string src, string dest, string looks, vector clr)`

Transform a three-component color into a new color space.

`vector4  ocio_transform(string dest, vector4 clr)`

`vector4  ocio_transform(string src, string dest, vector4 clr)`

`vector4  ocio_transform(string src, string dest, string looks, vector4 clr)`

Transform a four-component color into a new color space.

Show/hide arguments

`src`

The name of the color space to transform from. If this is not given, the function assumes the space assigned to `"data"`.

`dest`

The name of the color space to transform to.

`looks`

A comma separated list of color gradings (also known as “looks”).

`clr`

The color to transform.
