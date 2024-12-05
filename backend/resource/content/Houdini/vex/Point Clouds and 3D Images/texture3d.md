---
title: texture3d
order: 36
---
`<type> texture3d(string filename, string channel, vector P, ...)`

Returns the value of the 3d image at the position specified by P. If P
is outside of the bounding box of the image, the value returned will be
0\. If the channel specified contains more values than the return type
(i.e. a vector channel when a float return type is desired), the first
component of the vector will be returned. If the channel specified
contains fewer values than the return type, the missing components will
be filled with the last valid channel.

Texture files will be searched for in the path specified by the
`HOUDINI_TEXTURE_PATH` environment variable.

You can pass additional arguments to control the evaluation (see
[colormap](colormap.html "Looks up a (filtered) color from a texture file.")):

| `"filter"` | Specifies the filter for evaluation. |
| --- | --- |
| `"width"` | Specifies the filter width for evaluation. |
