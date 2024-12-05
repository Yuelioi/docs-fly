---
title: slice
order: 16
---

# slice

## 描述

Slices a sub-string or sub-array of a string or array. 切分数组。不过建议用[ : :] 。见最后示例

```c
string  slice(string s, int start, int end)
string  slice(string s, int start, int end, int step)
```

Extracts a string from a larger string.

从一个较大的字符串中提取一个字符串。

```c
<type>[] slice(<type>s[], int start, int end)
<type>[] slice(<type>s[], int start, int end, int step)
```

Extracts a sub-array from a larger array.

从一个较大的数组中提取一个子数组。

```c
string  slice(string s, int hasstart, int start, int hasend, int end, int hasstep, int step)

<type>[] slice(<type>array[], int hasstart, int start, int hasend, int end, int hasstep, int step)
```

A general-purpose signature to support the slicing syntax. If `hasstart` is `0`, it ignores `start` and uses `0`. If `hasend` is `0` it ignores `end` and
uses the length of the array. If `hasstep` is `0` it ignores `step` and uses `1`.

一个通用的签名，支持切片语法。如果 hasstart 是 0，它忽略 start 并使用 0。如果 hasend 是 0，它忽略 end 并使用数组的长度。如果 hasstep 是 0，它忽略 step，使用 1。

## 说明

- This is the function equivalent of using `value[start:end:step]` slicing syntax.
- If `start` or `end` are negative, they count from the end of the string/array.
- The computed range is clamped to the bounds of the original string/array.
- If step is zero, the function returns an empty string/array.
- If step is negative, the items are returned in reverse, and `end` should be less than `start`.

- 等同于 value[start :end: step]切分数。
- 如果 startorend 是负数，它们从字符串/数组的末端开始计算。
- 计算出的范围被夹在原始字符串/数组的边界内。
- 如果 step 为零，该函数返回一个空的字符串/数组。
- 如果步长为负数，则以相反的方式返回项目，并且 end 应该小于 start。

## 示例

```c
int[] nums = {10, 20, 30, 40, 50, 60};

slice(nums, 1, 3) // {20, 30}                  等价于nums[1:3]
slice(nums, 1, -1) // {20, 30, 30, 40, 50};    等价于nums[1:-1]
slice(nums, 0, len(nums), 2) // {20, 40, 60};  等价于 nums[0:len(nums):2]
slice(nums, 0, 0, 0, 0, 1, 2)// {20, 40, 60};  等价于 nums[::2]
```
