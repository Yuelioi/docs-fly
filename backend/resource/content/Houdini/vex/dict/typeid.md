---
title: typeid
order: 4
---
| Since | 20.0 |
| --- | --- |

`int  typeid(<type>value)`

`int  typeid(<type>value[])`

Returns a numeric code identifying the value’s type.

`int  typeid(dict dictionary, string key)`

Returns a numeric code identifying the type of the key’s value in the dictionary, or -1 if the key does not exist.
This can be compared against the `typeid()` of a specific VEX data type.

Examples

## examples

```vex
// Check if the value for "foo" is a matrix.
int type = typeid(d, "foo");
if (type == typeid(matrix())
{
    matrix m = d["foo"];
}

```
