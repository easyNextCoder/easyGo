## raft协议
[Raft是consoul和etcd的核心算法](https://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/Raft.html)
[动画演示](https://thesecretlivesofdata.com/raft/)

1. 一个raft集群包含多个服务器节点，每个节点可能有三种状态：
   * follower(跟随者): 所有节点都以follower的状态开始。如果没有收到leader消息则会变成candidate状态
   * candidate(候选人): 会向其他节点拉选票，如果得到大部分的票则变成leader，这个过程就叫做Leader选举(Leader Election)
   * leader(领导者): 所有对系统的修改都会先经过leader
2. 领导选举
   * Raft使用一种心跳机制来触发领导人选举
   * 当服务器程序启动时，节点都是follower跟随者的身份
   * 如果一个跟随者在一段时间里没有接收到任何消息，也就是选举超时，然后他就会认为系统中没有可用的领导者然后开始进行选举以选出新的领导者
   * 要开始一次选举过程，follower会给当前的term加1并且转换为candidate状态，然后它会并行的向集群中的其他服务器节点发送请求投票的RPCs来给自己投票
   * 候选人的状态维持到知道发生以下任何一个条件发生的时候
     * 他自己赢得了这次的选举
     * 其他的服务器成为了领导者
     * 一段时间之后没有任何一个获胜的人