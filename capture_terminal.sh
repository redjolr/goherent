#!/bin/sh
# capture_terminal.sh — render an escape-sequence experiment in a REAL terminal
# (via tmux) and dump the resulting grid, scrollback, and cursor position.
#
# This is the ground-truth oracle for making FakeAnsiTerminal behave like a real
# terminal. Run an experiment here, paste the output back, and it gets baked into
# a characterization test for the fake.
#
# Usage:
#   sh capture_terminal.sh WIDTH HEIGHT 'PRINTF_FORMAT'
#
# PRINTF_FORMAT is fed to printf(1), so use \n, \033, \r, etc.
#   - \033 = ESC               e.g. \033[0J   (erase cursor->end of screen)
#   - \033[<r>;<c>H            move cursor to row r, col c (1-based)
#   - \0337 / \0338            DECSC / DECRC  (save / restore cursor)
#   - DO NOT put a literal % in the format (printf treats it specially).
#
# Example:
#   sh capture_terminal.sh 6 3 'L1\nL2\nL3\nL4\nL5'
#
# Requires: tmux.

if [ "$#" -ne 3 ]; then
	echo "usage: sh capture_terminal.sh WIDTH HEIGHT 'PRINTF_FORMAT'" >&2
	exit 2
fi

w=$1
h=$2
fmt=$3

if ! command -v tmux >/dev/null 2>&1; then
	echo "tmux is required:  apk add tmux   (Alpine)  /  apt install tmux" >&2
	exit 1
fi

seqfile=$(mktemp)
# Materialize the exact bytes locally so we don't fight shell quoting inside tmux.
printf "$fmt" >"$seqfile"

tmux kill-session -t cap 2>/dev/null || true
tmux new-session -d -s cap -x "$w" -y "$h"
# Drop the status line so the pane is exactly H rows (not H-1).
tmux set-option -t cap status off

# clear -> known blank screen with cursor at home; cat -> our bytes; then keep
# the pane busy (sleep) so the shell prompt doesn't return and pollute the grid.
tmux send-keys -t cap "clear; cat '$seqfile'; tmux wait-for -S rendered; sleep 30" Enter
tmux wait-for rendered

echo "===== ${w}x${h} | format: ${fmt} ====="
echo "--- VISIBLE GRID (a trailing | marks end of each captured line) ---"
tmux capture-pane -t cap -p | cat -n | sed 's/$/|/'
echo "--- WITH SCROLLBACK (up to 8 lines above the screen included) ---"
tmux capture-pane -t cap -p -S -8 | cat -n | sed 's/$/|/'
echo "--- CURSOR (0-based) ---"
tmux display-message -p -t cap 'x=#{cursor_x} y=#{cursor_y}'
echo "==========================================="

tmux kill-session -t cap 2>/dev/null || true
rm -f "$seqfile"
