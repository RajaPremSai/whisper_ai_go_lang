# Whisper_GO

Whisper_GO is a Go client for transcribing audio files using OpenAI's Whisper model.  
It supports transcription via the OpenAI API and can be adapted for local Whisper usage.

## Features

- Transcribe audio files (e.g., `.m4a`) using OpenAI Whisper API.
- Simple Go interface.
- Structured transcription output.

## Requirements

- Go 1.18+
- OpenAI API key (for API usage)
- Internet connection (for API usage)

## Usage

### 1. Transcribe with OpenAI API

1. **Set your OpenAI API key:**

   On Windows (PowerShell):

   ```
   $env:OPENAI_API="your_openai_api_key"
   ```

2. **Run the program:**

   ```
   go run main.go
   ```

   The transcription will be printed to the console.

### 2. Transcribe Locally (Free, No API Key)

1. **Install Whisper (Python):**

   ```
   pip install git+https://github.com/openai/whisper.git
   ```

2. **Transcribe your audio file:**

   ```
   whisper file.m4a --model base --output_format json
   ```

   This creates `file.json` with the transcription.

3. **Modify `main.go` to read the local JSON file:**

   Replace API client usage with code that reads and prints the transcription from `file.json`.

## Project Structure

```
Whisper_GO/
├── api/
│   └── whisper/
├── models/
├── transcribe/
├── main.go
```

- `api/whisper/`: OpenAI API client code
- `models/`: Transcription response models
- `transcribe/`: Transcription options/config
- `main.go`: Entry point

## Things Learned

- How to use the OpenAI Whisper API for audio transcription in Go.
- Setting up and using environment variables for API authentication.
- Running Whisper locally with Python for free transcription.
- Reading and parsing JSON output from local Whisper runs in Go.
- Adapting Go code to support both API-based and local transcription workflows.
- Understanding project structure and modular code


## Example Output

```
Transcription : Hello, this is a test audio file.
```

## License

This project is provided for educational purposes.  
Check OpenAI Whisper and API terms for usage restrictions.

## References

- [OpenAI Whisper](https://github.com/openai/whisper)
- [OpenAI API Docs](https://platform.openai.com/docs/api-reference/audio/create)

