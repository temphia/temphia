package authed

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"golang.org/x/oauth2"
)

type OauthExchanger struct {
	auth     *entities.UserGroupAuth
	state    *claim.OauthState
	authCode string
}

type EData struct {
	email       string
	userName    string
	accessToken string
}

func (oe *OauthExchanger) Exchange() (*EData, error) {

	err := oe.checkState()
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     oe.auth.ClientId(),
		ClientSecret: oe.auth.ClientSecret(),
		RedirectURL:  `http://localhost:4000/z/auth/oauth_redirect`,
		Scopes:       strings.Split(oe.auth.Scopes, ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:   oe.auth.AuthURL(),
			TokenURL:  oe.auth.TokenURL(),
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	token, err := conf.Exchange(context.TODO(), oe.authCode)
	if err != nil {
		return nil, err
	}

	mc := jwt.MapClaims{}
	jparser := jwt.NewParser(jwt.WithoutClaimsValidation())

	_, _, err = jparser.ParseUnverified(token.Extra("id_token").(string), &mc)
	if err != nil {
		return nil, err
	}

	return &EData{
		email:       mc["email"].(string),
		userName:    "",
		accessToken: token.AccessToken,
	}, nil
}

func (oe *OauthExchanger) checkState() error {

	if oe.state.UserGroup != oe.auth.UserGroup {
		return easyerr.Error("Wrong user group")
	}

	return nil
}
