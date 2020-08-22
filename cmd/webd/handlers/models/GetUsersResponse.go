package handlermodels

import userpb "github.com/SmitSheth/Mini-twitter/internal/user/userpb"

type GetUsersResponse struct {
	Users []*userpb.AccountInformation
}
