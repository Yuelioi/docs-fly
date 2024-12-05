---
title: removevalue
order: 12
---

# removevalue

## 描述

Removes an item from an array. 从数组中移除一个元素

```c
int  removevalue(<type>&array[], <type>value)
```

Removes the first instance of `value` found from the array. Returns `1` if an
item was removed, or `0` otherwise.

从数组中删除 valuefound 的第一个实例。如果有项目被移除，则返回 1，否则返回 0。

## 示例：移除一个元素

```c
int arr1 [] = {1, 2, 3, 2};
int trg = removevalue(arr1,2);
printf("%s",trg);     // 返回1. 数组中第1个"2"的索引为1
// 原来的数组变为 {1,3,2};
```
