package data

import "time"

// Context represents the root of our data file (the "Sprint")
type Context struct {
	ProjectName      string    `yaml:"project_name"`
	CurrentIteration string    `yaml:"current_iteration"`
	Stories          []Story   `yaml:"stories"`
	ActiveStoryID    string    `yaml:"active_story_id"`
	ActiveTaskName   string    `yaml:"active_task_name"`
	LastActivity     time.Time `yaml:"last_activity"`
}

type Story struct {
	ID    string `yaml:"id"`
	Name  string `yaml:"name"`
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name         string  `yaml:"name"`
	EstimateMins int     `yaml:"estimate_mins"`
	Completed    bool    `yaml:"completed"`
	Entries      []Entry `yaml:"entries"`
}

type Entry struct {
	Timestamp    time.Time `yaml:"timestamp"`
	DurationMins int       `yaml:"duration_mins"`
	Note         string    `yaml:"note"`
	NeedsReview  bool      `yaml:"needs_review"` // The flag for "Audit Mode"
}
