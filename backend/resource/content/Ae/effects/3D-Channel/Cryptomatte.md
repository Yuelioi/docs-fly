---
title: Cryptomatte
order: 2
---

# Cryptomatte

## 简述

一种提取 exr 文件通道的效果, 可以单独提取某个 selection

## 效果展示

把 exr 文件导入 ae

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-0.png)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-cryptomatte_cn.png)

| Layer     | 图层 | レイヤー         |     |     |
| --------- | ---- | ---------------- | --- | --- |
| Selection | 选择 | 選択             |     |     |
| Manifest  | 显示 | ディスプレイ設定 |     |     |

## 参数详解

### 使用方法

单击：选择色块  
shift 单击某块：增选  
ctrl 单击某块：可以减选

### 默认为 Colors 状态

如下图所示。

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-2.png)

| 勾选 Colors          | （单击圆球部分时）                 | ![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-1.png) |
| -------------------- | ---------------------------------- | ------------------------------------------------------------------------------------------------- |
| 勾选 Matted Colors： | 把选择的颜色块删除（留下纯白色块） | ![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-3.png) |
| 勾选 Matted RGBA：   | 此时仅输出选中的部分               | ![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-4.png) |
| 勾选 Matte Only      | 丢掉选择的蒙版部分（纯白色块）     | ![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-cryptomatte-5.png) |
