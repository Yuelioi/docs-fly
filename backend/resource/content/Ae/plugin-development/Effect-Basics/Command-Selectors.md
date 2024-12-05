---
title: Command-Selectors
order: 1
---

# 入口指令

指令，简单地说，是 After Effects 希望你的效果做什么。除了一些必要的入口指令；大多数是可选的，尽管这些指令也有作用。

随着每个入口指令的发送，效果会从 After Effects 的[PF_InData](../effect-basics/PF_InData.html)中接收信息，在 PF_ParamDef[]中接收输入和参数值(一个包括输入图层参数的描述数组)，以及访问回调和函数套件。

在[PF_OutData](../effect-basics/PF_OutData.html)中向 After Effects 发送信息，并(在适当的时候)渲染输出到 PF_LayerDef，也称为[PF_EffectWorld](../effect-basics/PF_EffectWorld.html)。

在事件发生期间，会在[PF_EventExtra](../effect-ui-events/PF_EventExtra.html)中接收事件的具体信息。

## 调用序列

只有最初的几个入口指令是按序调用的；其余的调用顺序是由用户操作决定的。首次应用时，插件会收到 _PF_Cmd_GLOBAL_SETUP_ ,然后是 _PF_Cmd_PARAM_SETUP_ 。每次用户将效果添加到图层时，会发送 _PF_Cmd_SEQUENCE_SETUP_ 。

基本的非 SmartFX 效果渲染的每一帧，After Effects 会发送 _PF_Cmd_FRAME_SETUP_ , 再发送 _PF_Cmd_RENDER_ ,然后是 _PF_Cmd_FRAME_SETDOWN_ .

所有的效果插件必须响应 _PF_Cmd_RENDER_ .

对于 SmartFX，_PF_Cmd_SMART_PRE_RENDER_ 可以被发送任意次数，然后再发送单个 _PF_Cmd_SMART_RENDER_ .

当用户删除效果或关闭项目时, 会发送 _PF_Cmd_SEQUENCE_SETDOWN_ . 在项目加载时或应用图层发生变化时发送 _PF_Cmd_SEQUENCE_RESETUP_ . 在 After Effects 项目被写出到磁盘时发送 _PF_Cmd_SEQUENCE_FLATTEN_ .

在用户从效果控制窗口(ECW)选择 _关于..._ 时发送 _PF_Cmd_ABOUT_ .

_PF_Cmd_GLOBAL_SETDOWN_当 After Effects 关闭时，或者当效果的最后一个实例被删除时，将发送该信息。不要依靠这个消息来确定你的插件何时被从内存中删除，要使用操作系统特定的入口函数。

## 入口指令表

### 全局入口指令 Global Selectors

所有的插件都必须对这些入口指令做出反应。

| **入口指令**  | **响应**  |
| --- | --- |
| PF_Cmd_ABOUT  | 插件的"关于"对话框。在 out_data>return_msg 里, 接着 AE 会显示一个简单的对话框。<br/>在对话框中包含插件的版本信息。在 macOS 上，当前资源文件将在此选择期间设置为您的效果模块。  |
| PF_Cmd_GLOBAL_SETUP  | 设置任何需要的开关(flag)以及 _PF_OutData_ 描述插件行为的结构体成员(fields)(包括 out_data>my_version)。  |
| PF_Cmd_GLOBAL_SETDOWN | 释放所有全局数据(仅当分配了数据时才需要)。  |
| PF_Cmd_PARAM_SETUP  | 描述你的参数并用[PF_ADD_PARAM](../effect-details/interaction-callback-functions.html)注册参数.<br /> 并且注册自定义用户界面元素 <br />设置 [PF_OutData > num_params](../effect-basics/PF_OutData.html) 匹配你的参数数量 |

## 序列入口指令

控制序列数据的处理。

