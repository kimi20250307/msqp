package mj

import (
	"core/models/enums"
	"game/component/mj/mp"
	"game/component/proto"
)

type MessageReq struct {
	Type int         `json:"type"`
	Data MessageData `json:"data"`
}
type MessageData struct {
	ChairID     int         `json:"chairID"`
	Type        int         `json:"type"`
	Msg         string      `json:"msg"`
	RecipientID int         `json:"recipientID"`
	Card        mp.CardID   `json:"card"`
	Operate     OperateType `json:"operate"`
	Trust       bool        `json:"trust"`
}

type GameData struct {
	BankerChairID  int              `json:"bankerChairID"`  //庄家
	ChairCount     int              `json:"chairCount"`     //总座次人数
	CurBureau      int              `json:"curBureau"`      //当前局数
	GameStatus     GameStatus       `json:"gameStatus"`     //游戏状态
	GameStarted    bool             `json:"gameStarted"`    //是否已经开始
	Tick           int              `json:"tick"`           //倒计时
	MaxBureau      int              `json:"maxBureau"`      //最大局数
	CurChairID     int              `json:"curChairID"`     //当前玩家
	UserTrustArray []bool           `json:"userTrustArray"` //托管
	HandCards      [][]mp.CardID    `json:"handCards"`      //手牌
	OperateArrays  [][]OperateType  `json:"operateArrays"`  //操作
	OperateRecord  []*OperateRecord `json:"operateRecord"`  //操作记录
	RestCardsCount int              `json:"restCardsCount"` //剩余牌数
	Result         *GameResult      `json:"result"`         //结算
}
type UserWinRecord struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Score    int    `json:"score"`
}
type UserRecord struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	ChairID  int    `json:"chairID"`
}
type GameResult struct {
	Scores          []int         `json:"scores"`
	HandCards       [][]mp.CardID `json:"handCards"`
	MyMaCards       []MyMaCard    `json:"myMaCards"`
	RestCards       []mp.CardID   `json:"restCards"`
	WinChairIDArray []int         `json:"winChairIDArray"`
	GangChairID     int           `json:"gangChairID"`
	FangGangArray   []int         `json:"fangGangArray"`
	HuType          OperateType   `json:"huType"`
}
type MyMaCard struct {
	Card mp.CardID `json:"card"`
	Win  bool      `json:"win"`
}
type OperateRecord struct {
	ChairID int         `json:"chairID"`
	Card    *mp.CardID  `json:"card"`
	Operate OperateType `json:"operate"`
}

