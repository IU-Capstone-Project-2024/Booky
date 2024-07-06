package gpt

type AiModel interface {
	GetImprovedNote(note string) (string, error)
}
