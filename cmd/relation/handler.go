package main

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/relation/service"
	relation "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	invalidMsg := "invalid request"
	wrongMsg := "relation action failed"
	correctMsg := "relation action success"

	if req.ToUserId < 0 || req.ActionType < 0 || req.ActionType > 2 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return
	}
	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &wrongMsg
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	invalidMsg := "invalid request"
	failMsg := "relation follow list failed"
	correctMsg := "relation follow list success"
	// 没有找到用户和用户的关注列表
	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return resp, nil
	}

	var followList []*relation.User
	followList, err = service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	resp.UserList = followList
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	invalidMsg := "invalid request"
	failMsg := "Not Found your follower list"
	correctMsg := "Search your follower list success"

	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return resp, nil
	}
	var users []*relation.User
	users, err = service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &correctMsg
	resp.UserList = users

	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)

	inValidMsg := "Relation Friendlist request inValid"
	failMsg := "Not Found your friend list"
	successMsg := "Search your friend list success"

	if req.UserId < 0 || req.MUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &inValidMsg
		return resp, nil
	}

	var users []*relation.User
	users, err = service.NewRelationFriendListService(ctx).RelationFriendList(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = &successMsg
	resp.UserList = users

	return resp, nil
}

// RelationFollowCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowCount(ctx context.Context, req *relation.RelationFollowCountRequest) (resp *relation.RelationFollowCountResponse, err error) {
	resp = new(relation.RelationFollowCountResponse)

	invalidMsg := "Relation FollowCount request inValid"
	failMsg := "Relation FollowCount fail"
	successMsg := "Relation FollowCount success"
	if req.UserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &invalidMsg
		return resp, nil
	}

	followCount, err := service.NewRelationFollowCountService(ctx).RelationFollowCount(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = &successMsg
	resp.FollowCount = followCount

	return resp, nil
}

// RelationFollowerCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerCount(ctx context.Context, req *relation.RelationFollowCountRequest) (resp *relation.RelationFollowCountResponse, err error) {
	resp = new(relation.RelationFollowCountResponse)

	inValidMsg := "Relation FollowerCount request inValid"
	failMsg := "Relation FollowerCount fail"
	successMsg := "Relation FollowerCount success"

	if req.UserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = &inValidMsg
		return resp, nil
	}

	followerCount, err := service.NewRelationFollowerCountService(ctx).RelationFollowerCount(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = &failMsg
		return resp, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = &successMsg
	resp.FollowCount = followerCount

	return resp, nil
}

// RelationIsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationIsFollow(ctx context.Context, req *relation.RelationIsFollowRequest) (resp *relation.RelationIsFollowResponse, err error) {
	resp = new(relation.RelationIsFollowResponse)

	if req.UserId < 0 || req.ToUserId < 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "Relation IsFollow request inValid"
		return resp, nil
	}

	isFollow, err := service.NewRelationIsFollowService(ctx).RelationIsFollow(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Relation IsFollow fail"
		return resp, err
	}

	resp.StatusCode = 0
	resp.StatusMsg = "Relation isFollow success"
	resp.IsFollow = isFollow

	return resp, nil
}
