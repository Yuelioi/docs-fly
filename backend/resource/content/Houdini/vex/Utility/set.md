---
title: set
order: 17
---
The `set()` function has a **wide variety** of forms allowing you to perform many different types of conversions.

Tip
When filling a matrix in Houdini, the numbers go across the first row, then across the second row, and so on (“row-major”).

`vector2  set(float v1, float v2)`

`vector  set(float v1, float v2, float v3)`

`vector4  set(float v1, float v2, float v3, float v4)`

`matrix2  set(float v1, float v2, float v3, float v4)`

`matrix3  set(float v1, float v2, float v4, float v4, float v5, float v6, float v7, float v8, float v9)`

`matrix  set(float v1, float v2, float v3, float v4, float v5, float v6, float v7, float v8, float v9, float v10, float v11, float v12, float v13, float v14, float v15, float v16)`

Create a vector or matrix type from the values in the arguments.

```vex
vector4 v = set(1.0, 2.0, 3.0, 4.0);
matrix3 m = set(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0);

```

The [assign](assign.html "An efficient way of extracting the components of a vector or matrix into float variables.") function performs the opposite of this operation (pulling components out into variables).

`<vector> set(float nums[])`

`<matrix> set(float nums[])`

Creates a vector/matrix from an array of floats.

```vex
float[] nums = {1.0, 2.0, 3.0, 4.0};
vector4 v = set(nums);

```

`<vector> set(float|intv)`

`<matrix> set(float|intv)`

If you set a vector or matrix type from a single value, the vector/matrix is filled in with that value for all components.

```vex
vector4 v = set(2.0);  // -> {2.0, 2.0, 2.0, 2.0}

```

`matrix2  set(vector2 row1, vector2 row2)`

`matrix3  set(vector row1, vector row2, vector row3)`

`matrix  set(vector4 row1, vector4 row2, vector4 row3, vector4 row4)`

Creates a matrix type from a series of vector arguments of the same size representing the rows.

```vex
matrix3 m = set({0.0, 1.0, 0.0}, {1.0, 0.0, 1.0}, {0.0, 1.0, 0.0});

```

`matrix3  set(vector rows[])`

`matrix  set(vector4 rows[])`

You can also set a matrix type from an array of vectors representing rows. If there are not enough vectors in the array to fill the matrix, the extra rows will contain zeros.

`vector [] set(matrix3 m)`

`vector4 [] set(matrix m)`

Creates an array of vectors from the rows of a matrix of the same size.

```vex
matrix3 m3 = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0};
vector[] vs = set(m3);  // -> array [ {1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {7.0, 8.0, 9.0} ]

```

`vector  set(vector2 v)`

`vector4  set(vector2 v)`

`vector4  set(vector v)`

`matrix3  set(matrix2 m)`

`matrix  set(matrix2 m)`

`matrix  set(matrix3 m)`

If you set a larger vector or matrix type from a smaller type, the extra components will be zero.

```vex
vector2 v2 = {1.0, 2.0};
vector4 v4 = set(v2);  // -> {1.0, 2.0, 0.0, 0.0}

```

`float  set(vector v)`

`float  set(vector4 v)`

`vector2  set(vector4 v)`

`vector2  set(vector v)`

`vector  set(vector4 v)`

If you set a smaller vector with a larger vector, the smaller type will take components from the left.

```vex
vector4 v4 = {1.0, 2.0, 3.0, 4.0};
vector2 v2 = set(v4)  // -> {1.0, 2.0}

```

`matrix2  set(matrix3 m)`

`matrix2  set(matrix m)`

`matrix3  set(matrix m)`

If you set a smaller matrix with a larger matrix, the smaller type will take the top-left corner.

```vex
matrix3 m3 = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0};
matrix2 m2 = set(m3);  // -> {1.0, 2.0, 4.0, 5.0}

```

`float [] set(<vector>v)`

`float [] set(<matrix>m)`

Creates an array of floats from the components of a vector or matrix type.

```vex
matrix3 m3 = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0};
float[] nums = set(m3);  // -> array [ 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0 ]

```

`<vector>[] set(float nums[])`

`<matrix>[] set(float nums[])`

Creates an array of vector/matrix types by taking one component at a time from an array of floats. This is the same as using the [unserialize](unserialize.html "Turns a flat array of floats into an array of vectors or matrices.") function.

```vex
float[] nums = {1.0, 2.0, 3.0, 4.0};
vector2[] vs = set(nums);  // -> [ {1.0, 2.0}, {3.0, 4.0} ]

```

`int  set(float v)`

`float  set(int v)`

`float [] set(int vs[])`

`int [] set(float vs[])`

You can set a float from an int, or an int from a float, or an array of floats from an array of ints, or an array of ints from an array of floats.

```vex
float[] fracs = { 1.1, 2.2, 3.3, 4.4, 5.5, 6.6 };
int[] floored = set(fracs);  // -> array [ 1, 2, 3, 4, 5, 6 ]

```

`float [] set(float num)`

Sets all the items in an array of float to the same value.

`float  set(float nums[])`

Returns the first item in the array.

`dict  set(string key, ...)`

Initializes a dictionary with a series of key/value pairs. Every other parameter should
be a string for the key and be followed by the value for that key.

`<type> set(<type>v)`

`<type>[] set(<type>v[])`

If you call `set()` with the same argument type and return type, it simply returns the argument value.

```vex
string s = set("Hello");  // -> "Hello"

```
