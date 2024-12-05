---
title: removeindex
order: 11
---

# removeindex

## 描述

Removes an item at the given index from an array. 移除数组给定下标的元素

```c
<type> removeindex(<type>&array[], int index)
```

Removes the item at `index` from `array` and returns its value. This is the
same as pop (array, index) but has a slightly clearer name.

A negative index counts from the end of the array, so removeindex (array, -2) removes the second-to-last item.

从 array 中移除 index 处的项目并返回它。与 pop(array, index)相同，但是名字更好记一些。

负数从数组的末尾开始计算，所以 moveindex(array, -2)删除倒数第二项。

```c
int  removeindex(dict &dictionary, string index)
```

Removes the dictionary entry of `index` from `dictionary`.

Returns `0` if no such entry was present, `1` if something wasremoved.

从 dictionary 中删除 index 的字典条目。

如果没有这样的条目，返回 0，如果成功删除，返回 1。

## 示例

```c
int arr1[] = {1,2,3};
int trg = removeindex(arr1,2);
printf("%s",trg); // 3 数组第2个值为3

dict d1 = set("key1", 3, "key2", 5, "key3", 7);
int trg2 = removeindex(d1,"key2");

printf("%s",trg2); // 1 移除成功
```
