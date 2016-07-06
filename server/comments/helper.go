package comments

import "github.com/asaskevich/govalidator"

func (c *Comments) validateAdd(input *AddInput) *AddOutput {
	if len(input.Content) < 4 || len(input.Content) > 200 {
		return &AddOutput{false, "Content needs to be between 4 and 200 characters", "Content"}
	}
	if !govalidator.IsUUID(input.ProjectID) {
		return &AddOutput{false, "Please Provide ProjectID", "ProjectID"}
	}
	item := c.portfolio.GetByProjectID(input.ProjectID)
	if !item.Valid {
		return &AddOutput{false, "Project not found", "ProjectID"}
	}
	return &AddOutput{true, "", ""}
}
