---
title: TextLayer
order: 7
---

# TextLayer 文字图层

## TextLayer object

app.project.item(index).layer(index)

描述：TextLayer 对象表示合成中的文字图层。使用 LayerCollection.addText()方法创建。可以通过索引或名称在项目的图层集中对其进行访问。

父级关系：TextLayer 是 AVLayer 的子类，AVLayer 是 Layer 的子类。使用 TextLayer 时，AVLayer 和 Layer 的所有方法和属性均可用。

### AE 属性

TextLayer 没有定义其他属性，除了从 AVLayer 继承的属性外，还具有以下 AE 属性和属性组：

| 脚本名              | 中文介绍     |                        |            |
| ------------------- | ------------ | ---------------------- | ---------- |
| Text                | 文字         |                        |            |
| SourceText          | 源文本       |                        |            |
| PathOptions         | 路径选项     |                        |            |
| Path                | 路径         |                        |            |
| ReversePath         | 反转路径     |                        |            |
| PerpendicularToPath | 垂直于路径   |                        |            |
| ForceAlignment      | 强制对齐     |                        |            |
| FirstMargin         | 首字边距     |                        |            |
| LastMargin          | 末字边距     |                        |            |
| MoreOptions         | 更多选项     |                        |            |
|                     |              | AnchorPointGrouping    | 锚点分组   |
|                     |              | GroupingAlignment      | 分组对齐   |
|                     |              | Fill&Stroke            | 填充描边   |
|                     |              | InterCharacterBlending | 字符间混合 |
| Animators           | 动画制作工具 |                        |            |

### 未使用的属性

TimeRemapandMotionTrackers 从 AVLayer 继承的属性不适用于文本层，并且不使用其相关的 AVLayer 属性：

时间重映射、轨道遮罩

- canSetTimeRemapEnabled
- timeRemapEnabled
- trackMatteType
- isTrackMatte
- hasTrackMatte
