---
title: CC弯曲
order: 3
---

# CC Bend It - CC弯曲

## 简述

可以弯曲图层

## 效果展示

## 教程

0:00 ~ 7:20

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=56&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

### ![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-CC_Bend_It.png)

参数详解 #

### Bend 弯曲

决定围绕 **起始** 位置弯曲 **程度**

![](https://cdn.yuelili.com/20211222110815.png)

**如果** 效果跳出图层范围，将被切断，使用效果 ** _范围扩散_** 效果，可以避免。

![](https://cdn.yuelili.com/20211222110945.png)

### **Start，End (** 开始与结束)

开始与结束位置。还决定了开始和结束位置的 **宽度** 。

![](https://cdn.yuelili.com/20211222111104.png)

### Render Prestart 渲染预启动

**如何** 绘制" **开始** "以外的区域。

None：不要了。

Static：固定，也就是不变化

Bend：弯曲。

Mirror：镜像。

![](https://cdn.yuelili.com/20211222111159.png)

### Distort 扭曲

决定 " **结束** "以外的区域是否绘制。相当于是否扩展“竖向”空间。视频 3:20  
Legal（常规）是 **不** 进行绘制，Extended （扩展）则进行绘制。

![](https://cdn.yuelili.com/20211222111542.png)

## 案例

视频 6.20

给举哑铃增加动画效果

![](https://cdn.yuelili.com/20211222132132.png)

### 其他

普通图层，正常使用，但是对于矢量图层（比如形状图层），如果起始、结束位置不跟随形状本身，可以用表达式链接一下（视频：4：02）

起始、结束位置表达式

```javascript
value + transform.position;
```

### 友情效果

[CC Bend](https://www.yuelili.com/?p=15187) ：同样是扭曲图层

[范围扩散](https://www.yuelili.com/?p=15386) ：可以打破效果作用范围（默认是图层范围）
