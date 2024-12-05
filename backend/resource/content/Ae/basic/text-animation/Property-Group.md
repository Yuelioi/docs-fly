---
title: 属性组
order: 2
---

# 属性组

## 中英对照

| English                  | 中文           |                 |              |                        |                  |                                                                                                         |
| ------------------------ | -------------- | --------------- | ------------ | ---------------------- | ---------------- | ------------------------------------------------------------------------------------------------------- |
| Enable Per-character 3D  | 启用逐字 3D 化 |                 |              |                        |                  | 启用 3D 属性，比如 3D 位置、3D 旋转等                                                                   |
| Anchor Point             | 锚点           |                 |              |                        |                  | 整个文字的锚点                                                                                          |
| Position                 | 位置           |                 |              |                        |                  | 整个文字的位置                                                                                          |
| Rotation                 | 缩放           |                 |              |                        |                  | 整个文字的旋转                                                                                          |
| Scale                    | 旋转           |                 |              |                        |                  | 整个文字的缩放                                                                                          |
| Opacity                  | 不透明度       |                 |              |                        |                  | 整个文字的不透明度                                                                                      |
| All Transform Properties | 全部变换属性   |                 |              |                        |                  | 应用上面 5 个效果                                                                                       |
| Fill Color               | 填充颜色       |                 |              |                        |                  | 单独应用填充的颜色参数                                                                                  |
|                          |                | RGB             | RGB          |                        |                  |                                                                                                         |
|                          |                | Hue             | 色相         |                        |                  |                                                                                                         |
|                          |                | Saturation      | 饱和度       |                        |                  |                                                                                                         |
|                          |                | Brightness      | 亮度         |                        |                  |                                                                                                         |
|                          |                | Opacity         | 不透明度     |                        |                  |                                                                                                         |
| Stroke Color             | 描边颜色       |                 |              |                        |                  | 单独应用描边的颜色参数                                                                                  |
|                          |                | RGB             | RGB          |                        |                  |                                                                                                         |
|                          |                | Hue             | 色相         |                        |                  |                                                                                                         |
|                          |                | Saturation      | 饱和度       |                        |                  |                                                                                                         |
|                          |                | Brightness      | 亮度         |                        |                  |                                                                                                         |
|                          |                | Opacity         | 不透明度     |                        |                  |                                                                                                         |
| Stroke Width             | 描边宽度       |                 |              |                        |                  |                                                                                                         |
| Tracking                 | 字符间距       |                 |              |                        |                  | 每个字符间的间距                                                                                        |
|                          |                | Tracking Type   | 字符间距类型 |                        |                  | 啥时候会受效果器影响，形状为矩形无影响，以下案例形状为三角形                                            |
|                          |                |                 |              | After&After            | 之前&之后        | IIIII → I I I I I                                                                                       |
|                          |                |                 |              | After                  | 之前             | IIIII → I I I I I                                                                                       |
|                          |                |                 |              | After                  | 之后             | IIIII → I I I I I                                                                                       |
|                          |                | Tracking Amount | 字符间距大小 |                        |                  |                                                                                                         |
| Line Anchor              | 行锚点         |                 |              |                        |                  | 一整行的锚点，有点像段落对齐方式                                                                        |
| Line Spacing             | 行距           |                 |              |                        |                  | 多行之间的行间距`有横向间距+纵向间距                                                                    |     | Character Offset | 字符位移 |                     |              |     |     | 基于 Unicode 码把字符进行偏移，`比如偏移为 2 时，A 变成 C |
| Character Value          | 字符值         |                 |              |                        |                  | 所有文字变成 Unicode 码对应的文字，`比如字符值为 65 时，所有文字都变成 A                                |     |                  |          | Character Alignment | 字符对齐方式 |     |     |                                                           |     |     |     |     |     | Left or Top | 左侧或顶部 |     |     |     |     |     |     | Center | 中间 |     |     |     |     |     |     | Right or Bottom | 右侧或底部 |     |     |     |     |     |     | Adjust Kerning | 调整字偶间距 | 调整字形间距，因为英文字符形状不规则，有宽有窄，排列在一起的疏密不一致，勾选后会尽量一致`AIRPORT：很明显，I 的默认占位好小 |
|                          |                | Character Range | 字符范围     |                        |                  |                                                                                                         |
|                          |                |                 |              | Preserve Case & Digits | 保留大小写及数位 | 英文字母偏移为英文字母，不会越英文字母界，且大小写不变；``数字偏移成数字，不会越数字界，9 偏移 2 变成 1 |
|                          |                |                 |              | Full Unicode           | 完整的 Unicode   | 在几万个 Unicode 码中偏移，还可以偏移成汉字哦，不过会很怪                                               |
| Blur                     | 模糊           |                 |              |                        |                  | 增加模糊效果                                                                                            |

Unicode 码相关知识请自行百度
