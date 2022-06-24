package actions

type GetHttpSecurityHeadersData struct {
	Domain string
}

func (data GetHttpSecurityHeadersData) validatedUrl() bool {
	if data.Domain == "" {
		return false
	}

	return true
}
