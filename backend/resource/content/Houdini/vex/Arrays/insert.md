---
title: insert
order: 6
---

# insert

## 描述

Inserts an item, array, or string into an array or string.
将一个元素、数组或字符串插入一个数组或字符串中。

```c
void  insert(string &str, int index, string value)
```

Inserts the `value` into the string `str` at the given `index`.

If `index` is greater than the length of the string, the `value` will simply
be appended to the existing `str`.

将数值插入给定索引的字符串中。

如果 index 大于字符串的长度，value 将被简单地追加到现有的 str 中。

```c
string s = "hello world";

insert(s,2,"xxx");    // 正常，hexxxllo world
insert(s,99,"xxx");  // 正向越界， hello worldxxx
insert(s,-99,"xxx");  // 负向越界， xxxhello world
```

---

```c
void  insert(<type>&array[], int index, <type>value)

void  insert(<type>&array[], int index, <type>values[])
```

Inserts the items or items into the `array` starting at the given `index`.

If `index` is greater than the length current length of the `array`, the
function will fill the gaps with uninitialized values (for example, `0` or the
empty string).

If the `index` is negative, it counts from the _end_ of the string or array you're inserting into. (If the negative number is greater than the string/array length, it is clamped to `0`.)

将一个或多个项目插入到从给定索引开始的数组中。

如果 index 大于 array 的当前长度，该函数将用未初始化的值（例如，0 或空字符串）来填补空隙。

如果 index 是负数，它从你要插入的字符串或数组的末尾开始计算。(如果这个负数大于字符串/数组的长度，它就会被钳制为 0）。

For example, to insert the number `100` as the second-to-last item in an
array:

例如，在一个数组中插入数字 100 作为倒数第二项。

```c
insert (numbers; -1; 100)

int arr[] = {1,2,3};
int arr2[] = {4,5,6};

insert(arr,2,8); // {1, 2, 8, 3}
insert(arr,2,arr2); // {1, 2, 4, 5, 6, 3}

insert(arr,8,5); // {1, 2, 3, 0, 0, 0, 0, 0, 5} ，正向越界用0填充
insert (arr,-8, 5);// {5, 1, 2, 3} 负向越界，直接插开头
```

---

```c
int  insert(dict &dstdict, string dstkey, dict srcdict, string srckey)
```

Copies the value from `srcdict[srckey]` into `dstdict[dstkey]`

This will preserve the underlying type of the value.If the key is not in the source dictionary, it will be removed from the destination dictionary. The result is
`1` if the key existed in the destination dictionary before the update, and
`0` if it did not.

将 srcdict[srckey]中的值复制到 dstdict[dstkey]中。
这将保留值的基本类型。如果键不在源字典中，它将从目标字典中被删除。如果该键在更新前存在于目标字典中，结果是 1，如果不存在，结果是 0。

```c
//示例：字典插入新键值对。前面的是被插（目标字典）的。

dict d1 = set("key1", 3, "key2", 5, "key3", 7);
dict d2 = set("key3", 8, "key4", 6, "key5", 4);

insert(d2,"key3",d21,"key3");  //  都有key3。字典2的key3，被字典1的覆盖，变成7
printf ("%s\n", d2);     // ...key3: 7...

insert(d2,"key2",d21,"key3");  //  字典2没有key2，那么则新增一个key2，值为7
```

---

```c
void  insert(dict &dstdict, dict srcdict)
```

Merges `srcdict` into `dstdict`.Keys that match will be overwritten by values
in the source dictionary.

将 srcdict 合并到 dstdict 中。匹配的键将被源字典中的值所覆盖。

```c
//说明：字典合并

dict d1 = set("key1", 3, "key2", 5, "key3", 7);
dict d2 = set("key3", 8, "key4", 6, "key5", 4);



insert(d2,d1);  // d1 插入 d2, 如果有相同的键，d1会覆盖d2原来的。


printf("%s",d2);
/* 打印结果

{
 "key1":{
  "type":"int",
  "value":3
 },
 "key2":{
  "type":"int",
  "value":5
 },
 "key3":{
  "type":"int",
  "value":7
 },
 "key4":{
  "type":"int",
  "value":6
 },

 "key5":{
  "type":"int",
  "value":4
 }
}*/
```