| **入口指令**  | **响应**  |
| --- | --- |
| PF_Cmd_SEQUENCE_SETUP  | 分配和初始化任何序列的数据。在首次应用效果时发送。同时初始化[PF_InData](../effect-basics/PF_InData.html)也  |
| PF_Cmd_SEQUENCE_RESETUP | 重新创建序列数据(通常是用来解线性储存转换的(unflatten))。在从磁盘读取序列数据后、预合成或复制效果时发送；<br/>拷贝前 Ae 把序列数据进行线性储存转换。拷贝副本时，新旧序列都会收到 _PF_Cmd_SEQUENCE_RESETUP_ . <br/>在 _PF_Cmd_SEQUENCE_RESETUPs_ 期间不会有 _PF_Cmd_SEQUENCE_FLATTEN_ .  |
| PF_Cmd_SEQUENCE_FLATTEN | 保存和拷贝序列时会发送这个指令。线性储存转换(flatten)好的序列数据包含指针或句柄，所以可以被写入硬盘。<br/>这些数据会和工程一起保存。释放解转换的数据（unflat data），设置 _out_data>sequence_data_ 指向新的线性转换好的数据。为了存储，线性数据（flat data）必须按照正确的字节顺序排好。从 6.0 开始，如果一个效果的序列数据近期被线性储存转换(flatten)了，而插件可能没有收到附加的 _PF_Cmd_SEQUENCE_SETDOWN_ 就被删除了。在这种情况下，AE 会处理你的线性序列数据（flat sequence data）。 |
| PF_Cmd_SEQUENCE_SETDOWN | 释放所有序列数据  |

### 帧入口指令

传递给每一帧(或一组音频样本)，由你的插件来渲染。

| **入口指令**  | **Response**  |
| --- | ----- |
| PF_Cmd_FRAME_SETUP  | 可以利用这个指令分配些特定帧的数据——在每一帧被渲染前就会发送一次这个指令，以便设置特定帧的数据。如果你的效果改变输出缓冲区的大小，请指定新的输出高度，宽度和参考源(relative origin)。除输入层(input layer)外的所有参数都是有效的。<br/>如果将宽度和高度设置为 0，After Effects 将忽略对 _PF_Cmd_RENDER_ 的响应。<br/>提示：如果指定了对指令 _PF_Outflag_I_EXPAND_BUFFER_ 的响应函数你会收到两次该指令（和_PF_Cmd_FRAME_SETDOWN_指令），其中一次没有 _PF_Cmd_RENDER_ 指令。<br/>这样我们就知道提供的图层会不会可见。。<br/>帧数据可以追溯到机器可能有 8MB 内存的日子。给定调用序列(如上)，仅在_PF_Cmd_RENDER_ 期间分配效率要高得多。  |
| PF_Cmd_RENDER  | 根据输入帧和任何参数将效果渲染到输出中。<br/>此渲染调用只能支持每通道 8 位或 16 位渲染。每个通道 32 位渲染必须在 _PF_Cmd_SMART_RENDER_ 中处理。<br/> PF_InData 中的所有结构体成员都有效。<br/>如果您对此入口指令的响应被中断( _PF_ABORT_ 或 _PF_PROGRESS_ 的调用返回错误代码)，结果将不会被使用。<br/>在此入口指令期间不能删除 frame_data；必须等到 PF_Cmd_FRAME_SETDOWN。  |
| PF_Cmd_FRAME_SETDOWN  | 释放 _PF_Cmd_FRAME_SETUP_ 期间分配的任何帧数据。  |
| PF_Cmd_AUDIO_SETUP  | 在每次音频渲染之前发送。请求输入音频的时间跨度。分配和初始化任何特定于序列的数据。<br/>如果效果需要来自输出时间跨度以外的时间跨度的输入，请更新 PF_OutData 中的 startsampL 和 endsampL 结构体成员。  |
| PF_Cmd_AUDIO_RENDER  | 使用效果编辑的音频填充[PF_OutData.dest_snd](…/effect-basics/PF_OutData.html)。PF_InData 中的所有结构体成员都有效。<br/>如果对此入口指令的响应被中断(对 _PF_ABORT_ 或 _PF_PROGRESS_ 的调用返回错误代码)，结果将不会被使用。  |
| PF_Cmd_AUDIO_SETDOWN  | 释放 _PF_Cmd_AUDIO_SETUP_ 期间分配的内存  |
| PF_Cmd_SMART_PRE_RENDER | 仅限 SmartFX。根据效果实现的任何标准，确定效果产生输出所需的输入区域。<br/>在 MediaCore 托管时可能最多发送两次。第一次在 GetFrameDependency 期间来收集输入。<br/>源校验 (checkouts) 可以在此处返回完整的帧尺寸。一旦源被渲染，如果它们的大小与第一次调用不同，那么这个入口指令将与实际的源大小一起在第二次发出，以获得正确的输出大小。<br/>请注意，MediaCore 想要所有的输出，因此将使用 PF_PreRenderOutput::max_result_rect。<br/> **16.0 中的新功能** <br/>在 PF_PreRenderOutput 中设置 _PF_RenderOutputFlag_GPU_RENDER_POSSIBLE_ 以在 GPU 上渲染。<br/>如果未设置此标志，则请求的渲染不可能与请求的 GPU 一起使用，因为参数或渲染设置。<br/>主机可以使用另一个 what_gpu 选项(或 PF_GPU_Framework_None)重新调用 PreRender。 |
| PF_Cmd_SMART_RENDER  | 仅限 SmartFX。执行渲染并为要求渲染的效果区域提供输出。  |

