---
title: findattribval
order: 20
---
`int  findattribval(<geometry>geometry, string attribclass, string attribute_name, int|stringvalue, int which=0)`

`int [] findattribval(<geometry>geometry, string attribclass, string attribute_name, int|stringvalue)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

`attribute_name`

The name of the attribute to read.

`value`

The value to look for in the attribute.

`which`

If multiple elements have the given value in the attribute, this controls which match to return.

If you're searching for multiple elements with the same attribute value, you can use [findattribvalcount](findattribvalcount.html "Returns number of elements where an integer or string attribute has a certain value.") to get the total number of matches and then iterate through them by increasing the `which` argument to this function in a loop. See the examples below.
The array signature can also be used to return a list of all of the matching elements.

Returns

The number of the first point/primitive/vertex where the named attribute matches the given `value`. Returns `-1` if no element has the given attribute value.
The array signature returns the numbers of all of the points/primitives/vertices where the named attribute matches the given `value`.

Tip

The most common use cases (finding an point/primitive by its `name` or `id` attribute) have easier-to-use dedicated wrapper functions: [nametopoint](nametopoint.html "Finds a point by its name attribute."), [nametoprim](nametoprim.html "Finds a primitive by its name attribute."), [idtopoint](idtopoint.html "Finds a point by its id attribute."), and [idtoprim](idtoprim.html "Finds a primitive by its id attribute.").

- You can only search for integer or string values.

Examples

## examples

Find the primitive with `@id` == 10

```vex
int prim_num = findattribval(0, "prim", "id", 10);
// Note: you can use idtoprim(0, 10) instead

```

Find all points with `@age` == 10

```vex
for (int point_num : findattribval(0, "point", "age", 10))
{
    // ...do something with the point...
}

```

Find all points with `@age` == 10, using [findattribvalcount](findattribvalcount.html "Returns number of elements where an integer or string attribute has a certain value.").

```vex
int count = findattribvalcount(0, "point", "age", 10);
int point_num;
for (int i = 0; i < count; i++) {
    point_num = findattribval(0, "point", "age", 10, i);
    // ...do something with the point...
}

```
