package responses

import "github.com/hidenari-yuda/ai-market-go/domain/entity"

type Chat struct {
	Chat *entity.Chat `json:"chat"`
}

func NewChat(chat *entity.Chat) Chat {
	return Chat{
		Chat: chat,
	}
}

type ChatList struct {
	ChatList []*entity.Chat `json:"chat_list"`
}

func NewChatList(chats []*entity.Chat) ChatList {
	return ChatList{
		ChatList: chats,
	}
}
