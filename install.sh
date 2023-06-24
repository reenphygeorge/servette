#!/bin/sh

# Determine the default shell
SHELL_NAME=$(basename "$SHELL")
SERVETTE_PATH=/usr/local/bin/
ALIAS_NAME="srt"

cp "$(pwd)/srt" "$SERVETTE_PATH/srt"

if [ "$SHELL_NAME" = "bash" ] || [ "$SHELL_NAME" = "sh" ]; then
  echo 'export PATH="$SHELL_NAME:$SERVETTE_PATH"' >> ~/.bashrc
  echo 'export PATH="$SHELL_NAME:$SERVETTE_PATH"' >> ~/.profile
elif [ "$SHELL_NAME" = "zsh" ]; then
  echo 'export PATH="$SHELL_NAME:$SERVETTE_PATH"' >> ~/.zshrc
elif [ "$SHELL_NAME" = "fish" ]; then
  echo 'set -x PATH "$SHELL_NAME" $SERVETTE_PATH' >> ~/.config/fish/config.fish
else
  echo "Unsupported shell: $SHELL_NAME"
  exit 1
fi

echo "Installation completed!"
