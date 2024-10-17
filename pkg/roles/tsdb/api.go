package tsdb

import (
	"context"
	"strconv"
	"strings"
	"time"

	"beryju.io/gravity/pkg/roles/tsdb/types"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func (r *Role) APIMetrics() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input types.APIMetricsGetInput, output *types.APIMetricsGetOutput) error {
		pf := r.i.KV().Key(
			types.KeyRole,
			string(input.Role),
		)
		if input.Category != nil {
			pf = pf.Add(*input.Category)
		}
		prefix := pf.Prefix(true).String()
		rawMetrics, err := r.i.KV().Get(
			ctx,
			prefix,
			clientv3.WithPrefix(),
		)
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		for _, kv := range rawMetrics.Kvs {
			keyParts := strings.Split(strings.TrimPrefix(string(kv.Key), r.i.KV().Key(types.KeyRole).Prefix(true).String()), "/")
			ts, err := strconv.Atoi(keyParts[len(keyParts)-1])
			if err != nil {
				r.log.Warn("failed to parse timestamp", zap.Error(err), zap.String("key", string(kv.Key)))
				continue
			}
			if input.Since != nil && ts < int(input.Since.Unix()) {
				continue
			}
			value, err := strconv.ParseInt(string(kv.Value), 10, 0)
			if err != nil {
				r.log.Warn("failed to parse metric value", zap.Error(err))
				continue
			}
			output.Records = append(output.Records, types.APIMetricsRecord{
				Keys:  keyParts[:2],
				Time:  time.Unix(int64(ts), 0),
				Node:  keyParts[len(keyParts)-2],
				Value: value,
			})
		}
		return nil
	})
	u.SetName("tsdb.get_metrics")
	u.SetTitle("Retrieve Metrics")
	u.SetTags("roles/tsdb")
	u.SetExpectedErrors(status.Internal)
	return u
}
