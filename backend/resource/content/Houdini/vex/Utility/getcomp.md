---
title: getcomp
order: 5
---
`float  getcomp(<vector>v, int index)`

Returns the vector component at the given index.
This is the same as `v[index]`.

`float  getcomp(<matrix>m, int row, int column)`

Returns the matrix component at the given location.

`<type> getcomp(<type>array[], int index)`

Returns the array item at the given index.
This is the same as `array[index]`.

`<type> getcomp(<vector>array[], int i, int j)`

Returns the vector component at the given location and array index. This is the same as `getcomp(array[i], j)`.

`<type> getcomp(<matrix>array[], int i, int j, int k)`

Returns the matrix component at the given location and array index. This is the same as `getcomp(array[i], j, k)`.

`<type> getcomp(dict d, string index)`

`<type>[] getcomp(dict d, string index)`

Returns the dictionary item at the given index.
This is the same as `d[index]`.

`<type> getcomp(dict d, string index, <type>defvalue)`

`<type>[] getcomp(dict d, string index, <type>defvalue[])`

Returns the dictionary item at the given index. If it doesn’t exist,
return `defvalue`.
This is the same as `isvalidindex(d, index) ? d[index] : defvalue`.

`string  getcomp(string value, int index)`

Returns the *character* at the given index.
This is the same as `value[index]`.

Characters in VEX are strings as well. UTF-8 encoding is used,
so if the index is part way through a UTF-8 encoding, the result
is an empty string. Otherwise it is the entire valid UTF-8 character.
