# 错误描述

使用php的pdo操作tidb时，在事务提交时，虽然commit()返回成功，
但是实际事务是提交失败的，因为数据库找不到刚插入的数据。

# 注意事项

需要把脚本放到nginx下，进行并发。

测试使用ab即可重现，每次必现。
