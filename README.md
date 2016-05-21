#snowflake算法的实现

###说明
这里实现的思路非常简单,就是使用snowflake算法来生成递增的唯一id,使用snowflake生成id相比guid的好处，这里就不说了。

###依赖
下载地址:[https://godoc.org/launchpad.net/gozk/zookeeper](https://godoc.org/launchpad.net/gozk/zookeeper)

###原理与思路
我们都知道snowflake算法,生成递增id由三部分组成.
*	第一部分是由时间构成(42bit).
*	第二部分是由初始id构成(10bit).
*	第三部分是由递增序列构成(12bit).

其中这里面比较关键的是第二部分中的初始id生成方法.

这里的实现,提供了三种方法:

1.	递增id.
2.	文件id.
3.	zookeeper id.

>当然还可以使用mysql auto_increment特点
>redis的incr指令方式实现

###使用
计划使用[grpc](http://doc.oschina.net/grpc)实现一个分布式的id生成器,可以在需要使用的客户端机器上进行部署.

###总结
具体可以通过源代码查看
