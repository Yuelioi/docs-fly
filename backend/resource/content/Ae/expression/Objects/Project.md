---
title: Project
order: 11
---

# Project 工程

## fullPath

说明：项目的绝对文件路径，包括项目文件名。如果项目未保存，将返回空字符串。

类型：字符串

示例：

```javascript
thisProject.fullPath; //返回C:\Users\yl\Desktop\1.aep
```

## bitsPerChannel

说明：返回项目的颜色深度，"项目设置">"颜色管理"中设置。为 8、16 或 32 之一。与脚本项目属性 app.project.bitsPerChannel 等效。

类型：字符串

示例：

```javascript
thisProject.bitsPerChannel; //返回8。代表8位颜色深度
```

## linearBlending

说明：项目灰度系数。"项目设置">"颜色管理"中的"使用 1.0 灰度系数混合颜色"选项的状态。与脚本项目属性 app.project.linearBlending 等效。

类型：布尔值。

示例：

```javascript
thisProject.linearBlending; // 返回false
```

## 示例

以上三个示例：在文字图层的"源文本"建立表达式，效果如下

![](https://mir.yuelili.com/user/source/2020/07/exp-project1-1.png)![](https://mir.yuelili.com/user/source/2020/07/exp-project2-1.png)

Updated on 2021 年 9 月 9 日
