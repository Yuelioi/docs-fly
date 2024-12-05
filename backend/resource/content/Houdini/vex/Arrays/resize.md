---
title: resize
order: 14
---

# resize

## 描述

Sets the length of an array. 设置数组长度

```c
void  resize(<type>&array[], int size)
```

Changes the size of the given array. If `size` is greater than the current length of the array, the extra items at the end of the array will be initialized to default values of the given type (for example, `0`, empty string, `{0,0,0}`, etc.).

改变给定数组的大小。如果 size 大于当前数组的长度，那么数组末端的额外项将被初始化为给定类型的默认值（例如，0，空字符串，{0，0，0}，等等）。

```c
void  resize(<type>&array[], int size, <type>val)
```

Changes the size of the given array. If `size` is greater than the current length of the array, the extra items at the end of the array will be initialized to `val`.

改变给定数组的大小。如果 size 大于数组的当前长度，数组末尾的额外项将被初始化为 oval。

## 示例

```c
int arr1[] = {1,2,3,4};

resize(arr1,6);  // {1, 2, 3, 4, 0, 0}，长则补0
resize(arr1,3);  // {1, 2, 3}。短则裁剪

resize (arr1, 7);  // {1, 2, 3, 4, 8, 8, 8} // 长则补 8

```
