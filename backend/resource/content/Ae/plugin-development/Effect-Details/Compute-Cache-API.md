---
title: Compute-Cache-API
order: 6
---

# 计算缓存API

计算缓存 API 提供了一个线程安全的缓存，作为序列数据的替代或补充，效果可以在 Render 之前或期间计算、存储和读取数据。它应该被用来缓存那些计算起来很耗时的数据。对于多帧渲染的效果来说，它可以通过消除跨线程的冗余计算而带来很大的好处。该缓存与 After Effects 中的其他缓存是统一的，因此内存的使用在其他缓存中是平衡的。该模型还支持用户使用参数进行 A/B 测试，并且缓存状态在 A 和 B 状态下都会持续存在，从而加快了工作流程。最后两个设计特点对单帧和多帧的渲染效果都有好处。

计算缓存在 AEGP_ComputeCache 套件中实现，可以通过`AEGP_ComputeCacheSuite1`和`AEGP_ComputeCacheCallbacks`访问。

## AEGP_ComputeCacheSuite1

### AEGP_ClassRegister

使用计算类的全球唯一标识符注册缓存类型，例如 "adobe.ae.effect.test_effect.cache_v_1"。

一个 `AEGP_ComputeCacheCallbacks`类型的对象应该被设置为`AEGP_ComputeCacheSuite1`所要求的回调方法的函数指针。

这个函数通常在 PF_Cmd_GLOBAL_SETUP`期间调用，但也可以在任何时候调用。

```cpp
A_Err (*AEGP_ClassRegister)(
  AEGP_CCComputeClassIdP  compute_classP,
  const AEGP_ComputeCacheCallbacks  *callbacksP);
```

### AEGP_ClassUnRegister

使用计算类的全局唯一标识符，取消注册一个先前注册的缓存类型。

此时将通过调用delete_compute_value清除所有缓存值。

此函数通常会在PF_Cmd_GLOBAL_SETDOWN期间调用，但可以随时调用。

```cpp
A_Err(*AEGP_ClassUnregister)(
AEGP_CCComputeClassIdPcompute_classP);
```

### AEGP_ComputeIfNeededAndCheckout

这是主要的签出调用，用于计算和/或返回 AEGP_CCCheckoutReceiptP`的缓存条目的接收指针。

传入在 `AEGP_RegisterClass`方法中使用的`AEGP_CCComputeClassIdP`。

`AEGP_CCComputeOptionsRefconP`对象将根据需要传递给`AEGP_ComputeCacheCallbacks`、`generate_key`和`compute`方法。这个对象的类型对`AEGP_ComputeCacheSuite1`是不透明的，需要通过`generate_key`和`compute`的效果实现进行适当的铸造。

当需要计算缓存值时，`wait_for_other_threadB bool`会被使用。当设置为 "true "时，该方法将始终执行计算步骤或将完成的凭证返回到缓存中。当设置为 "false "时，该方法将完成计算步骤，除非其他线程已经在计算缓存条目，在这种情况下，将返回 A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING`。关于这个参数的更多信息，请参见[wait_for_other_threadB 对 AEGP_ComputeIfNeededAndCheckout 的影响](../effect-details/compute-cache-api.html)。

`CCCheckoutReceiptP`是一个不透明的指针，然后可以传入`AEGP_GetReceiptComputeValue`中，从缓存中获取计算值的指针。

```cpp
A_Err(*AEGP_ComputeIfNeededAndCheckout)(
AEGP_CCComputeClassIdPcompute_classP,
AEGP_CCComputeOptionsRefconPopaque_optionsP,
boolwait_for_other_threadB,
AEGP_CCCheckoutReceiptP*compute_receiptPP);
```

### AEGP_CheckoutCached

使用这个方法来检查缓存值是否已经被计算，如果有的话，返回 AEGP_CCCheckoutReceiptP`凭证。

