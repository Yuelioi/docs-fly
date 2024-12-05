---
title: Mask
order: 9
---

# Mask 蒙版与路径

可以针对以下对象使用表达式读写路径点或顶点的 x 坐标和 y 坐标：  
图层蒙版  
贝塞尔曲线形状  
针对"绘画"和"Roto 笔刷和优化边缘"效果的画笔描边。

## 路径示意图

![](https://mir.yuelili.com/user/AE/expression/a-z/createPath-sample1.png)

## name

说明：路径名称名称

类型：字符串

例如：

```javascript
thisComp.layer("Solid 1").mask("Mask 1").name;
//返回 Mask 1
```

## maskFeather

用法: Mask.maskFeather

说明: 蒙版羽化值

返回: 数值

## maskExpansion

用法: Mask.maskExpansion

说明: 蒙版扩展值

返回: 数值

## invert

用法: Mask.invert

说明: 蒙版是否反转

返回: 布尔值
