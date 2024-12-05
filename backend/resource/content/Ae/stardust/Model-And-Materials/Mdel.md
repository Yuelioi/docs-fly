---
title: 模型 Mdel
order: 2
---

# 模型 Mdel

## Model Source 模型源

- File 文件：用到导入外部模型
- Primitive 初始对象：创建基本几何体，如圆环、球、立方体、平面、圆柱。
- Text／mask 文字／遮罩：根据文字或者遮罩挤压模型。
- Volume 体积：新版本加的一种模式，可以导入 VDB 烟雾等等，需要 Volume 节点和这个模式配合。

### Material Click To Open 点击打开

### File Properties 文件属性

- File Name 文件名：用来导入外部模型
- OBJ Layer OBJ 图层：选择 OBJ 模型所在图层，导入外部模型时候会给加载好。
- Load Sequence 加载序列：动画记得勾选，默认会勾选
- Loop Sequence 循环序列：勾选了就会循环播放，不勾选动画结束就定格在最后一帧。
- Offset Sequence 偏移序列：用不着，一般都在三维中处理。
- Invert Normals 反转法线：将法线进行反转，一般用不着。

### Primitive Properties 初始对象属性

- Type 类型
  - Cube 立方体
  - Sphere 球
  - Plane 平面
  - Torus 圆环
  - Cylinder 圆柱
- Size X 大小 X：X 轴大小 单位像素
- Size Y 大小 Y：Y 轴大小
- Size Z 大小 Z：Z 轴大小
- Segments X X 轴分段：X 轴分段数，分段越高精度越高，stardust 不是三维软件一般不需要特别高的分段，需要 deform 模型或者动力学运算时可能需要提高分段精度。分段数不够对模型变形则就不够柔和。
- Segments Y Y 轴分段：Y 轴的分段数
- Segments Z Z 轴分段：Z 轴的分段数

