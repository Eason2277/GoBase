# Go 学习提交记录

这份记录用于把当前工作区的学习内容拆成可以独立提交的 commit。

## 提交原则

- 一次提交只对应一个主要学习主题。
- 提交顺序按照知识依赖关系，而不是按照文件名排序。
- 每个示例文件都保持独立运行，使用 `go run 文件名.go` 验证。
- 当前只记录提交计划，不在本次生成记录时执行 `git commit` 或 `git push`。

## 推荐提交顺序

### 1. `docs: add learning commit plan`

文件：`COMMIT_PLAN.md`

内容：记录 Go 学习示例的提交顺序、验证命令和当前暂存状态注意事项。

### 2. `feat: add if statement example`

文件：`9.if_01.go`

学习内容：条件判断、`if / else if / else`、逻辑运算符，以及库存状态的业务判断。

验证：

```bash
go run 9.if_01.go
```

### 3. `feat: add switch statement example`

文件：`switch01.go`

学习内容：根据订单状态码选择分支，理解 `switch` 与多层 `if` 的适用场景。

验证：

```bash
go run switch01.go
```

### 4. `feat: add for loop example`

文件：`for01.go`

学习内容：使用 `for` 模拟订单循环扣减库存，理解循环条件、循环变量和状态变化。

验证：

```bash
go run for01.go
```

### 5. `feat: add product struct example`

文件：`struct01.go`

学习内容：使用结构体描述商品，理解字段、嵌套结构体和业务对象建模。

验证：

```bash
go run struct01.go
```

### 6. `feat: add map lookup example`

文件：`map01.go`

学习内容：使用 `map` 保存商品或用户数据，理解键值查找、更新、删除，以及 comma-ok 判断键是否存在。

验证：

```bash
go run map01.go
```

### 7. `feat: add function and return value example`

文件：`function01.go`

学习内容：使用函数封装订单结算逻辑，理解参数、返回值，以及 `(value, ok)` 形式的结果判断。

验证：

```bash
go run function01.go
```

### 8. `feat: add interface polymorphism example`

文件：`interface.go`

学习内容：定义 `Speaker` 接口，让不同类型通过同一个函数完成行为调用，理解 Go 接口的隐式实现。

验证：

```bash
go run interface.go
```

## 独立资料提交

以下内容不属于 Go 示例主线，建议单独提交，避免学习代码和资料变更混在一起。

### 9. `docs: update learning guidelines`

文件：`AGENTS.md`

内容：补充数据结构并行学习、AI 时代学习取舍等协作与教学规则。

### 10. `docs: add algorithms reference`

文件：`《算法》[第4版].pdf`

内容：添加算法学习参考资料。

## 当前状态注意事项

当前 `git status --short` 中，7 个 Go 文件显示为 `AM`。这表示这些文件已有暂存版本，同时工作区还有未暂存修改。后续正式提交某个文件前，应先确认要提交的是当前文件的完整版本，再执行：

```bash
git add 文件名
git diff --cached -- 文件名
git commit -m "提交信息"
```

提交后使用下面的命令确认结果：

```bash
git status --short
git log --oneline -1
```
