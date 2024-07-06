package gpt

type GPT struct {
	AiModel AiModel
}

func (g *GPT) GetImprovedNote(note string) (string, error) {
	return g.AiModel.GetImprovedNote(note)
}
