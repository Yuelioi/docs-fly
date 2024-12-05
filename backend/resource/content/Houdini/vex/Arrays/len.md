---
title: len
order: 8
---

# len

## 描述

Returns the length of an array. 返回数组的长度

```c
int  len(<vector>v)
int  len(<matrix>m)
int  len(<type>array[])
int  len(string s)
int  len(dict d)
```

Returns the number of items/components in the given object. For an array, this
is the number of items in the array. For a matrix or vector, this is the
number of components.

For a string, this returns the number of _bytes_ (not characters).

For a dictionary, this returns the number of keys in the dictionary.

Do not confuse this function with [length](length.html) "Returns the magnitude
of a vector."), which returns the magnitude of a vector.

返回给定对象中的元素/组件的数量。对于数组来说，是数组的元素个数。对于矩阵或向量，这是其分量数量。

对于字符串，将返回 bytes 的数量（不是字符）。

对于字典，返回字典中的键的数量。

不要与 length 函数混淆，后者返回向量的大小。

## 示例

```c
vector v = {1.0, 2.0, 3.0};
matrix  m  = {{1,1,1,1},
     {2,2,2,2},
     {3,3,3,3},
     {4,4,4,4}};

int arr[] = {1,2,3};
string s = "hello world";
dict d = set("key1", 3, "key2", 5, "key3", 7);


printf("%d\n", len(v));     // 3 (3x1)
printf("%d\n", len(m));     // 16(4x4)
printf("%d\n", len(arr));   // 3
printf("%d\n", len(s));     // 11
printf("%d\n", len(d))
```
