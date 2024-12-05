---
order: 3
title: argsort
---

# argsort

## 描述

Returns the indices of a sorted version of an array.

返回数组有序排列下的索引列表

arg = 参数、属性 argument
sort = 排序
加起来 = 基于数组元素的属性排序

```c
int [] argsort(<type>value[])
```

Returns a list of indices that, if applied to the given array, will give a
sorted sequence in increasing order.

返回索引列表，如果应用于给定的数组，将给出一个按递增顺序排序的序列。

This lets sort an array by some property of the items in the array rather than
by the values themselves.

这样可以通过数组元素某些属性来排序（比如元素长度），而不是基于数值本身。

- argsort and sort use a stable sort. | argsort 和 [[sort]] 使用稳定排序。
- Use to reverse the order of the sort. | 使用 [[../../Houdini/Houdini/SOP/Polygon/reverse]] 来反转排序。

## 示例：数组自定义排序

```c
// 原数组：字符串数组
string colors[] = {"Red", "Green", "Blue", "Orange", "Violet", "Indigo"};


// 基于元素长度，创建新数组（{3, 5, 4, 6, 6, 6}）
int lengths[] = {};
foreach (string name; colors) {
 push(lengths, len(name));
}


// 获取索引列表（{0, 2, 1, 3, 4, 5}）
int ordering[] = argsort(lengths);


// 基于排序规则 给原来的数组排序
string colors_by_len[] = reorder(colors, ordering);


printf("%s\n", colors_by_len); // {Red, Blue, Green, Orange, Violet, Indigo}

```
