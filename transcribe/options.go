package transcribe

type TranscribeConfig struct {
	Model    string
	Language string
	File     string
}

type TranscribeOption func(*TranscribeConfig)

func WithModel(model string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.Model = model
	}
}

func WithLanguage(language string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.Language = language
	}
}

func WithFile(file string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.File = file
	}
}