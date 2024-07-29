package ecodes

// iam.auth 10.00.xx

const (
	IAM_USERNAME_NOT_FOUND                  = 100000 // 用户名不存在
	IAM_PASSWORD_ERROR                      = 100001 // 密码错误
	IAM_AUTHENTICATION_TOKEN_NOT_PROVIDED   = 100002 // 认证令牌未提供
	IAM_AUTHENTICATION_TOKEN_INVALID_FORMAT = 100003 // 无效的认证令牌格式：header 中的 Authorization 必须为 Bearer xxxxxxx
	IAM_AUTHENTICATION_TOKEN_INVALID        = 100004 // 无效的认证令牌：令牌格式正确但是内容无效或过期
	IAM_INVALID_USERNAME_LENGTH             = 100005 // 用户名长度不合法：需要在 1 到 30 之间
	IAM_INVALID_USERNAME_FORMAT             = 100006 // 用户名格式不合法：只能包含字母、数字、下划线
	IAM_USERNAME_ALREADY_EXISTS             = 100007 // 用户名已存在
	IAM_INVALID_EMAIL_FORMAT                = 100008 // 邮箱格式不合法
)
