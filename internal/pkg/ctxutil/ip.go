package ctxutil

import "context"

type ipKey struct{}

// GetSourceIP returns the source IP address from the context.
func GetSourceIP(ctx context.Context) string {
	ip, _ := ctx.Value(ipKey{}).(string)
	return ip
}

func AddIPToContext(ctx context.Context, ip string) context.Context {
	newContext := context.WithValue(ctx, ipKey{}, ip)
	return newContext
}
