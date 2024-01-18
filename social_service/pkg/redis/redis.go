package redis

import "strconv"

func FollowAction(userID, toUserID int64, actionType int32) error {
	// 先查询该用户是否关注这个人
	result, err := Redis.SIsMember(ctx, GenerateFollowKey(userID), toUserID).Result()
	if err != nil {
		return err
	}

	//不需要执行操作
	if (result == true && actionType == 1) || (result == true && actionType == 2) {
		return nil
	}

	//actionType 1: 关注 2: 取消关注
	pipe := Redis.TxPipeline()
	defer pipe.Close()
	if actionType == 1 { //关注操作
		pipe.SAdd(ctx, GenerateFollowKey(userID), toUserID)   //用户关注集合
		pipe.SAdd(ctx, GenerateFollowerKey(toUserID), userID) //被关注用户集合
		_, err := pipe.Exec(ctx)
		if err != nil {
			return err
		}
	} else { //取消关注操作
		pipe.SRem(ctx, GenerateFollowKey(userID), toUserID)
		pipe.SRem(ctx, GenerateFollowerKey(toUserID), userID)
		_, err := pipe.Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// FollowList 关注列表
func FollowList(userID int64, toUserIDS *[]int64) error {
	result, err := Redis.SMembers(ctx, GenerateFollowKey(userID)).Result()
	if err != nil {
		return err
	}

	for _, ID := range result {
		id, _ := strconv.ParseInt(ID, 10, 64)
		*toUserIDS = append(*toUserIDS, id)
	}

	return nil
}

// FollowerList 被关注列表
func FollowerList(toUserID int64, UserIDs *[]int64) error {
	result, err := Redis.SMembers(ctx, GenerateFollowKey(toUserID)).Result()
	if err != nil {
		return err
	}

	for _, ID := range result {
		id, _ := strconv.ParseInt(ID, 10, 64)
		*UserIDs = append(*UserIDs, id)
	}

	return nil
}

func FriendList(UserID int64, UserIDs *[]int64) error {
	_, err := Redis.Do(ctx, "SINTERSTORE", GenerateFriendKey(UserID), GenerateFollowKey(UserID), GenerateFollowerKey(UserID)).Result()
	if err != nil {
		return err
	}
	result, err := Redis.SMembers(ctx, GenerateFriendKey(UserID)).Result()
	if err != nil {
		return err
	}
	for _, r := range result {
		r64, _ := strconv.ParseInt(r, 10, 64)
		*UserIDs = append(*UserIDs, r64)
	}
	return nil
}

// FollowCount 关注人数
func FollowCount(UserID int64) (int64, error) {
	return Redis.SCard(ctx, GenerateFollowKey(UserID)).Result()
}

// FollowerCount 被关注人数
func FollowerCount(toUserID int64) (int64, error) {
	return Redis.SCard(ctx, GenerateFollowerKey(toUserID)).Result()
}

// IsFollow 是否被关注
func IsFollow(UserId int64, ToUserId int64) (bool, error) {
	result, err := Redis.SIsMember(ctx, GenerateFollowKey(UserId), ToUserId).Result()
	return result, err
}
