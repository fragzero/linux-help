# linux-help

A command-line interface (CLI) tool designed to assist Linux users in recalling and utilizing commands more effectively.  
`linux-help` leverages AI to understand your needs and provides relevant commands, which you can then easily paste into your terminal.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

# Installation

```sh
go build -ldflags "-w" -o linux-help
cp linux-help ~/.local/bin/
```

# Usage

## Get help with a specific task (example)
`linux-help "find files modified today"`

## Another example:
`linux-help "list all files in a directory"`

The AI will provide suggested commands. You can then choose to copy and paste the desired command in the terminal.

## Gemini key

To generate your Gemini API Key => https://ai.google.dev/gemini-api/docs/api-key

```sh
❯ export GEMINI_API_KEY=[your key]

❯ linux-help "to list 10 process consume most cpu and memory"
ps aux --sort=-%mem,-%cpu | head -n 11 | tail -n 10

❯ [SHIFT+CTRL+V] to paste the command
```

# Features

- AI-Powered Command Suggestions: Uses AI to understand your natural language queries and suggest relevant Linux commands.
- Easy Command Pasting: Streamlines the process of using the suggested command in your terminal. (Explain how this works - e.g., copying to clipboard, printing the command for manual copy, or direct execution).
- Command History (Optional): (If implemented) Keeps track of your previous queries and suggested commands for quick access.
- Extensible (Optional): (If applicable) Allows users to add custom commands or improve the AI's knowledge base.
- Cross-Platform (Linux): Designed specifically for Linux environments. (Mention any dependencies or specific Linux distributions if needed)

# Contributing

Contributions are welcome!
Please feel free to submit pull requests, bug reports, or feature suggestions.

- Fork the repository.
- Create a new branch for your feature or bug fix.
- Make your changes.
- Commit your changes and push the branch.   
- Submit a pull request.

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# Contact

- Author: FragZero
- Email: fragzero@fragzero.com.br
