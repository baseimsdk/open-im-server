package rpccache

import (
	"context"
	pbconversation "github.com/OpenIMSDK/protocol/conversation"
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/localcache"
	"github.com/openimsdk/open-im-server/v3/pkg/common/cachekey"
	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
	"github.com/openimsdk/open-im-server/v3/pkg/rpcclient"
	"github.com/redis/go-redis/v9"
)

func NewConversationLocalCache(client rpcclient.ConversationRpcClient, cli redis.UniversalClient) *ConversationLocalCache {
	lc := config.Config.LocalCache.Conversation
	log.ZDebug(context.Background(), "ConversationLocalCache", "topic", lc.Topic, "slotNum", lc.SlotNum, "slotSize", lc.SlotSize, "enable", lc.Enable())
	x := &ConversationLocalCache{
		client: client,
		local: localcache.New[any](
			localcache.WithLocalSlotNum(lc.SlotNum),
			localcache.WithLocalSlotSize(lc.SlotSize),
		),
	}
	if lc.Enable() {
		go subscriberRedisDeleteCache(context.Background(), cli, lc.Topic, x.local.DelLocal)
	}
	return x
}

type ConversationLocalCache struct {
	client rpcclient.ConversationRpcClient
	local  localcache.Cache[any]
}

func (c *ConversationLocalCache) GetConversationIDs(ctx context.Context, ownerUserID string) (val []string, err error) {
	log.ZDebug(ctx, "ConversationLocalCache GetConversationIDs req", "ownerUserID", ownerUserID)
	defer func() {
		if err == nil {
			log.ZDebug(ctx, "ConversationLocalCache GetConversationIDs return", "value", val)
		} else {
			log.ZError(ctx, "ConversationLocalCache GetConversationIDs return", err)
		}
	}()
	return localcache.AnyValue[[]string](c.local.Get(ctx, cachekey.GetConversationIDsKey(ownerUserID), func(ctx context.Context) (any, error) {
		log.ZDebug(ctx, "ConversationLocalCache GetConversationIDs rpc", "ownerUserID", ownerUserID)
		return c.client.GetConversationIDs(ctx, ownerUserID)
	}))
}

func (c *ConversationLocalCache) GetConversation(ctx context.Context, userID, conversationID string) (val *pbconversation.Conversation, err error) {
	log.ZDebug(ctx, "ConversationLocalCache GetConversation req", "userID", userID, "conversationID", conversationID)
	defer func() {
		if err == nil {
			log.ZDebug(ctx, "ConversationLocalCache GetConversation return", "value", val)
		} else {
			log.ZError(ctx, "ConversationLocalCache GetConversation return", err)
		}
	}()
	return localcache.AnyValue[*pbconversation.Conversation](c.local.Get(ctx, cachekey.GetConversationKey(userID, conversationID), func(ctx context.Context) (any, error) {
		log.ZDebug(ctx, "ConversationLocalCache GetConversation rpc", "userID", userID, "conversationID", conversationID)
		return c.client.GetConversation(ctx, userID, conversationID)
	}))
}

func (c *ConversationLocalCache) GetSingleConversationRecvMsgOpt(ctx context.Context, userID, conversationID string) (int32, error) {
	conv, err := c.GetConversation(ctx, userID, conversationID)
	if err != nil {
		return 0, err
	}
	return conv.RecvMsgOpt, nil
}
