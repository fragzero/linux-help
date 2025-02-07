# linux-help

Command line interface to help linux users to remember the commands using AI and paste the command. 

## Example

```sh
❯ export GEMINI_API_KEY=[your key]

❯ linux-help "to list 10 process consume most cpu and memory"
ps aux --sort=-%mem,-%cpu | head -n 11 | tail -n 10

❯ [SHIFT+CTRL+V] to paste the command
```

## Gemini key

To generate your Gemini API Key => https://ai.google.dev/gemini-api/docs/api-key

## Build

```sh
go build -ldflags "-w" -o linux-help
cp linux-help ~/.local/bin/
```
