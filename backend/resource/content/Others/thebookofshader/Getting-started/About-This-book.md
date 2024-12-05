---
title: About-This-book
order: 1
---

# 介绍

上面的图像是以不同的方式制作的。第一个是梵高用手涂上一层又一层的颜料。他花了几个小时。第二个是在几秒钟内通过四个像素矩阵的组合生成的：一个用于青色，一个用于品红色，一个用于黄色，一个用于黑色。关键区别在于第二张图像是以非连续方式生成的（这意味着不是一步一步，而是同时进行）。

这本书是关于革命性的计算技术，_片段着色器_，它将数字生成的图像提升到一个新的水平。您可以将其视为相当于古腾堡的图形印刷机。

![古腾堡出版社](https://thebookofshaders.com/00/gutenpress.jpg)

古腾堡出版社

片段着色器让您能够以超快的速度完全控制屏幕上呈现的像素。这就是为什么它们被用于各种情况，从手机上的视频过滤器到令人难以置信的 3D 视频游戏。

![That Game Company的旅程](https://thebookofshaders.com/00/journey.jpg)

That Game Company 的旅程

在接下来的章节中，您将发现这种技术的速度和强大程度令人难以置信，以及如何将其应用到您的专业和个人工作中。

## 这本书是给谁的？

本书是为具有编码经验、线性代数和三角学基础知识并希望将他们的工作提升到令人兴奋的新图形质量水平的创意编码员、游戏开发人员和工程师而写的。（如果你想学习如何编码，我强烈建议你从[Processing](https://processing.org/)开始，等你熟悉了再回来。）

本书将教您如何使用着色器并将其集成到您的项目中，从而提高它们的性能和图形质量。因为 GLSL（OpenGL 着色语言）着色器可以在各种平台上编译和运行，所以您可以将在这里学到的知识应用到任何使用 OpenGL、OpenGL ES 或 WebGL 的环境中。换句话说，您将能够将您的知识应用于[处理](https://processing.org/)草图、[openFrameworks](http://openframeworks.cc/)应用程序、[Cinder](http://libcinder.org/)交互式安装、[Three.js](http://threejs.org/)网站或 iOS/Android 游戏。

## 这本书涵盖了哪些内容？

本书将重点介绍 GLSL 像素着色器的使用。首先，我们将定义什么是着色器；然后我们将学习如何使用它们制作程序形状、图案、纹理和动画。您将学习着色语言的基础并将其应用到更有用的场景中，例如：图像处理（图像操作、矩阵卷积、模糊、滤色器、查找表和其他效果）和模拟（Conway 的人生游戏、Gray-Scott 的反应扩散、水波纹、水彩效果、Voronoi 细胞等）。在本书的末尾，我们将看到一组基于 Ray Marching 的高级技术。

*每章都有交互式示例供您使用。*当您更改代码时，您会立即看到更改。这些概念可能是抽象的和令人困惑的，因此交互式示例对于帮助您学习材料至关重要。将概念付诸实践的速度越快，学习过程就越容易。

本书未涵盖的内容：

- 这_不是_一本 openGL 或 webGL 的书。OpenGL/webGL 是一个比 GLSL 或片段着色器更大的主题。要了解有关 openGL/webGL 的更多信息，我建议您查看： [OpenGL 简介](https://open.gl/introduction)，[OpenGL 编程指南第 8 版](http://www.amazon.com/OpenGL-Programming-Guide-Official-Learning/dp/0321773039/ref=sr_1_1?s=books&ie=UTF8&qid=1424007417&sr=1-1&keywords=open+gl+programming+guide)（也称为红皮书）或[WebGL：启动和运行](http://www.amazon.com/WebGL-Up-Running-Tony-Parisi/dp/144932357X/ref=sr_1_4?s=books&ie=UTF8&qid=1425147254&sr=1-4&keywords=webgl)

- 这_不是_一本数学书。尽管我们将涵盖许多依赖于对代数和三角学的理解的算法和技术，但我们不会详细解释它们。对于有关数学的问题，我建议在附近保留以下书籍之一：[3rd Edition of Mathematics for 3D Game Programming and computer Graphics](http://www.amazon.com/Mathematics-Programming-Computer-Graphics-Third/dp/1435458869/ref=sr_1_1?ie=UTF8&qid=1424007839&sr=8-1&keywords=mathematics+for+games)或[2nd Edition of Essential Mathematics for Games and Interactive Applications](http://www.amazon.com/Essential-Mathematics-Games-Interactive-Applications/dp/0123742978/ref=sr_1_1?ie=UTF8&qid=1424007889&sr=8-1&keywords=essentials+mathematics+for+developers)。

## 你需要什么开始？

不多！如果您有支持 WebGL 的现代浏览器（如 Chrome、Firefox 或 Safari）和互联网连接，请单击本页末尾的“下一章”按钮开始。

或者，根据您拥有的或需要从本书中获得的内容，您可以：

- [制作本书的离线版本](https://thebookofshaders.com/appendix/00/)

- [在没有浏览器的情况下在 Raspberry Pi 上运行示例](https://thebookofshaders.com/appendix/01/)

- [制作要打印的图书的 PDF](https://thebookofshaders.com/appendix/02/)

- 查看本书的[GitHub 存储库以帮助解决问题和共享代码。](https://github.com/patriciogonzalezvivo/thebookofshaders)
