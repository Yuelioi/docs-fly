---
title: Multi-Frame-Rendering-in-AE
order: 15
---

# AE中的多帧渲染

为了利用具有更多 CPU 核心和线程，After Effects 2022 及以上版本现已支持多帧渲染。多帧渲染(MFR)允许多帧同时渲染，从而加快合成的渲染和输出。

第三方特效可以通过 AE 特效 SDK 设置以下 PF_OutFlag 支持多帧渲染。

```cpp
PF_OutFlag2_SUPPORTS_THREADED_RENDERING
```

这个开关表示该效果支持多线程并发渲染。一个图层上的这个效果的单个或多个应用可以在多个线程上同时被调用渲染。在该开关被设置之前，效果必须线程安全。更多信息参见下面的 <什么是效果的线程安全> 部分。

::: tip
当 After Effects 使用多帧渲染时，没有线程安全的效果，如果没有设置这个开关，将强制每个渲染线程一次性进入和退出效果代码。这将大大降低 MFR 提供的性能优化，因此在效果控制窗口会显示警告图标，提醒用户注意性能影响。
:::

对于在渲染过程中需要写入 sequence_data 的效果，为了向后兼容，可以设置这个开关。

```cpp
PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER
```

每个渲染线程都有自己的 sequence_data 实例，不与其他渲染线程共享或同步。如果存储在 sequence_data 中的数据计算起来很费时，应该利用新的多帧渲染 "计算缓存"(见后)。

::: tip

使用`PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER`开关需要 2021 年 3 月或更高版本的SDK进行编译。

:::

## 多帧渲染在2021 3月的SDK上架

2021 年 3 月的 SDK 引入了新的`sequence_data`行为，从 AE beta builds 22.0x6(2021 年 6 月 29 日发布)开始启用。任何用 2020 年 6 月的 SDK 编译的效果必须用 2021 年 3 月的 SDK 重新编译以支持多帧渲染。虽然建议使用 SDK 常量 PF_AE_PLUG_IN_VERSION 和 PF_AE_PLUG_IN_SUBVERS 来自动设置相关的 SDK，但效果也必须向 AE 报告,它们是用 13.25 及以上版本编译的。

下表概述效果为支持新行为需要做的改变。

| MFR 以及新序列数据用法 | 需要在新SDK上的变动
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 效果没有设置 PF_OutFlag2_SUPPORTS_THREADED_RENDERING                                                        | 无需更改。效果和sequence_data照常工作。                                                                                                                                                                                                                                                                                                                                                                |
| 效果设置 PF_OutFlag2_SUPPORTS_THREADED_RENDERING 但在渲染期间没有读取和写入sequence_data | 使用2021年3月的SDK重新编译插件，无需其他代码更改。<br/>如果插件未使用2021年3月的SDK编译，则插件将从AE 22.0x6开始停止使用MFR。|                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| 插件设置 PF_OutFlag2_SUPPORTS_THREADED_RENDERING 但只在渲染期间读取sequence_data                     | 使用2021年3月的SDK重新编译插件，通过`PF_EffectSequenceDataSuite1` 更新读取 sequence_data 用于线程安全 .详见[在渲染时使用多帧渲染访问sequence_data](global-sequence-frame-data.html)                                                                                                                                                                    |
| 插件设置 PF_OutFlag2_SUPPORTS_THREADED_RENDERING 并在渲染期间读取和写入sequence_data            | 使用2021年3月的SDK重新编译插件并将插件修改为：<br />1.利用[计算缓存API](compute-cache-api.html) 用于线程安全的缓存访问，而不是直接读取/写入sequence_data。详情请参阅[计算多帧渲染缓存]。<br />AND / OR <br />2. 添加 `PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER` 恢复对sequence_data的直接读/写访问。 |

::: tip

使用 2021 年 3 月的 SDK 编译的效果，并使用 PF_OutFlag2_SUPPORTS_THREADED_RENDERING 开关，以及 PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER 开关，将与从 18.0 开始的 After Effects beta 版本一起工作，当时引入了`PF_EffectSequeceDataSuite1`。如果你需要支持两种 sequence_data 行为，请检查这个套件是否存在。

