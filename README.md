# tsk

A small task tracker written in Go, built as a playground for learning the Claude Code harness:
context (CLAUDE.md), hooks and guardrails, skills, subagents, MCP, plugins, and Claude working in GitHub.

## Build & run

```bash
go build -o tsk .
./tsk version
```

## Roadmap

| Days  | Build                                     | Harness concept                         |
|-------|-------------------------------------------|-----------------------------------------|
| 1     | Repo, Cobra CLI skeleton                  | Baseline: no config                     |
| 2–3   | SQLite storage, `add` / `list` / `done`   | CLAUDE.md, `/context`, memory           |
| 4–5   | Tests, golangci-lint                      | Hooks and guardrails                    |
| 6–7   | HTTP API                                  | Skills                                  |
| 8–9   | Refactors, features                       | Subagents                               |
| 10–11 | MCP server                                | Tool design for agents                  |
| 12–14 | `@claude` on issues and PRs, auto-review  | Headless mode / GitHub Actions          |
| 15    | Bundle as a plugin, write-up              | Plugins / skill packs                   |

See [LEARNINGS.md](LEARNINGS.md) for the journal.
