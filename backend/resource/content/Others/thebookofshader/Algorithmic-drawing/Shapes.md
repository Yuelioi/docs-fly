---
title: Shapes
order: 4
---

## 形状

![Alice Hubbard, Providence, United States, ca. 1892. Photo: Zindman/Freemont.](https://thebookofshaders.com/07/froebel.jpg)

Alice Hubbard, Providence, United States, ca. 1892. Photo: Zindman/Freemont.

终于！我们一直学习的技能就等着这一刻！你已经学习过 GLSL 的大部分基础，类型和函数。你一遍又一遍的练习你的造型方程。是时候把他们整合起来了。你就是为了这个挑战而来的！在这一章里，你会学习到如何以一种并行处理方式来画简单的图形。

### 长方形

想象我们有张数学课上使用的方格纸，而我们的作业是画一个正方形。纸的大小是 10 x 10 而正方形应该是 8 x 8。你会怎么做？

![](https://thebookofshaders.com/07/grid_paper.jpg)

你是不是会涂满除了第一行第一列和最后一行和最后一列的所有格点？

这和着色器有什么关系？方格纸上的每个小方形格点就是一个线程（一个像素）。每个格点有它的位置，就想棋盘上的坐标一样。在之前的章节我们将 x 和 y 映射到 rgb 通道，并且我们学习了如何将二维边界限制在 0 和 1 之间。我们如何用这些来画一个中心点位于屏幕中心的正方形？

我们从空间角度来判别的 if 语句伪代码开始。这个原理和我们思考方格纸的策略异曲同工。

```
if ( (X GREATER THAN 1) AND (Y GREATER THAN 1) )
        paint white
    else
        paint black
```

现在我们有个更好的主意让这个想法实现，来试试把 if 语句换成 step（），并用 0 到 1 代替 10 x 10 的范围。

```
uniform vec2 u_resolution;

void main(){
    vec2 st = gl_FragCoord.xy/u_resolution.xy;
    vec3 color = vec3(0.0);

    // Each result will return 1.0 (white) or 0.0 (black).
    float left = step(0.1,st.x); // Similar to ( X greater than 0.1 )
    float bottom = step(0.1,st.y); // Similar to ( Y greater than 0.1 )

    // The multiplication of left*bottom will be similar to the logical AND.
    color = vec3( left * bottom );

    gl_FragColor = vec4(color,1.0);
}
```

step（）函数会让每一个小于 0.1 的像素变成黑色（vec3（0.0））并将其余的变成白色（vec3（1.0））。`left` 乘 `bottom` 效果相当于逻辑 AND —— 当 x y 都为 1.0 时乘积才能是 1.0。这样做的效果就是画了两条黑线，一个在画布的底边另一个在左边。

![](https://thebookofshaders.com/07/rect-01.jpg)

在前一例代码中我们重复每个像素的结构（左边和底边）。我们可以把原来的一个值换成两个值直接给 step（）来精减代码。就像这样：

```
vec2 borders = step(vec2(0.1),st);
    float pct = borders.x * borders.y;
```

目前为止，我们只画了长方形的两条边（左边和底面）。看下下面的例子：

```
// Author @patriciogv - 2015
// http://patriciogonzalezvivo.com

#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

void main(){
    vec2 st = gl_FragCoord.xy/u_resolution.xy;
    vec3 color = vec3(0.0);

    // bottom-left
    vec2 bl = step(vec2(0.1),st);
    float pct = bl.x * bl.y;

    // top-right
    // vec2 tr = step(vec2(0.1),1.0-st);
    // pct *= tr.x * tr.y;

    color = vec3(pct);

    gl_FragColor = vec4(color,1.0);
}

```

取消 21~22 行的注释来看看如何转置坐标的同时重复使用 `step()` 函数。这样二维向量 vec2(0.0,0.0) 会被变换到右上角。这就是转置页面和重复过程的数字等价。

![](https://thebookofshaders.com/07/rect-02.jpg)

注意在 18 行和 22 行，所有的边（宽）都被放大了。等价于这样写：

```
vec2 bl = step(vec2(0.1),st); // bottom-left
    vec2 tr = step(vec2(0.1),1.0-st); // top-right
    color = vec3(bl.x * bl.y * tr.x * tr.y);
```

是不是很有趣？这种都是关于运用 step() 函数、逻辑运算和转置坐标的结合。

再进行下一个环节之前，挑战下下面的练习：

- 改变长方形的比例和大小。

- 用 smoothstep() 函数代替 step() 函数，试试在相同的代码下会有什么不同。注意通过改变取值，你可以不仅可以得到模糊边界也可以由漂亮的顺滑边界。

- 应用 floor() 做个另外的案例。

- 挑个你最喜欢的做成函数，这样未来你可以调用它。并且让它灵活高效。

- 写一个只画长方形四边的函数。

- 想一下如何在一个画板上移动并放置不同的长方形？如果你做出来了，试着像[Piet Mondrian](http://en.wikipedia.org/wiki/Piet_Mondrian)一样创作以长方形和色彩的图画。

![Piet Mondria - Tableau (1921)](https://thebookofshaders.com/07/mondrian.jpg)

Piet Mondria - Tableau (1921)

### 圆

在笛卡尔坐标系下，用方格纸来画正方形和长方形是很容易的。但是画圆就需要另一种方式了，尤其我们需要一个对“每个像素”的算法。一种解决办法是用[`step()`](https://thebookofshaders.com/glossary/?search=step)函数将重新映射的空间坐标来画圆。

如何实现？让我们重新回顾一下数学课上的方格纸：我们把圆规展开到半径的长度，把一个针脚戳在圆圆心上，旋转着把圆的边界留下来。

![](https://thebookofshaders.com/07/compass.jpg)

将这个过程翻译给 shader 意味着纸上的每个方形格点都会隐含着问每个像素（线程）是否在圆的区域以内。我们通过计算像素到中心的距离来实现（这个判断）。

![](https://thebookofshaders.com/07/circle.jpg)

There are several ways to calculate that distance. The easiest one uses the [`distance()`](https://thebookofshaders.com/glossary/?search=distance) function, which internally computes the [`length()`](https://thebookofshaders.com/glossary/?search=length) of the difference between two points (in our case the pixel coordinate and the center of the canvas). The `length()` function is nothing but a shortcut of the [hypotenuse equation](http://en.wikipedia.org/wiki/Hypotenuse) that uses square root ([`sqrt()`](https://thebookofshaders.com/glossary/?search=sqrt)) internally.

有几种方法来计算距离。最简单的是用[`distance()`](https://thebookofshaders.com/glossary/?search=distance)函数，这个函数其实内部调用 [`length()`](https://thebookofshaders.com/glossary/?search=length)函数，计算不同两点的距离（在此例中是像素坐标和画布中心的距离）。length（）函数内部只不过是用平方根([`sqrt()`](https://thebookofshaders.com/glossary/?search=sqrt))计算斜边的方程。

![](https://thebookofshaders.com/07/hypotenuse.png)

你可以使用[`distance()`](https://thebookofshaders.com/glossary/?search=distance), [`length()`](https://thebookofshaders.com/glossary/?search=length) 或 [`sqrt()`](https://thebookofshaders.com/glossary/?search=sqrt)到计算屏幕的中心的距离。下面的代码包含着三个函数，毫无悬念的他们返回相同的结果。

- 注释和取消某行的注释来试试看用不同方式得到相同的结果。

```
// Author @patriciogv - 2015
// http://patriciogonzalezvivo.com

#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

void main(){
 vec2 st = gl_FragCoord.xy/u_resolution;
    float pct = 0.0;

    // a. The DISTANCE from the pixel to the center
    pct = distance(st,vec2(0.5));

    // b. The LENGTH of the vector
    //    from the pixel to the center
    // vec2 toCenter = vec2(0.5)-st;
    // pct = length(toCenter);

    // c. The SQUARE ROOT of the vector
    //    from the pixel to the center
    // vec2 tC = vec2(0.5)-st;
    // pct = sqrt(tC.x*tC.x+tC.y*tC.y);

    vec3 color = vec3(pct);

 gl_FragColor = vec4( color, 1.0 );
}
```

上回我们把到中心的距离映射为颜色亮度。离中心越近的越暗。注意到映射值不宜过高，因为从中心（vec2(0.5, 0.5)）到最远距离才刚刚超过 0.5 一点。仔细考察这个映射：

- 你能从中推断出什么？

- 我们怎么用这个方法来画圆？

- 试试有没有其他方法来实现这样画布内圆形渐变的效果。

### 距离场

我们可也可以从另外的角度思考上面的例子：把它当做海拔地图（等高线图）——越黑的地方意味着海拔越高。想象下，你就在圆锥的顶端，那么这里的渐变就和圆锥的等高线图有些相似。到圆锥的水平距离是一个常数 0.5。这个距离值在每个方向上都是相等的。通过选择从那里截取这个圆锥，你就会得到或大或小的圆纹面。

![](https://thebookofshaders.com/07/distance-field.jpg)

其实我们是通过“空间距离”来重新解释什么是图形。这种技巧被称之为“距离场”，从字体轮廓到 3D 图形被广泛应用。

来小试下牛刀：

- 用[`step()`](https://thebookofshaders.com/glossary/?search=step)函数把所有大于 0.5 的像素点变成白色，并把小于的变成黑色（0.0）。

- 反转前景色和背景色。

- 使用[`smoothstep()`](https://thebookofshaders.com/glossary/?search=smoothstep)函数，通过修改不同的值来试着做出一个边界顺滑的圆。

- 一旦遇到令你满意的应用，把他写成一个函数，这样将来就可以调用了。

- 给这个圆来些缤纷的颜色吧！

- 再加点动画？一闪一闪亮晶晶？或者是砰砰跳动的心脏？（或许你可以从上一章汲取一些灵感）

- 让它动起来？能不能移动它并且在同一个屏幕上放置多个圆？

- 如果你结合函数来混合不同的距离场，会发生什么呢？

```
pct = distance(st,vec2(0.4)) + distance(st,vec2(0.6));
pct = distance(st,vec2(0.4)) * distance(st,vec2(0.6));
pct = min(distance(st,vec2(0.4)),distance(st,vec2(0.6)));
pct = max(distance(st,vec2(0.4)),distance(st,vec2(0.6)));
pct = pow(distance(st,vec2(0.4)),distance(st,vec2(0.6)));
```

- 用这种技巧制作三个元素，如果它们是运动的，那就再好不过啦！

#### 添加自己的工具箱

就计算效率而言，[`sqrt()`](https://thebookofshaders.com/glossary/?search=sqrt)函数，以及所有依赖它的运算，都耗时耗力。[`dot()`](https://thebookofshaders.com/glossary/?search=dot)点乘是另外一种用来高效计算圆形距离场的方式。

```
// Author @patriciogv - 2015
// http://patriciogonzalezvivo.com

#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

float circle(in vec2 _st, in float _radius){
    vec2 dist = _st-vec2(0.5);
 return 1.-smoothstep(_radius-(_radius*0.01),
                         _radius+(_radius*0.01),
                         dot(dist,dist)*4.0);
}

void main(){
 vec2 st = gl_FragCoord.xy/u_resolution.xy;

 vec3 color = vec3(circle(st,0.9));

 gl_FragColor = vec4( color, 1.0 );
}
```

### 距离场的特点

![Zen garden](https://thebookofshaders.com/07/zen-garden.jpg)

Zen garden

距离场几乎可以用来画任何东西。显然，图形越复杂，方程也越复杂。但是一旦你找到某个特定图形的公式，就很容易添加图形或应用像过渡边界的效果。正因如此，距离场经常用于字体渲染，例如[Mapbox GL Labels](https://blog.mapbox.com/drawing-text-with-signed-distance-fields-in-mapbox-gl-b0933af6f817), [Matt DesLauriers](https://twitter.com/mattdesl) [Material Design Fonts](http://mattdesl.svbtle.com/material-design-on-the-gpu) 和 [as is describe on Chapter 7 of iPhone 3D Programming, O’Reilly](http://chimera.labs.oreilly.com/books/1234000001814/ch07.html#ch07_id36000921).

看看下面的代码：

```
#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

void main(){
  vec2 st = gl_FragCoord.xy/u_resolution.xy;
  st.x *= u_resolution.x/u_resolution.y;
  vec3 color = vec3(0.0);
  float d = 0.0;

  // Remap the space to -1. to 1.
  st = st *2.-1.;

  // Make the distance field
  d = length( abs(st)-.3 );
  // d = length( min(abs(st)-.3,0.) );
  // d = length( max(abs(st)-.3,0.) );

  // Visualize the distance field
  gl_FragColor = vec4(vec3(fract(d*10.0)),1.0);

  // Drawing with the distance field
  // gl_FragColor = vec4(vec3( step(.3,d) ),1.0);
  // gl_FragColor = vec4(vec3( step(.3,d) * step(d,.4)),1.0);
  // gl_FragColor = vec4(vec3( smoothstep(.3,.4,d)* smoothstep(.6,.5,d)) ,1.0);
}
```

我们一开始把坐标系移到中心并把它映射到-1 到 1 之间。在 _24 行_ 这儿，我们用一个[`fract()`](https://thebookofshaders.com/glossary/?search=fract) 函数来呈现这个距离场产生的图案。这个距离场不断重复，就像在禅花园看到的环一样。

现在我们来看下 _19 行_ 的距离场方程。这里我们在计算点 `(.3,.3)` 或 `vec3(.3)`到所有四象限的距离（这就是 [`abs()`](https://thebookofshaders.com/glossary/?search=abs) 在起作用）。

如果你取消第 _20 行_ 的注释，你会发现我们把到四个点的距离用[`min()`](https://thebookofshaders.com/glossary/?search=min) 函数合并到 0，并产生了一个有趣的图案。

现在再试着取消第 _21 行_ 的注释，我们做的和之前一样，只不过这次用的是 [`max()`](https://thebookofshaders.com/glossary/?search=max) 函数。这次的记过是圆角矩形。注意距离场的环形是如何离中心越远越光滑的。

最后从_27 行到 29 行_一行行地取消注释，思考距离场的不同用途。

### 极坐标下的图形

![Robert Mangold - Untitled (2008)](https://thebookofshaders.com/07/mangold.jpg)

Robert Mangold - Untitled (2008)

在关于颜色的章节我们通过如下的方程把每个像素的 **半径** 和 **角度** 笛卡尔坐标映射到极坐标。

```
vec2 pos = vec2(0.5)-st;
    float r = length(pos)*2.0;
    float a = atan(pos.y,pos.x);
```

我们用了部分方程在这章的开头来画圆，即用 [`length()`](https://thebookofshaders.com/glossary/?search=length) 计算到中心的距离。现在我们可以用极坐标来画圆。

极坐标这种方式虽然有所限制但却十分简单。

下面你会看到在同样在笛卡尔坐标下图形在极坐标下的着色器案例（在 _lines 21 和 25_之间）。对这些函数一个个取消注释，看看两坐标系之间的联系。

```
// Author @patriciogv - 2015
// http://patriciogonzalezvivo.com

#ifdef GL_ES
precision mediump float;
#endif

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

void main(){
    vec2 st = gl_FragCoord.xy/u_resolution.xy;
    vec3 color = vec3(0.0);

    vec2 pos = vec2(0.5)-st;

    float r = length(pos)*2.0;
    float a = atan(pos.y,pos.x);

    float f = cos(a*3.);
    // f = abs(cos(a*3.));
    // f = abs(cos(a*2.5))*.5+.3;
    // f = abs(cos(a*12.)*sin(a*3.))*.8+.1;
    // f = smoothstep(-.5,1., cos(a*10.))*0.2+0.5;

    color = vec3( 1.-smoothstep(f,f+0.02,r) );

    gl_FragColor = vec4(color, 1.0);
}
```

试着：

- 让这些图形动起来。
- 结合不同的造型函数来 **雕刻** 图形，制作诸如花，雪花和齿轮。
- 用我们在 **造型函数** 章节的 `plot()` 函数画等高线。

### 整合的魅力

到目前为止，我们知道如何用[`atan()`](https://thebookofshaders.com/glossary/?search=atan)函数来根据角度调整半径以获得不同的图形，以及如何用`atan()`结合所以和距离场有关的技巧得到可能的效果。

看下下面来自[Andrew Baldwin](https://twitter.com/baldand)的例子。这里的技巧是用极坐标的方式通过定义多边形的边数来构建一个距离场。

```
#ifdef GL_ES
precision mediump float;
#endif

#define PI 3.14159265359
#define TWO_PI 6.28318530718

uniform vec2 u_resolution;
uniform vec2 u_mouse;
uniform float u_time;

// Reference to
// http://thndl.com/square-shaped-shaders.html

void main(){
  vec2 st = gl_FragCoord.xy/u_resolution.xy;
  st.x *= u_resolution.x/u_resolution.y;
  vec3 color = vec3(0.0);
  float d = 0.0;

  // Remap the space to -1. to 1.
  st = st *2.-1.;

  // Number of sides of your shape
  int N = 3;

  // Angle and radius from the current pixel
  float a = atan(st.x,st.y)+PI;
  float r = TWO_PI/float(N);

  // Shaping function that modulate the distance
  d = cos(floor(.5+a/r)*r-a)*length(st);

  color = vec3(1.0-smoothstep(.4,.41,d));
  // color = vec3(d);

  gl_FragColor = vec4(color,1.0);
}

```

- 用这个例子，改造一个输入位置，指定图形（形状）的顶点数来返回一个距离场（的值）。

- 结合使用 [`min()`](https://thebookofshaders.com/glossary/?search=min) 和 [`max()`](https://thebookofshaders.com/glossary/?search=max) 函数混合距离场。

- 用距离场画个自己感兴趣的 logo。

恭喜！你完成了最艰难的部分！休息下让这些概念沉淀一下吧 —— 用 Processing 来画简单的形状很容易，但却不到火候。在 shader 的世界里，画形状是很纠结，而且适应这种新的编程范式会有些累人。

在本章的最后，您将找到\[PixelSpirit Deck\]（[https://patriciogonzalezvivo.github.io/PixelSpiritDeck/）的链接，这组卡片将帮助您学习新的 SDF 功能，将其组合到您的设计和使用中](https://patriciogonzalezvivo.github.io/PixelSpiritDeck/%EF%BC%89%E7%9A%84%E9%93%BE%E6%8E%A5%EF%BC%8C%E8%BF%99%E7%BB%84%E5%8D%A1%E7%89%87%E5%B0%86%E5%B8%AE%E5%8A%A9%E6%82%A8%E5%AD%A6%E4%B9%A0%E6%96%B0%E7%9A%84SDF%E5%8A%9F%E8%83%BD%EF%BC%8C%E5%B0%86%E5%85%B6%E7%BB%84%E5%90%88%E5%88%B0%E6%82%A8%E7%9A%84%E8%AE%BE%E8%AE%A1%E5%92%8C%E4%BD%BF%E7%94%A8%E4%B8%AD) 在您的着色器上。 卡片组具有预先的学习曲线，因此，每天拿一张卡片并进行练习将推动并挑战您几个月的技能。

既然现在你知道了如何画形状，我十分肯定你脑袋里已经充满了新的点子。在接下来的章节里你会学习到怎么移动，旋转以及缩放图形。这将使你的创作如虎添翼！
