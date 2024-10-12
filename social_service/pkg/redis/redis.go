package redis

import "strconv"

func FollowAction(userID, toUserID int64, actionType int32) error {
	// 先查询该用户是否关注这个人，检查用户关注信息的 Redis 集合中是否存在某个特定的用户（toUserID）
	result, err := Redis.SIsMember(ctx, GenerateFollowKey(userID), toUserID).Result()
	if err != nil {
		return err
	}

	//不需要执行操作
	if (result == true && actionType == 1) || (result == true && actionType == 2) {
		return nil
	}

	//actionType 1: 关注 2: 取消关注
	// 创建了一个事务管道，允许你将多个 Redis 命令打包成一个事务，并通过一次网络往返发送到 Redis 服务器。
	// 这样可以减少网络延迟，并且可以在服务器端一次性执行所有的命令，而不需要等待每个命令的回复。
	pipe := Redis.TxPipeline()
	defer pipe.Close()
	if actionType == 1 { //关注操作
		pipe.SAdd(ctx, GenerateFollowKey(userID), toUserID)   //用户关注集合
		pipe.SAdd(ctx, GenerateFollowerKey(toUserID), userID) //被关注用户集合
		// 执行了事务管道中的所有命令
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
	// 调用了 Redis 的 Do 方法，它允许你执行任意 Redis 命令。"SINTERSTORE" 是 Redis 的命令名，用于计算集合的交集并存储到一个新的集合中。
	// 后面跟着的三个参数是要计算交集的集合，这里分别是用户的关注集合、被关注者的粉丝集合以及用户自身。
	// 这样做的目的可能是为了获取用户的朋友（即关注了用户并且被用户关注的用户），并将结果存储在一个新的集合中。
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
	// 调用了 Redis 的 SCard 方法，它用于获取指定集合的成员数量。
	return Redis.SCard(ctx, GenerateFollowKey(UserID)).Result()
}

// FollowerCount 被关注人数
func FollowerCount(toUserID int64) (int64, error) {
	return Redis.SCard(ctx, GenerateFollowerKey(toUserID)).Result()
}

// IsFollow 是否被关注
func IsFollow(UserId int64, ToUserId int64) (bool, error) {
	// 获取当前用户的关注列表中有没有（被关注的）这个用户的 ID在里面？
	result, err := Redis.SIsMember(ctx, GenerateFollowKey(UserId), ToUserId).Result()
	return result, err
}
