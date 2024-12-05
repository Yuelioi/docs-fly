---
title: Text-Style
order: 2
---

# Text-Style

## setText()

用法：style.setText(value)

说明：设置文本的内容

参数：value 为字符串

示例：

```javascript
text.sourceText.style.setText("hhhh"); //此时屏幕上会显示 hhhh
```

## fontSize

用法：style.fontSize

说明：获取文本的字体大小

类型：数值

示例：

```javascript
text.sourceText.style.fontSize; //比如当前我的字体大小为：24
```

## setFontSize()

用法：style.setFontSize(fontsize)

说明：设置文本的字体大小

参数：fontsize 为数值

示例：

```javascript
text.sourceText.style.setFontSize(50); //比如当前我的字体大小为：24。使用表达式后变成：50。用此表达式也可做文字逐步变大的效果
```

## font

用法：style.font

说明：文本的字体名称。不过是字体英文名

类型：字符串

示例：

```javascript
text.sourceText.style.font; //返回当前字体：SourceHanSansTC-Medium （中文叫 思源黑体)
```

## setFont()

用法：setFont(font)

说明：设置文本的字体

参数：font 为字符串

示例：

```javascript
text.sourceText.style.setFont("Simsum"); //把当前字体设置为：Simsum
```

## isFauxBold

用法：style.isFauxBold

说明：检测当文本是否为仿粗体。单击下方按钮进行设置

类型：布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/faux-bold.png)

```javascript
text.sourceText.style.isFauxBold; //是仿粗体则类型：truetrue
```

## setFauxBold()

用法：style.setFauxBold(bold)

说明：字体仿粗体的设置。仿粗体见下图

参数：bold 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/faux-bold.png)

```javascript
text.sourceText.style.setFauxBold(true); //效果：把当前文本设置为仿粗体
```

## isFauxItalic

用法：style.isFauxItalic

说明：文本是否为仿斜体

类型：布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/faux-italic.png)

```javascript
text.sourceText.style.isFauxItalic; //是仿斜体则类型：true
```

## setFauxItalic()

用法：style.setFauxItalic(value)

说明：设置文本的仿斜体

参数：value 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/faux-italic.png)

```javascript
text.sourceText.style.setFauxItalic(true); // 效果：把当前文本设置为仿斜体
```

## isAllCaps

用法：style.isAllCaps

说明：文本是否为大写化

类型：布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/all-caps.png)

```javascript
text.sourceText.style.isAllCaps; //当前文本为大写化，故而类型：true
```

## setAllCaps()

用法：style.setAllCaps(value)

说明：把文本设置为大写化

参数：value 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/all-caps.png)

```javascript
text.sourceText.style.setAllCaps(true); //原文字：Abc。设置表达式后：ABC
```

## isSmallCaps

用法：style.isSmallCaps

说明：文本是否为小写化

类型：布尔值

示例：见下图示例

![](https://mir.yuelili.com/user/AE/expression/a-z/small-caps.png)

![](https://mir.yuelili.com/user/AE/expression/a-z/small-caps2.png)

```javascript
text.sourceText.style.isSmallCaps; //原文本：Abc 并非全为小写，故类型：false
```

## setSmallCaps()

用法：style.setSmallCaps(value)

说明：文本设置为小写化

参数：value 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/small-caps.png)

![](https://mir.yuelili.com/user/AE/expression/a-z/small-caps2.png)

```javascript
text.sourceText.style.setSmallCaps(true); //原文本：Abe。设置表达式后：ABE(其中BE只有正常字母一半高）
```

## tracking

用法：style.tracking

说明：文本的字间距

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tracking.png)

```javascript
text.sourceText.style.tracking; //比如返回80
```

## setTracking()

用法：style.setTracking(value)

说明：设置文本的字间距

参数：value 为数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tracking.png)

```javascript
text.sourceText.style.setTracking(80); //设置为80
```

## leading

用法：style.leading

说明：文本的段间距

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/leading.png)

