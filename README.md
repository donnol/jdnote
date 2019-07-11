# jdnote

简单笔记。

- [√] 消息队列：[nsq](https://github.com/nsqio/nsq)
- [√] 运行时监控：[pprof](https://golang.org/pkg/net/http/pprof/)，[prometheus](https://github.com/prometheus/prometheus)和[grafana](https://grafana.com/grafana)
- [√] 全文搜索：[bleve](https://github.com/blevesearch/bleve)
- [] 数据库的表关系图生成

> 表关联：结构体通过 rel 标签指向具体的表的字段，如：ContentID int `rel:"content.id"`
>
> 图表示：graphviz
