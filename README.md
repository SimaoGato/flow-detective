# Flow Detective 🕵️‍♂️

> **Stop breaking your flow.** Automate your time tracking directly from your Git workflow.

Flow Detective is a CLI sidecar that "listens" to your development activity. It uses Git hooks to automatically log time against your tasks, visualizes your budget with a TUI, and generates "Copy-Paste" reports for your timesheet platform.

![Go Version](https://img.shields.io/badge/go-1.25-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Build Status](https://github.com/SimaoGato/flow-detective/actions/workflows/ci.yaml/badge.svg)

## 🚀 Features

* **🧠 Context Aware:** Knows what Story/Task you are working on.
* **👻 Passive Tracking:** Logs time automatically when you run `git commit`.
* **❤️ Health Bar UI:** Visualizes budget vs. actual time using a beautiful terminal interface.
* **📋 Instant Reporting:** Generates markdown-formatted reports for standups or timesheets.
* **🛡️ Idle Detection:** (Coming Soon) Flags entries if you were idle for > 4 hours.

## 📦 Installation

### From Source
Prerequisites: Go 1.22+

```bash
# 1. Clone the repo
git clone [https://github.com/SimaoGato/flow-detective.git](https://github.com/SimaoGato/flow-detective.git)
cd flow-detective

# 2. Build the binary
make build

# 3. (Optional) Move to your path
sudo mv bin/flow /usr/local/bin/
```

## 🕵️ Usage Guide

1. Initialize a Project
Go to your project's root directory and create the local context.
```
flow init
```
Creates .flow/context.yaml (Add this folder to your .gitignore!)

2. Start a Task
Tell the detective what you are working on. This sets the "Short-Term Memory".
```
flow start "Refactor Login API"
```

3. The "Magic" (Work & Commit)
Just do your work. When you commit, Flow Detective calculates the time since your last activity and logs it.
```
git add .
git commit -m "feat: login logic"
# Output: 🕵️  Logged 45m to 'Refactor Login API'
```
*Important*: To enable this, install the hook once per repo:
```
flow hook
```

4. Check Status
See if you are burning through your budget.
```
flow status
```

5. End of Day Report
Generate your daily summary.
```
flow report
```

## 🛠️ Configuration (.flow/context.yaml)
The data is stored in a human-readable YAML file. You can manually edit estimates or fix time logs here.
```
project_name: "Revolution Backend"
current_iteration: "Sprint 42"
stories:
  - id: "REV-101"
    name: "Core Infrastructure"
    tasks:
      - name: "Implement Middleware"
        estimate_mins: 120
        entries:
          - duration_mins: 45
            note: "Git Commit"
```

## 💻 Development
We use make to automate standard tasks.
```
make audit   # Runs fmt, vet, and tests (CI equivalent)
make build   # Compiles the binary to ./bin/flow
make run     # Builds and runs the app
```
