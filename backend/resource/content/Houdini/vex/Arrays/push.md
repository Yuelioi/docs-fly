---
title: push
order: 10
---

# push

## 描述

Adds an item to an array. 在数组中添加元素

```c
void  push(<type>&array[], <type>value)
```

Appends the given value to the end of the array. Increases the size of `array`
by 1. This is the same as [append(array, value)](append.html "Adds an item to
an array or string.").

将给定的值追加到数组的末端。数组大小增加 1，与 append(array, value)相同。

```c
void  push(<type>&array[], <type>values[])
```

Concatenates the values from the `values` array to the end of `array`.
Increases the size of `array` by `len(values)`. This is the same as
[append(array, values)](append.html) () ("Adds an item to an array or string.").

将 values 中的值串联到 array 的末尾。array 大小增加 len(values)个。与 append(array, values)相同。

Tip

You can set an individual item in an array using `array[n] = x`.

可以用 array[n] = x 来设置一个数组中的单个项目。

## 示例

```c
// 数组尾部追加元素
int arr[]={1,2};
push(arr,3);// {1, 2, 3}


// 数组尾部追加数组
int k[]={1,2,3};
int j[]={4,5};
push(k,j);// {1, 2, 3, 4, 5}


// 修改数组一个元素
int arr2[] = {5,6,7};
arr2[0] = 0;//{0, 6, 7}
```
