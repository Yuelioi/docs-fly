---
title: 调整柔和遮罩
order: 3
---

# Refine Soft Matte - 调整柔和遮罩

## 简述

- 可以用来调整被裁剪过的边缘，例如，用绿幕抠像后。
- 精化细节，如头发，会自动计算并调整裁剪区域。

## 效果展示

## 教程

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Matte-Refine_Soft_Matte.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Matte-Refine_Soft_Matte_cn.png)

| Refine Soft Matte         | 调整柔和遮罩 | ソフトマットを調整     |                       |                |                            |     |     |     |
| ------------------------- | ------------ | ---------------------- | --------------------- | -------------- | -------------------------- | --- | --- | --- |
| Calculate Edge Details    | 计算边缘细节 | エッジ詳細を計算       |                       |                |                            |     |     |     |
| Additional Edge Radius    | 其他边缘半径 |                        |                       |                |                            |     |     |     |
| View Edge Region          | 查看边缘区域 | エッジ領域を表示       |                       |                |                            |     |     |     |
| Smooth                    | 平滑         | スムージング           |                       |                |                            |     |     |     |
| Feather                   | 羽化         | 境界のぼかし           |                       |                |                            |     |     |     |
| Contrast                  | 对比度       | コントラスト           |                       |                |                            |     |     |     |
| Shift Dae                 | 移动边缘     | エッジをシフト         |                       |                |                            |     |     |     |
| Chatter Reduction         | 震颤减少     | エッジのガタつきを軽減 | Off                   | 关闭           | 閉じる                     |     |     |     |
|                           |              |                        | More Detailed         | 更详细         |                            |     |     |     |
|                           |              |                        | Smoother (Slower)     | 更平滑（更慢） |                            |     |     |     |
| Reduce Chatter            | 减少震颤     | エッジのガタつき軽減値 |                       |                |                            |     |     |     |
| More Motion Blur          | 更多运动模糊 | 他のモーションブラー   |                       |                |                            |     |     |     |
| Motion Blur               | 运动模糊     | モーションブラー       |                       |                |                            |     |     |     |
|                           |              |                        | Samples Per Frame     | 每帧样本       | フレームあたりのサンプル数 |     |     |     |
|                           |              |                        | Shutter Angle         | 快门角度       | シャッター角度             |     |     |     |
|                           |              |                        | Higher Quality        | 较高品质       | より高品質                 |     |     |     |
| Decontaminate Edae Colors | 净化边缘颜色 | エッジカラーを除去     |                       |                |                            |     |     |     |
|                           |              |                        | Amount                | 净化数量       | 除去量                     |     |     |     |
|                           |              |                        | Extend Where Smoothed | 扩展平滑的地方 | 滑らかさを拡張             |     |     |     |
|                           |              |                        | Increase Radius       | 增加净化半径   | 除去する半径を増加         |     |     |     |
|                           |              |                        | View Map              | 查看净化地图   | 除去マップを表示           |     |     |     |

## 参数详解

### 其他边缘半径

你可以调整边缘调整的宽度。  
数值越高，从边缘开始应用的效果就越广。

![](https://cdn.yuelili.com/20211225204815.png)

### 震颤减少

你可以调整播放过程中的边缘响声的处理。  
默认设置为[关闭]，这意味着不应用边缘震荡。

**更平滑（更慢）。**  
渲染时间稍长，但更流畅。

技巧：处理的程度由边缘响声减少值控制。 [你可以为 "高级 "指定最多 100%，为 "平滑 "指定最多 400%（较慢）。

### 移除

你可以调整去除边缘的背景颜色

### 增加半径

调整背景颜色的去除程度。  
数值越高，要移除的区域越广。

## 案例：去除绿幕抠像边缘杂色

A 图为 keylight 抠出来的，有很多毛糙边缘。

B 图 使用本效果

![](https://cdn.yuelili.com/20211225211456.png)

![](https://cdn.yuelili.com/20211225211339.png)