type ReviewRecord struct {
	RoomID        string               `json:"roomID"`
	HandCards     [][]mp.CardID        `json:"handCards"`
	OperateRecord []*OperateRecord     `json:"operateRecord"`
	UserArray     []proto.UserRoomData `json:"userArray"`
	CardsCount    int                  `json:"cardsCount"`
	MaxBureau     int                  `json:"maxBureau"`
	Qidui         bool                 `json:"qidui"`
	Result        *GameResult          `json:"result"`
}
type BureauReview struct {
	Uid      string `json:"uid"`
	WinScore int    `json:"winScore"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	IsBanker bool   `json:"isBanker"`
}

const operateTm1 = 30 // 弃牌操作时间
const operateTm2 = 30 //  碰杠操作时间

type OperateType int

const (
	OperateTypeNone OperateType = iota
	HuChi                       //吃胡
	HuZi                        //自摸
	Peng                        //碰
	GangChi                     //吃杠
	GangBu                      //补杠
	GangZi                      //自摸杠
	Guo                         //过
	Qi                          //弃
	Get                         //拿牌
)

type GameStatus int

const (
	GameStatusNone GameStatus = iota
	Dices                     //掷骰子
	SendCards                 //发牌
	Playing                   //游戏
	ZhaMa                     //扎码
	Result                    //结算
)

type GameStatusTm int

const (
	GameStatusTmNone   GameStatusTm = 0
	GameStatusTmDices               = 3 //掷骰子
	GameStatusTmSend                = 3 //发牌
	GameStatusTmPlay                = 0 //游戏
	GameStatusTmZha                 = 5 //扎码
	GameStatusTmResult              = 5 //结算
)

type GameType int

const (
	HongZhong4 GameType = 1
	HongZhong8          = 2
)

const OperateTime int = 30 //操作时间
const OperateQi int = 30   //弃牌操作时间
const OperatePg int = 30   //碰杠操作时间

const (
	GameStatusPush         = 401 //游戏状态推送
	GameBankerPush         = 402 //庄家推送
	GameDicesPush          = 403 //骰子推送
	GameSendCardsPush      = 404 //发牌推送
	GameRestCardsCountPush = 405 //剩余牌数推送
	GameTurnPush           = 406 //操作推送 轮到谁出牌了
	GameTurnOperateNotify  = 307 //操作通知
	GameTurnOperatePush    = 407 //操作推送
	GameResultPush         = 408 //结果推送
	GameBureauPush         = 409 //局数推送
	GameEndPush            = 410 //结束推送
	GameChatNotify         = 311 //游戏聊天
	GameChatPush           = 411
	GameTrustNotify        = 312 //托管通知
	GameTrustPush          = 412 //托管推送
	GameReviewNotify       = 313 //游戏回顾通知
	GameReviewPush         = 413 //游戏回顾推送
	GameDismissPush        = 414 //解散推送
	GameGetCardNotify      = 315 //拿牌通知
	GameGetCardPush        = 415 //拿牌推送
)

type DismissUser struct {
	Uid           string `json:"uid"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	HuCount       int    `json:"huCount"`
	GongGangCount int    `json:"gongGangCount"`
	AnGangCount   int    `json:"anGangCount"`
	MaCount       int    `json:"maCount"`
	WinScore      int    `json:"winScore"`
}
type Creator struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func GameDismissPushData(
	userArray []*DismissUser,
	creator *Creator,
	reason enums.RoomDismissReason,
	hongBaoList any) any {
	return map[string]any{
		"type": GameDismissPush,
		"data": map[string]any{
			"userArray":   userArray,
			"creator":     creator,
			"reason":      reason,
			"hongBaoList": hongBaoList,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameStatusPushData(gameStatus GameStatus, tick int) any {
	return map[string]any{
		"type": GameStatusPush,
		"data": map[string]any{
			"gameStatus": gameStatus,
			"tick":       tick,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameBankerPushData(bankerChairID int) any {
	return map[string]any{
		"type": GameBankerPush,
		"data": map[string]any{
			"bankerChairID": bankerChairID,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameDicesPushData(dice1, dice2 int) any {
	return map[string]any{
		"type": GameDicesPush,
		"data": map[string]any{
			"dice1": dice1,
			"dice2": dice2,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameSendCardsPushData(handCards [][]mp.CardID, chairID int) any {
	return map[string]any{
		"type": GameSendCardsPush,
		"data": map[string]any{
			"handCards": handCards,
			"chairID":   chairID,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameRestCardsCountPushData(restCardsCount int) any {
	return map[string]any{
		"type": GameRestCardsCountPush,
		"data": map[string]any{
			"restCardsCount": restCardsCount,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameBureauPushData(curBureau int) any {
	return map[string]any{
		"type": GameBureauPush,
		"data": map[string]any{
			"curBureau": curBureau,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameTurnPushData(chairID int, card mp.CardID, tick int, operateArray []OperateType) any {
	//card 如果小于等于0 代表 不存在 需要返回null 客户端是识别null 会做处理
	var c any
	if card > 0 {
		c = card
	}
	return map[string]any{
		"type": GameTurnPush,
		"data": map[string]any{
			"chairID":      chairID,
			"card":         c,
			"tick":         tick,
			"operateArray": operateArray,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameChatPushData(chairID, t int, msg string, recipientID int) any {
	return map[string]any{
		"type": GameChatPush,
		"data": map[string]any{
			"chairID":     chairID,
			"type":        t,
			"stream":      msg,
			"recipientID": recipientID,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameTurnOperatePushData(chairID int, card mp.CardID, operate OperateType, success bool) any {
	var c any
	if card > 0 && card < 36 {
		c = card
	}
	return map[string]any{
		"type": GameTurnOperatePush,
		"data": map[string]any{
			"chairID": chairID,
			"card":    c,
			"operate": operate,
			"success": success,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameResultPushData(result *GameResult) any {
	return map[string]any{
		"type": GameResultPush,
		"data": map[string]any{
			"result": result,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameTrustPushData(chairID int, trust bool) any {
	return map[string]any{
		"type": GameTrustPush,
		"data": map[string]any{
			"chairID": chairID,
			"trust":   trust,
		},
		"pushRouter": "GameMessagePush",
	}
}
func GameReviewPushData(record []*ReviewRecord) any {
	return map[string]any{
		"type": GameReviewPush,
		"data": map[string]any{
			"reviewRecord": record,
		},
		"pushRouter": "GameMessagePush",
	}
}
