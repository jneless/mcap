go-sdk-quick-start.md

本文介绍如何通过 TOS Go SDK 来完成常见的操作，如创建桶，上传、下载和删除对象等。
<span id="前提条件"></span>
## 前提条件

1. [安装 SDK](/docs/6349/93476)
2. [初始化客户端](/docs/6349/93477)

<span id="客户端通用示例"></span>
## 客户端通用示例
使用 TosClient 的通用示例如下。
```go
package main

import (
   "context"
   "fmt"
   "strings"

   "github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

func main() {
   var (
      accessKey = os.Getenv("TOS_ACCESS_KEY")
      secretKey = os.Getenv("TOS_SECRET_KEY")
      // Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"

      // 将文件上传到 example_dir 目录下的 example.txt 文件
      objectKey = "example_dir/example.txt"
      ctx       = context.Background()
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   if err != nil {
      panic(err)
   }
   // 将字符串 “Hello TOS” 上传到指定 example_dir 目录下的 example.txt
   body := strings.NewReader("Hello TOS")
   output, err := client.PutObjectV2(ctx, &tos.PutObjectV2Input{
      PutObjectBasicInput: tos.PutObjectBasicInput{
         Bucket: bucketName,
         Key:    objectKey,
      },
      Content: body,
   })
   // 判断错误
if err != nil {
   // 判断服务端错误
   if serverErr, ok := err.(*tos.TosServerError); ok {
      fmt.Println("Error:", serverErr.Error())
      fmt.Println("Request ID:", serverErr.RequestID)
      fmt.Println("Response Status Code:", serverErr.StatusCode)
      fmt.Println("Response Header:", serverErr.Header)
      fmt.Println("Response Err Code:", serverErr.Code)
      fmt.Println("Response Err Msg:", serverErr.Message)
   } else if clientErr, ok := err.(*tos.TosClientError); ok {
      // 判断客户端错误 
      fmt.Println("Error:", clientErr.Error())
      fmt.Println("Client Cause Err:", clientErr.Cause.Error())
   } else {
      fmt.Println("Error:", err)
   }
   panic(err)
}
   fmt.Println("PutObjectV2 Request ID:", output.RequestID)
}
```

<span id="创建桶"></span>
## 创建桶
存储桶（Bucket）是 TOS 的全局唯一的命名空间，相当于数据的容器，用来储存对象数据。如下代码展示如何使用 CreateBucket 方法创建一个新存储桶。
```go
package main

import (
   "context"
   "fmt"
   "os"

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
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   // 创建存储桶
   resp, err := client.CreateBucketV2(context.Background(), &tos.CreateBucketV2Input{
      Bucket: bucketName,
   })
   checkErr(err)
   fmt.Println("create bucket request id: ", resp.RequestID)
   fmt.Println("create bucket response status code: ", resp.StatusCode)
}
```

:::tip
关于创建桶的更多信息，请参见[创建桶](/docs/6349/93454)。
:::
<span id="上传对象"></span>
## 上传对象
新建存储桶成功后，可以上传对象对象到存储桶中。以下代码将字符串上传到 TOS 存储桶。
```go
package main

import (
   "context"
   "fmt"
   "os"
   "strings"

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
      // Bucket 对于的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"
      // 填写对象名
      objectKey = "*** Provide your object key ***"
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)

   // 上传对象 Body ， 以 string 对象为例
   body := strings.NewReader("object content")
   // 上传对象
   output, err := client.PutObjectV2(context.Background(), &tos.PutObjectV2Input{
      PutObjectBasicInput: tos.PutObjectBasicInput{
         Bucket: bucketName,
         Key:    objectKey,
      },
      Content: body,
   })
   checkErr(err)
   fmt.Println("Put Object Request ID: ", output.RequestID)
   fmt.Println("Put Object Response Status Code: ", output.StatusCode)
}
```

:::tip
关于上传对象的更多示例链接，请参见[上传对象](/docs/6349/132396)。
:::
<span id="下载对象"></span>
## 下载对象
以下展示从存储桶中下载一个对象到内存中。
```go
package main

import (
   "context"
   "fmt"
   "io/ioutil"
   "os"

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
      // Bucket 对于的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"
      // 填写对象名
      objectKey = "*** Provide your object key ***"
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)

   // 下载对象
   output, err := client.GetObjectV2(context.Background(), &tos.GetObjectV2Input{
      Bucket: bucketName,
      Key:    objectKey,
   })
   checkErr(err)
   defer output.Content.Close()
   fmt.Println("Get object request id: ", output.RequestID)
   fmt.Println("Get object response status code: ", output.StatusCode)
   body, err := ioutil.ReadAll(output.Content)
   checkErr(err)
   fmt.Println("Get Object body: ", body)
}
```

:::tip
关于下载对象的更多示例链接，请参见[下载对象](/docs/6349/132398)。
:::
<span id="列举对象"></span>
## 列举对象
以下展示从存储桶中列举已经上传的对象。
```go
package main

import (
   "context"
   "fmt"
   "os"

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
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   // 列举对象
   output, err := client.ListObjectsV2(context.Background(), &tos.ListObjectsV2Input{
      Bucket: bucketName,
   })
   checkErr(err)
   fmt.Println("List Object Request Id:", output.RequestID)
   // 通过判断 IsTruncated 判断是否需要继续列举对象，若需要继续列举对象下次列举时需要传入返回的 NextMarker
   if !output.IsTruncated {
      output, err = client.ListObjectsV2(context.Background(), &tos.ListObjectsV2Input{
         Bucket:           bucketName,
         ListObjectsInput: tos.ListObjectsInput{Marker: output.NextMarker},
      })
      checkErr(err)
      fmt.Println("List Object Request Id:", output.RequestID)
   }
}
```

:::tip
关于列举对象的更多示例链接，请参见[列举对象](/docs/6349/93468)。
:::
<span id="删除对象"></span>
## 删除对象
以下代码展示从存储桶中删除一个对象。
```go
package main

import (
   "context"
   "fmt"
   "os"

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
      // Bucket 对于的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"
      // 填写 BucketName
      bucketName = "*** Provide your bucket name ***"
      // 填写对象名
      objectKey = "*** Provide your object key ***"
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)

   // 删除对象
   output, err := client.DeleteObjectV2(context.Background(), &tos.DeleteObjectV2Input{
      Bucket: bucketName,
      Key:    objectKey,
   })
   checkErr(err)
   fmt.Println("Delete object request id: ", output.RequestID)
   fmt.Println("Delete object response status code: ", output.StatusCode)
}
```

:::tip
关于删除对象的更多示例链接，请参见[删除对象](/docs/6349/93465)。
:::
<span id="关闭-client"></span>
## 关闭 Client
TOS Go SDK 默认开启长连接，您创建 Client 实例访问 TOS 后，可使用以下代码关闭不再使用的 Client 实例，释放连接资源。
```go
package main

import (
   "context"
   "fmt"
   "strings"

   "github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

func main() {
   var (
      accessKey = os.Getenv("TOS_ACCESS_KEY")
      secretKey = os.Getenv("TOS_SECRET_KEY")
      // Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
      endpoint = "https://tos-cn-beijing.volces.com"
      region   = "cn-beijing"

   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   if err != nil {
      panic(err)
   }
   // 关闭客户端
   client.Close()
}
```


