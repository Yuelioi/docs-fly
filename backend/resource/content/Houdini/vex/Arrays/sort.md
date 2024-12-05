---
title: sort
order: 17
---

# sort

## 描述

Returns the array sorted in increasing order.

```c
int [] sort(int values[])
float [] sort(float values[])
string [] sort(string values[])
```

Returns a version of the given array sorted in increasing order.

返回一个按递增顺序排序的给定数组的版本。

- argsort and sort use a stable sort.
- Use reverse to reverse the order of the sort.

- argsort and sort 使用一个稳定的排序。
- 使用 reverse 来逆转排顺序。

## 示例

```c
int numbers[] = {5, 2, 90, 3, 1};
int descending_nums[] = reverse (sort (numbers)); // {90, 5, 3, 2, 1}

int arr1[] = {1,4,3};
int arr2[]  = sort(arr1);
printf("%s", arr2);  // {1, 3, 4}

string s[] = {"alice","david","cici"};
string s2[] = sort(s);
printf("%s", s2); // {alice, cici, david}
```
