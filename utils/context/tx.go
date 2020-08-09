package context

import "github.com/donnol/jdnote/utils/store/db"

func WithTx(ctx Context, f func(ctx Context) error) (err error) {
	if err = db.WithTx(ctx.DB(), func(tx db.DB) error {
		// 新建变量，不要修改了原有变量
		var ctx = ctx.NewWithTx(tx)

		if err := f(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
