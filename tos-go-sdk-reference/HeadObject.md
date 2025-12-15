HeadObject.md

您可以通过 Go SDK 的 `HeadObject` 接口判断对象是否存在。
<span id="391f9225"></span>
## 注意事项
要判断对象是否存在，您的账号必须具备 `tos:GetObject` 权限，具体操作请参见[权限配置指南](/docs/6349/102120)。
<span id="e6928431"></span>
## 示例代码
以下代码展示如何判断对象是否存在。

```Go
package main

import (
   "context"
   "fmt"
   "net/http"

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

      // 存储桶中的对象名
      objectKey = "example_dir/example.txt"
      ctx       = context.Background()
   )
   // 初始化客户端
   client, err := tos.NewClientV2(endpoint, tos.WithRegion(region), tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))
   checkErr(err)
   output, err := client.HeadObjectV2(ctx, &tos.HeadObjectV2Input{
      Bucket: bucketName,
      Key:    objectKey,
   })
   if err != nil {
      if serverErr, ok := err.(*tos.TosServerError); ok {
         // 判断对象是否存在
         if serverErr.StatusCode == http.StatusNotFound {
            fmt.Println("Object not found.")
         } else {
            fmt.Println("Error:", serverErr.Error())
            fmt.Println("Request ID:", serverErr.RequestID)
            fmt.Println("Response Status Code:", serverErr.StatusCode)
            fmt.Println("Response Header:", serverErr.Header)
            fmt.Println("Response Err Code:", serverErr.Code)
            fmt.Println("Response Err Msg:", serverErr.Message)
            panic(err)
         }
      } else {
         panic(err)
      }
   }
   fmt.Println("HeadObjectV2 Request ID:", output.RequestID)
   // 查看内容语言格式
   fmt.Println("HeadObjectV2 Response ContentLanguage:", output.ContentLanguage)
   // 查看下载时的名称
   fmt.Println("HeadObjectV2 Response ContentDisposition:", output.ContentDisposition)
   // 查看编码类型
   fmt.Println("HeadObjectV2 Response ContentEncoding:", output.ContentEncoding)
   // 查看缓存策略
   fmt.Println("HeadObjectV2 Response CacheControl:", output.CacheControl)
   // 查看对象类型
   fmt.Println("HeadObjectV2 Response ContentType:", output.ContentType)
   // 查看缓存过期时间
   fmt.Println("HeadObjectV2 Response Expires:", output.Expires)
}
```

<span id="b125d336"></span>
## **相关文档**
关于判断对象是否存在的 API 接口，请参见 [HeadObject](/docs/6349/74864)。
