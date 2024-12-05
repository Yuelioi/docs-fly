---
title: 旋涡条纹
order: 30
---

# Smear - 旋涡条纹

## 简述

通过 2 个蒙版，一个负责作用范围，一个负责发力区域，进行扭曲、旋转扭曲、位移扭曲

## 效果展示

## 教程

1.25~

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=103&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Smear.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Smear_cn.png)

| Smear                | 旋涡条纹 | 波形ワープ         |                   |            |              |
| -------------------- | -------- | ------------------ | ----------------- | ---------- | ------------ |
| Source Mask          | 源蒙版   | リシェープ元マスク |                   |            |              |
| Boundary Mask        | 边界蒙版 | 境界マスク         |                   |            |              |
| Mask offset          | 蒙版位移 | マスクオフセット   |                   |            |              |
| Mask Rotation        | 蒙版旋转 | マスク回転         |                   |            |              |
| Mask scale           | 蒙版缩放 | マスクスケール     |                   |            |              |
| Percent              | 百分比   | パーセント         |                   |            |              |
| Elasticity           | 弹性     | 弾力性             | Stiff             | 生硬       |              |
|                      |          |                    | Less stiff        | 不那么生硬 |              |
|                      |          |                    | Below normal      | 标准以下   |              |
|                      |          |                    | Normal            | 正常       | 標準         |
|                      |          |                    | Absolutely Normal | 绝对正常   |              |
|                      |          |                    | Above Average     | 平均以上   |              |
|                      |          |                    | Loose             | 松散       |              |
|                      |          |                    | Liquid            | 液体       |              |
|                      |          |                    | Super Liquid      | 超级流体   |              |
| Interpolation Method | 计算密度 | 補間法             | Discrete          | 分离       | 不連続       |
|                      |          |                    | Linear            | 线性       | リニア       |
|                      |          |                    | Smooth            | 平滑       | スムージング |

## 参数详解

![](https://cdn.yuelili.com/20211223005436.png)

源蒙版：扭曲中心区域

边界蒙版：扭曲边界

蒙版位移：中心区域移动

蒙版旋转：中心区域旋转

蒙版缩放：中心区域缩放

百分比：扭曲程度，0 为不扭曲，可以做动画。

![](https://cdn.yuelili.com/20211223012858.png)

弹性：其实就是画质，越好品质越高，渲染越慢。用正常就差不多了（3.30）

计算密度：如果做动画，普通 F9 对它的动画不生效，只能从这里设置动画方式（3.50）

- 分离：提高精度，计算关键帧缓动
- 线性： 忽视关键帧的缓动
- 平滑：不同关键帧平滑缓动

![](https://cdn.yuelili.com/20211223012938.gif)

## 案例
