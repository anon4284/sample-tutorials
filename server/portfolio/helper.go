package portfolio

func validateAdd(input *AddInput) *AddOutput {
	if len(input.Content) < 100 || len(input.Content) > 4000 {
		return &AddOutput{false, "Content needs to be between 100 and 4000 characters", "Content"}
	}
	if len(input.Title) < 4 || len(input.Title) > 50 {
		return &AddOutput{false, "Title needs to be between 4 and 50 characters", "Title"}
	}
	if len(input.Description) < 10 || len(input.Description) > 200 {
		return &AddOutput{false, "Description needs to be between 10 and 200 characters", "Description"}
	}
	return &AddOutput{true, "", ""}
}
