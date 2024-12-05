---
title: setprimintrinsic
order: 72
---
`int  setprimintrinsic(int geohandle, string name, int prim_num, <type>value, string mode="set")`

`int  setprimintrinsic(int geohandle, string name, int prim_num, <type>value[], string mode="set")`

Despite the name, some “intrinsic” attributes on primitives are writeable.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`name`

The name of the intrinsic to set.

`prim_num`

The number of the primitive to change the value on.

`mode`

(Optional) if given, this controls how the function modifies any existing value in the attribute.

| `"set"` | Overwrite the attribute with the given value. |
| --- | --- |
| `"add"` | Add to the attribute the value. |
| `"min"`, `"minimum"` | Set the attribute to the minimum of itself and the value. |
| `"max"`, `"maximum"` | Set the attribute to the maximum of itself and the value. |
| `"mult"`, `"multiply"` | Multiply the attribute by the value. For matrices, this will do matrix multiplication. For vectors, component-wise. |
| `"toggle"` | Toggles the attribute, independent of the source value. Useful for toggling group membership. |
| `"append"` | Valid for string, dict, and array attributes. For strings and  arrays, appends the source value to the end of the original  value. For dictionaries, updates the original dictionary with  the source dictionary, replacing any matching keys. |
