---
title: Arrays
order: 1
---
| On this page | * [Overview](#overview) * [Declaring array types](#declaring-array-types) * [Accessing and setting array values](#accessing-and-setting-array-values) * [Slicing Arrays](#slicing-arrays) * [Copying between arrays and vectors/matrices](#copying-between-arrays-and-vectors-matrices) * [Looping over an array](#looping-over-an-array) * [Working with arrays](#working-with-arrays) * [VCC pragmas](#vcc-pragmas) * [Limitations](#limitations) |
| --- | --- |

Overview

## overview

VEX includes an array datatype. This is useful in several places:

- Supporting ramp parameters.
- Reading capture data from surface nodes using the [import()](functions/import.html) function.
- General programming, wherever arrays would be useful.

Note
Currently VEX does not support multi-dimensional arrays.

Note
A function that returns an array defined inside a Snippet,
Wrangle, or other function should have the function keyword.

This example shows off some of the crazy things that you can do with arrays:

```vex
surface
crazy(
      string maps[] = { "Mandril.rat", "default.pic" };
      export float alength = 0;
      )
{
    vector texclr, av[];

    texclr = texture(maps[s+t > 1], s, t);
    av = array( {1,0,0}, vector(nrandom()), t, texclr, {.5,0,0});

    if (fit(noise(s*8), 0, 1, .3, .7) > t)
        av = array(1, {0,1,0}, 0);

    Cf = spline("linear", s, av);
    alength = len(av);
}

```

Declaring array types

## declaring-array-types

To declare an array variable, the general form is
`member_type var_name[]`:

```vex
// my_array is an array of floats
float   my_array[];

// v is a single vector, vector_array is an array of vectors
vector  v, vector_array[];

// str_array is an array of strings
string  str_array[];

```

You can optionally put a size inside the square brackets, but the
VEX compiler currently ignores it.

To declare a function that returns an array:

```vex
// A function which returns an array of vectors
// DOES NOT WORK IN A WRANGLE/SNIPPET: use function keyword then.
vector[] rgb_array()
{
...
};    

```

It is ambiguous if you are a nested function, however. Note that
Wrangles and Snippets are always implicitly nested. To declare a
nested function that returns an array:

```vex
// A function which returns an array of vectors
cvex
foo()
{
    // Use the optional 'function' keyword to avoid type ambiguity
    function vector[] rgb_array()
    {
    ...
    };    
}

```

To specify a literal array, use curly braces, with the array members
separated by commas:

```vex
vector an_array[] = { {1, 2, 3}, {2, 3, 4}, {4, 5, 6} };

vector[] rgb_array()
{
    return { {1, 0, 0}, {0, 1, 0}, {0, 0, 1} };
}

```

Note
Literal arrays are constructed *at compile time*, so they **cannot
include variables**.

For example, this is an error:

```vex
int arr[] = { my_var, other_var + 2 }; // Error

```

To avoid this error, use the `array()` function which constructs
the array at runtime from any number of arguments:

```vex
int arr[] = array( my_var, other_var + 2 );

```

If you specify scalars where a vector is expected, the compiler assigns
the scalar value to all components of the vector:

```vex
vector an_array[] = { 1, 2, 3};
// an_array[] == { {1, 1, 1}, {2, 2, 2}, {3, 3, 3} }

```

The `array()` function creates an array from its arguments.

```vex
int my_array[] = array(1, 2, 3, 4, 5);

```

You can use `array()` to generate an array of any type.
To force `array()` to generate vectors (for example):

```vex
vector (array (value1, value2, ...) );

```

Accessing and setting array values

## accessing-and-setting-array-values

Use `arrayname[index]` to look up a value by its position in the
array.

```vex
vector bw[] = { 0, 1 };
// bw[] == { {0, 0, 0}, {1, 1, 1} }
Cf = bw[index];

```

Array bounds are checked at run time. Reading out of bounds will return `0` or
`""`. This may generate a warning or optional run-time error in the future.
Writing past the end of an array will resize the array to include the
index written to. The new entries will be set to `0` or `""`.

Python-style indexing is used. This means negative indices refer to positions
from the end of the array.

```vex
int nums[] = { 0, 1, 2, 3, 4, 5 };
int n = nums[10];  // Returns 0
int b = nums[-2];  // Returns 4

string strs[] = { };
string s = strs[20];  // Returns ""

```

You can also assign values using the square brackets notation:

```vex
float nums[] = { };
nums[0] = 3.14;

```

(The [getcomp](functions/getcomp.html "Extracts a single component of a vector type, matrix type, or array.") and [setcomp](functions/setcomp.html "Sets a single component of a vector or matrix type, or an item in an array.") functions are equivalents for using
the square brackets notation.)

Note
The square-brackets operator also works on vectors. You can use it with
matrices as well using a pair of brackets: `float a = m3[0][1];`

Slicing Arrays

## slicing-arrays

The square-brackets can be used to extract sub-arrays using the Python slicing notation.

```vex
int nums[] = { 0, 1, 2, 3, 4, 5 };
int start[] = nums[0:2];  // { 0, 1 }
int end[] = nums[-2:];  // { 4, 5 }
int rev[] = nums[::-1];  // { 5, 4, 3, 2, 1, 0 }
int odd[] = nums[1::2]; // { 1, 3, 5 }

```

The [slice](functions/slice.html "Slices a sub-string or sub-array of a string or array.") function is the equivalent for using the slice-based square
brackets notation.

Copying between arrays and vectors/matrices

## copying-between-arrays-and-vectors-matrices

The assignment operator supports assigning values between vector types and
arrays of floats:

```vex
float x[];
// Cf and P are vectors

x = set(P);   // Assigns the components of P to the corresponding
              // members of the array x

Cf = set(x);  // Assigns the first 3 members of x as the
              // components of the vector Cf

```

If the array is not long enough to fill the vector/matrix, the last member is
repeated as often as necessary.

```vex
float x[] = {1, 2} // Not long enough to fill a vector
Cf = set(x);  // Cf == {1, 2, 2}

```

You can also assign between matrix types and arrays of `vector2`/`vector`/`vector4`:

```vex
vector2     v2[];
vector      v[];
vector4     v4[];
matrix2     m2 = 1;
matrix3     m3 = 1;
matrix      m4 = 1;

v = set(m3);   // Each row of the 3x3 matrix is put into a vector
m3 = set(v);   // Copy the vectors into the row vectors of the matrix
v4 = set(m4);  // Extract the rows of the matrix into the vector4 array
m4 = set(v4);  // Create a matrix using the vector4's in the array as row vectors

```

In summary:

| Left side = | Right side | Notes |
| --- | --- | --- |
| `vector2` | `float[]` | E.g. `vector2 v = {1,2}` |
| `vector` | `float[]` | E.g. `vector v = {1,2,3}` |
| `vector4` | `float[]` | E.g. `vector4 v = {1,2,3,4};` |
| `matrix2` | `float[]` | E.g. `matrix2 m = {1,2,3,4};` |
| `matrix2` | `vector2[]` | E.g. `matrix2 m = { {1,2}, {4,5} };` |
| `matrix3` | `float[]` | E.g. `matrix3 m = {1,2,3,4,5,6,7,8,9};` |
| `matrix3` | `vector[]` | E.g. `matrix3 m = { {1,2,3}, {4,5,6}, {7,8,9}};` |
| `matrix` | `float[]` | E.g. `matrix m = {1,2,3,4,5,6,7,8,9.., 16};` |
| `matrix` | `vector4[]` | E.g. `matrix m = { {1,2,3,4}, {5,6,7,8}, ... {13,14,15,16}};` |
| `float[]` | `vector2` | Create an array of 2 floats from the components |
| `float[]` | `vector` | Create an array of 3 floats from the components |
| `float[]` | `vector4` | Create an array of 4 floats from the components |
| `float[]` | `matrix2` | Create an array of 4 floats from the matrix2 |
| `vector2[]` | `matrix2` | Create an array of 2 vector2s from the matrix2 |
| `float[]` | `matrix3` | Create an array of 9 floats from the matrix3 |
| `vector[]` | `matrix3` | Create an array of 3 vectors from the matrix3 |
| `float[]` | `matrix4` | Create an array of 16 floats |
| `vector4[]` | `matrix4` | Create an array of 4 `vector4`s. |

Looping over an array

## looping-over-an-array

See [foreach](functions/foreach.html "Loops over the items in an array, with optional enumeration.").

Working with arrays

## working-with-arrays

The following functions let you query and manipulate arrays.

[resize](functions/resize.html "Sets the length of an array.")

Sets the length of the array. If the array is enlarged, intermediate values
will be `0` or `""`.

[len](functions/len.html "Returns the length of an array.")

Returns the length of an array.

[pop](functions/pop.html "Removes the last element of an array and returns it.")

Removes the last item from the array (decreasing the size of the array by 1)
and returns it.

[removevalue](functions/removevalue.html "Removes an item from an array.")

Removes the first instance of a value in the array. Returns `1` if an item was removed, or `0` otherwise.

[removeindex](functions/removeindex.html "Removes an item at the given index from an array.")

Removes an item at the given index and returns it.

[push](functions/push.html "Adds an item to an array.")

Adds an item to the end of an array (increasing the size of the array by 1).

[getcomp](functions/getcomp.html "Extracts a single component of a vector type, matrix type, or array.")

Gets the value of an array component, the same as `array[num]`.

[setcomp](functions/setcomp.html "Sets a single component of a vector or matrix type, or an item in an array.")

Sets the value of an array component, the same as `array[num] = value`.

[array](functions/array.html "Efficiently creates an array from its arguments.")

Efficiently creates an array from its arguments.

[serialize](functions/serialize.html "Flattens an array of vector or matrix types into an array of floats.")

Flattens an array of vectors or matrices into an array of floats.

[unserialize](functions/unserialize.html "Turns a flat array of floats into an array of vectors or matrices.")

Reverses the effect of [serialize](functions/serialize.html "Flattens an array of vector or matrix types into an array of floats."): assembles a flat array of floats
into an array of vectors or matrices.

[neighbours](functions/neighbours.html "Returns an array of the point numbers of the neighbours of a point.")

An array-based replacement for the [neighbourcount](functions/neighbourcount.html "Returns the number of points that are connected to the specified point.")/[neighbour](functions/neighbour.html "Returns the point number of the next point connected to a given point.")
combo. Returns an array of the point numbers of the
neighbors of a given point.

In addition, the following functions work with arrays:

- [min](functions/min.html)
- [avg](functions/avg.html "Returns the average value of the input(s)")
- [spline](functions/spline.html "Samples a value along a polyline or spline curve.")
- [import()](functions/import.html)
- [addattribute()](functions/addattribute.html)
- [metaimport](functions/metaimport.html "Once you get a handle to a metaball using metastart and metanext, you
  can query attributes of the metaball with metaimport.")

VCC pragmas

## vcc-pragmas

The `ramp` pragma lets you specify a ramp user interface for a set of
parameters.

```vex
#pragma ramp <ramp_parm> <basis_parm> <keys_parm> <values_parm>
```

See [VCC pragmas](pragmas.html) for more information.

Limitations

## limitations

- Currently VEX does not support multi-dimensional arrays.
- Arrays cannot be passed between shaders (through [simport](functions/simport.html "Imports a variable sent by a surface shader in an illuminance loop."), etc.).
- Arrays cannot be written to image planes.
