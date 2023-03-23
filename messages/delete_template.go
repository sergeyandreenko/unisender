package messages

import (
	"strconv"

	"github.com/sergeyandreenko/unisender/api"
)

// DeleteTemplateRequest request to delete a template.
type DeleteTemplateRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteTemplateRequest) Execute() (err error) {
	err = r.request.Execute("deleteTemplate", nil)
	return
}

// DeleteTemplate returns request to delete a template.
func DeleteTemplate(request *api.Request, templateID int64) *DeleteTemplateRequest {
	request.Add("template_id", strconv.FormatInt(templateID, 10))

	return &DeleteTemplateRequest{
		request: request,
	}
}
