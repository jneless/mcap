getobject-range.md

如果您只需获取对象中的部分数据，您可以使用范围下载，下载指定范围内的数据，本文介绍范围下载。
<span id="注意事项"></span>
## **注意事项**

* 下载对象前，您必须具有 `tos:GetObject` 权限，具体操作，请参见[权限配置指南](/docs/6349/102120)。
* 对于开启多版本的桶，下载指定版本对象时，您必须具有 `tos:GetObjectVersion` 权限，具体操作，请参见[权限配置指南](/docs/6349/102120)。
* 如果应用程序会在同一时刻大量下载同一个对象，您的访问速度会受到 TOS 带宽及地域的限制。建议您使用 CDN 产品，提升性能的同时也能降低您的成本。通过 CDN 访问 TOS 的详细信息，请参见[使用 CDN 加速访问 TOS 资源](/docs/6349/106079)。

<span id="15538313"></span>
## 下载范围说明
指定对象范围时，起始值最小为 0，最大为对象长度减 1，例如大小为 100 字节的对象，则指定的正常下载范围为 0~99。
<span id="示例代码"></span>
## 示例代码
<span id="指定-start-end-下载对象"></span>
### **指定 Start/End 下载对象**
以下代码用于指定 Start/End 下载桶中对象的部分数据。
```go
package main

import (
   "context"
   "fmt"
   "io/ioutil"

   "github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

func checkErr(err error) {
   if err != nil {
      if serverErr, ok := err.(*tos.TosServerError); ok {
         fmt.Println("Error:", serverErr.Error())
         fmt.Println("Request ID:", serverErr.RequestID)
         fmt.Println("Response Status Code:", serverErr.StatusCode)
         fmt.Println("Response Header:", serverErr.Header)
         fmt.Println("Response Err Code:", serverErr.Code)
         fmt.Println("Response Err Msg:", serverErr.Message)
      } else if clientErr, ok := err.(*tos.TosClientError); ok {
         fmt.Println("Error:", clientErr.Error())
         fmt.Println("Client Cause Err:", clientErr.Cause.Error())
      } else {
         fmt.Println("Error:", err)
      }
      panic(err)
   }
}

func main() {
   var (
      accessKey = os.Getenv("TOS_ACCESS_KEY")
      secretKey = os.Getenv("TOS_SECRET_KEY")
      // Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"

      // 下载指定的对象
      objectKey = "example_dir/example.txt"
      ctx       = context.Background()
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   // 获取 32-64 字节，包含 32 和 64 共 33 字节
   getOutput, err := client.GetObjectV2(ctx, &tos.GetObjectV2Input{
      Bucket:     bucketName,
      Key:        objectKey,
      RangeStart: 32,
      RangeEnd:   64,
   })
   checkErr(err)
   fmt.Println("GetObjectV2 Request ID:", getOutput.RequestID)
   // 下载时前设置的 response content type
   fmt.Println("GetObjectV2 Response ContentType:", getOutput.ContentType)
   // 下载时前设置的 response content encoding
   fmt.Println("GetObjectV2 Response ContentEncoding:", getOutput.ContentEncoding)
   // 下载数据大小
   fmt.Println("GetObjectV2 Response ContentLength", getOutput.ContentLength)
   defer getOutput.Content.Close()
   body, err := ioutil.ReadAll(getOutput.Content)
   checkErr(err)
   // 完成下载
   fmt.Println("Get Object Content:", body)
}
```

<span id="配置进度条"></span>
### **配置进度条**
下载时可通过实现 tos.DataTransferStatusChange 接口接收下载进度，代码示例如下。
```go
package main

import (
   "context"
   "fmt"
   "io/ioutil"

   "github.com/volcengine/ve-tos-golang-sdk/v2/tos"
   "github.com/volcengine/ve-tos-golang-sdk/v2/tos/enum"
)

// 自定义进度回调，需要实现 tos.DataTransferStatusChange 接口
type listener struct {
}

func (l *listener) DataTransferStatusChange(event *tos.DataTransferStatus) {
   switch event.Type {
   case enum.DataTransferStarted:
      fmt.Println("Data transfer started")
   case enum.DataTransferRW:
      // Chunk 模式下 TotalBytes 值为 -1
      if event.TotalBytes != -1 {
         fmt.Printf("Once Read:%d,ConsumerBytes/TotalBytes: %d/%d,%d%%\n", event.RWOnceBytes, event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
      } else {
         fmt.Printf("Once Read:%d,ConsumerBytes:%d\n", event.RWOnceBytes, event.ConsumedBytes)
      }
   case enum.DataTransferSucceed:
      fmt.Printf("Data Transfer Succeed, ConsumerBytes/TotalBytes: %d/%d,%d%%\n", event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
   case enum.DataTransferFailed:
      fmt.Printf("Data Transfer Failed\n")
   }
}

func checkErr(err error) {
   if err != nil {
      if serverErr, ok := err.(*tos.TosServerError); ok {
         fmt.Println("Error:", serverErr.Error())
         fmt.Println("Request ID:", serverErr.RequestID)
         fmt.Println("Response Status Code:", serverErr.StatusCode)
         fmt.Println("Response Header:", serverErr.Header)
         fmt.Println("Response Err Code:", serverErr.Code)
         fmt.Println("Response Err Msg:", serverErr.Message)
      } else if clientErr, ok := err.(*tos.TosClientError); ok {
         fmt.Println("Error:", clientErr.Error())
         fmt.Println("Client Cause Err:", clientErr.Cause.Error())
      } else {
         fmt.Println("Error:", err)
      }
      panic(err)
   }
}

func main() {
   var (
      accessKey = os.Getenv("TOS_ACCESS_KEY")
      secretKey = os.Getenv("TOS_SECRET_KEY")
      // Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"

      // 下载指定的对象
      objectKey = "example_dir/example.txt"
      ctx       = context.Background()
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   // 获取 32-64 字节，包含 32 和 64 共 33 字节
   getOutput, err := client.GetObjectV2(ctx, &tos.GetObjectV2Input{
      Bucket:     bucketName,
      Key:        objectKey,
      RangeStart: 32,
      RangeEnd:   64,
      // 获取当前下载进度
      DataTransferListener: &listener{},
   })
   checkErr(err)
   fmt.Println("GetObjectV2 Request ID:", getOutput.RequestID)
   // 下载时前设置的 response content type
   fmt.Println("GetObjectV2 Response ContentType:", getOutput.ContentType)
   // 下载时前设置的 response content encoding
   fmt.Println("GetObjectV2 Response ContentEncoding:", getOutput.ContentEncoding)
   // 下载数据大小
   fmt.Println("GetObjectV2 Response ContentLength", getOutput.ContentLength)
   defer getOutput.Content.Close()
   body, err := ioutil.ReadAll(getOutput.Content)
   checkErr(err)
   // 完成下载
   fmt.Println("Get Object Content:", body)
}
```

