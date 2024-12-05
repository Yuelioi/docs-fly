---
title: TextDocument
order: 8
---

# TextDocument 文本

## TextDocument object

new TextDocument(docText)  
app.project.item(index).layer(index).property("Source Text").value  
描述：TextDocument 对象存储 TextLayer 的 Source Text 属性的值。使用构造函数创建它，并传递要封装的字符串。

示例

这将设置一些源文本的值，并显示一个警报，显示新值

```javascript
var myTextDocument = new TextDocument("HappyCake");
myTextLayer.property("Source Text").setValue(myTextDocument);
alert(myTextLayer.property("Source Text").value);
```

这将设置文本的关键帧值，以随着时间的推移显示不同的单词

```javascript
var textProp = myTextLayer.property("Source Text");
textProp.setValueAtTime(0, newTextDocument("Happy"));
textProp.setValueAtTime(0.33, newTextDocument("cake"));
textProp.setValueAtTime(0.66, newTextDocument("is"));
textProp.setValueAtTime(1, newTextDocument("yummy!"));
```

这会为某些文本设置各种字符和段落设置

```javascript
var textProp = myTextLayer.property("Source Text");
var textDocument = textProp.value;
myString = "Happy holidays!";
textDocument.resetCharStyle();
textDocument.fontSize = 60;
textDocument.fillColor = [1, 0, 0];
textDocument.strokeColor = [0, 1, 0];
textDocument.strokeWidth = 2;
textDocument.font = "Times New Roman PSMT";
textDocument.strokeOverFill = true;
textDocument.applyStroke = true;
textDocument.applyFill = true;
textDocument.text = myString;
textDocument.justification = ParagraphJustification.CENTER_JUSTIFY;
textDocument.tracking = 50;
textProp.setValue(textDocument);
```

属性

### allCaps

textDocument.allCaps

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本图层启用了全大写，则为 true；否则为 true。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：布尔值 只读。

### applyFill

textDocument.applyFill

描述：为 true 时，文本层显示填充。访问 fillColor 实际颜色的属性。如果为 false，则仅显示中风。

类型：布尔值 读/写。

### applyStroke

textDocument.applyStroke

描述：如果为 true，则文字层显示笔触。访问 strokeColor 实际颜色 strokeWidth 及其厚度的属性。如果为 false，则仅显示填充。

类型：布尔值 读/写。

### baselineLocs

textDocument.baselineLocs

::: info 提示
在 After Effects 13.6(CC 2015)中添加了此功能
:::

描述：文本层的基线(x，y)位置。段落文本框中的换行被视为多行。

::: info 提示
如果一行无字符，则开始和结束的 x 和 y 值将是最大浮点值(3.402823466e + 38F)。
:::

类型：浮点值数组，形式为

[
line0.start_x,
line0.start_y,
line0.end_x,
line0.end_y,
line1.start_x,
line1.start_y,
line1.end_x,
line1.end_y,
...
lineN-1.start_x,
lineN-1.start_y,
lineN-1.end_x,
lineN-1.end_y
]

### baselineShift

textDocument.baselineShift

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：此文本层的基线偏移量(以像素为单位)。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：浮点值；只读。

### boxText

textDocument.boxText

描述：如果文本层是段落(有界)文本层，则为 true；否则为 true。否则为假。

类型：布尔值 只读。

### boxTextPos

textDocument.boxTextPos

::: info 提示
在 After Effects 13.2(CC 2014.2)和 After Effects 14(CC2017)中添加了此功能，看来这也是可写的。
:::

描述：该层从一个段落(框)文本层的锚点开始，以像素尺寸的[width，height]数组为坐标。

:::danger
此属性仅适用于段落文本图层。该值仅反映当前时间文本层中的第一个字符。
:::

类型：([X，Y])个位置坐标的数组；只读。

示例

```javascript
// For a paragraph text layer returns [x, y] position from layer anchor point in layer coordinates.
// e.g. approximately [0, -25] with default character panel settings.
var boxTextLayerPos = myTextLayer.sourceText.value.boxTextPos;
```

### boxTextSize

textDocument.boxTextSize

描述：段落(框)文本图层的大小，以像素尺寸的[width，height]数组为单位。

类型：两个整数的数组(最小值为 1)；读/写。

### fauxBold

textDocument.fauxBold

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本层启用了人造粗体，则为 true；否则为 false。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：布尔值 只读。

示例

```javascript
var isFauxBold = myTextLayer.sourceText.value.fauxBold;
```

### fauxItalic

textDocument.fauxItalic

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本层启用了仿斜体，则为 true；否则为 true。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：布尔值 只读。

### fillColor

textDocument.fillColor

描述：文本层的填充颜色，作为浮点值的数组。例如，在一个 8bpc 项目中，红色值为 255 将为 1.0，而在一个 32bpc 项目中，蓝色过度设置值可为 3.2。[r, g, b]

:::danger

该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::

类型：浮点值数组；读/写。[r, g, b]

