---
title: 匹配颗粒
order: 4
---

# Match Grain - 匹配颗粒

## 简述

匹配颗粒效果可匹配两个图像之间的杂色。此效果对合成和蓝屏/绿屏工作特别有用。匹配颗粒效果只能添加杂色，不能移除杂色，因此如果目标图像的杂色比源图像多，则不能实现精确匹配。在此情况下，可以先使用移除颗粒效果清理目标图像，然后将匹配颗粒效果应用到结果，以此实现完美匹配。

匹配颗粒效果先从杂色采样这一步开始。基本上，通过合成整帧新杂色来匹配杂色样本。您可以在此效果应用到新图象之前以许多方式修改杂色，例如，从某图像复制杂色，在将杂色应用到其他图像之前，使其更大、更红。

与添加颗粒很多参数大差不错。（请参阅[**添加颗粒**](https://www.yuelili.com/docs/ae-effect/add-grain/)效果。）

## 效果展示

## 教程

## 中英日对照

![](https://cdn.yuelili.com/20220103221225.png) |
![](https://cdn.yuelili.com/20220103221231.png)  

|Match Grain | 匹配颗粒 | グレイン（マッチ） | | | | | ||
|---|---|---|---|---|---|---|---|---|
|Viewing Mode | 查看模式 | 表示モード | Preview | 预览 | プレビュー | | ||
|| | | Noise Samples | 杂色样本 | | | ||
|| | | Compensation Samples | 补偿范例 | | | ||
|| | | Blending Matte | 混合遮罩 | | | ||
|| | | Final Output | 最终输出 | | | ||
|Noise Source Layer | 杂色源图层 | ノイズソースレイヤー | | | | | ||
|Preview Region | 预览区域 | プレビュー範囲 | Center | 中心 | 中心 | | ||
|| | | Width | 宽度 | 幅 | | ||
|| | | Height | 高度 | 高さ | | ||
|| | | Show Box | 显示方框 | ボックスを表示 | | ||
|| | | Box Color | 方框颜色 | ボックスのカラー | | ||
|Compensate for Existing Noise | 补偿现有杂色 | 既存ノイズを補正 | | | | | ||
|Tweaking | 微调 | 調整 | | | | | ||
|| | | Intensity | 强度 | 密度 | | ||
|| | | Size | 大小 | サイズ | | ||
|| | | Softness | 柔和度 | 柔らかさ | | ||
|| | | Aspect Ratio | 长宽比 | 縦横比 | | ||
|| | | Channel Intensities | 通道强度 | チャンネルの強度 | | ||
|| | | | | | Red Intensity | 红色强度 ||
|| | | | | | Green Intensity | 绿色强度 ||
|| | | | | | Blue Intensity | 蓝色强度 ||
|| | | Channel Size | 通道大小 | チャンネルのサイズ | | ||
|| | | | | | Red Size | 红色大小 ||
|| | | | | | Green Size | 绿色大小 ||
|| | | | | | Blue Size | 蓝色大小 ||
|Color | 颜色 | カラー | | | | | ||
|| | | Monochromatic | 单色 | モノクロ | | ||
|| | | Saturation | 饱和度 | 彩度 | | ||
|| | | Tint Amount | 色调量 | 色合いの量 | | ||
|| | | Tint Color | 色调颜色 | 色合いのカラー | | ||
|Application | 应用 | 効果 | | | | | ||
|| | | Blending Mode | 混合模式 | 描画モード | | ||
|| | | Shadows | 阴影 | シャドウ | | ||
|| | | Midtones | 中间调 | ミッドトーン | | ||
|| | | Highlights | 高光 | ハイライト | | ||
|| | | Midpoint | 中点 | 中間点 | | ||
|| | | Channel Balance | 通道平衡 | チャンネルバランス | | ||
|| | | | | | Red Shadows | 红色阴影 ||
|| | | | | | Red Midtones | 红色中间调 ||
|| | | | | | Red Highlights | 红色高光 ||
|| | | | | | Green Shadows | 绿色阴影 ||
|| | | | | | Green Midtones | 绿色中间调 ||
|| | | | | | Green Highlights | 绿色高光 ||
|| | | | | | Blue Shadows | 蓝色阴影 ||
|| | | | | | Blue Midtones | 蓝色中间调 ||
|| | | | | | Blue Highlights | 蓝色高光 ||
|Sampling | 采样 | サンプリング | | | | | ||
|| | | Source Frame | 源帧 | ソースフレーム | | ||
|| | | Number of Samples | 样本数量 | サンプル数 | | ||
|| | | Sample Size | 样本大小 | サンプルサイズ | | ||
|| | | Sample Box Color | 采样框颜色 | サンプルボックスのカラー | ||
|| | | Sample Selection | 样本选择 | サンプルの選択 | Automatic | 自动 | オートメーション|
|| | | | | | Manual | 手动 | カスタム|
|| | | Noise Sample Points | 杂色样本点 | ノイズサンプルポイント | ||
|| | | Intensity | 强度 | 密度 | | ||
|Animation | 动画 | アニメーション | | | | | ||
|| | | Animation Speed | 动面速度 | | | ||
|| | | Animate Smoothly | 动画流畅 | 滑らかなアニメーション | ||
|| | | Random Seed | 随机植入 | ランダムシード | | ||
|Blend with Original | 与原始图像混合 | 元の画像とブレンド | | | | | ||
|| | | Amount | 数量 | 量 | | ||
|| | | Combine Match and Mask Using | 结合匹配和蒙版使用 | マッチとマスクを合成するモード | Multiply|
| 相乘 | 乗算  
|| | | | | | Screen | 滤色 | スクリーン|
|| | | Blur Matte | 模模遮罩 | | | ||
|| | | Color Matching | 颜色匹配 | カラーマッチ | | ||
|| | | | | | Match Color Using | 匹配颜色使用 ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | Matching Color | 匹配颜色 | マッチングカラー|
|| | | | | | Matching Tolerance | 匹配容差 | マッチングの許容度|
|| | | | | | Matching Softness | 匹配和度 ||
|| | | | | | Invert Match | 反转匹配 ||
|| | | Masking Layer | 蒙版图层 | マスクレイヤー | | ||
|| | | | | | Mask Layer | 蒙版图层 | マスクレイヤー|
|| | | | | | Masking Mode | 蒙版模式 ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | | ||
|| | | | | | If Mask Size Differs | 如果蒙版大小不同 ||
|| | | | | | | ||

## 参数详解

## 案例
