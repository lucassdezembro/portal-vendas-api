package enum

type LoginKeyType string

const (
	LOGIN_KEY_TYPE_DOCUMENT LoginKeyType = "DOCUMENT"
	LOGIN_KEY_TYPE_EMAIL    LoginKeyType = "EMAIL"
)

func (lkt LoginKeyType) String() string {
	return string(lkt)
}
