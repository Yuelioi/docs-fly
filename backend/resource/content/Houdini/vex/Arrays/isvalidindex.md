---
title: isvalidindex
order: 7
---

# isvalidindex

## 描述

检查数组或字符串中的索引是否合法

```c
int  isvalidindex(<type>&array[], int index)
int  isvalidindex(string str, int index)
```

Returns `1` if `index` is within range for the given string/array, or `0`
otherwise.

This is equivalent to `index < len (array) && index >= -len (array)`

如果在给定字符串/数组的范围内有 index，则返回 1，否则返回 0。

相当于 `index < len (array) && index >= -len (array)`。

---

```c
int  isvalidindex(dict d, string key)
```

Returns `1` if the key is in the dictionary, or `0` otherwise.

如果键在字典中，返回 1，否则返回 0。

## 示例

```c
int arr[] = {1,2,3};
string s = "hello world";
int i = isvalidindex(arr,2); // 2没超，返回1
int j =isvalidindex(s,11); // 11超了，返回0


dict d = set("key1", 3, "key2", 5, "key3", 7);
int k =isvalidindex(d,"key1"); // 字典有这个键，返回1

printf("%d", i);
printf("%d", j);
printf ("%d", k);
```