<span id="配置客户端限速"></span>
### **配置客户端限速**
下载对象时可以通过客户端使用 tos.RateLimiter 接口对下载数据所占用的带宽进行限制，代码如下所示。
```go
package main

import (
   "context"
   "fmt"
   "io/ioutil"
   "sync"
   "time"

   "github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

type rateLimit struct {
   rate     int64
   capacity int64

   currentAmount int64
   sync.Mutex
   lastConsumeTime time.Time
}

func NewDefaultRateLimit(rate int64, capacity int64) tos.RateLimiter {
   return &rateLimit{
      rate:            rate,
      capacity:        capacity,
      lastConsumeTime: time.Now(),
      currentAmount:   capacity,
      Mutex:           sync.Mutex{},
   }
}

func (d *rateLimit) Acquire(want int64) (ok bool, timeToWait time.Duration) {
   d.Lock()
   defer d.Unlock()
   if want > d.capacity {
      want = d.capacity
   }
   increment := int64(time.Now().Sub(d.lastConsumeTime).Seconds() * float64(d.rate))
   if increment+d.currentAmount > d.capacity {
      d.currentAmount = d.capacity
   } else {
      d.currentAmount += increment
   }
   if want > d.currentAmount {
      timeToWaitSec := float64(want-d.currentAmount) / float64(d.rate)
      return false, time.Duration(timeToWaitSec * float64(time.Second))
   }
   d.lastConsumeTime = time.Now()
   d.currentAmount -= want
   return true, 0
}

func checkErr(err error) {
   if err != nil {
      if serverErr, ok := err.(*tos.TosServerError); ok {
         fmt.Println("Error:", serverErr.Error())
         fmt.Println("Request ID:", serverErr.RequestID)
         fmt.Println("Response Status Code:", serverErr.StatusCode)
         fmt.Println("Response Header:", serverErr.Header)
         fmt.Println("Response Err Code:", serverErr.Code)
         fmt.Println("Response Err Msg:", serverErr.Message)
      } else if clientErr, ok := err.(*tos.TosClientError); ok {
         fmt.Println("Error:", clientErr.Error())
         fmt.Println("Client Cause Err:", clientErr.Cause.Error())
      } else {
         fmt.Println("Error:", err)
      }
      panic(err)
   }
}

func main() {
   var (
      accessKey = os.Getenv("TOS_ACCESS_KEY")
      secretKey = os.Getenv("TOS_SECRET_KEY")
      // Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"

      // 下载指定的对象
      objectKey = "example_dir/example.txt"
      ctx       = context.Background()
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   rateLimit1m := int64(1024 * 1024)
   // 获取 32-64 字节，包含 32 和 64 共 33 字节
   getOutput, err := client.GetObjectV2(ctx, &tos.GetObjectV2Input{
      Bucket:     bucketName,
      Key:        objectKey,
      RangeStart: 32,
      RangeEnd:   64,
      // 配置客户端限制
      RateLimiter: NewDefaultRateLimit(rateLimit1m, rateLimit1m),
   })
   checkErr(err)
   fmt.Println("GetObjectV2 Request ID:", getOutput.RequestID)
   // 下载时前设置的 response content type
   fmt.Println("GetObjectV2 Response ContentType:", getOutput.ContentType)
   // 下载时前设置的 response content encoding
   fmt.Println("GetObjectV2 Response ContentEncoding:", getOutput.ContentEncoding)
   // 下载数据大小
   fmt.Println("GetObjectV2 Response ContentLength", getOutput.ContentLength)
   defer getOutput.Content.Close()
   body, err := ioutil.ReadAll(getOutput.Content)
   checkErr(err)
   // 完成下载
   fmt.Println("Get Object Content:", body)
}
```

<span id="相关文档"></span>
## **相关文档**
关于下载对象的 API 文档，请参见 [GetObject](/docs/6349/74856)。
