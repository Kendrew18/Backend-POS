package request

type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type To struct {
	Email string `json:"email"`
}

type TemplateVariables struct {
	VerificationCode string `json:"verification_code"`
	UserName         string `json:"user_name"`
	AppsName         string `json:"apps_name"`
}

type Payload struct {
	From              From              `json:"from"`
	To                []To              `json:"to"`
	TemplateUUID      string            `json:"template_uuid"`
	TemplateVariables TemplateVariables `json:"template_variables"`
}
