---
title: append
order: 2
---

# append

## 描述

Adds an item to an array or string.

添加一个元素到数组或者字符串到原数组/字符串

## 函数分类

Appends the second string to the first.

将第二个字符串追加到第一个字符串上。

```c
void  append(string &array, string value)
```

Appends the given value to the end of the array. Increases the size of `array`
by 1. This is the same as [[push]] (array, value)

将给定值追加到数组末尾。此时 array 的大小增加 1，与 push(array, value)相同。

```c
void  append(<type>&array[], <type>value)
```

Concatenates the values from the `values` array to the end of `array`.
Increases the size of `array` by `len(values)`. This is the same as
[[push]] (array, values)

将 value sarray 中的值串联到 array 的末尾。将 array 的大小增加到 len(values)。与 push(array, values)相同。

```c
void  append(<type>&array[], <type>values[])
```

::: tip

You can set an individual item in an array using `array[n] = x`.

可以用 array[n] = x 来设置数组中的单个元素。
:::

## 示例

```c
// 数组尾部添加 1 个元素
int arr[]={1,2};
append(arr,3);  // {1, 2, 3}


// 字符串尾部添加新字符串
string s = "hello";
append(s," world");  // hello world


// 数组尾部追加数组
int k[]={1,2,3};
int j[]={4,5};
append(k,j);  // k → {1, 2, 3, 4, 5} 。j不变


// 修改数组一个元素
int arr2[] = {5,6,7};
arr2[0] = 0;  //{0, 6, 7}
```
