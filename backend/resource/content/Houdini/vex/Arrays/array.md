---
title: _Array总表
order: 1
---

# _Array总表

## 增

### append 追加值

Adds an item to an array or string.

向字符串尾部追加新的字符串

向数组尾部追加元素、其他数组

```c
void  append(string &array, string value)
```

```c
void  append(<type>&array[], <type>value)
```

```c
void  append(<type>&array[], <type>values[])
```

### array 创建

```c
<type>[] array(...)
```

### insert 插入

```c
void  insert(string &str, int index, string value)
```

```c
void  insert(<type>&array[], int index, <type>value)
```

```c
void  insert(<type>&array[], int index, <type>values[])
```

```c
int  insert(dict &dstdict, string dstkey, dict srcdict, string srckey)
```

```c
void  insert(dict &dstdict, dict srcdict)
```

### push 追加

```c
void  push(<type>&array[], <type>value)
```

```c
void  push(<type>&array[], <type>values[])
```

## 删

### pop 移除元素

```c
<type> pop(<type>&array[])
```

```c
<type> pop(<type>&array[], int index)
```

### removeindex 基于索引移除

```c
<type> removeindex(<type>&array[], int index)
```

```c
int  removeindex(dict &dictionary, string index)
```

### removevalue 基于值移除

```c
int  removevalue(<type>&array[], <type>value)
```

## 查

### foreach 遍历

```c
    foreach ([element_type] value; array) {

}

    foreach (index, value; array) statement;

    foreach (int index; element_type value; array) statement;
```

### isvalidindex 判断索引越界

```c
int  isvalidindex(<type>&array[], int index)
```

```c
int  isvalidindex(string str, int index)
```

```c
int  isvalidindex(dict d, string key)
```

### len 元素个数

```c
int  len(<vector>v)
```

```c
int  len(<matrix>m)
```

```c
int  len(<type>array[])
```

```c
int  len(string s)
```

```c
int  len(dict d)
```

## 排序

### argsort 排序

```c
int [] argsort(<type>value[])
```

### reorder 重新排序

```c
string  reorder(string value, int indices[])
```

```c
<type>[] reorder(<type>values[], int indices[])
```

### resize 重新规划尺寸

```c
void  resize(<type>&array[], int size)
```

```c
void  resize(<type>&array[], int size, <type>val)
```

### reverse 反转

```c
string  reverse(string str)
```

```c
<type>[] reverse(<type>values[])
```

### slice 分割

```c
string  slice(string s, int start, int end)
```

```c
string  slice(string s, int start, int end, int step)
```

```c
<type>[] slice(<type>s[], int start, int end)
```

```c
<type>[] slice(<type>s[], int start, int end, int step)
```

`string slice(string s, int hasstart, int start, int hasend, int end, int hasstep, int step)`

`<type>[] slice(<type>array[], int hasstart, int start, int hasend, int end, int hasstep, int step)`

### sort 排序

```c
int [] sort(int values[])
```

```c
float [] sort(float values[])
```

```c
string [] sort(string values[])
```

### upush 用于 SIMD 数组

```c
void  upush(<type>&array[], <type>value)
```
