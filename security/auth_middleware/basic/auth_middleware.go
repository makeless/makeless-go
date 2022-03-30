package basic

import (
	"context"
	"fmt"
	"github.com/makeless/makeless-go/v2/security/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"regexp"
	"strings"
)

type AuthMiddleware struct {
	Auth           makeless_go_auth.Auth
	AuthMethods    map[string]bool
	NonAuthMethods map[string]bool
}

type AuthMiddlewareContextKey struct{}

func (authMiddleware *AuthMiddleware) AuthFunc(ctx context.Context) (context.Context, error) {
	var err error
	var ok bool
	var token, method string
	var claim *makeless_go_auth.Claim

	if method, ok = grpc.Method(ctx); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid method")
	}

	if _, ok = authMiddleware.NonAuthMethods[method]; ok {
		return ctx, nil
	}

	if _, ok = authMiddleware.AuthMethods[method]; !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unknown method")
	}

	if token, ok, err = authMiddleware.TokenLookup(ctx); err != nil || !ok {
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "no token")
		}
	}

	if claim, err = authMiddleware.Auth.Verify(token); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return context.WithValue(ctx, AuthMiddlewareContextKey{}, claim), nil
}

func (authMiddleware *AuthMiddleware) TokenLookup(ctx context.Context) (string, bool, error) {
	var err error
	var ok bool
	var md metadata.MD

	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		return "", false, nil
	}

	if authorizationMd := md.Get("authorization"); len(authorizationMd) == 1 {
		splits := strings.SplitN(authorizationMd[0], " ", 2)

		if len(splits) < 2 {
			return "", true, fmt.Errorf("bad authorization string")
		}

		if !strings.EqualFold(splits[0], "bearer") {
			return "", true, fmt.Errorf("request unauthenticated with %s", "bearer")
		}

		return splits[1], true, nil
	}

	if cookieMd := md.Get("cookies"); len(cookieMd) > 0 {
		var regex *regexp.Regexp

		if regex, err = regexp.Compile(`jwt=([\w\.\-]+)`); err != nil {
			return "", false, err
		}

		for _, cookie := range cookieMd {
			var match = regex.FindStringSubmatch(cookie)

			if len(match) < 2 {
				continue
			}

			return match[1], true, nil
		}
	}

	return "", false, nil
}
