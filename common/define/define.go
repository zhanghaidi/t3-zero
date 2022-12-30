package define

// PageSize 分页的默认参数
var (
	PageSize            = 20
	MaxPageSize         = 200
	TokenExpire         = 3600 * 24 * 7
	HotNum              = 10
	MobileVerifyMinutes = 3
)

const (
	APP_URL        = "https://www.evdo.vip"
	OSS_URL        = "https://attachment.evdo.vip"
	EXAM_PATH      = "/exam/index?examCode="
	DB_AUTH_CODE   = "wfE5RYtwIzMiCJd4SD"
	JWT_SECRET_KEY = ""
)

const (
	DefaultGiveType  uint = iota
	RegisterGiveType      //新用户注册赠送
	SignGiveType          //每日签到赠送
	PayGiveType           //开通会员支付赠送
	InviteGiveType        //邀请用户注册赠送
	SystemGiveType        //系统赠送
)

const (
	Default      uint = iota
	SuperAdmin        //超级管理员
	SchoolAdmin       //院校管理员
	TeacherAdmin      //教师管理员
	ContentAdmin      //内容管理员
)

const (
	NoneSystemNotice uint = iota
	SystemNotice          //后台推送系统通知
	VipGiveNotice         //vip获取提醒
)

const (
	NoneUserNotice          uint = iota
	TopicCommentNotice           //话题评论通知
	TopicCommentReplyNotice      //话题评论回复通知
	TopicLikeNotice              //话题点赞通知
	TopicCommentLikeNotice       //话题评论点赞通知
)

const (
	NotJoin uint = iota
	WaitingJoin
	AlreadyJoin
)

const (
	MaxTopicContentWords = 1500 //话题内容字数限制
	MaxTopicCommentWords = 1500 //话题评论字数限制
	MaxTopicThumbs       = 2    //话题图片数量限制
	MaxTopicTitleWords   = 50   //话题标题字数限制
)
