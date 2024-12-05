---
title: SolidSource
order: 4
---

# SolidSource 纯色源

## 纯色源对象

app.project.item(index).mainSource  
app.project.item(index).proxySource

描述：纯色源（SolidSource）对象代表纯色素材源。

父级关系：纯色源是素材源（FootageSource）的子类。因此素材源的方法和属性均可用。

## 属性

### color

solidSource.color

描述：纯色的颜色，红色，绿色和蓝色。

类型：三个浮点值的数组；读/写。[R, G, B]均为 0.0~1.0
