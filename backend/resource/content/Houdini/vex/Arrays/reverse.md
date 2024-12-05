---
title: reverse
order: 15
---

# reverse

## 描述

Returns an array or string in reverse order.

```c
string  reverse(string str)
```

Returns a UTF-8 encoded string with the reversed _characters_ (not bytes) from `str`. This is different from what `str[::-1]` returns.

返回一个 UTF-8 编码的字符串，其中包含 str 的反转字符（不是字节）。这与 str[::-1]的返回内容不同。

```c
<type>[] reverse(<type>values[])
```

Returns a reversed copy of the given array.
返回给定数组的反转副本。

## Examples

```c
reverse ("hello") == "olleh";
reverse ({1, 2, 3, 4}) == {4, 3, 2, 1};
```