PF_Cmd_SMART_PRE_RENDER 的 struct

```cpp
typedef struct {
PF_RenderRequest output_request; // what the effect is being asked to render
short bitdepth; // bitdepth the effect is being driven in (in bpc)
const void *gpu_data; // (new AE 16.0)
PF_GPU_Framework what_gpu; // (new AE 16.0)
A_u_long device_index; // (new AE 16.0) For use in conjunction with PrSDKGPUDeviceSuite
} PF_PreRenderInput;
```

### 通信

After Effects 和插件之间的通信通道。

| 入口指令  | Response  |
| --- | ----- |
| PF_Cmd_EVENT  | 此入口指令使用额外的参数；要处理的事件类型由 e_type 字段指示，该字段是额外指向的结构体成员。<br/>请参阅[效果 UI 和事件](…/effect-ui-events/effect-ui-events.html)。  |
| PF_Cmd_USER_CHANGED_PARAM  | 用户更改了参数值。只有设置了 _PF_ParamFlag_SUPERVISE_ 标志，才会收到此命令。<br/>修改参数以控制值，或者使一个参数值影响其他参数值。参数可以通过不同的操作来修改。<br/>`in_data.current_time`设置为用户在用户界面中查看的帧的时间(内部，将合成的当前时间转换为图层时间)，同时他们正在更改触发 _PF_Cmd_USER_CHANGED_PARAM_ 的参数。<br/>它也是自动添加关键帧的时间(如果还没有关键帧，并且启用了秒表)。<br/>这通常与传递给紧接其后的 PF_Cmd_RENDER 的值相同(除非大写锁定关闭)，但不一定---可能会打开其他合成窗口，从而在不同的时间响应更改参数  |
| PF_Cmd_UPDATE_PARAMS_UI  | 效果控制面板(ECP)需要更新。可能发生在打开 ECP 或移动到合成中的新时间之后。<br/>可以通过调用 _PF_UpdateParamUI()_ 来修改参数特征(例如启用或禁用它们)。<br/>响应此命令只能进行表面更改。响应 _PF_Cmd_UPDATE_PARAMS_UI_ 时不要更改参数值；而是在 _PF_Cmd_USER_CHANGED_PARAM_ 期间更改。<br/>如果在 PiPL 中在 _PF_Cmd_GLOBAL_SETUP_ 期间设置了 _PF_OutFlag_SEND_UPDATE_PARAMS_UI_ ,此命令只会定期发送.<br/>注意：在此入口指令期间永远不要校验(check out)参数。几乎肯定会导致递归不良。  |
| PF_Cmd_DO_DIALOG  | 显示选项对话框。这是在单击选项按钮(或选择了菜单命令)时发送的。<br/>只有当效果先前指示它有一个对话框<br/>(通过设置全局 _PF_OutFlag_I_DO_DIALOG_ 标志来响应 _PF_Cmd_GLOBAL_SETUP_ )时，才会发送此入口指令。<br/>在 3.x 版中， _PF_Cmd_DO_DIALOG_ 传递的参数无效。<br/>不再是这种情况；插件可以访问非层参数，在其他时间签出参数，并在 _PF_Cmd_DO_DIALOG_ 期间执行 UI 更新。<br/>他们仍然可能不会更改参数值。  |
| PF_Cmd_ARBITRARY_CALLBACK  | 管理任意数据类型。只有在注册了自定义数据类型参数时，才会收到这个参数。<br/>额外的参数指示正在调用哪个处理函数。<br/>自定义数据类型将在实现任意数据中进一步讨论。  |
| PF_Cmd_GET_EXTERNAL_DEPENDENCIES | 仅当 _PF_Cmd_GLOBAL_SETUP_ 期间设置了 _PF_OutFlag_I_HAVE_EXTERNAL_DEPENDENCIES_ 时才发送。<br/>填充一个字符串句柄(在额外指向的 _PF_ExtDependenciesExtra_ 中)，其中包含插件依赖项的描述，确保为终止的 NULL 字符分配空间。<br/>如果没有要报告的依赖项，则仅返回字符串句柄的 NULL 指针。<br/>如果检查类型 _PF_DepCheckType_ALL_DEPENDENCIES_ ,则报告插件渲染所需的一切。<br/>如果检查类型 PF_DepCheckType_MISSING_DEPENDENCIES，则只报告缺失的项目(如果没有缺失，则报告空字符串)。  |
| PF_Cmd_COMPLETELY_GENERAL  | 响应 AEGP。额外的参数指向 AEGP 发送的任何参数。<br/>AEGP 只能与响应此入口指令的效果进行通信。  |
| PF_Cmd_QUERY_DYNAMIC_FLAGS  | 仅发送给在 _PF_OutFlags2_ 、PiPL 和 _PF_Cmd_GLOBAL_SETUP_ 期间指定 _PF_OutFlag2_SUPPORTS_QUERY_DYNAMIC_FLAGS_ 的插件。<br/>对于所有动态标志，如果您在此命令期间更改它们，您必须在 PF_Cmd_GLOBAL_SETUP 期间设置标志。<br/>此入口指令将在任意时间发送。<br/>作为响应，效果应使用 _PF_CHECKOUT_PARAM_ 访问其(非层)参数，并决定是否应设置任何支持 _PF_Cmd_QUERY_DYNAMIC_FLAGS_ 的标志，例如：<br/>_PF_OutFlag_WIDE_TIME_INPUT_<br/> _PF_OutFlag_NON_PARAM_VARY_<br/> _PF_OutFlag_PIX_INDEPENDENT_<br/> _PF_OutFlag_I_USE_SHUTTER_ANGLE_ <br/> _PF_OutFlag2_I_USE_3D_CAMERA_ <br/>_PF_OutFlag2_I_USE_3D_LIGHTS_<br/> _PF_OutFlag2_DOESNT_NEED_EMPTY_PIXELS_<br/>_PF_OutFlag2_REVEALS_ZERO_ALPHA_<br/> _PF_OutFlag2_DEPENDS_ON_UNREFERENCED_MASKS_ <br/>_PF_OutFlag2_OUTPUT_IS_WATERMARKED_<br/>'After Effects 将此信息用于缓存和优化目的，因此请尝试尽快响应。 |
| PF_Cmd_GPU_DEVICE_SETUP  | 主机可以随时调用此入口指令。每个 GPU 设备只能调用一次。<br/>多个 GPU 设备可能同时处于设置状态。<br/>它将在 GlobalSetup 之后和 SequenceSetup 之前调用。<br/>其目的是让效果在必要时进行 GPU 初始化，并让效果有机会仅根据该设备的属性而不是任何渲染上下文(帧大小等)选择退出 GPU 设备。<br/>如果效果拒绝 GPU 设备，它将被调用进行 CPU 渲染。<br/>_PF_InData::what_gpu！=PF_GPU_Framework_None_ 是预期的。<br/>如果 what_gpu 中的设备和框架受支持，效果预计会在 PF_OutData::中设置一个或两个 PF_OutFlag2_SUPPORTS_GPU_RENDER_Fxx 标志 out_flags2。<br/>请注意，AE 16.0 中只有 PF_OutFlag2_SUPPORTS_GPU_RENDER_F32。<br/>此处不设置标志的效果将不被视为支持这些设备中任何一个的 GPU 渲染。<br/>PF_GPUDeviceSetupOutput::gpu_data 是插件拥有的指针，必须使用 _PF_Cmd_GPU_DEVICE_SETDOWN_ 入口指令释放。<br/>此指针在渲染时也可用。  |
| PF_Cmd_GPU_DEVICE_SETDOWN  | 释放与 gpu_data 相关的任何资源。在 AE 中，这将在 GPU 设备发布之前调用。  |
| PF_Cmd_GPU_SMART_RENDER_GPU  | GPU 等效于现有的 _PF_Cmd_SMART_RENDER_ 入口指令。<br/>在渲染时，_PF_Cmd_SMART_RENDER_ 或 _PF_Cmd_SMART_RENDER_GPU_ 入口指令将被调用，具体取决于效果是否会产生 CPU 或 GPU 帧作为输出。<br/>PF_Cmd_SMART_RENDER_GPU 只会在 what_gpu！=PF_GPU_Framework_None 时被调用，并对任何输入/输出 PF_LayerDef 有影响。<br/>当此入口指令正在进行时，所有帧签入和签出都将在 GPU 帧上操作。注意 PF_Cmd_SMART_RENDER 共享额外的结构。<br/>what_gpu 和 device_index 字段在 GPU 相关入口指令的额外输入中，向插件指示要用于渲染的 GPU 框架。<br/>输入和输出缓冲区将在此框架和设备上准备。<br/>设备、上下文、命令队列和其他相关的 GPU 状态可以用 PrSDKGPUDeviceSuite::GetDeviceInfo 查询。<br/> _PF_Cmd_SMART_PRE_RENDER_ 和 _PF_Cmd_SMART_RENDER_GPU_ 入口指令调用之间 what_gpu 相同。  |

