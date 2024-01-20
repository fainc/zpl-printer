## Windows ZPL printer sdk for go
## 热敏标签小票打印机 Windows ZPL 指令打印 SDK

- 热敏标签打印机大部分支持 ZPL 等指令，自测霍尼韦尔OD480d通过。
- 使用指令打印免驱。
- 目前仅实现必要打印功能。
- 通过 windows call dll 实现。
- 不是所有厂商都使用相同的方案和 sdk，如果当前 dll sdk 不适用，请自行寻找官方 dll sdk 和 指令文档，并且需要更改源代码的 Proc 实现。
- 目前（2024-1-20）已知支持打印机：霍尼韦尔OD480d（同品牌或系列大概率可行），如果帮助到你，可以反馈打印机支持列表。
```
 func main() {
    var err error
    sdk, err = zpl.NewSDK("printer.sdk.dll") // 自定义 dll sdk 文件
    if err != nil {
       fmt.Println(err.Error())
       return
    }
    printer, err = zpl.NewPrinter(sdk, "", "USB,") // 默认值就可以打开USB端口并连接
    if err != nil {
       fmt.Println(err)
       return
    }
    err = printer.PrintConfigurationLabel() // 测试直接打印配置标签，正常应该直接打印出来
    if err != nil {
       fmt.Println(err)
       return
    }
    err = printer.Release() // 不需要时释放打印机对象
    if err != nil {
       fmt.Println(err)
       return
    }
    err = sdk.Release() // 不需要时释放 windows dll 资源
    if err != nil {
       fmt.Println(err)
       return
    }
    return
}
```