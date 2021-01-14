package process

import (
	"context"
	"fmt"
	"testing"
)

type successImpl struct {
	name string
}

func (i *successImpl) Do(ctx context.Context) error {
	fmt.Println(i.name, "Do")
	return nil
}

func (i *successImpl) Rollback(ctx context.Context) {
	fmt.Println(i.name, "Rollback")
}

func (i *successImpl) WhenFail(ctx context.Context, err error) {
	fmt.Printf("When failed: %+v\n", err)
}

type dofailImpl struct {
	name string
}

func (i *dofailImpl) Do(ctx context.Context) error {
	return fmt.Errorf("do failed")
}

func (i *dofailImpl) Rollback(ctx context.Context) {
	fmt.Println(i.name, "Rollback")
}

func (i *dofailImpl) WhenFail(ctx context.Context, err error) {
	fmt.Printf("When failed: %+v\n", err)
}

type panicDoImpl struct {
	name string
}

func (i *panicDoImpl) Do(ctx context.Context) error {
	panic("panic do")
}

func (i *panicDoImpl) Rollback(ctx context.Context) {
	fmt.Println(i.name, "Rollback")
}

func (i *panicDoImpl) WhenFail(ctx context.Context, err error) {
	fmt.Printf("When failed: %+v\n", err)
}

type panicRollbackImpl struct {
	name string
}

func (i *panicRollbackImpl) Do(ctx context.Context) error {
	fmt.Println(i.name, "Do")
	return nil
}

func (i *panicRollbackImpl) Rollback(ctx context.Context) {
	// 如果回滚失败了，自己想好重试策略
	panic("panic rollback")
}

func (i *panicRollbackImpl) WhenFail(ctx context.Context, err error) {
	fmt.Printf("When failed: %+v\n", err)
}

func TestAtomicRun(t *testing.T) {
	var ctx = context.Background()

	success, dofail := &successImpl{"jj"}, &dofailImpl{"jd"}
	panicDo, panicRollback := &panicDoImpl{"jp"}, &panicRollbackImpl{"jpr"}

	for i, cas := range []struct {
		p1 AtomicProcessor
		p2 AtomicProcessor
	}{
		{success, success},
		{success, dofail},
		{success, panicDo},
		{panicRollback, dofail},
	} {
		no := i + 1
		fmt.Printf("=== No.%d\n", no)

		if err := AtomicRun(ctx, cas.p1, cas.p2); err != nil {
			t.Logf("Run failed: %+v\n", err)
		} else {
			t.Logf("Run success.\n")
		}
	}
}
