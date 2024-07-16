package log

import (
	"context"

	"github.com/ecodeclub/webook/internal/ai/internal/service/gpt/handler"

	"github.com/ecodeclub/webook/internal/ai/internal/domain"
	"github.com/gotomicro/ego/core/elog"
)

type HandlerBuilder struct {
	logger *elog.Component
}

var _ handler.Builder = &HandlerBuilder{}

func NewHandler() *HandlerBuilder {
	return &HandlerBuilder{
		logger: elog.DefaultLogger,
	}
}

func (h *HandlerBuilder) Name() string {
	return "log"
}

func (h *HandlerBuilder) Next(next handler.Handler) handler.Handler {
	return handler.HandleFunc(func(ctx context.Context, req domain.GPTRequest) (domain.GPTResponse, error) {
		logger := h.logger.With(elog.String("tid", req.Tid),
			elog.Int64("uid", req.Uid),
			elog.String("biz", req.Biz))
		// 记录请求
		logger.Info("请求 GPT")
		resp, err := next.Handle(ctx, req)
		if err != nil {
			// 记录错误
			logger.Error("请求gpt服务失败", elog.FieldErr(err))
			return resp, err
		}
		// 记录响应
		logger.Info("请求gpt服务响应成功", elog.Int64("tokens", resp.Tokens))
		return resp, err
	})
}
