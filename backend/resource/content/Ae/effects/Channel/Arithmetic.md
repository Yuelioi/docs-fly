---
title: 算术
order: 1
---

# Arithmetic - 算术

## 简述

算术效果可在图像的红色、绿色和蓝色通道上执行各种简单的数学运算。

此效果适用于 8-bpc 颜色。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00410.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=115&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Arithmetic.png)

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Arithmetic_cn.png)

| Arithmetic         | 算术       | アリスマチック |             |        |              |
| ------------------ | ---------- | -------------- | ----------- | ------ | ------------ |
| Operator           | 运算符     | 演算子         | And         | 与     |              |
|                    |            |                | Or          | 或     |              |
|                    |            |                | Xor         | 亦或   |              |
|                    |            |                | Add         | 相加   | 加算         |
|                    |            |                | Subtract    | 相减   | 減算         |
|                    |            |                | Difference  | 差值   | 差           |
|                    |            |                | Min         | 最小值 | 最小         |
|                    |            |                | Max         | 最大值 | 最大         |
|                    |            |                | Block Above | 上界   | 上のプロック |
|                    |            |                | Block Below | 下界   | 下のプロック |
|                    |            |                | Slice       | 限制   | 制限         |
|                    |            |                | Multiply    | 相乘   | 乗算         |
|                    |            |                | Screen      | 滤色   | スクリーン   |
| Red Value          | 红色值     | 赤の値         |             |        |              |
| Green Value        | 绿色值     | 緑の値         |             |        |              |
| Blue Value         | 蓝色值     | 青の値         |             |        |              |
| Clipping           | 剪切       | クリッピング   |             |        |              |
| Clip Result Values | 剪切结果值 | クリップ結果値 |             |        |              |

## 参数详解

### 运算符(选择算法)

为图像中的每个像素的每个通道原始值和效果取值进行运算：

以下示例 A 为原值 B 为本效果取值 C 为最终值

**-And Or Xor 与、或、异或**

各通道独立进行位逻辑运算

如何运算:拿出你电脑自带的计算器 切换到程序员模式,

以 And 为例 红色通道的 78and26 = 10

78 在二进制下为 ‭01001110‬

26 在二进制下为 ‭00011010‬

And 为同 1 取 1 其余取 0 返回 00001010 也就是 10

其他同理

### 示例

原始图层为 78,43,85 调用值为 26,45,91 时

![](https://cdn.yuelili.com/20220103162713.png)

| 78,43,85 | 26,45,91 | 78=01001110                     |
| -------- | -------- | ------------------------------- |
| And      | 10,41,81 | 10=1010(同 1 得 1 其余情况为 0) |
| Or       | 94,47,95 | Or 是同 0 得 0 其余为 1         |
| Xor      | 84,6,14  | Xor 是 10 01 得 1 其余为 0      |

**-Add /Subtract /Difference /Multiply 相加、相减、差值、相乘**

应用基本数学函数。

C = A + B

C = A - B

C = |A - B|

C= A \* B/256 (见附表的正片叠底)

**-Max 最大值**

效果值和原值取最大。

C = max(A,B)

**-Min 最小值**

效果值和原值取最小

C = min(A,B)

**-Block Above 上界**

原值大于指定值，设置为零；否则，保留原值。

**-Block Below 下界**

原值小于指定值，设置为零；否则，保留原值。

**-限制 Slice**

原值高于指定值，通道值设置为 1.0；否则设置为零。在这两种情况下，均将其他颜色通道的值设置为 1.0。

**-滤色 Screen**

乘以通道值的补色，然后获取结果的补色。结果颜色绝不会比任一输入颜色深。

见附表的滤色

C=255-(A 反相\*B 反相)/255

### 剪切结果值 Clip Result Values

防止所有函数创建超过有效范围的颜色值。如果未选择此选项，则某些颜色值可能会环绕。

详细教程戳:[混合模式计图解](https://www.yuelili.com/?p=5665)

A 为上层 B 为下层 C 为终值

![](https://mir.yuelili.com/user/source/2020/06/Channel-Arithmetic3.png)
