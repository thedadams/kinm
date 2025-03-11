package strategy

import (
	"context"
	"time"

	"github.com/obot-platform/kinm/pkg/db/glogrus"
	"github.com/obot-platform/kinm/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var log = glogrus.New(glogrus.Config{})

var _ rest.Getter = (*GetAdapter)(nil)

type Getter interface {
	Get(ctx context.Context, namespace, name string) (types.Object, error)
}

func NewGet(strategy Getter) *GetAdapter {
	return &GetAdapter{
		strategy: strategy,
	}
}

type GetAdapter struct {
	strategy Getter
}

func (a *GetAdapter) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	ns, _ := request.NamespaceFrom(ctx)
	start := time.Now()
	defer func() {
		log.Info(ctx, "Get %s/%s took %s", ns, name, time.Since(start))
	}()
	return a.strategy.Get(ctx, ns, name)
}
