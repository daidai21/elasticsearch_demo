# 基本样例

设置下表的Type配置
    存几份、刷新时间、压缩 之类的配置

primaries 主分片数量
replicas 备份数

设置Mapping
    存储表结构定义

sdk
    一般就是增、删、改、查

TODO:
    记录下流程
    看下配置要注意的
    sdk run下

mvn archetype:generate -DgroupId=com.example.demo -DartifactId=base -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
