package consts

const (
	// PostPrefix Post相关数据 eg:
	//  post:'post_id' : {
	// 						collection_count:1,
	//						upvote_count:1,
	//							}
	PostPrefix = "post:"
	// CollectPrefix redis 收藏相关的 eg: collect:'user_id' = zset()
	CollectPrefix       = "collect:"
	PostCollectCountKey = "collection_count"
	// UpvotePrefix redis 点赞相关的 eg: upvote:'user_id' = zset()
	UpvotePrefix       = "upvote:"
	PostUpvoteCountKey = "upvote_count"
)