如果缓存没有被计算，将返回 A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING`。

```cpp
A_Err(*AEGP_CheckoutCached)(
AEGP_CCComputeClassIdPcompute_classP,
AEGP_CCComputeOptionsRefconPopaque_optionsP,
AEGP_CCCheckoutReceiptP*compute_receiptPP);
```

### AEGP_GetReceiptComputeValue

使用此方法从计算方法中获取缓存值。

传入从 `AEGP_ComputeIfNeededAndCheckout`或`AEGP_CheckoutCached`收到的凭证。

返回的 `CCCputeValueRefconP`应该被投到计算方法中使用的正确对象类型。

```cpp
A_Err(*AEGP_GetReceiptComputeValue)(
constAEGP_CCCheckoutReceiptPcompute_receiptP,
AEGP_CCComputeValueRefconP*compute_valuePP);
```

### AEGP_CheckinComputeReceipt

在使用检查出来的、计算出来的缓存值完成效果代码后，在返回主机之前，调用这个方法，传入从 `AEGP_ComputeIfNeededAndCheckout`或`AEGP_CheckoutCached`返回的凭证。

如果传入的凭证是无效的，将返回错误 `A_Err_STRUCT`。一个弹出式的错误对话框也将显示，**"试图检入无效的凭证。请确保你没有重复签入或签入无效的凭证。"

```cpp
A_Err(*AEGP_CheckinComputeReceipt)(
AEGP_CCCheckoutReceiptPcompute_receiptP);
```

## AEGP_ComputeCacheCallbacks

效果必须为这些回调提供实现。

### generate_key

在创建缓存条目和进行缓存查询时被调用。应该是快速计算的。所有需要唯一地址的缓存条目的输入都必须被散列到密钥中。如果需要层层检查来计算缓存值，比如用直方图，那么必须包括该输入的哈希值。参见 PF_ParamUtilsSuite::PF_GetCurrentState`以获得一个层参数的哈希值。注意这是生成画面所需的输入的哈希值，而不是画面中的像素的哈希值，因此在进行这个调用时不会触发渲染。

`AEGP_CCComputeOptionsRefconP`将包含传入`AEGP_ComputeIfNeededAndCheckout`或 `AEGP_CheckoutCached`方法的数据。

`AEGP_CCComputeKeyP` `out_keyP`返回散列键值，类型定义见 `AE_ComputeCacheSuite.h`中的`AEGP_CCComputeKey`定义。

Note

传入 `generate_key`和`compute`的 `AEGP_CCComputeOptionsRefconP`参数必须包含计算缓存值的哈希密钥/计算缓存值本身的所有输入。这将经常包括许多或所有的效果参数和任何计算缓存值所需的层参数。更多细节见[Real-world Integration Example](../effect-details/compute-cache-api.html)。

```cpp
A_Err(*generate_key)(
AEGP_CCComputeOptionsRefconPoptionsP,
AEGP_CCComputeKeyPout_keyP);
```

### compute

当需要计算一个缓存值时，由 `AEGP_ComputeIfNeededAndCheckout`调用。

`AEGP_CCComputeOptionsRefconP`将包含传入`AEGP_ComputeIfNeededAndCheckout`方法的数据。

设置 `out_valuePP`指向计算缓存值的结果，并投向`AEGP_CCComputeValueRefconP`类型。

比如说。

```cpp
*out_valuePP = reinterpret_cast<AEGP_CCComputeValueRefconP>(myComputedResultP);
```

```cpp
A_Err(*compute)(
AEGP_CCComputeOptionsRefconPoptionsP,
AEGP_CCComputeValueRefconP*out_valuePP);
```

### approx_size_value

由缓存系统调用，以确定计算出的缓存值所使用的内存总面积。计算的值不需要是一个平面结构。

大小是缓存清除启发式的一个输入。

`AEGP_CCComputeValueRefconP`是计算出的缓存值，可以用来生成要返回的大小值。

```cpp
size_t(*approx_size_value)(
AEGP_CCComputeValueRefconPvalueP);
```

### delete_compute_value

当缓存条目需要被清除时，会调用这个值来释放。缓存值拥有的所有资源必须在这里释放。

```cpp
void(*delete_compute_value)(
AEGP_CCComputeValueRefconPvalueP);
```

