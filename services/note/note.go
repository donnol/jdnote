package note

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/donnol/jdnote/models/note"
	"github.com/donnol/jdnote/utils/context"
)

// noteImpl 笔记
type noteImpl struct {
	noteModel note.Noter
}

// GetPage 获取分页
func (n *noteImpl) GetPage(ctx context.Context, param PageParam) (r PageResult, err error) {
	entity := note.Entity{
		Title:  param.Title,
		Detail: param.Detail,
	}
	res, total, err := n.noteModel.GetPage(ctx, entity, param.Param)
	if err != nil {
		return
	}
	r.Total = total

	r.List = make([]Result, 0, len(res))
	var tmp Result
	for _, single := range res {
		tmp = Result{}

		tmp, err = tmp.Init(single)
		if err != nil {
			return
		}

		r.List = append(r.List, tmp)
	}

	return
}

// Get 获取
func (n *noteImpl) Get(ctx context.Context, id int) (r Result, err error) {
	e, err := n.noteModel.Get(ctx, id)
	if err != nil {
		return
	}
	r, err = r.Init(e)
	if err != nil {
		return
	}

	return
}

// AddOne 添加
func (n *noteImpl) AddOne(ctx context.Context) (id int, err error) {
	id, err = n.noteModel.AddOne(ctx)
	if err != nil {
		return
	}

	return
}

// Mod 修改
func (n *noteImpl) Mod(ctx context.Context, id int, p Param) (err error) {
	if err = n.noteModel.Mod(ctx, id, note.Entity{
		Title:  p.Title,
		Detail: p.Detail,
	}); err != nil {
		return
	}

	return
}

// Del 删除
func (n *noteImpl) Del(ctx context.Context, id int) (err error) {
	err = n.noteModel.Del(ctx, id)
	if err != nil {
		return
	}

	return
}

// Publish 发布
func (n *noteImpl) Publish(ctx context.Context, id int) error {
	// 获取内容
	data, err := n.noteModel.Get(ctx, id)
	if err != nil {
		return err
	}

	// 生成md文件
	now := time.Now()
	content := n.getHugoContent(data.Title, data.Detail, now.Format("2006-01-02 15:04:05"), true, []string{}, []string{}, []string{})

	// 重新生成网页
	filename := strings.ReplaceAll(data.Title, " ", "_")
	filename += ".md"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (n *noteImpl) getHugoContent(title, detail, date string, isDraft bool, categories, tags, keywords []string) string {
	var content string

	headFormat := `---
title: "%s"
date: %s
draft: %s`
	var draftStr string
	if isDraft {
		draftStr = "true"
	} else {
		draftStr = "false"
	}
	content += fmt.Sprintf(headFormat, title, date, draftStr)

	for i, single := range categories {
		if i == 0 {
			content += `
categories:`
		}

		content += fmt.Sprintf(`
- %s`, single)
	}

	for i, single := range tags {
		if i == 0 {
			content += `
tags:`
		}

		content += fmt.Sprintf(`
- %s`, single)
	}

	for i, single := range keywords {
		if i == 0 {
			content += `
keywords:`
		}

		content += fmt.Sprintf(`
- %s`, single)
	}

	content += fmt.Sprintf(`
---

%s
`, detail)

	return content
}

// Hide TODO:隐藏
func (n *noteImpl) Hide(ctx context.Context, id int) error {

	return nil
}
