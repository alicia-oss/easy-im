package service

import (
	"easy_im/internal/domain/im/model"
	"easy_im/internal/domain/im/pkg"
	"easy_im/internal/domain/im/repo"
)

type IGroupService interface {
	GetGroupByUserId(userId uint) ([]*model.Group, error)
	GetUserIdInGroup(groupId uint) ([]uint, error)
	AddUserToGroup(userId, groupId uint) error
	RemoveUserFromGroup(userId, groupId uint) error

	CreateGroup(group *model.Group) error
	DeleteGroup(groupId uint) error
	UpdateGroupInfo(group *model.Group) error
}

type groupServiceImpl struct {
	groupRepo     repo.IGroupRepo
	groupUserRepo repo.IGroupUserRepo
}

func (g *groupServiceImpl) GetGroupByUserId(userId uint) ([]*model.Group, error) {
	gus, err := g.groupUserRepo.GetByUserId(userId)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	gIds := make([]uint, len(gus))
	for i, gu := range gus {
		gIds[i] = gu.GroupId
	}
	groups, err := g.groupRepo.GetByIds(gIds)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	return groups, nil
}

func (g *groupServiceImpl) GetUserIdInGroup(groupId uint) ([]uint, error) {
	gus, err := g.groupUserRepo.GetByGroupId(groupId)
	if err != nil {
		return nil, pkg.ErrUnknown
	}
	uIds := make([]uint, len(gus))
	for i, gu := range gus {
		uIds[i] = gu.UserId
	}
	return uIds, nil
}

func (g *groupServiceImpl) AddUserToGroup(userId, groupId uint) error {
	err := g.groupUserRepo.Add(&model.GroupUser{GroupId: groupId, UserId: userId})
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (g *groupServiceImpl) RemoveUserFromGroup(userId, groupId uint) error {
	err := g.groupUserRepo.Delete(&model.GroupUser{GroupId: groupId, UserId: userId})
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (g *groupServiceImpl) CreateGroup(group *model.Group) error {
	err := g.groupRepo.Add(group)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (g *groupServiceImpl) DeleteGroup(groupId uint) error {
	err := g.groupRepo.Delete(groupId)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}

func (g *groupServiceImpl) UpdateGroupInfo(group *model.Group) error {
	err := g.groupRepo.Save(group)
	if err != nil {
		return pkg.ErrUnknown
	}
	return nil
}
