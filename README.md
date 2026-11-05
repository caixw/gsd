# gsd

[![Go](https://github.com/caixw/gostructdoc/actions/workflows/go.yml/badge.svg)](https://github.com/caixw/gostructdoc/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/caixw/gostructdoc/branch/master/graph/badge.svg)](https://codecov.io/gh/caixw/gostructdoc)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/caixw/gostructdoc)](https://pkg.go.dev/github.com/caixw/gostructdoc)
![Go version](https://img.shields.io/github/go-mod/go-version/caixw/gostructdoc)
![License](https://img.shields.io/github/license/caixw/gostructdoc)

将 Go 的结构体字段提取出来形成 markdown 文档

安装
---

```bash
go install github.com/caixw/gostructdoc
```

如果是在 `go:generate` 中使用，也可以直接使用以下命令写入 go.mod 中：

```bash
go get -tool github.com/caixw/gostructdoc
```

版权
---

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