```javascript
text.sourceText.style.leading; //比如返回50
```

## setLeading()

用法：style.setLeading(value)

说明：设置文本的段间距

参数：value 为数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/leading.png)

```javascript
text.sourceText.style.setLeading(50); //设置为50
```

## autoLeading

用法：style.autoLeading

说明：文本是否为自动间距

类型：布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/auto-leading.png)

```javascript
text.sourceText.style.autoLeading; //是则类型：true
```

## setAutoLeading()

用法：style.setAutoLeading(value)

说明：设置文本为自动间距

参数：value 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/auto-leading.png)

```javascript
text.sourceText.style.setAutoLeading(true);
```

## baselineShift

用法：style.baselineShift;

说明：文本基线偏移值

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/baseline-shift.png)

```javascript
text.sourceText.style.baselineShift; //比如50
```

## setBaselineShift()

用法：style.setBaselineShift(value)

说明：设置文本基线偏移值

参数：value 为布尔值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/baseline-shift.png)

```javascript
text.sourceText.style.setBaselineShift(50); //比如设置为80
```

## applyFill

用法：style.applyFill

说明：文字是否有填充

类型：布尔值

示例：

```javascript
text.sourceText.style.applyFill; //有则类型：true
thisComp.layer("月离离真帅").text.sourceText.getStyleAt(2).applyFill; //当前合成的"月离离真帅"文字图层的第3个文字（真）有填充效果，故返回True
```

## setApplyFill()

用法：style.setApplyFill(value)

说明：设置文本是否填充

参数：value 为布尔值

示例：

```javascript
text.sourceText.style.setApplyFill(false); //关闭文字填充
```

## fillColor

用法：style.fillColor

说明：文本的填充颜色

类型：RGB 三维数组，不过颜色的 255 在这里为 1

示例：

```javascript
text.sourceText.style.fillColor; //比如填充为红色时，类型：[1,0,0]，也就是红色的[255,0,0]
```

## setFillColor()

用法：style.setFillColor(value)

说明：设置文本的填充颜色

参数：value 为 RGB 三维数组，不过颜色的 255 在这里为 1

示例：

```javascript
text.sourceText.style.setFillColor([0, 1, 0]); //比如我想把文字填充改为绿色 [0,255,0] → [0,1,0]。注意需要先开启填充
```

## applyStroke

用法：style.applyStroke

说明：文本是否有描边。

类型：布尔值

示例：

```javascript
text.sourceText.style.applyStroke; //有描边则类型：true
thisComp.layer("月离离真帅").text.sourceText.getStyleAt(2).applyStroke; //当前合成的"月离离真帅"文字图层的第3个文字（真）有描边，故返回True
```

## setApllyStroke()

用法：style.setApllyStroke(value)

说明：设置文本的描边

参数：value 为布尔值

示例：

```javascript
text.sourceText.style.setApplyStroke(true); //使文字增加描边效果
```

## strokeColor

用法：style.strokeColor

说明：文本的描边颜色

类型：RGB 三维数组，只不过颜色的 255 在这里为 1

示例：

```javascript
text.sourceText.style.strokeColor; //比如我的描边是红色，则返回[1,0,0]
```

## setStrokeColor()

用法：style.setStrokeColor(value)

说明：设置文本的描边颜色

参数：value 为 RGB 三维数组，不过颜色的 255 在这里为 1

示例：

```javascript
text.sourceText.style.setStrokeColor([0, 0, 1]); //把描边设置为蓝色，注意需要先开启描边
```

## strokeWidth

用法：style.strokeWidth

说明：获取描边宽度

类型：数值

示例：

```javascript
text.sourceText.style.strokeWidth; // 当前描边宽度为50
```

## setStrokeWidth()

用法：style.setStrokeWidth(value)

说明：设置描边宽度

参数：value 为数值

示例：

```javascript
text.sourceText.style.setStrokeWidth(20); //把描边设置为20宽，注意需要先开启描边
```
