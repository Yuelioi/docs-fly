---
title: 更多选项
order: 1
---

# 更多选项

![](https://mir.yuelili.com/2021/07/cada62ab21ba4f63082af50884e87594.png)

![img](https://mir.yuelili.com/2021/07/1abeca39d6a7e7107a7f53457eef7e5f.png)

## 中英对照

| 序号 | 英文                     | 中文       |                            |                        |                       |                          |
| ---- | ------------------------ | ---------- | -------------------------- | ---------------------- | --------------------- | ------------------------ |
| 0    | 反转路径                 | 晚点更新   |                            |                        |                       |                          |
|      |                          | 垂直于路径 |                            |                        |                       |                          |
|      |                          | 强制对齐   |                            |                        |                       |                          |
|      |                          | 首字边距   |                            |                        |                       |                          |
|      |                          | 末字边距   |                            |                        |                       |                          |
|      | More Options             | 更多选项   |                            |                        |                       |                          |
| ①    | Anchor Point Grouping    | 锚点分组   | Character                  | 字符                   |                       |                          |
|      |                          |            | Word                       | 词                     |                       |                          |
|      |                          |            | Line                       | 行                     |                       |                          |
|      |                          |            | All                        | 全部                   |                       |                          |
| ②    | Grouping Alignment       | 分组对齐   | 0.0,0.0%                   | 0.0,0.0%               |                       |                          |
| ③    | Fill & stroke            | 填充和描边 | Per Character Palette      | 每字符调板             |                       |                          |
|      |                          |            | All Fills Over All Strokes | 全部填充在全部描边之上 |                       |                          |
|      |                          |            | All Strokes Over All fills | 全部描边在全部填充之上 |                       |                          |
|      | Inter-Character Blending | 字符间混合 | Normal                     | 正常                   | 一堆混合模式 不截图了 | 混合模式本站有超详细教程 |

## ① 锚点分组

设置锚点的分类方式，锚点就是下图的小红叉 x

| Character | 字符 | ![](https://mir.yuelili.com/user/AE/text/basic/Anchor-Grouping-Character.png) | 分组为字符时, 锚点默认位置在每个字符的中下部                 |
| --------- | ---- | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------ |
| Word      | 词   | ![](https://mir.yuelili.com/user/AE/text/basic/Anchor-Grouping-1.png)         | 分组为词时, 锚点默认位置在每个词的中下部(区分是不是词：空格) |
| Line      | 行   | ![](https://mir.yuelili.com/user/AE/text/basic/Anchor-Grouping-Line.png)      | 分组为行时, 锚点默认位置在每行的中下部(区分是不是行：回车)   |
| All       | 全部 | ![](https://mir.yuelili.com/user/AE/text/basic/Anchor-Grouping-All.png)       | 分组为全部时, 锚点默认位置在全部字符的中心                   |

## ② 分组对齐

|                    | 0.0,0.0% | 40.0,-40.0%                                                                              |
| ------------------ | -------- | ---------------------------------------------------------------------------------------- |
| Grouping Alignment | 分组对齐 | ![](https://mir.yuelili.com/user/AE/text/basic/Anchor-Grouping-1.png) |

## ③ Fill & stroke 填充和描边

![](https://mir.yuelili.com/2021/07/d33e633d5b391a6225bac8a8f7e10b65-1.png)

每字符调板: 每个字的填充描边设置，可以在字符面板里自定义

![](https://mir.yuelili.com/2021/07/a8a291e943c11eaf0207bb84a2439b5d-1.png)

全部填充在全部描边之上：全局设置！所有文字填充都在描边上方

![](https://mir.yuelili.com/2021/07/a0b5c4292153172a85922717d60437d2-1.png)

全部描边在全部填充之上：全局设置！所有文字描边都在填充上方

![](https://mir.yuelili.com/2021/07/77ef02f60b937db6dfa046aeb49861df-1.png)

填充是从文字边缘向 **内** 填充, 而描边是文字边缘向 **内外** 扩散, 因此会有交叉区域

此选项决定谁上谁下

![](https://mir.yuelili.com/user/AE/text/basic/Fill+stroke.png)
