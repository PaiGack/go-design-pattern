命令模式
===

# 概念

将请求转化为一个保护与请求相关的所有信息的独立对象。该转化能根据不同的请求，将方法参数化、延迟请求活将其放入队列中，且能实现可撤销操作。

方法参数化是指将每个请求参数传入具体命令的工厂方法创建命令，同时具体的命令会会默认设置好接受对象。

优点
- 不管请求参数个数及类型，还是调用对象由几个，都会被封装到具体命令对象的成员字段上
- 统一的 Execute 接口方法进行调用，屏蔽各个请求的差异，便于命令拓展、多命令组装、回滚等


# 示例
[控制电饭煲做饭](electriccooker.go)

电饭煲的控制面板会提供设置煮粥、蒸饭模式，及开始和停止按钮，电饭煲控制系统会根据模式的不同设置相应的火力，压强及时间等参数。
煮粥，蒸饭就相当于不同的命令，开始按钮就相当命令触发器，设置好做饭模式，点击开始按钮电饭煲就开始运行，同时还支持停止命令。


