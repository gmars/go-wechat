package message

import (
	"encoding/xml"
	"fmt"
	"go-wechat/core"
)

// 开放平台info类消息解析器
func (m *Message) parseInfo(baseInfo *TypeParse, data []byte) (*ReplyMsg, error) {
	var (
		replay = ReplyMsg{
			ReplayType: ReplyTypeInfo,
			TypeValue:  baseInfo.InfoType,
		}
		err     error
		resType interface{}
	)

	switch baseInfo.InfoType {
	case InfoTypeAuthorized:
		resType = new(AuthorizedInfo)
		break
	case InfoTypeComponentVerifyTicket:
		resType = new(ComponentVerifyTicketInfo)
		break
	case InfoTypeUnauthorized:
		resType = new(UnauthorizedInfo)
		break
	case InfoTypeUpdateAuthorized:
		resType = new(UpdateAuthorizedInfo)
		break
	case InfoTypeNotifyThirdFasterRegister:
		resType = new(NotifyThirdFasterRegisterInfo)
		break
	case InfoTypeNotifyThirdFastVerifyBetaApp:
		resType = new(NotifyThirdFastVerifyBetaAppInfo)
		break
	default:
		return nil, core.NewError(404, fmt.Sprintf("消息类型%s的解析器还未添加", baseInfo.MsgType))
	}

	if err = xml.Unmarshal(data, &resType); err != nil {
		return nil, err
	}
	replay.Data = resType
	return &replay, nil
}

// 解析事件类型
func (m *Message) parseMsg(baseInfo *TypeParse, data []byte) (*ReplyMsg, error) {
	var (
		replay = ReplyMsg{
			ReplayType: ReplyTypeStandard,
			TypeValue:  baseInfo.MsgType,
		}
		err     error
		resType interface{}
	)

	switch baseInfo.MsgType {
	case MsgText:
		resType = new(TextMsg)
		break
	case MsgImage:
		resType = new(ImageMsg)
		break
	case MsgVoice:
		resType = new(VoiceMsg)
		break
	case MsgVideo:
		resType = new(VideoMsg)
		break
	case MsgShortVideo:
		resType = new(ShortVideoMsg)
		break
	case MsgLocation:
		resType = new(LocationMsg)
		break
	case MsgLink:
		resType = new(LinkMsg)
		break
	default:
		return nil, core.NewError(404, fmt.Sprintf("消息类型%s的解析器还未添加", baseInfo.MsgType))
	}

	if err = xml.Unmarshal(data, &resType); err != nil {
		return nil, err
	}
	replay.Data = resType
	return &replay, nil
}

// 解析事件类型
func (m *Message) parseEvent(baseInfo *TypeParse, data []byte) (*ReplyMsg, error) {
	var (
		replay = ReplyMsg{
			ReplayType: ReplyTypeEvent,
			TypeValue:  baseInfo.Event,
		}
		err     error
		resType interface{}
	)
	if baseInfo.MsgType != "event" {
		return nil, core.NewError(500, fmt.Sprintf("event处理器无法处理类型为%s的消息", baseInfo.MsgType))
	}
	switch baseInfo.Event {
	case EventSubscribe:
		resType = new(SubscribeEvent)
		break
	case EventScan:
		resType = new(ScanEvent)
		break
	case EventLocation:
		resType = new(LocationEvent)
		break
	case EventClick:
		resType = new(ClickEvent)
		break
	case EventView:
		resType = new(ViewEvent)
		break
	case EventScanCodePush:
		resType = new(ScanCodePushEvent)
		break
	case EventScanCodeWaitMsg:
		resType = new(ScanCodeWaitMsgEvent)
		break
	case EventPicSysPhoto:
		resType = new(PicSysPhotoEvent)
		break
	case EventPicPhotoOrAlbum:
		resType = new(PicPhotoOrAlbumEvent)
		break
	case EventPicWeixin:
		resType = new(PicWeixinEvent)
		break
	case EventLocationSelect:
		resType = new(LocationSelectEvent)
		break
	case EventViewMiniProgram:
		resType = new(ViewMiniProgramEvent)
		break
	case EventTemplateSendJobFinish:
		resType = new(TemplateSendJobFinishEvent)
		break
	case EventMassSendJobFinish:
		resType = new(MassSendJobFinishEvent)
		break
	case EventSubscribeMsgPopup:
		resType = new(SubscribeMsgPopupEvent)
		break
	case EventSubscribeMsgChange:
		resType = new(SubscribeMsgChangeEvent)
		break
	case EventSubscribeMsgSent:
		resType = new(SubscribeMsgSendEvent)
		break
	case EventPublishJobFinish:
		resType = new(PublishJobFinishEvent)
		break
	case EventQualificationVerifySuccess:
		resType = new(QualificationVerifySuccessEvent)
		break
	case EventQualificationVerifyFail:
		resType = new(QualificationVerifyFailEvent)
		break
	case EventNamingVerifySuccess:
		resType = new(NamingVerifySuccessEvent)
		break
	case EventNamingVerifyFail:
		resType = new(NamingVerifyFailEvent)
		break
	case EventAnnualRenew:
		resType = new(AnnualRenewEvent)
		break
	case EventVerifyExpired:
		resType = new(VerifyExpiredEvent)
		break
	case EventPoiCheckNotify:
		resType = new(PoiCheckNotifyEvent)
		break
	case EventUserAuthorizeInvoice:
		resType = new(UserAuthorizeInvoiceEvent)
		break
	case EventUpdateInvoiceStatus:
		resType = new(UpdateInvoiceStatusEvent)
		break
	case EventSubmitInvoiceTitle:
		resType = new(SubmitInvoiceTitleEvent)
		break
	case EventUserInfoModified:
		resType = new(UserInfoModifiedEvent)
		break
	default:
		return nil, core.NewError(404, fmt.Sprintf("事件%s的解析器还未添加", baseInfo.Event))
	}

	if err = xml.Unmarshal(data, &resType); err != nil {
		return nil, err
	}
	replay.Data = resType
	return &replay, nil
}
