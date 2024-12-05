---
title: Using VEX expressions
order: 5
---
| On this page | * [Overview](#overview) * [Why VEX for ad-hoc modifications?](#why) * [Syntax](#syntax) * [Gotchas](#gotchas) * [Accessing parameter values](#parameters) * [Accessing geometry attributes and info](#attributes)   + [Non-float attributes with known types](#known)  + [Accessing attributes on other inputs](#other_inputs) * [Setting geometry attributes](#setgeoattrs) * [Indexing variables](#indexing) * [Arrays](#arrays) * [Declaring attributes](#declare) * [Accessing globals](#accessing-globals) * [Creating geometry](#creating-geometry) * [Geometry traversal functions](#traverse) * [Accessing group membership](#groups) * [User-defined functions](#user-defined-functions) * [Includes](#includes) * [Tips](#tips) * [Troubleshooting error messages](#troubleshooting-error-messages) |
| --- | --- |

Overview

## overview

Several nodes in Houdini let you write short VEX expressions or snippets of VEX code. For example, the [![](../icons/SOP/attribwrangle.svg)Attrib Wrangle](../nodes/sop/attribwrangle.html "Runs a VEX snippet to modify attribute values.") geometry nodes; [![](../icons/DOP/geometrywrangle.svg)Geometry Wrangle](../nodes/dop/geometrywrangle.html "Runs a VEX snippet to modify attribute values."), and [![](../icons/DOP/gas.svg)Gas Field Wrangle](../nodes/dop/gasfieldwrangle.html "Runs CVEX on a set of fields.") dynamics nodes, and [particle dynamics nodes](../dopparticles/index.html "How to create particle simulations.").

These VEX expressions run on each element (point, particle, primitive, voxel, depending on the node type) passing through the node. The code can read the values of node parameters and geometry attributes, and set special variables to change values in the input geometry.

Tip
The [![](../icons/MISC/python.svg)Python SOP](../nodes/sop/python.html "Runs a Python snippet to modify the incoming geometry.") is similar but lets you edit geometry using a Python snippet.

Why VEX for ad-hoc modifications?

## why

For performance reasons, Houdini is moving toward doing ad-hoc geometry modifications with VEX operating on attributes, rather than HScript expressions operating on local variables and external channel references.

- Using VEX and attributes has major performance benefits over HScript expressions and local variables. It runs faster and automatically supports threading and parallel computation.
- Working directly on attributes instead of local variables actually has some ease-of-use advantages, since the naming of local variables could be inconsistent with the underlying attribute’s name, and inconsistent from node to node.
- In HScript expressions, getting the value of an attribute that didn’t already have a local variable mapping set up in the node was a chore (for example, `point(opinputpath(".",0), $PT, "my_var", 0)`). In VEX this is much easier: `v@my_var`. Since technical work in Houdini often revolves around attributes, this can actually make VEX expressions a lot simpler than the equivalent HScript expressions.
- Passing information down the network on attributes is inherently friendlier to parallel processing than using external references on later nodes to data on earlier nodes.
- Currently, VEX operations are supported inside [compiled SOP blocks](../model/compile.html "You can mark a chain of compilable SOPs as a compiled block which can execute much faster in parallel."), but HScript expressions using local variables cannot be compiled.
- VEX has gained equivalents of most HScript expression functions, and is easier to use for things like array and string processing, with conveniences such as Python-like array/string slicing and Python like dictionaries.

As users work on ever-larger and more complex geometry, threading and parallel processing become more and more important to get acceptable performance. This simple fact is the reason why VEX will only become more widely used to replace HScript expressions for ad-hoc geometry manipulation.

HScript will probably always be available for certain jobs where it’s handier than VEX. For geometry manipulation, however, wrangling and VEX/VOPs is the way forward, and it’s worthwhile to learn the new workflow.

Syntax

## syntax

A **VEX snippet** parameter lets you enter a snippet of [VEX code](index.html "VEX is a high-performance expression language used in many places in Houdini, such as writing shaders."). See the list of [VEX functions](functions/index.html).

VEX has a concept of “contexts”. Some functions are only available in certain contexts (for example, functions for accessing geometry information in the SOP context). A VEX snippet runs in the `CVEX` context.

Gotchas

## gotchas

- Each statement **must** end with a semicolon (`;`)!
- `//` and `/* ... */` can be used for comments.
- In VEX, trigonometry functions such as [sin](functions/sin.html "Returns the sine of the argument.") and [cos](functions/cos.html "Returns the cosine of the argument.") use radians, not degrees.
- Vector attributes are handled as `@v.x` rather than `$VX`. That is, you get one `@v` vector value of which you access the `x`, `y`, or `z` component using dot notation, rather than getting three separate variable `$VX`, `$VY`, and `$VZ`.
- [rand](functions/rand.html "Creates a random number between 0 and 1 from a seed.") produces vector noise when applied to a vector variable. This can be unexpected, like in the force example, where you may have expected all components of the force to be randomized equally by `@id`. Use a `float()` cast to force it scalar.
- [Setting geometry attributes](snippets.html#setgeoattrs) in different ways has different effects that can be confusing if you don’t know what’s going on. See [the explanation of how setting geometry attributes works](snippets.html#setgeoattrs) for more information.

Accessing parameter values

## parameters

In the snippet, you can read/write the value of a parameter on the node using the internal name of a parameter as a variable name.

To get the internal ID of a parameter, hover over the parameter name in the parameter editor.
The tooltip will show `Parameter: id`.

For example, in the [![](../icons/POP/color.svg)Particle Color DOP](../nodes/dop/popcolor.html "A POP node that colors particles."), the ID of the **Color** parameter is `color`.

```vex
color = @Cd;

```

Multi-component parameters are accessed as vectors. For example, the **Position** parameter has the internal name `t`:

```vex
// Set the translation
t = {0, 1, 0};

```

You can use the dot operator to access individual components of the parameter:

```vex
// Move one unit along the X axis
t.x = t.x + 1;

```

Note
To access the value of a user created parameter, use the [chv](functions/chv.html "Evaluates a channel or parameter, and return its value.") vex function.

Accessing geometry attributes and info

## attributes

In the snippet, you can read/write the value of an attribute using `@attribute_name`.
For example, to read or write to the `P` (position) attribute, use `@P` in the VEX code.

- Particle DOPs can access particle attributes but *can’t modify them*. Instead they affect the particles by varying the parameter values per-particle. See [writing particle VEX expressions](../dopparticles/vexpressions.html).
- In the Volume Wrangle node, you can use `@volume_name` to read or write to a volume.
- If you write to a `@attribute` in the VEX code and the attribute does not exist, Houdini will create it. (The Volume Wrangle node will *not* create new volumes this way.)
- Houdini provides some attribute-like variables you can use in the snippet. `@elemnum` contains the number of the current element being processed. `@numelem` contains the total number of elements in the geometry/list. See [indexing variables](snippets.html#indexing) below.
- Some nodes may provide additional attribute-like variables you can read to get computed information about the geometry. For example, in the [![](../icons/SOP/volumewrangle.svg)Volume Wrangle node](../nodes/sop/volumewrangle.html "Runs a VEX snippet to modify voxel values in a volume."), you can use `@center` to get the center of the volume.
- Houdini knows to cast [some commonly used attributes](snippets.html#known) using the appropriate VEX datatype.
  The table of [known attributes below](snippets.html#known) lists attributes that Houdini can automatically cast.
  Houdini assumes all other `@` references are **float** unless you manually specify a different type.
  To manually specify the VEX datatype for an attribute, add a character representing the type before the `@` sign.
  For example, to cast the `foo` attribute as a string, you would use `s@foo`.

Note
Automatic casting does **not** work if you use `@opinputn_name` to [access a different input](snippets.html#other_inputs). In that case, you must always specify the type.

The following table lists the available datatypes and the corresponding characters.

| VEX type | Syntax |
| --- | --- |
| `float` | `f@name` |
| `vector2` (2 floats) | `u@name` |
| `vector` (3 floats) | `v@name` |
| `vector4` (4 floats) | `p@name` |
| `int` | `i@name` |
| `matrix2` (2×2 floats) | `2@name` |
| `matrix3` (3×3 floats) | `3@name` |
| `matrix` (4×4 floats) | `4@name` |
| `string` | `s@name` |
| `dict` | `d@name` |

### Non-float attributes with known types

known
As a convenience, you don’t need to specify the type of the following commonly used attributes
(Houdini knows what type they should be).
Other attributes are assumed to be float unless you specify a type (see above).

For example, in a VEX snippet you can just type `@Cd` instead of having to type `v@Cd` to specify that it’s a vector.

Tip
See the [attributes page](../model/attributes.html "Describes how Houdini represents geometry using details, primitives, points, vertices, and attributes.") for information on commonly used attributes.

| VEX type | Attribute names |
| --- | --- |
| `vector` (3 floats) | [@P](../model/attributes.html#p),  [@accel](../model/attributes.html#accel),  [@Cd](../model/attributes.html#cd),  [@N](../model/attributes.html#n),  [@scale](../model/attributes.html#scale),  [@force](../model/attributes.html#force),  [@rest](../model/attributes.html#rest),  [@torque](../model/attributes.html#torque),  [@up](../model/attributes.html#up),  [@uv](../model/attributes.html#uv),  [@v](../model/attributes.html#v) `@center`, `@dPdx`, `@dPdy`, `@dPdz` (see [Volume Wrangle](../nodes/sop/volumewrangle.html "Runs a VEX snippet to modify voxel values in a volume.")). |
| `vector4` (4 floats) | [@backtrack](../model/attributes.html#backtrack),  [@orient](../model/attributes.html#orient),  [@rot](../model/attributes.html#rot) |
| `int` | [@id](../model/attributes.html#id),  [@nextid](../model/attributes.html#nextid),  [@pstate](../model/attributes.html#pstate) `@elemnum`, `@ptnum`, `@primnum`, `@vtxnum`, `@numelem`, `@numpt`, `@numprim`, `@numvtx` (see [indexing variables](snippets.html#indexing) below). `@group_*` (see [accessing group membership](snippets.html#groups) below). `@ix`, `@iy`, `@iz`, `@resx`, `@resy`, `@resz` (see [Volume Wrangle](../nodes/sop/volumewrangle.html "Runs a VEX snippet to modify voxel values in a volume.")). |
| `string` | [@name](../model/attributes.html#name),  [@instance](../model/attributes.html#instance) |

### Accessing attributes on other inputs

other_inputs
If the node has more than one input, you can get an attribute from a different input by prefixing the name with `opinputinputnum_`, for example `v@opinput1_P`. This reads the named attribute from the same element (point/primitive/vertex) on the numbered input (where the first input is input 0, the second input is 1, and so on).

The “same element” may be the element with the same index in the other input (for example, when you're processing point number 10, `@opinput1_P` would give you the `P` attribute on point number 10 in the second input).

However, nodes can have an “Attribute to Match” parameter that lets you match up “same” elements based on the value of an attribute. For example, if you used `id` as the “attribute to match”, and you were processing a polygon with attribute `id` set to `12`, then `@opinput1_P` would give you the `P` attribute on the polygon in the second input that also has `id` set to `12`. Check the parameters of the node in which you're writing the snippet.

Setting geometry attributes

## setgeoattrs

You can set geometry attributes using `@` syntax, for example `@foo = 12.5`. This is the preferred method for setting geometry attributes in a VEX snippet.

There are logically three different geometries available during a VEX function:

- The *input geometry*. This is what you read with [point(0, …)](functions/point.html "Reads a point attribute value from a geometry."), for example.
- The *current geometry*. This is the current point/prim of the points/prims the VEX function is running over. This is what read/modify using the `@foo` syntax.
- The *output geometry*. This is what you write to with [setpointattrib(0, …)](functions/setpointattrib.html "Sets a point attribute in a geometry."), for example.

This is why reads from `point()` will not show any changes you write with `@foo` or `setpointattrib()`. And it is also why any changes you make with `setpointattrib` override any changes using the `@foo =` syntax.

Using these separate copies of the geometry allows us to re-order the operations. Otherwise we would need either a lot of locking (hurting performance), or there would be race conditions in your code. This approach allows us to stay lock-free and also have changes be deterministic.

Indexing variables

## indexing

Most snippets involve looping over all the points/primitives in a geometry. You can also loop over a list of numbers from zero to some limit. It is often useful to know the number of the current element in the list you are looping over, and the total number of elements in the list.

`@elemnum`

The number of the current element.

You can use `@elemnum` to be generic (or when you are iterating over numbers). If you know you are operating on points (for example), you could use `@ptnum` instead to be clearer, but at the risk that the code won’t work if you change to operating on primitives or vertices.

`@numelem`

The total number of elements in the current geometry/list.

You can use `@numelem` to be generic (or when you are iterating over numbers). If you know you are operating on points (for example), you could use `@numpt` instead to be clearer, but at the risk that the code won’t work if you change to operating on primitives or vertices.

`@ptnum`

The point number of the current point, when the snippet is looping over points. If looping over vertices, this is the point that the vertex is wired to. If looping over primitives, this is the point of the 0th vertex on the primitive.

`@primnum`

The primitive number of the current [primitive](../model/primitives.html), when the snippet is looping over primitives. If looping over vertices, this is the primitive that owns the vertex. If looping over points, this is a primitive that contains the point, -1 if no primitive does. Note that if a point is in more than one primitive, it is arbitrary which one is returned.

`@vtxnum`

The vertex number of the current vertex, when the snippet is looping over vertices. If looping over points, this is a vertex that wires to this point, -1 if no vertex does. If more than one vertex wires to this point, it is arbitrary which vertex is returned. If looping over primitives, it it is the 0th vertex of the primitive.

The *linear vertex number*. This counts over *all vertices in the geometry*, from 0 to the total number of vertices in the geometry - 1). It is different from the vertex’s *primitive index*, which is the vertex’s number within the primitive it is a part of.

The primitive number of the primitive this vertex is on is in `@primnum`. To get the primitive number an arbitrary linear vertex index is on, use [vertexprim](functions/vertexprim.html "Returns the number of the primitive containing a given vertex."). To convert the linear vertex index into a vertex index within the primitive it is a part of, use [vertexprimindex](functions/vertexprimindex.html "Converts a linear vertex index into a primitive vertex number."). When looping over vertices, the number of vertices in the current primitive is in `@numvtx`. To get the number of vertices on an arbitrary primitive, call [primvertexcount](functions/primvertexcount.html "Returns number of vertices in a primitive in a geometry.") with the primitive number.

For example, if you wanted to set a vertex attribute on the vertices of a polycurve to the proportional value along the curve, you would say:

```vex
# Take linear vertex index and convert to primitive index
int vtx = vertexprimindex(0, @vtxnum)

# Set an attribute to the vertex number divided by the last index, so the
# values range from 0.0 to 1.0 along the length of the polycurve
f@prop = vtx / (float(@numvtx) - 1)

```

`@numpt`

The total number of points in the current geometry, when the snippet is looping over points.

`@numprim`

The total number of [primitives](../model/primitives.html) in the current geometry, when the snippet is looping over primitives.

`@numvtx`

The number of vertices in the current primitive, when looping over
vertices, primitives, or points. When in detail mode, the total
number of vertices.

This means @vtxnum is not in the range of `0 - @numvtx-1`
for most iteration types!

Note
The element number of the *last* element is `@numelem - 1`, because the first item is numbered `0`.

For example, to get a list of the points near the current point in a snippet:

```vex
int nbors[] = neighbours(0, @ptnum)

```

To read an attribute from the point *opposite* the current point on a curve:

```vex
vector opposite_color = point(0, "Cd", (@numpt - 1) - @ptnum )

```

Arrays

## arrays

You can bind arrays by appending `[]`, as in

```vex
i[]@connected_pts = neighbours(0, @ptnum);

```

For example, the following code loads the `foo` attribute as a vector
and copies it to the `P` (position) attribute.
You don’t need to specify the type of the `P` attribute because
it’s one of the known attributes Houdini casts automatically.

```vex
@P = v@foo;

```

The following code sets the `x` component of the `Cd` attribute to
the value of the `whitewater` attribute.
You don’t need to specify the type of the `Cd` attribute because
it’s one of the known attributes.
You don’t need to specify the type of the `whitewater` attribute because
it’s a float and unknown attributes are cast as float automatically.

```vex
@Cd.x = @whitewater;

```

Tip
You only have to specify the type character the *first time*
you refer to the attribute in the code.

You can also explicitly prototype attribute bindings. This allows
you to also specify the default value of the attribute which will
be used if the attribute isn’t bound. If an attribute is created,
it will be also set to this default value.

Note
String attributes do not currently set their defaults properly when
created.

This is done by declaring them as a variable. The declaration
must start at the start of the line. Only one variable can
be declared in a line. The default value must be a constant value,
computed values like `3*5` will fail as they are not valid initializers
in the parameter list.

The following will create a foo attribute of type vector. If it doesn’t
exist on the input, the default value will be set to `{ 1, 3, 5 }`.

```vex
vector @foo = { 1, 3, 5 };

```

After declaring it in this manner it is not necessary to use the `v@foo`
syntax, `@foo` will suffice as the type has been specified.

Attributes prototyped in this fashion will take precedence over any
inline definitions (such as `v@foo`). In the future mismatched types
or mismatched defaults may be considered an error.

For more information, see the [POP Attributes](../dopparticles/attributes.html) page.

Declaring attributes

## declare

You can specify the type and default value of attributes before you use them like this:

```vex
float @mass = 1;
vector @up = {0, 1, 0};

```

This can be useful in two ways:

- It gives a default value to the variable: if the attribute (for example, `@mass`) exists, the assignment is ignored. If the attribute doesn’t exist, it uses the assignment.
- It specifies the data type of the attribute. After declaring the type of the `@up` attribute like this, you could just use `@up` instead of `v@up`.

You **cannot** do any computation on the right side of the equals sign (`=`). The following are syntax errors:

```vex
float @mass = 1 / area;  // Error
vector @up = set(0, 1, 0);  // Error

```

Accessing globals

## accessing-globals

Unlike in HScript expressions, you cannot use global variables such as `$F`.

In a VOP, you can wire variables such as `Time` and `Frame` from the Globals
node to use them in a VEX snippet.

You can use the following implicit variables:

`@Time`

Float time (`$T`)

`@Frame`

Float frame (`$FF`)

`@SimTime`

Float simulation time (`$ST`), only present in DOP contexts.

`@SimFrame`

Float simulation frame (`$SF`), only present in DOP contexts.

`@TimeInc`

Float time step (`1/$FPS`)

Creating geometry

## creating-geometry

You can create geometry with VEX snippets in certain nodes, such as [![](../icons/SOP/attribwrangle.svg)Attrib Wrangle](../nodes/sop/attribwrangle.html "Runs a VEX snippet to modify attribute values.").

The [addpoint](functions/addpoint.html "Adds a point to the geometry."), [addprim](functions/addprim.html "Adds a primitive to the geometry."), and [addvertex](functions/addvertex.html "Adds a vertex to a primitive in a geometry.") functions let you create points, primitives, and vertices. You can alter geometry using [setattrib](functions/setattrib.html "Writes an attribute value to geometry.") and [setprimvertex](functions/setprimvertex.html "Rewires a vertex in the geometry to a different point."). You can remove geometry using [removepoint](functions/removepoint.html "Removes a point from the geometry.") and [removeprim](functions/removeprim.html "Removes a primitive from the geometry."). To set group membership use [setpointgroup](functions/setpointgroup.html "Adds or removes a point to/from a group in a geometry.") and [setprimgroup](functions/setprimgroup.html "Adds or removes a primitive to/from a group in a geometry."). Use [setprimintrinsic](functions/setprimintrinsic.html "Sets the value of a writeable primitive intrinsic attribute.") to modify things like the transform of sphere primitives.

It’s faster to set attributes on the current element using bound variables (for example `@name = val`) rather than [setattrib](functions/setattrib.html "Writes an attribute value to geometry."). Only use [setattrib](functions/setattrib.html "Writes an attribute value to geometry.") if you need to do set an attribute on *other* elements. If you are using [setattrib](functions/setattrib.html "Writes an attribute value to geometry.") and are modifying a point from different source points, set the `mode` argument to `"add"` to composite the results.

The geometry creation functions can run in parallel. All changes are queued and applied *after* your VEX code has iterated over all existing geometry. This means [setattrib](functions/setattrib.html "Writes an attribute value to geometry.") will *overwrite* changes you make through bound variables (for example `@name = val`).

The first argument to the geometry creation functions is a “geometry handle”, which specifies a destination for the created (this is intended to support writing to a file as an alternative to writing to the current geometry). Use `geoself()` as the first argument to specify the current geometry.

```vex
addpoint(geoself(), {0, 1, 0});

```

The [addprim](functions/addprim.html "Adds a primitive to the geometry.") function can currently generate a polygon (`"poly"`) or polyline (`"polyline"`). If you create a polygon, you *must* add vertices to the points using [addvertex](functions/addvertex.html "Adds a vertex to a primitive in a geometry."). Houdini will likely crash on a polygon that has points but not vertices.

```vex
int p1 = addpoint(geoself(), {0, 1, 0});
int p2 = addpoint(geoself(), {1, 1, 0});
int p3 = addpoint(geoself(), {1, 1, 1});

int prim = addprim(geoself(), "poly");
addvertex(geoself(), prim, p1);
addvertex(geoself(), prim, p2);
addvertex(geoself(), prim, p3);

```

- You can inspect the topology of the geometry using [vertexpoint](functions/vertexpoint.html "Returns the point number of linear vertex in a geometry."), [pointvertex](functions/pointvertex.html "Returns a linear vertex number of a point in a geometry."), [vertexprim](functions/vertexprim.html "Returns the number of the primitive containing a given vertex."), [vertexnext](functions/vertexnext.html "Returns the linear vertex number of the next vertex sharing a point with a given vertex."), [vertexprev](functions/vertexprev.html "Returns the linear vertex number of the previous vertex sharing a point with a given vertex."), and [primvertexcount](functions/primvertexcount.html "Returns number of vertices in a primitive in a geometry.").
- You can read from point cloud files using the `pc*` functions ([pcopen](functions/pcopen.html "Returns a handle to a point cloud file."), [pcnumfound](functions/pcnumfound.html "This node returns the number of points found by pcopen."), [pciterate](functions/pciterate.html "This function can be used to iterate over all the points which were
  found in the pcopen query."), [pcimport](functions/pcimport.html "Imports channel data from a point cloud inside a pciterate or a pcunshaded loop."), and so on).

Geometry traversal functions

## traverse

See [VEX geometry functions](geometry.html).

Accessing group membership

## groups

A special virtual attribute of the form `@group_groupname` lets you get or set group membership for the current element.

You can check if the current point/primitive/particle is in a named group by checking if `@group_name == 1`.

You can add or remove the current point/primitive to a group by setting the virtual `@group_name` attribute. Setting the attribute to `1` (or any non-zero value) puts the current element in that group. Setting the attribute to `0` removes the current element from that group.

User-defined functions

## user-defined-functions

You can define your own functions as part of a VEX snippet using the [VEX function syntax](lang.html#functions).
For example:

```vex
float mysquare(float a)
{
  return a * a;
}

windresist = mysquare(windresist);

```

Includes

## includes

Any `#include` directives found in the code snippet will be automatically moved
outside of the generated VEX function, so will behave as expected.

Determining if a parameter is present for attribute binding is done by a
simple scan of the code after pre-processing is done. This pre-processing
is done **only** on the code snippet; however, and does **not** process any
`#include` files. It can therefore be confused by `#ifdef` directives that
depend on `#includes`.

Tips

## tips

- When you're editing in the multi-line editor you can press `⌃ Ctrl` + `Enter` to “commit” the changes and update Houdini.
- The VEX snippet is run at every frame (or in a simulation network, every time step).
- You can exit the snippet early using the `return` statement. For example, the following VEX will only set `windresist` to `0` on brand new particles:

```vex
if (@age > 0) return;

// If @age was > 0, we returned above, so this line
// only runs for particles where @age == 0
windresist = 0;

```

Troubleshooting error messages

## troubleshooting-error-messages

| Your code | Problem |
| --- | --- |
| `force += 2` | `Syntax error, unexpected '}', expecting ';'` Each statement must end with a semicolon (`;`) |
| `@v += force;` | `SRead-only expression given for read/write parameter` Particle nodes cannot modify particle attributes. |
| `x = { 0, @y, 0};` | `Syntax error, unexpected identifier, expecting '}'.` You cannot have varying arguments to the `{}` vector constructor. Use `set()` instead: `x = set(0, @y, 0);`. |
| `x = set(0, $F, 0);` | `Doesn't animate.` While `$F` will be evaluated, VOP networks are not time dependent so it won’t animate. Use `@Frame` instead. |
