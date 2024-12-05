---
title: ocio_transformview
order: 8
---
`vector  ocio_transformview(string src, string display, string view, vector clr)`

Transform a three-component color into a new view.

`vector4  ocio_transformview(string src, string display, string view, vector4 clr)`

Transform a four-component color into a new view.

Show/hide arguments

`src`

The name of the color space to transform from.

`display`

The name of the Display containing the desired View.

`view`

The View to use as the destination to transform to.

`clr`

The color to transform.
