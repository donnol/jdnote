package notestore

import (
	"time"

	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type noteImpl struct {
}

// AddOne 添加一条记录，并返回它的id
func (note *noteImpl) AddOne(ctx context.Context) (id int, err error) {
	err = ctx.DB().GetContext(ctx, &id, `INSERT INTO t_note(user_id, title, detail)
		VALUES($1, '', '')
		RETURNING id`,
		context.MustGetUserValue(ctx),
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加笔记
func (note *noteImpl) Add(ctx context.Context, entity notemodel.Entity) (id int, err error) {
	err = ctx.DB().GetContext(ctx, &id, `INSERT INTO t_note(user_id, title, detail)
		VALUES($1, $2, $3)
		RETURNING id
		`,
		entity.UserID,
		entity.Title,
		entity.Detail,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// Mod 修改笔记
func (note *noteImpl) Mod(ctx context.Context, id int, entity *notemodel.Entity) (err error) {
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
		err = errors.WithStack(err)
		return
	}
	return
}

func (note *noteImpl) ModStatus(ctx context.Context, id int, status notemodel.Status) (err error) {
	_, err = ctx.DB().NamedExecContext(ctx, `Update t_note set
		status = :status
		Where id = :id
		`,
		map[string]interface{}{
			"status": status,
			"id":     id,
		},
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// Del 删除笔记
func (note *noteImpl) Del(ctx context.Context, id int) (err error) {
	_, err = ctx.DB().NamedExecContext(ctx, `Delete FROM t_note
		Where id = :id
		`,
		map[string]interface{}{
			"id": id,
		},
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// GetPage 获取笔记分页
func (note *noteImpl) GetPage(ctx context.Context, entity notemodel.Entity, param common.Param) (
	res []notemodel.EntityWithTotal,
	err error,
) {

	err = ctx.DB().SelectContext(ctx, &res, `
		SELECT 
			id,
			title,
			detail,
			created_at,
			status,
			COUNT(*) OVER () AS total
		FROM t_note
		WHERE true
		
		AND CASE WHEN $3 <> '' THEN
			title ~* $3
		ELSE true END
		AND CASE WHEN $4 <> 0 THEN
			id = $4
		ELSE true END
		AND CASE WHEN $5 <> '' THEN
			detail ~* $5
		ELSE true END
		AND CASE WHEN $6 THEN
			created_at >= $7::timestamp
		ELSE true END
		AND CASE WHEN $8 THEN
			created_at <= $9::timestamp
		ELSE true END
		AND CASE WHEN $10 THEN
			status = 2
		ELSE true END

		ORDER BY id DESC
		LIMIT $1
		OFFSET $2
		`,
		param.PageSize,
		param.PageIndex,
		entity.Title,
		entity.ID,
		entity.Detail,

		param.BeginTime != 0,
		time.Unix(param.BeginTime, 0),
		param.EndTime != 0,
		time.Unix(param.EndTime, 0),
		param.OnlyPublish,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Get 获取笔记
func (note *noteImpl) Get(ctx context.Context, id int) (entity notemodel.Entity, err error) {
	err = ctx.DB().GetContext(ctx, &entity, `
		SELECT id, user_id, title, detail, status, created_at
		FROM t_note
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// GetList 获取笔记列表
func (note *noteImpl) GetList(ctx context.Context, ids []int64) (entitys []notemodel.Entity, err error) {
	if err = ctx.DB().SelectContext(ctx, &entitys, `
		SELECT id, user_id, title, detail, created_at
		FROM t_note
		WHERE id = any($1)
		ORDER BY id DESC
		`,
		pq.Int64Array(ids),
	); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
