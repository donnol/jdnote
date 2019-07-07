package note

import (
	"encoding/json"

	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
)

// Note 笔记
type Note struct {
	model.Base
}

// AddNote 添加笔记
func (note *Note) AddNote(ctx context.Context, entity Entity) (id int, err error) {
	err = ctx.DB().GetContext(ctx, &id, `INSERT INTO t_note(user_id, title, detail)
		VALUES($1, $2, $3)
		RETURNING id
		`,
		entity.UserID,
		entity.Title,
		entity.Detail,
	)
	if err != nil {
		return
	}
	return
}

// ModifyNote 修改笔记
func (note *Note) ModifyNote(ctx context.Context, id int, entity Entity) (err error) {
	_, err = ctx.DB().NamedExecContext(ctx, `Update t_note set
		title = :title,
		detail = :detail
		Where id = :id
		`,
		map[string]interface{}{
			"title":  entity.Title,
			"detail": entity.Detail,
			"id":     id,
		},
	)
	if err != nil {
		return
	}
	return
}

// GetNoteList 获取笔记列表
func (note *Note) GetNoteList(ctx context.Context, entity Entity, param model.CommonParam) (
	res struct {
		Data  []json.RawMessage
		Total int
	},
	err error,
) {
	var dbResult []struct {
		Data  json.RawMessage
		Total int
	}
	err = ctx.DB().Select(&dbResult, `
		SELECT json_build_object(
			'ID', id,
			'Title', title,
			'Detail', detail,
			'CreatedAt', created_at
		) AS data,
		COUNT(*) OVER () AS total
		FROM t_note
		WHERE true
		
		AND CASE WHEN $3 <> '' THEN
			title ~* $3
		ELSE true END
		AND CASE WHEN $4 <> 0 THEN
			id = $4
		ELSE true END
		ORDER BY id DESC
		LIMIT $1
		OFFSET $2
		`,
		param.Size,
		param.Offset,
		entity.Title,
		entity.ID,
	)
	if err != nil {
		return
	}
	for i, single := range dbResult {
		if i == 0 {
			res.Total = single.Total
		}
		res.Data = append(res.Data, single.Data)
	}
	return
}

// GetNote 获取笔记
func (note *Note) GetNote(ctx context.Context, id int) (entity Entity, err error) {
	err = ctx.DB().GetContext(ctx, &entity, `
		SELECT id, user_id, title, detail, created_at
		FROM t_note
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return
	}
	return
}
