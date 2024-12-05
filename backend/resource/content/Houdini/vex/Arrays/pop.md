---
order: 9
title: pop
---

# pop

## 描述

Removes the last element of an array and returns it. 移除最后一个元素并返回

```c
<type> pop(<type>&array[])
```

Removes the last item from the array and returns it.

移除数组中的最后一个项目并返回。

```c
<type> pop(<type>&array[], int index)
```

Removes the item at `index` from the array and returns its value.

从数组中移除 index 处的元素并返回。

A negative index counts from the end of the list, so `pop (array, -2)` would remove the second-to-last value.

负数从列表的末尾开始计算，因此 pop(array, -2)将删除倒数第二的值。

## 示例

```c
int arr[] = {1,2,3};
int d = pop(arr); // 3,最后一个元素
int d = pop(arr,1); // 2,第2个元素

printf("%s",d);
```
