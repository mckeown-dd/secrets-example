package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type secretsPayload struct {
	Secrets []string `json:secrets`
	Version int      `json:version`
}

func main() {
	data, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read from stdin: %s", err)
		os.Exit(1)
	}
	secrets := secretsPayload{}
	json.Unmarshal(data, &secrets)

	res := map[string]map[string]string{}
	for _, handle := range secrets.Secrets {
		res[handle] = map[string]string{
			"value": os.Getenv(handle),
		}
	}

	output, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not serialize res: %s", err)
		os.Exit(1)
	}
	fmt.Printf(string(output))
}
