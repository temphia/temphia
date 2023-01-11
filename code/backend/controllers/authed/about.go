package authed

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type About struct {
	UserInfo User `json:"user_info,omitempty"`
	OrgInfo  Org  `json:"org_info,omitempty"`
}

type Org struct {
	Name   string `json:"name,omitempty"`
	Slug   string `json:"slug,omitempty"`
	OrgBio string `json:"org_bio,omitempty"`
}

type User struct {
	UserId    string `json:"user_id,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Bio       string `json:"bio,omitempty"`
	GroupID   string `json:"group_id,omitempty"`
	PublicKey string `json:"pub_key,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}

func (c *Controller) About(uclaim *claim.User) (*About, error) {

	user, err := c.coredb.GetUserByID(uclaim.TenentId, uclaim.UserID)
	if err != nil {
		return nil, err
	}

	tenant, err := c.coredb.GetTenant(uclaim.TenentId)
	if err != nil {
		return nil, err
	}

	group, err := c.coredb.GetUserGroup(uclaim.TenentId, uclaim.UserGroup)
	if err != nil {
		return nil, err
	}

	about := &About{
		UserInfo: User{
			UserId:    user.UserId,
			FullName:  user.FullName,
			Bio:       user.Bio,
			GroupID:   user.GroupID,
			PublicKey: user.PublicKey,
			GroupName: group.Name,
		},
		OrgInfo: Org{
			Name:   tenant.Name,
			Slug:   tenant.Slug,
			OrgBio: tenant.OrgBio,
		},
	}

	return about, nil

}