![](http://cdn.yuelili.com/202121311427-6.png)

（以 C4D 为例：X 轴分段 10 Y 轴分段 5 Z 轴分段 2）

- Bevel Segments 倒角分段
- Bevel Size 倒角大小

![](http://cdn.yuelili.com/202121311428-x.png)

模型的倒角，则就是边和边之间的过渡，分段越多越圆滑，在三维制作模型中都需要好看的倒角，在 stardust 中用的偏少

### Extrude Properties 挤出属性

- Extrude Layer 挤出图层

选择挤出哪一层（AE 的图层）进行挤出模型，可以是固态层 MASK 也可以是文字，形状层不行。如果在 stardust 本层进行画 MSK，记得把模式改成 none。不然会影响形状和观看。并且画 MASK 的时候不要进行交叉。

- Type 类型
  - Default 默认：根据 mask 的面来挤出

![](http://cdn.yuelili.com/202121311430-9.png)

- Edge 边：根据 MASK 的边缘进行挤出

![](http://cdn.yuelili.com/202121311430-j.png)

- Path Samples 路径样本：挤出的分段数，如果是圆滑的图形挤出需要比较高的分段，分段少则会有棱有角。

![](http://cdn.yuelili.com/202121311431-l.png)

- Path Offset 路径偏移：对路径进行偏移，正数是往外扩张，负值收缩，反应在模型上就是正值就是扩大了，负值缩小。
- Edges Width 边缘宽度：在 edge 模式下生效，边缘挤压的宽度。

![](http://cdn.yuelili.com/202121311433-h.png)

- Extrude Depth 边缘深度：挤压的高度

![](http://cdn.yuelili.com/202121311432-1.png)

- Extrude Segments 挤出分段：stardust 基本上采用正方体之类的简单模型，分段基本上没用，一般配合 deform 进行模型变形使用
- Bevel Profile 倒角：用曲线控制倒角的方式，和下面的参数一起控制挤出模型的倒角
- Bevel Segments 倒角分段：分段数越大，倒角越圆滑。
- Bevel Depth 倒角厚度：反应在模型上就是模型，倒角厚度越大模型变的越大
- Bevel Distance 倒角距离：倒角距离越大模型也会因为倒角被撑大了。

![](http://cdn.yuelili.com/202121311436-w.png)

stardust 一般并不会做多么复杂的倒角，用这几个参数简单调调就行，真正的三维模型必然不会在 stardust 做，金属字一般也用 E3D 解决，stardus 也提供了许多预设需要可以自己调用。

![](http://cdn.yuelili.com/202121311437-e.png)

### Model Properties 模型属性

- Inherit Motion 继承运动

继承空物体的运动，创建一个空物体打成三维图层可以用来控制这个模型位移、旋转、缩放。（视口操作可能看着没效果，拖拉一下进度条即可）

- Origin XY 起点 XY
- Origin Z 起点 Z
- Angle X X 轴角度
- Angle Y Y 轴角度
- Angle Z Z 轴角度

![](http://cdn.yuelili.com/202121311438-B.png)

- Size 大小
- Center 中心：强制把模型归到 AE 的世界中心，此时控制模型位置也不好使了，因为模型会被强制到中心
- Flip X X 轴反转：

以 X 轴模型就行反转，详情参考 AE 中的水平反转，垂直反转，只不过一个是对视频图片，stardust 中是对模型。

- Flip Y Y 轴反转
- Flip Z Z 轴反转
- Normalize Scale 标准大小：标准大小：将模型变成标准大小。和 Align Model 配合可以用来观察模型。其次，可用于校正模型过大过小
- Align Model 对齐模型
  - Default 默认
  - Front 正面
  - Back 背面
  - Left 左面
  - Right 右面
  - UP 顶面
  - Down 底面
- Texture Mapping 贴图纹理
  - UV Map：

模型贴图的投射方式（首先要连接材质节点）有模型（就是默认）球、平直三种投射方式，就是三种不同的投射方式而已，比较不是三维软件用的不多，下面有 UV 控制这张图平铺几次和偏移。

![](http://cdn.yuelili.com/202121311442-3.png)

![](http://cdn.yuelili.com/202121311442-m.png)

- Map Orientation：使用空物体控制贴图方向，不能使用默认的 UV 投射方式只可以使用剩下两种。
- Apply Deformers：UV 贴图是否应用于 deform，可以理解为 C4D 的粘滞纹理标签，本身 deform 这个功能用的目前使用偏少。
- Refine Model 模型精度

模型的精度，如果导入模型可能太卡，可能选择下面的选项进行优化一下。

- OFF 关
- LOW 低
- Medium 中等
- High 高
- Double Sided 双面

比如一张纸会有两个面，一般来说模型里面都是看不着的，但是如果面的法线反了（就是反面在外面，正面在里面，一般来说默认贴图都是贴在你的正面，在里面看不到面不贴），就可以看到贴图材质了。当然不仅代表双面也代表内外，用来做一些管道动画是非常有用的。

![](https://cdn.yuelili.com/20211116162327.png)

- Cast Shadows 投射阴影：模型是否投射阴影
- Accept Shadows 接受阴影：模型是否接受阴影
- Smooth Normals 平滑法线：和 Smooth Angle Thres 配合使用，用少量的分段做出圆滑的模型，比如球体。
- Smooth Angle Thres 平滑角度

![](http://cdn.yuelili.com/202121311446-Z.png)

- Create Null 创建空对象：自动创建一个空对象（和模型对其的空对象）进行控制模型

### Reflection Plane 反射面

都是用来控制一个反射面，点击开启就会有一个反射面（是一个蓝色的面，默认和你平面平行，一般开启不需要再调整其他属性，调整材质就行了），首先这个功能不可以单独模型进行使用（只有一个物体你没反射的对象），还要有一个立方体或者平面，你开启反射面之后
你的反射面和你的立方体或者平面会形成一定的视觉效果（类似镜面反射的效果），调整上面的偏移 XYZ 等形成不同的效果，毫无卵用的效果。（调整反射面要在你遮挡的那个正方体或者平面到你这个模型之间
超过其中一个就没有效果）一般来说还可以开启配合材质做一种假的反射材质。但是实际上不会用 stardust 做一些写实反射效果所以实用性很小。

![](http://cdn.yuelili.com/202121311447-Y.png)

- Enabled 启用
- Offset 偏移
- Angle X 角度 X
- Angle Y 角度 Y
- Angle Z 角度 Z
