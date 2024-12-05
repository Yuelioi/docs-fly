---
title: Key
order: 6
---

# Key 关键帧

当您访问 Key 对象（主要指关键帧和标记）时，可获得 time、index 和 value 属性。标记具体请看下一章

## value

说明：关键帧/标记的值。

类型：数值或数组。

示例：

```javascript
transform.position.key(1).value; //返回本图层位置的第1个关键帧数值
```

## time

说明：关键帧/标记的时间。

返回类型：数值。

示例：

```javascript
transform.position.key(1).time; //返回本图层位置的第1个关键帧时间
```

## index

说明：关键帧/标记的索引。

类型：数值。

示例：

```javascript
transform.position.key(1).index; //返回本图层位置的第1个关键帧索引
```
