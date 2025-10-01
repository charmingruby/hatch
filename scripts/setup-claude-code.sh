#!/bin/bash
set -e

echo "Setting up Hatch for Claude Code..."

if [ -d "docs" ]; then
    rm -rf .claude
    cp -r docs .claude
else
    echo "⚠️  docs not found, skipping duplication..."
fi

if [ -f "AGENTS.md" ]; then
    mv AGENTS.md CLAUDE.md
else
    echo "⚠️  AGENTS.md not found, skipping..."
fi

echo "✅ Claude Code setup complete!"