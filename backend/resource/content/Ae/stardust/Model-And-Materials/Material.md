---
title: 材质 Material
order: 1
---

# 材质 Material

### Material 材质

定义 3D 模型材质。

#### Type 类型

- Solid 固有色：默认就是类似漫射材质的一种实体材质
- Transparent 透明：可以穿透过去的一种透明材质，并不是真正意义上的透明材质，实际上是把模型的面片进行 add 模式叠加（就是图层的那种混合模式）

![](http://cdn.yuelili.com/202121311448-s.png)

- Volume 体积：体积的材质，必须和体积节点以及模型体积一块用。

![](http://cdn.yuelili.com/202121311449-1.png)

### Solid & Transparent

以下为前 2 个选项的参数

#### Diffuse 漫反射

- Diffuse Amount 漫射量：漫射对颜色的影响，为 0 那就没有漫射了就成黑色了
- Ambient Amount 环境量：环境对颜色的影响。

![](http://cdn.yuelili.com/202121311450-Q.png)

环境是 stardust 最上面进行选择，我将环境弄成蓝色，调节材质中的参数环境的蓝色对材质影响就越大

- Color 颜色：漫射的颜色，可以理解为本体的颜色。
- Color From Particle 颜色从粒子那里：继承粒子的颜色，如果模型是用粒子发射出来，继承粒子节点那里颜色影响，单个模型把这个改成 0，不然不好使。
- Diffuse Texture 漫射贴图：用一张图做材质
- Texture Opacity 贴图不透明度：贴图透明度降低原本漫射的颜色就会出来。

![](http://cdn.yuelili.com/202121311453-5.png)

#### Roughness 粗糙

- Roughness Amount 粗糙量：粗糙度越强，模型就越粗糙，磨砂玻璃肯定比普通玻璃更粗糙一些。
- Roughness Texture 粗糙贴图：用一种黑白贴图去控制哪里粗糙度高哪里不高，白色为 100%黑色为 0%。
- Gloss／Roughness 镜面光泽／粗糙：默认是金属度，上面那个数值越高对应则是粗糙度越低，取消就反过来。

#### Reflection／Metal 反射／金属

- Amount 量：反射的强度，比如镜子，金属等等
- Reflection Texture 反射贴图
- Reflections Amount 反射量
- Environment Amount 环境量：环境对反射的影响
- Specular Amount 高光量：反射中高光的强度。

![](http://cdn.yuelili.com/202121311454-y.png)

这里以 C4D 作为参考右边是一个完全反射的材质，然后是在完全反射情况下提高了粗糙度，最后一个则是用了一个黑白贴图进行控制粗糙度，有的地方粗糙有的地方不粗糙，其实想在 st 里做到非常好看的反射是一件很不现实的事，基本上就是调整调整反射高光做出一个光感就可以了。

#### Refraction 折射

- Amount 量：折射的强度，在透明材质下使用配合环境模拟玻璃。
- Distortion 形变

#### Subsurface Scattering 次表面散射

- Amount 量：强度
- Color 颜色：透光的颜色，但是你在调节时候你会发现会出现一种相反的颜色，因为散射现象，是把蓝光散射出去，在散射很强的情况下，你光线出去之后你选择的这个颜色全部消减掉了，所以会呈现他的反色。
- Scatter 散射：散射的强度，光线打入物体，光线在物体中会进行一种散射现象。

![](http://cdn.yuelili.com/202121311457-2.png)

- Density 密度：越高，越不容易透光

![](http://cdn.yuelili.com/202121311458-2.png)

使用 3S 的时候一点过要打开开关，不然不好使，3S 效果就是一种透光不透明的效果，比如蜡烛，玉石，人的皮肤一样，具有透光性，不过 stardust 是一种作假的方式，效果一般般，一般也用不上。

#### Emissive 放射

- Amount 量
- Color From Particle：继承粒子颜色，不确定是什么版本加进来的默认为 0。就一直为 0 就行了
- Emissive Texture 放射贴图：并且必须使用贴图才会有明显的效果

![](http://cdn.yuelili.com/202121311500-W.png)

#### Bump 凸起

- Amount 量：凹凸的强度
- Bump Texture 凸起贴图
- Normal／Bump 正常／凸起

#### Transparent material 透明材质

- Opacity 不透明度：要在透明材质下影响透明度
- Opacity From Particle 透明度来自粒子
- Normal Affects Opacity 正常影响不透明度
- Affect DOF 影响景深

#### Shadow Catcher 阴影捕捉

![](http://cdn.yuelili.com/202121311503-x.png)

![](http://cdn.yuelili.com/202121311504-D.png)

- Enable 启用

勾选之后模型就消失但是投影下来阴影还会保留，也可以使用进行模型遮挡（其他模型和他交接部分回合勾选了的模型一块消失）

- stencil mode 模具模式

#### Texture Time Sample

- Current time 当前时间
- Random Still Frame 随机静帧

#### Material ID 材质编号

只有导入模型时候带贴图分不同的编号

**Scatter** **散射**

**Input Channel** **输入通道**

- densitey 密度
- Temperature 温度

### Volume 参数

**Scatter** **散射**

- density amount 密度值：数值越大密度越高，看起来就越厚重越实，类似透明度
- scatter color 散射颜色：
- sctter color gradient 散射颜色渐变：用渐变这个图标影响散射的颜色，从左到右对应实体的由外到内。

![](http://cdn.yuelili.com/202121311510-P.png)

- Gradient input scale 渐变输入比例：默认是 1，小于偏向图表右边，大于偏向图表左边颜色。
- input minimum 输入最小值：调整会偏向右边的颜色
- input maximum 输入最大值：调整会偏向左边的颜色。这两个参数其实就类似色阶左右拉。
- Enable Remap Density 启用重映射密度：启用之后可以使用参数来调整浓度默认是很糟糕的，所有地方浓度都是 1 就全都是纯色了。

![](http://cdn.yuelili.com/202121311518-W.png)

- Remap Density 重映射密度：根据曲线调整密度

![](http://cdn.yuelili.com/202121311519-M.png)

- Shadow Density 阴影密度：字面意思（就是阴影透明度），默认就比较合适。
- Shadow Darkness 阴影黑暗：就是阴影多黑的意思
- anisotropy 个向异性：测试完感觉没什么用
- Brightiness 亮度：字面意思
- Gamma 伽马：简单来说调低变亮，调高变暗

#### **Absorption** **吸收**

- Absorption Amount 吸收强度：字面意思，吸收越强表现则会越黑，因为光被吸收了阴影也会越多，尤其背光的地方。

![](http://cdn.yuelili.com/202121311548-J.png)

- Enable Remap Absorptio 启用重映射吸收：和密度一样不说了
- Remap Absorption Densit：和密度一样的曲线

**Emission** **发光**

- Input Channel 输入通道：选择哪个通道进行发光 （一般来说都是用温度通道）
- Emission Amount 发光强度：字面意思，。
- Emission color 发光颜色：字面意思

![](http://cdn.yuelili.com/202121311606-W.png)

- Emisssion Color Gradient 发光颜色渐变：用渐变控制发光颜色
- Gradient input scale 发光输入比例：和密度曲线一样
- Auto Range 自动范围：测试没什么区别，默认勾着就行
- input minimum 最小输入：
- Input maximum 最大输入：都和密度一样。

总之 stardust 使用 VDB 烟雾这种东西，体积，模型，和三个节点是都要使用，如果不会连接那就打开预设里面的看看。
