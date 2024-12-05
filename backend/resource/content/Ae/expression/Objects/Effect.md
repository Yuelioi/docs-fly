---
title: Effect
order: 3
---

# Effect 效果

## active

返回类型：布尔值。

如果效果打开（"效果"开关 处于选定状态），则返回 true。

示例

```javascript
thisComp.layer("Solid 2").effect("Bulge").active;
```

## param(name)

类型：属性。

参数：name 是一个字符串。

说明：返回效果中的属性。效果控制里的点始终位于图层空间中。

示例：

```javascript
thisComp.layer("Solid 2").effect("Bulge").param("Bulge Height"); //返回凸出高度值
```

## param(index)

类型：属性。

参数：index 是数值

说明：返回效果中的属性。效果控制里的点始终位于图层空间中。

示例：

```javascript
thisComp.layer("Solid 2").effect("Bulge").param(4); //返回凸出高度值
```

## name

说明：效果属性的名称

类型：字符串

例如：

```javascript
thisComp.layer("Solid 2").effect("Bulge").name; //返回Bulge
```

Updated on 2021 年 9 月 9 日
