---
title: 表达式选择器
order: 5
---

# 表达式选择器

## 参数一览

| English           | 中文        |                             |                  |
| ----------------- | ----------- | --------------------------- | ---------------- |
| Based On          | 依据        |                             |                  |
|                   |             | Characters                  | 字符             |
|                   |             | Characters Excluding Spaces | 不包含空格的字符 |
|                   |             | Words                       | 词               |
|                   |             | Lines                       | 行               |
| Amount            | 数量        |                             |                  |
| Expression:Amount | 表达式:数量 |                             |                  |

## Based On 依据

以哪种方式计算文字。

| English                     | 中文             | 图示                                                                                | 说明                                                                                                                         |
| --------------------------- | ---------------- | ----------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| Characters                  | 字符             | ![img](https://mir.yuelili.com/user/AE/text/basic/Based-on2.png) | 一个字符算一个整体，效果对每个字符单独生效。<br />英文: 5 + 1(空格) + 5 = 11 个字符<br />中文: 2+1(空格) + 2 = 5 个字符      |
| Characters Excluding Spaces | 不包含空格的字符 | ![img](https://mir.yuelili.com/user/AE/text/basic/Based-on2.png) | 一个**字符**算一个整体，效果对每个字符单独生效。<br />英文: 5 + 5 = 10 个字符(空格不算)<br />中文: 2+ 2 = 4 个字符(空格不算) |
| Words                       | 词               | ![img](https://mir.yuelili.com/user/AE/text/basic/Based-on2.png) | 一个**单词**算一个整体，效果对每个单词单独生效。<br />中文: 1+ 1 = 2 个词(空格作为词分隔符)                                  |
| Lines                       | 行               | ![img](https://mir.yuelili.com/user/AE/text/basic/Based-on2.png) | 一**行**算一个整体，效果对每行单独生效。<br />1 行+1 行 = 2 行                                                               |

## Amount 数量

默认：100%，100%，100%，

效果影响的数量，百分比显示，可以为负。

举例：当前效果为 Y 位置偏移 100，如果数量=60%，则实际偏移 60px，如果数量=-30%，则实际偏移-30px.

注意：当开启下面的表达式时，这里的值无效！

## Expression:Amount 表达式:数量

表达式选择器一般会与范围选择器一同使用
默认显示为：selectorValue \* textIndex/textTotal
[selectorValue]

上一个选择器的数量值。比如上个范围选择器数量= 80，那么此时 selectorValue =80%

[textIndex]

每个字符的索引值，从 **1** 开始。（如果依据是"词"，则是每个词的索引值）

PS：这里之前写成 0 了，包括推荐的文字教程翻译也错了，后经熊喵的提醒，故改正之。如有其它错误请戳扣扣

![](https://mir.yuelili.com/user/AE/text/basic/textIndex-1-1.jpg)

[textTotal]

总字符个数，上图总字符=9 个

## 示例一：Y 轴逐个偏移

![](https://mir.yuelili.com/user/AE/text/basic/textIndex-2-2.jpg)

### 添加文字效果，Y 轴移动 100 像素

位置 1 的 A：

```javascript
影响量 = selectorValue _ textIndex/textTotal = selectorValue（上个选择器的数量，默认 100%）_ textIndex （索引为 1）/ textTotal（总字符数）=100%\*1/9 = 11%

影响结果 = 影响数值 _ 设置的效果 = 11% _ 100 = 11 （向下偏移 11 像素）
```

位置 5 的 A：

```javascript
影响量 = selectorValue _ textIndex/textTotal = selectorValue（上个选择器的数量，默认 100%）_
textIndex （索引为 5）/ textTotal（总字符数）=100%_5/9 = 55%
影响结果 = 影响数值 _ 设置的效果 = 55% \* 100 = 55 （向下偏移 55 像素）
```

## 示例二：奇数上偏移 偶数下偏移

![](https://mir.yuelili.com/user/AE/text/basic/textIndex-3.png)

表达式如下：

```javascript
if (textIndex % 2 == 0) {
  //如果文字索引为偶数
  selectorValue; //实际影响数值 = 上个选择器影响数值
} else {
  -selectorValue; //实际影响数值 = 上个选择器影响数值的相反数
}
```

## 示例三：自定义文字序号生效

![](https://mir.yuelili.com/2021/07/773620bc58c53cb78b3ba24f069bf2bd.png)

添加一个位置向下的文字效果，然后范围表达式如下。（由[艾露露](https://space.bilibili.com/292690)提供）

结果只有 1/3/4 生效

```javascript
if ([1, 3, 4].indexOf(textIndex) != -1) {
  selectorValue;
} else {
  0 * selectorValue;
}
```