## 生成一个 key

`generate_key`回调必须在注册类中返回一个唯一的密钥，作为缓存中的条目的缓存密钥，但为了适应未来，我们强烈建议该密钥在所有注册类中是全球唯一的。AE SDK 提供了 "AEGP_HashSuite1 "套件来帮助生成一个 GUID，作为密钥使用。

`generate_key`的结果必须以`AEGP_CCComputeKey`对象的形式提供，其类型由以下结构定义。

```cpp
typedef struct AEGP_GUID {
 A_long bytes[4];
} AEGP_GUID;

```

## AEGP_HashSuite1

`AEGP_HashSuite1`可以用来生成一个唯一的密钥，在 `AEGP_ComputeCacheCallbacks` `generate_key()`回调方法中使用。

在获得套件后，用一个缓冲区调用`AEGP_CreateHashFromPtr()`方法；我们建议用一个可识别的字符串的字符数组，这样你可以很容易地回忆起缓存条目中储存的内容。然后调用`AEGP_HashMixInPtr()`，加入任何效果参数，层层检查哈希结果等，应该会产生一个不同的缓存密钥和条目。

### AEGP_CreateHashFromPtr

调用这个开始创建哈希值，这个哈希值将在 `hashP`中返回，可用于从`generate_key`返回。

```cpp
A_Err(*AEGP_CreateHashFromPtr)(
constA_u_longlongbuf_sizeLu,
constvoid*bufPV,
AEGP_GUID*hashP);
```

### AEGP_HashMixInPtr

为每个效果参数、层检查哈希或其他将用于计算缓存条目的数据调用这个。

```cpp
A_Err(*AEGP_HashMixInPtr)(
constA_u_longlongbuf_sizeLu,
constvoid*bufPV,
AEGP_GUID*hashP);
```

下面是一个使用 `AEGP_HashSuite1`的例子，Levels2Histo_generate_key_cb()是为 `generate_key()`调用的回调。

```cpp
A_Err Levels2Histo_generate_key_cb(AEGP_CCComputeOptionsRefconP opaque_optionsP, AEGP_CCComputeKeyP out_keyP)
{
  try
  {
    const Levels2Histo_options&  histo_op( *reinterpret_cast<Levels2Histo_options*>(opaque_optionsP));
    A_Err err = Err_NONE;

    AEFX_SuiteScoper<AEGP_HashSuite1> hash_suite = AEFX_SuiteScoper<AEGP_HashSuite1>(
        in_dataP,
        kAEGPHashSuite,
        kAEGPHashSuiteVersion1,
        out_dataP);

    // define a simple buffer that is easy to recognize as a starting hash
    const char* hash_buffer = "Level2Histo";
    err = hash_suite->AEGP_CreateHashFromPtr(sizeof(hash_buffer), hash_buffer, out_keyP);

    // Mix in effect parameters that would create a different compute result and should generate a different cache entry and key.
    if (!err) {
      err = hash_suite->AEGP_HashMixInPtr(sizeof(histo_op.depthL), &histo_op.depthL, out_keyP);
    }

    if (!err) {
      err = hash_suite->AEGP_HashMixInPtr(sizeof(histo_op.bB), &histo_op.bB, out_keyP);
    }

    // mix in any other effect parameters that should affect the cache key
    // ...

    // out_keyP is returned as the generated key for use as the cache key.
  }
  catch (...)
  {
      /* return most appropriate PF_Err */
  }
}
```

## 计算或签出缓存值

当添加缓存支持时，首先要回答的问题是一个渲染调用是否需要签出一个以上的缓存值。如果需要一个以上的缓存值来完成渲染，那么可以应用多签出模式，在多个渲染调用中并发计算缓存，从而避免计算的序列化。

### 单个缓存值

如果一个渲染调用在渲染一帧时只需要一个缓存值，那么将`AEGP_ComputeIfNeededAndCheckout`中的`wait_for_other_threadB`参数设置为`true`。签出调用将返回一个凭证，可能会调用计算回调来填充缓存；或者等待另一个已经开始需要计算的线程。