:::

## 多帧渲染对入口指令的影响

UI 入口指令仍然在主线程上发送，然而`PF_Cmd_SEQUENCE_SETUP`、`PF_Cmd_SEQUENCE_RESETUP`、`PF_Cmd_SEQUENCE_SETDOWN`、`PF_Cmd_SMART_PRE_RENDER`、`PF_Cmd_RENDER`和`PF_Cmd_SMART_RENDER`可能在处理 UI入口指令的同时在多个线程上发送，因此所有这些入口指令必须是线程安全。

`PF_Cmd_GLOBAL_SETUP`和`PF_Cmd_GLOBAL_SETDOWN`入口指令将只在主线程上发送，不会与任何其他入口指令同时发送。

## 多帧渲染中的序列数据

`sequence_data`对象和相关的序列入口指令多年来一直被用来提供一种在效果生命周期内存储数据的方法。多帧渲染引入了一些需要注意的变化。

### 2020 年 6 月起的变化

多帧渲染需要 After Effects 将 "序列数据 "传递给渲染线程。为了使需要用`PF_OutFlag_SEQUENCE_DATA_NEEDS_FLATTENING`开关的线性数据`sequence_data`生效，现在也必须设置`PF_OutFlag2_SUPPORTS_GET_FLATTENED_SEQUENCE_DATA'开关。

::: tip
在未来的 After Effects 版本中，强制要求设置`PF_OutFlag2_SUPPORTS_GET_FLATTENED_SEQUENCE_DATA`开关，并在插件中处理相关的入口指令。在加载不符合这一要求的效果时，会添加警告对话框。
:::

### 从 2021 年 3 月起的变化

- 在渲染时，"sequence_data "对象现在是常量，应该通过 "PF_EffectSequenceDataSuite "接口访问。
- 默认情况下，禁止在渲染时写入`seqeunce_data`，如果在渲染时试图写入`sequence_data`，结果将为undefined。
- 如果一个效果在渲染时必须写入 sequence_data，必须设置`PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER`开关，这将告诉 After Effects 允许写入`sequence_data`，但这会牺牲性能。当渲染开始时，"sequence_data "对象将被复制到每个渲染线程，每个渲染线程将有自己独立的 "sequence_data "副本，在渲染过程中进行管理。出于性能考虑，最好利用 <计算多帧渲染缓存> 来写入效果所需的任何数据。

## 计算多帧渲染缓存

计算缓存提供了一个线程安全的缓存，作为序列数据的替代或补充，效果可以在 Render 之前或期间计算、存储和读取数据。

### 何时需要计算缓存

- 如果你的效果使用`sequence_data`并且需要在 Render 过程中写入或更新`sequence_data`时，应该使用计算缓存，特别是所需数据计算很耗时。
- 如果没有计算缓存，效果需要添加`PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER`开关，此时每个渲染线程创建唯一的 sequence_data 副本。然后，每个渲染线程需要独立地进行计算，并且无法在渲染线程之间共享结果。
- 通过使用计算缓存，渲染线程可以分享计算数据的任务，并从已经计算过的数据中获益。
- 根据效果的需要，计算缓存 API 支持单核和多核的计算任务。参见[计算缓存 API](compute-cache-api.html)。

### 如何开启计算缓存

计算缓存 API 从 2021 年 3 月的 SDK 中开始提供，在 After Effects 2022 及以上版本中默认启用该套件。

参见[计算缓存 API](compute-cache-api.html) 文档，了解细节和示例代码。

## 什么是效果的线程安全

**当实现和共享数据保证没有竞争条件,并且在并发访问时始终处于正确状态时，效果是线程安全的。**。

更具体地说，效果：

1. 没有静态或全局变量,或者有"无竞争条件"的静态或全局变量。
2. 在渲染时不向`in_data->global_data`写入。可以进行读取。只在`PF_Cmd_GLOBAL_SETUP`和`PF_Cmd_GLOBAL_SETDOWN`中写入。
3. 在渲染时或`PF_Cmd_UPDATE_PARAMS_UI`事件中不向`in_data->sequence_data`写入。读取可以通过 _PF_EffectSequenceDataSuite_ 接口完成。

::: tip
如果效果使用任何阻塞同步机制，如互斥锁或门，则在回调主机时不得保留这些机制。常见的调用是在使用套件或进行签出调用时。否则很可能会导致死锁。
:::

## 如何定位效果中的静态和全局变量

为了帮助你定位效果中的静态变量和全局变量，我们开发了一个**静态分析工具**供你使用。
可以在[https://github.com/adobe/ae-plugin-thread-safety](https://github.com/adobe/ae-plugin-thread-safety)找到这个工具。

### 如果在 Mac 上开发

1.在上面提供的 URL 中克隆/下载 Git Repo

2.在**Mac**文件夹中找到 bash 脚本`check_symbols_for_thread_safety.sh`

3.在插件或效果的包内容中，找到二进制文件。(例如，**Curves.plugin**的二进制文件在`/Applications/Adobe After Effects [your AE version]/Plug-ins/Effects/Curves.plugin/Contents/MacOS/Curves`)

4.分析二进制文件，运行。

```cpp
check_symbols_for_thread_safety.sh [Binary location]
For example, check_symbols_for_thread_safety.sh /Applications/Adobe After Effects [your AE version]/Plug-ins/Effects/Curves.plugin/Contents/MacOS/Curves)
```

5.你会看到该工具的输出是这样的格式。

```cpp
[symbol type]; [symbol name]
```

6.`[symbol type]`是一个区分大小写的字母，表示变量的类型。可以在这里找到所有的类型信息。[https://linux.die.net/man/1/nm](https://linux.die.net/man/1/nm)

7.下面是一个输出的例子。

```cpp
b; Deform::FindSilEdges()::new_kInfinite
```

`b`显示这个符号在未初始化的数据部分，这表明它可能是一个静态变量。

`Deform::FindSilEdges()::new_kInfinite`是符号名称，其中`Deform`是变量所在的名字空间的名称。

`FindSilEdges()`是该变量所在的函数名称。

`new_kInfinite`是实际的变量名称。名称空间和函数名称可能不会根据变量所在的位置来显示。

8.在你的代码中搜索每个符号，修复它(见[这里]中的方法)，并对你的解决方案中的每个二进制文件进行重复。

### 在 Windows 上开发

#### 1.准备

0. **为了运行这个工具，需要先安装 Visual Studio**。
1. 克隆/下载上面的 Git Repo
2. 在**Win**文件夹中找到`register_msdia.cmd`脚本。
3. 从**开始菜单中搜索** "x64 Native Tools Command Prompt for VS...."。
4. 右击 -> 以管理员身份运行
5. 在终端，`cd`到`register_msdia.cmd`所在的目录。
6. 运行`.\register_msdia.cmd`。
7. 这个脚本将为你注册**DIA SDK**和其他一些必要的依赖项
8. 现在静态分析器应该可以工作了

#### 2. 使用 Windows 静态分析器

1. 在**Win**文件夹中找到可执行文件`CheckThreadSafeSymbols.exe`
2. 在**调试**模式下编译你的效果，找到其**.pdb**文件
3. 如果你没有修改你的项目构建设置，你也应该在同一个构建目录中找到一些**.obj**文件
4. 你有两个选择：使用`-objfile`或`-source`开关扫描二进制文件或源文件。

- 你可以从任何一个选项中获得相同的符号。
- 如果不知道源代码最终会出现在什么二进制文件中，或者想在每个源文件的基础上跟踪线程安全，请使用`-source`选项。
- 如果想更精细地控制你的项目的哪些部分被扫描，请使用`-objfile`选项。

5.要分析一个对象文件中的符号，请运行。

```cpp
CheckThreadSafeSymbols.exe -objfile [要分析的二进制文件绝对路径] [相对于.pdb的绝对路径]
```

6.要分析源文件中的符号，请运行。

```cpp
CheckThreadSafeSymbols.exe -source [要分析的源文件绝对路径] [相对于.pdb的绝对路径]
```

7.全局变量并不局限于 pdbs 中的一个文件或二进制文件的范围，所以你必须检查所有项目的全局变量列表，而不需要过滤。使用-g 输出来获得所有的列表。

```cpp
CheckThreadSafeSymbols.exe -g [相对于.pdb的绝对路径]
```

8.如果不确定效果输出什么二进制文件，该工具也可以输出一个**(noisy)**的二进制文件列表，以及每个二进制文件提取数据的源文件。修改过的文件可能会在最上面的位置。要看到这个列表，请运行。

```cpp
CheckThreadSafeSymbols.exe -sf [相对于.pdb的绝对路径]
```

9.输出的符号将采取以下形式。

```cpp
[Symbol name], [Symbol type], [Datakind], ([Section type of data location], [Binary Address][Binary Address Offset])
```

10.下面是一个输出样例。

```cpp
menuBuf, Type: char[0x1000], File Static, (static, [0008FCD0][0003:00001CD0])
```

> `menuBuf`是实际的变量名称。
>
> `Type: char[0x1000]`显示它是什么类型的变量。这里的数据是一个`char'。
>
> `File Static`显示它是什么类型的数据。这里的数据是一个**文件范围的静态变量。**可以在这个页面上找到所有的数据种类和含义[https://docs.microsoft.com/en-us/visualstudio/debugger/debug-interface-access/datakind?view=vs-2019](https://docs.microsoft.com/en-us/visualstudio/debugger/debug-interface-access/datakind?view=vs-2019)
>
> `static`表明该数据在内存的静态部分。
>
> `[0008FCD0][0003:00001CD0]`显示数据的二进制地址和二进制地址偏移。

11.在代码中搜索每个符号，修复它(关于如何修复，见下)，并对解决方案中的每个二进制/源文件重复进行。

## 如果效果中有静态和全局变量，该怎么办？

当你看到一个静态或全局变量时，如果可能的话，最好是把它变成一个局部变量。但这个变量必须是静态或全局的呢？

下面是一些标准的处理方法：

1.是否可以在不改变行为的情况下，在函数之间轻松地传递数据呢？

```cpp
// 非线程安全代码的例子

static int should_just_be_local;

void UseState() {
  DoComputation(should_just_be_local);
}

void SetAndUseState() {
  should_just_be_local = DoComputation();
  UseState();
}
```

**要么把它添加到一个结构中，要么扩展函数参数以包括它**。

```cpp
// 我们可以通过函数参数传递should_just_be_local变量来修复上面的代码

void UseState(int should_just_be_local) {
  DoComputation(should_just_be_local);
}

void SetAndUseState() {
  int should_just_be_local = DoComputation();
  UseState(should_just_be_local);
}
```

2.能否在执行代码前对数据进行初始化(如查找表、常量变量)

```cpp
// 非线程安全代码的例子

// 代码中的许多地方需要读取这个表，但不会写到这个表上
static int state_with_initializer[64];

static bool state_was_initialized = false;

void InitializeState() {
  for (int i = 0; i < 64; ++i) {
    state_with_initializer[i] = i * i;
  }
  state_was_initialized = true;
}

void Main() {
if (!state_was_initialized) {
InitializeState();
}
DoComputation(state_with_initializer);
}
```

 **使之成为** `const` **或用宏代替它**

 ```cpp
std::array<int, 64> InitializeState() {
  std::array<int, 64> temp;
  for (int i = 0; i < 64; ++i) {
    temp[i] = i * i;
  }
  return temp;
}

// We can fix the above code by making this table a const and initialize it before using it
static const std::array<int, 64> state_with_initializer = InitializeState();

void Main() {
  DoComputation(state_with_initializer);
}
```

3.数据是否在运行时初始化一次，基于在后续渲染中不会改变的数据?

```cpp
// Example of a non Thread-Safe code
static int depends_on_unchanging_runtime_state;

void UseState() {
  DoComputation(depends_on_unchanging_runtime_state);
}

void SetAndUseState() {
  depends_on_unchanging_runtime_state = DoComputationThatNeedsStateOnlyOnce();
  UseState();
}
```

**仔细检查在代码执行之前是否知道此状态（案例2），但是如果您必须在运行时初始化，请使用const静态本地。（注意，静态本地对象的线程安全初始化是C++规范的一部分）**

```cpp
void UseState(int depends_on_unchanging_runtime_state) {
              DoComputation(depends_on_unchanging_runtime_state);
      }

void SetAndUseState() {

  // We can fix the above code by making the variable a const static local
  static const int depends_on_unchanging_runtime_state = DoComputationThatNeedsStateOnlyOnce();

  UseState(depends_on_unchanging_runtime_state);
}
```

4.数据必须保持静态/全局，而不是一个常量。但每个渲染线程可以有自己的数据副本。

```cpp
// This variable has to be static and not a const
static int this_thread_needs_access;

void SetState(int new_state) {
  this_thread_needs_access = new_state;
}

void UseState() {
  DoComputation(this_thread_needs_access);
}
```

**只要让变量thread_local**。

```cpp
// Make this variable a thread_local variable
thread_local static int this_thread_needs_access;

void SetState(int new_state) {
  this_thread_needs_access = new_state;
}

void UseState() {
  DoComputation(this_thread_needs_access);
}
```

**5. 数据必须保持静态/全局而不是常量，每个线程都需要从最新的状态中读写。(罕见)**

```cpp
// This variable has to be static and not a const
// It also needs to be shared across several threads
static int every_thread_needs_latest_state;

void SetState(int new_state) {
  every_thread_needs_latest_state = new_state;
}

void UseState() {
  DoComputation(every_thread_needs_latest_state);
}
```

**在这种情况下，用一个mutex来保护访问。

```cpp
// Add a mutex (lock)
static std::mutex ex_lock;

static int every_thread_needs_latest_state;

void SetState(int new_state) {
  {
    // Protect the access with the mutex (lock)
    std::lock_guard<std::mutex> lock(ex_lock);
    every_thread_needs_latest_state = new_state;
  }
}

void UseState() {
  int state_capture;
  {
    // Protect the access with the mutex (lock)
    std::lock_guard<std::mutex> lock(ex_lock);
    state_capture = every_thread_needs_latest_state;
  }
  DoComputation(state_capture);
}
```

::: tip

以上例子是效果中的常见情况。你可以随时想出其他方法来处理你的静态和全局，以最适合你的需要。

:::

## 将效果设置为线程安全

- 在 GlobalSetup 中设置`PF_OutFlag2_SUPPORTS_THREADED_RENDERING`开关，告诉 After Effects 你的效果是线程安全的，支持多帧渲染。
- 如果需要，添加`PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER`以允许在 Render 阶段写入 sequence_data。
- 更新`AE_Effect_Global_OutFlags_2`魔法数字。在不改变魔法数字的情况下第一次用你的效果启动 AE，应用你的效果，AE 会给你一个正确的数字来输入。
- 如果你使用`PF_OutFlag_SEQUENCE_DATA_NEEDS_FLATTENING`开关，记得也要设置`PF_OutFlag2_SUPPORTS_GET_FLATTENED_SEQUENCE_DATA`开关。

## 如何测试效果是否为线程安全

一旦完成了上述步骤，使你的效果成为线程安全的，现在应该准备做一些测试。

### 开启多帧渲染

- 多帧渲染在 After Effects 2022 中是默认启用的。
- 要打开或关闭多帧渲染，请导航到首选项>内存和性能>性能，然后切换多帧渲染复选框。

### 测试效果

一旦完成了上述准备步骤，请彻底测试你的效果。你应该能够测试简单和复杂的合成，并看到性能的改善，因为效果利用了多帧渲染。

- 仔细检查你现有的所有手动和自动测试计划。
- 测试所有的效果参数，确保它们能正常工作。
- 适当地加入一些已经被做成线程安全的 AE 效果。参见[线程安全的第一方效果]部分。
- 确保在启用多帧渲染时，不会出现崩溃、挂起、渲染差异或其他意外变化。

## 线程安全第一方效果

请访问[https://helpx.adobe.com/after-effects/user-guide.html/after-effects/using/effect-list.ug.html](https://helpx.adobe.com/after-effects/user-guide.html/after-effects/using/effect-list.ug.html)了解 MFR 支持的效果的完整列表。每周都会有更多的效果加入。
