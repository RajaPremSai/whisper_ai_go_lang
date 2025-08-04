package main

import (
	"fmt"
	"log"
	"os"
	"whisper_go/api/whisper"
)

func main(){
	client := whisper.NewClient(whisper.WithKey(os.Getenv("OPENAI_API")))

	response,err:=client.TranscribeFile("file.m4a")

	if err!=nil{
		log.Fatalf("error transcribing file %v\n",err)
	}

	fmt.Printf("Transcription : %s\n",response.Text)
}