# micmall
简介：从0到1自建微服务架构，采用Web和Service两层架构设计，基本分为用户服务、用户操作服务、商品服务、库存服务、购物车和订单服务。

技术栈：Go语言， Gorm， Gin，gRPC， Redis， MySQL， Consul， Nacos， Jaeger， Sentinel， Docker，Viper，Zap

核心功能实现

用户服务：通过JWT Token技术实现用户认证
库存服务：使用Redis实现分布式锁来防止超卖
商品服务：使用Redis对商品信息进行缓存，减缓数据库压力，并且对缓存设置超时时间和删除缓存时采用双发策略尽可能保证数据库和缓存的数据一致性
订单服务：需要跨服务调用商品详情服务和库存扣减的服务，并且执行本地的事务，采用了RocketMQ提供的事务消息来解决分布式事务问题
服务治理：使用Consul充当注册中心，使用Nacos充当配置中心，使用Jaeger对调用链进行链路追踪，Sentinal实现服务熔断和限流
