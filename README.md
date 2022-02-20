# go-chevereto 命令行上传图片到chevereto图床

## 特性

- 支持Windows、Linux
- 自动复制上传后的图片完整url
- 单文件，无多余依赖

## 使用帮助

上传后会自动将上传后的图片完整url复制到剪贴板，便于写作时粘贴使用

使用时需要将 `conf.example.ini` 文件重命名为 `conf.ini` 并修改url值和key值

```shell
λ go-chevereto.exe -h
Usage of go-chevereto.exe:
  -f string
        待上传的图片的路径
```