### 多个签出缓存值

如果一个渲染调用需要多个缓存值，那么可以使用多签出模式来保持渲染线程的利用率，从而避免计算的序列化。

使用多检查的概念是让一个渲染线程(例如渲染帧 3)利用任何其他渲染线程(例如帧 1、2)，这些线程正在与该线程同时计算所需的缓存值(例如，帧 3 需要来自帧 1 和 2 的数据)。如果没有其他线程正在计算要求的缓存值，那么渲染线程(帧 3)将执行计算。一旦所有的缓存值检查调用完成，渲染线程(帧 3)就可以等待其他线程(帧 1、2)完成计算，然后再执行像素渲染。一旦像素渲染完成，请确保签入任何被签出的缓存值(帧 1、2 和 3)。

下面是说明这种方法的伪代码示例。

```cpp
Render()
{
    // Make a request for each cache value that is needed to complete the render
    bool first_err = AEGP_ComputeIfNeededAndCheckout(first_options, do_not_wait, first_cache_receipt);
    bool second_err = AEGP_ComputeIfNeededAndCheckout(second_options, do_not_wait, second_cache_receipt);
    // Add as many additional do_not_wait checkout calls here as needed.

    // Once all the requests have been made, check to see if any of the Checkouts did not return
    // a valid checkout receipt.
    if(first_err == A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING) {
        AEGP_ComputeIfNeededAndCheckout(wait, first_cache_receipt);
    }
    if(second_err == A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING) {
        AEGP_ComputeIfNeededAndCheckout(wait, second_cache_receipt);
    }
    // Add as many additional waiting checkout calls here as needed

    // All cache values are now available via AEGP_GetReceiptComputeValue for use in the Render

    // ... complete the render steps

    // Check in all cache values now
    AEGP_CheckinComputeReceipt(first_cache_receipt);
    AEGP_CheckinComputeReceipt(second_cache_receipt);
}
```

## wait_for_other_threadB 对 AEGP_ComputeIfNeededAndCheckout 的影响

对`AEGP_ComputeIfNeededAndCheckout`的调用在几乎所有的参数排列中都会返回缓存值的签出凭证，除非`wait_for_other_threadB`被设置为`false`，而另一个线程已经在渲染请求的缓存值。

| **Cache State**                                                                                          | **wait_for_other_threadB set to False**                             | **wait_for_other_threadB set to True** |
| -------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- | -------------------------------------- |
| _No cache for key_                                                                                       | Compute and checkout receipt returned                               | Compute and checkout receipt returned  |
| _Being computed by another thread_                                                                       | Returns A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING                       |                                        |
| Note that After Effects will not report this error to the user, it is only for the effect to respond to. | Wait for another thread and return checkout receipt upon completion |                                        |
| _Cached_                                                                                                 | Checkout receipt returned                                           | Checkout receipt returned              |

## 检查缓存状态

- 可能会有这样的情况：一个效果需要检查一个缓存值是否已经被计算出来，但又不想在等待另一个线程完成计算时实际执行或阻塞。这可以通过`AEGP_CheckoutCached()`方法实现。
- 这个调用可以用来实现一个轮询模式，在这个模式下，另一段代码被期望填充缓存。例如，一个 UI 线程可以轮询缓存，以获取渲染线程上生成的直方图。
- 如果缓存值可用，`AEGP_CCCheckoutReceiptP`参数将返回一个签出凭证，可以传递给`AEGP_GetReceiptComputeValue()`来检索缓存值。如果缓存值不可用，该方法将返回一个`A_Err_NOT_IN_CACHE_OR_COMPUTE_PENDING`错误代码。

## 缓存的持久性

