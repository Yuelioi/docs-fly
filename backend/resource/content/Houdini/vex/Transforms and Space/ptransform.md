---
title: ptransform
order: 20
---
`vector  ptransform(vector vec, matrix transform)`

`vector4  ptransform(vector4 vec, matrix transform)`

Transforms the vector using the given transform matrix.

`vector  ptransform(string tospace, vector vec)`

`vector4  ptransform(string tospace, vector4 vec)`

Transforms from `"space:current"`.

`vector  ptransform(string fromspace, string tospace, vector vec)`

`vector4  ptransform(string fromspace, string tospace, vector4 vec)`

Transforms the vector from `fromspace` into `tospace`.

Show/hide arguments

`fromspace`, `tospace`

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

- [ptransform](ptransform.html "Transforms a vector from one space to another.") interprets the vector as a position.
- [vtransform](vtransform.html "Transforms a directional vector.") interprets the vector as a direction
  vector, and so doesn’t apply the translations from the matrix.
- [ntransform](ntransform.html "Transforms a normal vector.") interprets the vector as a normal vector,
  and so multiplies by the inverse transpose of the matrix (ignoring the
  translations).
  Examples

## examples

The version with only a tospace argument assumes fromspace is
`"space:current"`. For example:

```vex
Pworld = ptransform("space:world", P);

```

…is equivalent to:

```vex
Pworld = ptransform("space:current", "space:world", P);

```

Transform a vector from its current space to object space:

```vex
ospace = ptransform("space:object", P)

```

Transform a vector from object space to mantra’s natural coordinate
space (“camera” space):

```vex
ospace = ptransform("space:object", "space:current", P)

```
