// 全局common包，收集通用方法
package common

const (
	DefaultPageSiz = 1000
)

// Iterator 分页迭代执行，可用于大量查询时的分批次执行
// pageSize 每页页码，total 为数据总量
func Iterator(pageSize, total int, handler func(offset, limit int)) {
	pageIte := pageIterator{
		pageMaxNum: pageSize,
		total:      total,
		curPage:    0,
	}
	for pageIte.HasNext() {
		handler(pageIte.Next())
	}
}

// PageIterator 分页查询迭代器接口
type PageIterator interface {
	// HasNext 是否还有下一页
	HasNext() bool
	// Next 获取下一页的截取段
	Next() (int, int)
}

type pageIterator struct {
	total      int
	pageMaxNum int
	curPage    int
}

// HasNext 是否还有下一页
func (p *pageIterator) HasNext() bool {
	return p.curPage*p.pageMaxNum < p.total
}

// Next 获取下一页的截取段
func (p *pageIterator) Next() (offset, limit int) {
	offset, limit = p.curPage*p.pageMaxNum, min((p.curPage+1)*p.pageMaxNum, p.total)
	p.curPage++
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
