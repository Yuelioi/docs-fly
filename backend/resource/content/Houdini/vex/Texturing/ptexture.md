---
title: ptexture
order: 8
---
`<type> ptexture(string map, int face_id, ...)`

`<type> ptexture(string map, int face_id, float s, float t, ...)`

This function is deprecated since ptex support has been integrated into the `texture()` function.
Optional arguments

## optional-arguments

| Keyword | Values |
| --- | --- |
|  |`channel`
td>>
An integer value indicating which channel of the ptex image to use.

|  |`filter`
td>>

| `filtersharp` | A floating point value indicating the filter sharpness. This is only valid for the bicubic filter. The range is 0-1 (the default is 1.0). |
| `lerp` | An boolean value indicating whether to interpolate between Mip maps. The default is true. |
| `blur` | The texture blur value used for evaluation (default 0). |
| `width` | The texture width value used for evaluation (default 1). |
