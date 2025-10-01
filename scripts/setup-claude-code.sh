#!/bin/bash
set -e

echo "🚀 Setting up Hatch for Claude Code..."

# 1. Move agents directory
if [ -d "docs/agents" ]; then
    echo "📁 Moving docs/agents → .claude/"
    mv docs/agents .claude
else
    echo "⚠️  docs/agents not found, skipping..."
fi

# 2. Rename AGENTS.md → CLAUDE.md
if [ -f "AGENTS.md" ]; then
    echo "📝 Renaming AGENTS.md → CLAUDE.md"
    mv AGENTS.md CLAUDE.md
else
    echo "⚠️  AGENTS.md not found, skipping..."
fi

# 3. Update links in CLAUDE.md
if [ -f "CLAUDE.md" ]; then
    echo "🔗 Updating links in CLAUDE.md"
    sed -i.bak 's|docs/agents/|.claude/|g' CLAUDE.md
    sed -i.bak 's|APPLICATION.MD|docs/application.md|g' CLAUDE.md
    sed -i.bak 's|LAYOUT.MD|docs/layout.md|g' CLAUDE.md
    rm CLAUDE.md.bak
fi

# 4. Update links in agent files
if [ -d ".claude" ]; then
    echo "🔗 Updating links in agent files"

    for file in .claude/*.md; do
        if [ -f "$file" ]; then
            # Update app references (../../app/ → ../app/)
            sed -i.bak 's|../../app/|../app/|g' "$file"

            # Update docs references (../application.md → docs/application.md)
            sed -i.bak 's|\.\./application\.md|docs/application.md|g' "$file"
            sed -i.bak 's|\.\./layout\.md|docs/layout.md|g' "$file"

            # Remove backup files
            rm "${file}.bak"
        fi
    done
fi

echo "✅ Claude Code setup complete!"
echo ""
echo "Structure:"
echo "  CLAUDE.md         # AI agent context"
echo "  .claude/          # Agent commands"
echo "  docs/             # Documentation"
echo ""
echo "Next steps:"
echo "  1. Update HATCH_APP in your codebase"
echo "  2. Configure .env file"
echo "  3. Start coding with Claude Code!"
