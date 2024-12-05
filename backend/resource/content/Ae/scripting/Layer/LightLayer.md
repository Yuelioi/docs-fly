---
title: LightLayer
order: 5
---

# LightLayer 灯光图层

## LightLayer object

app.project.item(index).layer(index)

描述：LightLayer 对象表示合成中的灯光图层。使用 LayerCollection.addLight()方法创建。可以通过索引或名称在项目的图层集中对其进行访问。

父级关系：LightLayer 是 Layer 对象的子类。使用 Light Layer 时，Layer 的所有方法和属性均可用。

AE 属性

LightLayer 没有定义其他属性，但是具有与其他图层类型不同的 AE 属性。它具有以下属性和属性组：

| Marker       | 标记     |                    |            |
| ------------ | -------- | ------------------ | ---------- |
| Transform    | 变换     |                    |            |
|              |          | PointofInterest    | 目标点     |
|              |          | Position           | 位置       |
|              |          | Scale              | 缩放       |
|              |          | Orientation        | 方向       |
|              |          | XRotation          | X 旋转     |
|              |          | YRotation          | Y 旋转     |
|              |          | Rotation           | 旋转       |
|              |          | Opacity            | 不透明度   |
| LightOptions | 灯光选项 |                    |            |
|              |          | Casts shadows      | 投影       |
|              |          | Light Transmission | 透光率     |
|              |          | Acepts Shadows     | 接受阴影   |
|              |          | Acepts Lights      | 接受灯光   |
|              |          | Ambient            | 环境       |
|              |          | Diffuse            | 漫射       |
|              |          | Specular Intensity | 镜面强度   |
|              |          | Specular Shininess | 镜面反光度 |
|              |          | Metal              | 金属质感   |

## 属性

### LightLayer.lightType

app.project.item(index).layer(index).lightType

描述：灯光类型

类型：LightType，枚举值; 读/写：

- LightType.PARALLEL
- LightType.SPOT
- LightType.POINT
- LightType.AMBIENT

示例：把选择的灯光图层改为点光

```javascript
app.project.activeItem.selectedLayers[0].lightType = LightType.POINT;
```
