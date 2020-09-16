package request_id

import "context"

func GetRequestIDFromCtx(ctx context.Context) string {
	return ctx.Value("request-id").(string)
}
