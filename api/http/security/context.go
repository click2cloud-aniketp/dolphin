package security

import (
	"context"
	"net/http"

	"dolphin/api"
)

type (
	contextKey int
)

const (
	contextAuthenticationKey contextKey = iota
	contextRestrictedRequest
)

// storeTokenData stores a TokenData object inside the request context and returns the enhanced context.
func storeTokenData(request *http.Request, tokenData *dockm.TokenData) context.Context {
	return context.WithValue(request.Context(), contextAuthenticationKey, tokenData)
}

// RetrieveTokenData returns the TokenData object stored in the request context.
func RetrieveTokenData(request *http.Request) (*dockm.TokenData, error) {
	contextData := request.Context().Value(contextAuthenticationKey)
	if contextData == nil {
		return nil, dockm.ErrMissingContextData
	}

	tokenData := contextData.(*dockm.TokenData)
	return tokenData, nil
}

// storeRestrictedRequestContext stores a RestrictedRequestContext object inside the request context
// and returns the enhanced context.
func storeRestrictedRequestContext(request *http.Request, requestContext *RestrictedRequestContext) context.Context {
	return context.WithValue(request.Context(), contextRestrictedRequest, requestContext)
}

// RetrieveRestrictedRequestContext returns the RestrictedRequestContext object stored in the request context.
func RetrieveRestrictedRequestContext(request *http.Request) (*RestrictedRequestContext, error) {
	contextData := request.Context().Value(contextRestrictedRequest)
	if contextData == nil {
		return nil, dockm.ErrMissingSecurityContext
	}

	requestContext := contextData.(*RestrictedRequestContext)
	return requestContext, nil
}
