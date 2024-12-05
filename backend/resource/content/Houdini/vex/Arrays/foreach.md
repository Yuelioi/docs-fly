---
title: foreach
order: 5
---

# foreach

## foreach

Loops over the items in an array, with optional enumeration.
循环处理一个数组中的元素，可以选择枚举法。

The length of the array is determined before the first iteration, so if the
array is changed during the foreach this will not be reflected in the number
of iterations.

数组长度是在第一次迭代前确定，如果数组在 foreach 过程中被改变，不会反馈到迭代次数上。

## Simple form 简单形式

```c
foreach ([element_type] value; array) {

}
```

This loops over the members of array. For each iteration, it **copies** the
current member to value and then executes statement. For example:

循环数组成员。对于每次迭代，会将复制当前成员到 value，然后执行语句。比如：

```c
int an_array[] = {1, 2}

// 逐次打印元素（共2次）
 foreach (int num; an_array) {
 printf("%d", num);
}
```

## Enumerated form 枚举

The second form lets you specify an enumeration variable:

指定一个枚举变量。

```c
foreach (index, value; array) statement;
foreach (int index; element_type value; array) statement;
```

For each iteration, this form assigns the current _position_ in the arrayto
index, **copies** the current member to value, and executesstatement. For
example:

对于每次迭代，会将数组中的当前位置（索引）分配给 index，将当前成员复制到 value，然后执行语句。

可以在遍历中使用各种数组信息：

- index：当前遍历的数组索引
- value：当前数组元素
- array：要遍历的数组

![](https://cdn.yuelili.com/20220227234450.png)

示例:

```c
// 打印时间表

string days[] = { "Mon", "Tue", "Wed", "Thu", "Fri" };
foreach (int i; string name; days) {
 printf("%d\t%s\n", i, name);
}

```

和 python 的语法有点像

```python
for i, x in enumerate(xs):
```
