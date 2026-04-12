package main

import (
	"context"
	"fmt"
)

type contextKey string

const requestId contextKey = "requestId"

const spanID contextKey = "spanID"

func enter() {

	ctx := context.WithValue(context.Background(), requestId, 123)

	fmt.Printf("Result: %v", svcA(ctx))
}

func svcA(ctx context.Context) int {
	fmt.Printf("From svc A, requestId:%v\n", ctx.Value(requestId))
	return svcB(context.WithValue(ctx, spanID, 1))
}

func svcB(ctx context.Context) int {
	fmt.Printf("From svc B, requestId:%v\n", ctx.Value(requestId))
	return svcC(context.WithValue(ctx, spanID, 1))
}

func svcC(ctx context.Context) int {
	fmt.Printf("From svc C, requestId:%v\n", ctx.Value(requestId))
	return 456
}