## 区别

 _PF_Cmd_USER_CHANGED_PARAM_ 和 _PF_Cmd_UPDATE_PARAMS_UI_ 之间有一个微妙的区别。

效果需要区分用户实际改变参数值(_PF_Cmd_USER_CHANGED_PARAM_)，和只是在时间轴上刷屏(_PF_Cmd_UPDATE_PARAMS_UI_)，这也是在插件第一次加载时发送的。

只有前几个入口指令是可预测的；其余的调用顺序由用户操作决定。

第一次应用时，插件会收到_PF_Cmd_GLOBAL_SETUP_ ,然后是_PF_Cmd_PARAM_SETUP_ . 每次用户将效果添加到一个图层时，_PF_Cmd_SEQUENCE_SETUP_被发送。

对于基本的非 SmartFX 效果渲染的每一帧，After Effects 会发送_PF_Cmd_FRAME_SETUP_ ,然后是_PF_Cmd_RENDER_ ,然后是_PF_Cmd_FRAME_SETDOWN_ . 所有的效果插件必须响应_PF_Cmd_RENDER_ .

对于 SmartFX，_PF_Cmd_SMART_PRE_RENDER'可以被发送任意次数，然后再发送单个_PF_Cmd_SMART_RENDER'。

_PF_Cmd_SEQUENCE_SETDOWN_是在退出时发送的，当用户删除效果或关闭项目时。_PF_Cmd_SEQUENCE_RESETUP_在项目加载时或应用的层发生变化时发送。_PF_Cmd_SEQUENCE_FLATTEN_是在 After Effects 项目被写入磁盘时发送的。