- 与扁平化的序列数据不同，计算缓存的内容不与项目一起存储，当项目重新打开时，任何计算的内容都需要重新计算。
- 如果 After Effects 的其他操作需要内存，缓存中的条目会自动清除。在编写依赖缓存值的代码时，应假设每次都需要完成计算步骤。
- `approx_size_value`回调应该快速返回，但要提供一个合理准确的缓存数据测量值。这将使 After Effects 能够更好地决定何时清除哪些数据。
- 取消注册缓存类将从缓存中删除该类的所有数据。它将导致对缓存中与该缓存类相关的每个条目进行`delete_compute_value`回调。
- `delete_compute_value`回调应该释放与缓存条目相关的任何资源。计算缓存只包含一个指向资源的 void\*指针，不能代表效果释放资源。

## Real-world 的集成示例

After Effects 中的 Auto Color 插件是一个效果，现在利用计算缓存和`HashSuite1`套件来缓存直方图和水平数据，当效果参数 Temporal Smoothing 被设置为大于 0 的值时使用。

整合缓存和哈希套件的最初步骤是确定哪些数据是由 Auto Color 的时间平滑计算的，哪些部分的计算是耗时的，然后是哪些效果参数会导致需要重新计算。

::: tip

每个效果都需要计算和缓存不同的数据，所以你需要为你的效果做独特的审查。
:::
对于 Auto Color 的 Temporal Smoothing，被渲染的帧需要来自它周围的帧的直方图和电平数据。需要的周围帧的数量是基于时间平滑的参数值。直方图和电平数据的计算都很昂贵，但一般来说，可以为每一帧计算一次，缓存起来，然后根据需要重新使用。

然而，在自动色彩效果中，有许多其他参数被用来计算缓存值，包括黑色剪辑、白色剪辑、中间色调和自动色彩模式。因此，这些参数需要包含在`generate_key`和`compute`方法中。

有了这些信息，我们开始了计算缓存的整合。

1. 定义类的注册 ID，并添加注册和取消注册结帐缓存类和回调的调用

> - 对 AEGP_ClassRegister 的调用在`PF_Cmd_GLOBAL_SETUP`期间执行。
> - 对 AEGP_ClassUnregister 的调用在`PF_Cmd_GLOBAL_SETDOWN`期间执行。

2. 实现`generate_key`, `compute`, `approx_size_value`和`delete_compute_value`的回调函数。

> - `generate_key`利用`AEGP_HashSuite1`在黑色剪辑、白色剪辑、中间色调和自动水平模式下生成一个独特的密钥混合。它还混合了帧的时间和时间步长，以确保缓存对正在计算的特定帧是唯一的。
> - `compute`计算直方图和水平，并将这两个数据结构存储到一个结构中，作为计算回调的`out_valuePP`参数来设置。
> - `approx_size_value`增加了缓存值中的直方图和级别数据结构的`sizeof()`，以返回被缓存条目使用的内存大小。
> - `delete_compute_value`清除缓存条目的直方图和级别数据结构所占用的内存。

3. 将计算/签出调用整合到时间平滑中

> - Temporal Smoothing 代码已经更新，包括对`AEGP_ComputeIfNeededAndCheckout`的调用。对每一帧时间/时间步长的时间平滑算法进行调用，利用其他渲染线程计算周围帧直方图和水平数据的结果。

4. 整合缓存的签出和签入

> - 一旦计算出一帧所需的所有缓存值，效果代码就会使用`AEGP_GetReceiptComputeValue`检查出所需的缓存值。
> - 然后缓存值被用作时间平滑算法的一部分，对帧的颜色进行调整。
> - 一旦当前帧不再需要缓存值，就会为每个缓存值的接收调用`AEGP_CheckinComputeReceipt`。
> - Auto Color 此时不使用`AEGP_CheckoutCached`。

5. 测试 sequence_data 与 Compute Cache 的实现方式

> - Auto Color 使用序列数据来存储直方图和水平数据，在使用计算缓存之前，它将在每个渲染线程上有一个唯一的 sequence_data 副本。这意味着每一帧所需的直方图和色阶都需要在每个线程上进行渲染。
> - 随着使用计算缓存的改变，正在渲染的每一帧都获得了其他渲染线程计算直方图和级别数据并将其存储起来供将来使用的性能优势。
> - 使用计算缓存渲染自动色彩效果的改进，使得渲染速度比 sequence_data 版本至少快 3 倍。
