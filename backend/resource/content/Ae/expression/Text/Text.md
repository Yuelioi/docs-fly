---
title: Text
order: 3
---

# Text

## Text

指的是文字对象

After Effects 使用 Get 和 Set 函数以在表达式引擎读写文本属性值，这样您就可以使用表达式来驱动和设置字体样式，例如跨文本图层的文本属性。

[戳我查看官方文档](https://helpx.adobe.com/cn/after-effects/using/expressions-text-properties.html#%E5%85%B6%E4%BB%96%E8%A6%81%E7%82%B9)

说明：表达式均在文字图层-源文本中书写，请参考示例 1 的格式说明，案例格式统一使用格式 2

示例 1：大致格式说明

```javascript
style.fontSize; //格式1：本图层使用，极致简约
text.sourceText.style.fontSize; //格式2：本图层使用
thisLayer.text.sourceText.style.fontSize; //格式3：还是本图层使用
comp("Comp 1").layer(1).text.sourceText.style.fontSize; //格式4：跨合成使用
```

示例 2：多个样式叠加设置，只需用.进行连接即可

```javascript
style.setFillColor(hexToRgb("00FF00")).setApplyStroke(1).setStrokeWidth(5).setStrokeColor([1, 0, 0]).setFontSize(200); //设置以上5种样式
```

## sourceText

使用方法: Text.sourceText

说明: 文字内容(文字图层的)

类型: 源文本内容的字符串(AE17.0+)

## Font…

使用方法: Text.Font…

说明: 文字样式. 启动一个对话框窗口，供用户指定字体名称和粗细。 选择后，内部字体名称将作为字符串注入表达式编辑器。

类型: 字符串
