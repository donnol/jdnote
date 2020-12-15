package notemodel

type Status int

const (
	StatusUnknown = 0
	StatusDraft   = 1
	StatusPublish = 2
)

func (status Status) IsValid() bool {
	switch status {
	case StatusDraft, StatusPublish:
		return true
	}
	return false
}

func (status Status) IsPublish() bool {
	return status == StatusPublish
}
