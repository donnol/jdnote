package notesrv

import (
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/timer"
	"github.com/donnol/tools/log"
	"github.com/pkg/errors"
)

// noteImpl 笔记
type noteImpl struct {
	noteStore notestore.Noter
}

// GetPage 获取分页
func (n *noteImpl) GetPage(ctx context.Context, param PageParam) (r PageResult, err error) {
	entity := notemodel.Entity{
		Title:  param.Title,
		Detail: param.Detail,
	}
	res, err := n.noteStore.GetPage(ctx, entity, param.Param)
	if err != nil {
		return
	}

	r.List = make([]Result, 0, len(res))
	var tmp Result
	for i, single := range res {
		if i == 0 {
			r.Total = single.Total
		}
		tmp = Result{}

		tmp, err = tmp.Init(single.Entity)
		if err != nil {
			return
		}

		// 详情截取前30个字符
		limit := 30
		if len([]rune(single.Detail)) > limit {
			tmp.Detail = string([]rune(single.Detail)[:limit])
		}

		r.List = append(r.List, tmp)
	}

	return
}

// Get 获取
func (n *noteImpl) Get(ctx context.Context, id int) (r Result, err error) {
	e, err := n.noteStore.Get(ctx, id)
	if err != nil {
		return
	}
	r, err = r.Init(e)
	if err != nil {
		return
	}

	return
}

func (n *noteImpl) GetPublish(ctx context.Context, id int) (r Result, err error) {
	e, err := n.noteStore.Get(ctx, id)
	if err != nil {
		return
	}
	// 检查是否publish状态
	if !e.Status.IsPublish() {
		return r, errors.Errorf("不存在该笔记")
	}
	r, err = r.Init(e)
	if err != nil {
		return
	}

	return
}

// AddOne 添加
func (n *noteImpl) AddOne(ctx context.Context) (id int, err error) {
	id, err = n.noteStore.AddOne(ctx)
	if err != nil {
		return
	}

	return
}

// Mod 修改
func (n *noteImpl) Mod(ctx context.Context, id int, p *Param) (err error) {
	if err = n.noteStore.Mod(ctx, id, &notemodel.Entity{
		Title:  p.Title,
		Detail: p.Detail,
	}); err != nil {
		return
	}

	return
}

// Del 删除
func (n *noteImpl) Del(ctx context.Context, id int) (err error) {
	err = n.noteStore.Del(ctx, id)
	if err != nil {
		return
	}

	return
}

// Publish 发布
func (n *noteImpl) Publish(ctx context.Context, id int) error {
	// 获取内容
	_, err := n.noteStore.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := n.noteStore.ModStatus(ctx, id, notemodel.StatusPublish); err != nil {
		return err
	}

	return nil
}

// Hide 隐藏
func (n *noteImpl) Hide(ctx context.Context, id int) error {
	// 获取内容
	_, err := n.noteStore.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := n.noteStore.ModStatus(ctx, id, notemodel.StatusDraft); err != nil {
		return err
	}

	return nil
}

func (n *noteImpl) Timer(ctx context.Context) timer.FuncJob {
	mark := " | noteImpl Timer | "
	return timer.WithTimeConsuming(mark, func() {
		log.Default().Infof("Timer begin.")
		defer log.Default().Infof("Timer finish.")

		// do something
	})
}