_PF_Cmd_ABOUT_在用户从效果控制窗口(ECW)选择_About..._时发送。

_PF_Cmd_GLOBAL_SETDOWN_当 After Effects 关闭时，或者当效果的最后一个实例被删除时，将发送该信息。不要依靠这个消息来确定你的插件何时从内存中被删除；使用操作系统特定的入口函数。

## 示例

```cpp
PF_Err
EffectMain(
 PF_Cmd   cmd,
 PF_InData  *in_data,
 PF_OutData  *out_data,
 PF_ParamDef  *params[],
 PF_LayerDef  *output,
 void   *extra)
{
PF_Err  err = PF_Err_NONE;
 
try {
  switch (cmd) {
    case PF_Cmd_ABOUT:
    err = About(in_data,
       out_data,
       params,
       output);
    break;
    
    case PF_Cmd_GLOBAL_SETUP:
    err = GlobalSetup( in_data,
         out_data,
         params,
         output);
    break;
    
    case PF_Cmd_PARAMS_SETUP:
    err = ParamsSetup( in_data,
         out_data,
         params,
         output);
    break;
    
    case PF_Cmd_RENDER:
    err = Render( in_data,
        out_data,
        params,
        output);
    break;

    case PF_Cmd_USER_CHANGED_PARAM:
    err = paramChanged( in_data,  // 自己写参数变化函数
    out_data,
    (PF_SmartRenderExtra*)extra);
    break;
  }
 }
 catch(PF_Err &thrown_err){
  err = thrown_err;
 }
 return err;
}

```
