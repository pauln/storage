package request

import (
	"context"
	"github.com/ory/fosite"
)

/* These functions provide a concrete implementation of fosite.AccessTokenStorage */

// CreateAccessTokenSession creates a new session for an Access Token in mongo
func (m *MongoManager) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) (err error) {
	return m.createSession(signature, request, mongoCollectionAccessTokens)
}

// GetAccessTokenSession returns a session if it can be found by signature in mongo
func (m MongoManager) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (request fosite.Requester, err error) {
	return m.findSessionBySignature(signature, session, mongoCollectionAccessTokens)
}

// DeleteAccessTokenSession removes an Access Tokens current session from mongo
func (m *MongoManager) DeleteAccessTokenSession(ctx context.Context, signature string) (err error) {
	return m.deleteSession(signature, mongoCollectionAccessTokens)
}
