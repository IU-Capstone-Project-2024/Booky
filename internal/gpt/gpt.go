package gpt

type GPT interface {
	GetImprovedNote(note string) (string, error)
}
