---
title: fromNDC
order: 2
---
`vector  fromNDC(vector v)`

Transforms the vector from NDC space to the current space.

`vector  fromNDC(string space, vector v)`

Transforms the vector from NDC space to the named space.

Show/hide arguments

`space`

The possible values for the space arguments are:

| An object path | Use the object space of an object specified by a path string.   Tip  In some cases, such as point instancing, mantra may  automatically mangle object paths. You can generate an `.ifd`  file and look inside to try to find what mantra is calling  the object you want. |
| --- | --- |
| `"space:object"` | Object space of the *current* object. |
| `"space:light"` | Object space of the *current* light when executing a shadow or light shader. |
| `"space:world"` | Houdini world space. |
| `"space:camera"` | mantra camera space. |
| `"space:ndc"` | Normal Device Coordinate space. |
| `"space:lightndc"` | Normal Device Coordinate space for the *current* light when executing a shadow or light shader. |
| `"space:current"` | The current space the vector is in. |