### font

textDocument.font

描述：文本层的字体由其 PostScript 名称指定。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::

类型：字符串; 读/写。

### fontFamily

textDocument.fontFamily

::: info 提示
在 After Effects 13.1(CC 2014.1)中添加了此功能。
:::

描述：带有字体家族名称的字符串。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：字符串; 只读。

### fontLocation

textDocument.fontLocation

::: info 提示
在 After Effects 13.1(CC 2014.1)中添加了此功能。
:::

描述：字体文件的路径，提供其在磁盘上的位置。
:::danger
不保证会为所有字体类型返回；对于某些字体，返回值可能为空字符串。

该值仅反映当前时间文本层中的第一个字符。
:::

类型：字符串; 只读。

### fontSize

textDocument.fontSize

描述：文本层的字体大小(以像素为单位)。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::
类型：浮点值(包括 0.1 到 1296)；读/写。

### fontStyle

textDocument.fontStyle

::: info 提示
在 After Effects 13.1(CC 2014.1)中添加了此功能。
:::

描述：带有样式信息的字符串，例如“粗体”，“斜体”

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::
类型：字符串; 只读。

### horizo​​ntalScale

textDocument.horizontalScale

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：此文本层的水平比例(以像素为单位)。
:::danger
该值仅反映当前时间文本层中的第一个字符。
:::
类型：浮点值；只读。

示例

var valOfHScale = myTextLayer.sourceText.value.horizontalScale;

### justification

textDocument.justification

描述：文字层的段落对齐方式。

类型：甲 ParagraphJustification 枚举值; 只读。之一：

- ParagraphJustification.LEFT_JUSTIFY
- ParagraphJustification.RIGHT_JUSTIFY
- ParagraphJustification.CENTER_JUSTIFY
- ParagraphJustification.FULL_JUSTIFY_LASTLINE_LEFT
- ParagraphJustification.FULL_JUSTIFY_LASTLINE_RIGHT
- ParagraphJustification.FULL_JUSTIFY_LASTLINE_CENTER
- ParagraphJustification.FULL_JUSTIFY_LASTLINE_FULL

### leading

textDocument.leading

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：文本层的行间距。

:::danger
如果文本图层的每行前导设置都不同，则此属性返回第一行的设置。另外，如果您更改该值，则会将文本图层中的所有行重置为指定的设置。
:::

类型：浮点值；读/写。

示例

```javascript
// This creates a text layer and sets the leading to 100

var composition = app.project.activeItem;
var myTextLayer = comp.layers.addText("Spring\nSummer\nAutumn\nWinter");
var myTextSource = myTextLayer.sourceText;
var myTextDocument = myTextSource.value;
myTextDocument.leading = 100;
myTextSource.setValue(myTextDocument);
```

### pointText

textDocument.pointText

描述：如果文本层是点(无边界)文本层，则为 true；否则为 true。否则为假。

类型：布尔值 只读。

### smallCaps

textDocument.smallCaps

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本图层启用了小写字母，则为 true；否则为 true。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::
类型：布尔值 只读。

### strokeColor

textDocument.strokeColor

描述：文本层的笔触颜色，为[r，g，b]浮点值的数组。例如，在一个 8bpc 项目中，红色值为 255 将为 1.0，而在一个 32bpc 项目中，蓝色过度设置值可为 3.2。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::
类型：浮点值数组[r，g，b]；读/写。

### strokeOverFill

textDocument.strokeOverFill

描述：指示文本图层的填充和描边的渲染顺序。如果为 true，则笔划会显示在填充上方。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::
类型：布尔值 读/写。

### strokeWidth

textDocument.strokeWidth

描述：文本层的笔触粗细(以像素为单位)。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::
类型：浮点值(0 到 1000，含)；读/写。

### subscript

textDocument.subscript

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本层启用了下标，则为 true；否则为 false。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::
类型：布尔值 只读。

### superscript

textDocument.superscript

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：如果文本层启用了上标，则为 true；否则为 false。否则为假。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：布尔值 只读。

### text

textDocument.text

描述：文本层的“源文本”属性的文本值。

类型：字符串; 读/写。

### tracking

textDocument.tracking

描述：文本层字符之间的间距。

:::danger
该值仅反映当前时间文本层中的第一个字符。如果更改此值，它将文本层中的所有字符重置为指定的设置。
:::
类型：浮点值；读/写。

### tsume

textDocument.tsume

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：此文本层的 tsume 值。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::
类型：浮点值；只读。

### verticalScale

textDocument.verticalScale

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：此文本层的垂直比例(以像素为单位)。

:::danger
该值仅反映当前时间文本层中的第一个字符。
:::

类型：浮点值；只读。

方法

### resetCharStyle()

textDocument.resetCharStyle()

描述：恢复“字符”面板中的默认文本字符特征。

参数：无。

返回：无。

### resetParagraphStyle()

描述：恢复“段落”面板中的默认文本段落特征。

参数：无。

返回：无。
