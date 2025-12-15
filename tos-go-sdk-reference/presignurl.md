presignurl.md

预签名 URL 可在不公开对象权限的情况下，临时授权他人访问私有资源。在有效期内，任何用户均可通过预签名链接下载文件。您可以通过`presign` 命令生成下载对象的预签名 URL，安全便捷地分享私有对象。
<span id="6b29a390"></span>
## 命令语法

```mixin-react
return (<Tabs>
<Tabs.TabPane title="Linux/macOS" key="XX3NwbAkH1"><RenderMd content={`* 生成单个对象的预签名 URL
   \`\`\`Bash
   ./tosutil presign tos://bucket/key [-vp=xxx] [-qp=xxx] [-versionId=xxx] [-e=xxx] [-re=xxx] [-i=xxx] [-k=xxx] [-t=xxx] [-conf=xxx]
   \`\`\`

* 批量生成指定前缀对象的预签名 URL
   \`\`\`Bash
   ./tosutil presign tos://bucket/[prefix] -r [-vp=xxx] [-qp=xxx] [-include=*.xxx] [-exclude=*.xxx] [-d] [-v] [-marker=xxx] [-versionIdMarker=xxx] [-limit=1] [-bt=xxx] [-e=xxx] [-re=xxx] [-i=xxx] [-k=xxx] [-t=xxx] [-conf=xxx]
   \`\`\`

`}></RenderMd></Tabs.TabPane>
<Tabs.TabPane title="Windows" key="eQeX7mQh9X"><RenderMd content={`* 生成单个对象的预签名 URL
   \`\`\`Bash
   tosutil presign tos://bucket/key [-vp=xxx] [-qp=xxx] [-versionId=xxx] [-e=xxx] [-re=xxx] [-i=xxx] [-k=xxx] [-t=xxx] [-conf=xxx]
   \`\`\`

* 批量生成指定前缀对象的预签名 URL
   \`\`\`Bash
   tosutil presign tos://bucket/[prefix] -r [-vp=xxx] [-qp=xxx] [-include=*.xxx] [-exclude=*.xxx] [-d] [-v] [-marker=xxx] [-versionIdMarker=xxx] [-limit=1] [-bt=xxx] [-e=xxx] [-re=xxx] [-i=xxx] [-k=xxx] [-t=xxx] [-conf=xxx]
   \`\`\`

`}></RenderMd></Tabs.TabPane></Tabs>);
 ```

<span id="5a3e4a98"></span>
## 参数说明
通用可选参数的说明，请参见[通用可选参数](/docs/6349/152744#通用可选参数)。

| | | | \
|**参数名称** |**参数类型** |**描述** |
|---|---|---|
| | | | \
|-vp |String |预签名 URL 的有效期，配置后，预签名 URL 将在指定时长后自动失效。有效期取值范围为 1 分钟 ~ 30 天。如果您配置 tosutil 工具时使用的是临时访问密钥，则授权码有效期可选最大值与临时访问密钥的过期时间保持一致。 |\
| | |默认值为 1 天，表示生成的授权码将在 1 天后失效，该参数支持以下配置方式： |\
| | | |\
| | |* w：表示设置单位为周。 |\
| | |* d：表示设置单位为天。 |\
| | |* h：表示设置单位为小时。 |\
| | |* min：表示设置单位为分钟。 |\
| | |* s：表示设置单位为秒。 |\
| | | |\
| | |例如 `1d` 表示有效期为 1 天，`1w` 表示有效期为 1 周，`1h` 表示有效期为 1 小时。 |\
| | |:::tip |\
| | |如果您没有选择时间单位，则默认单位为秒，例如 `2400` 表示有效期为 2400秒。 |\
| | |::: |
| | | | \
|-qp |String |预签名 URL 中包含的 Query 参数。 |
| | | | \
| -versionId |String |待生成预签名 URL 的单个对象的版本号，默认为最新版本对象。 |
| | | | \
|  -r |String |按指定的对象前缀批量生成预签名 URL 。 |
| | | | \
| -include |String |批量生成预签名 URL 时，使用名称匹配模式指定授权访问的对象，支持以下字符： |\
| | | |\
| | |* `?` ：匹配单个任意字符。 |\
| | |* `*` ：匹配多个任意字符。 |\
| | |* `#` ：作为分隔符。 |\
| | | |\
| | |例如`-include=*.png#*.txt` 表示匹配所有以 `.png` 和 `.txt `结尾的文件。 |\
| | |:::tip |\
| | |为了避免因操作系统转义特殊符号的导致的解析失败等问题，建议您使用引号设置名称匹配模式。 |\
| | |::: |
| | | | \
| -exclude |String |批量生成预签名 URL 时，使用名称匹配模式排除指定对象，支持以下字符： |\
| | | |\
| | |* `?` ：匹配单个任意字符。 |\
| | |* `*` ：匹配多个任意字符。 |\
| | |* `#` ：作为分隔符。 |\
| | | |\
| | |例如 `-exclude=*.png#*.txt` 表示匹配所有以 `.png` 和 `.txt `结尾的文件。 |\
| | |:::tip |\
| | |为了避免因操作系统转义特殊符号导致的解析失败等问题，建议您使用引号设置名称匹配模式。 |\
| | |::: |
| | | | \
|-d |String |仅为当前目录下的对象生成预签名 URL，而非递归生成所有对象的预签名 URL。 |
| | | | \
| -v |String |为桶内多版本对象生成预签名 URL，生成结果包含最新版本的对象和历史版本的对象。 |
| | | | \
| -marker |String |批量生成预签名 URL 时，对象的起始位置。系统将按对象名称的字典序排序，并为该参数以后的所有对象逐一生成预签名 URL。 |
| | | | \
|-versionIdMarker |String |批量生成预签名 URL 时，多版本对象的起始位置，必须与 `-marker` 参数配合使用。 |\
| | |系统将按照对象名称和版本号的字典序排序，并为该参数以后的所有对象逐一生成预签名 URL。 |
| | | | \
| -limit |String |批量生成预签名 URL 时的最大返回数量，默认值为 1000。 |
| | | | \
|-bt |String | 存储桶的类型。取值说明如下： |\
| | | |\
| | |* `fns`：扁平桶。 |\
| | |* `hns`：分层桶。 |\
| | | |\
| | |如果未指定存储桶的类型，则默认从桶元数据获取存储桶的类型。 |

<span id="6f940e2c"></span>
## 使用示例
**为 bucketname 桶内 object.png 对象，生成有效期为 2 小时的预签名 URL**

* 命令
   ```Bash
   ./tosutil presign tos://bucketname/object.png -vp=2h
   ```

* 返回
   ```Bash
   https://bucketname.tos-cn-beijing.volces.com/object.png?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxxxx%2F20250916%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20250916T071704Z&X-Tos-Expires=7200&X-Tos-Signature=xxx&X-Tos-SignedHeaders=host
   ```


**为 bucketname 桶内前缀为 prefix 的对象，生成有效期为 2 小时的预签名 URL**

* 命令
   ```Bash
   ./tosutil presign tos://bucketname/prefix -r -vp=2h
   ```

* 返回
   ```Bash
   tos://bucketname/prefix-new.png
   https://bucketname.tos-cn-beijing.volces.com/prefix-new.png?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxxxxx%2F20250916%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20250916T071945Z&X-Tos-Expires=7200&X-Tos-Signature=3dc7be30794ee2ad55xxxxx0a26da96b4567b9da32011****&X-Tos-SignedHeaders=host

   tos://bucketname/prefix-new.pptx
   https://bucketname.tos-cn-beijing.volces.com/prefix-new.pptx?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxxxx%2F20250916%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20250916T071945Z&X-Tos-Expires=7200&X-Tos-Signature=xxxxxxx&X-Tos-SignedHeaders=host
   ```
