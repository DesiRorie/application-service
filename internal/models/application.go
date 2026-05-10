package models

type Application struct {
	Company  string `json:"company"`
	Position string `json:"position"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
	ID       int    `json:"id"`
}

var Applications = []Application{

	{
		ID:       1,
		Company:  "Bank of America",
		Position: "QA Automation Engineer",
		Status:   "Interview Scheduled",
		Notes:    "Technical interview on Tuesday covering API testing and Go basics",
	},
	{
		ID:       2,
		Company:  "Red Ventures",
		Position: "Software Engineer I",
		Status:   "Applied",
		Notes:    "Applied through company careers page",
	},
	{
		ID:       3,
		Company:  "Lowe's",
		Position: "Backend Go Developer",
		Status:   "Rejected",
		Notes:    "Position required 5+ years of Go experience",
	},
	{
		ID:       4,
		Company:  "Ally",
		Position: "SDET",
		Status:   "Assessment Completed",
		Notes:    "Finished HackerRank assessment",
	},
	{
		ID:       5,
		Company:  "Microsoft",
		Position: "Software Engineer",
		Status:   "Recruiter Screen",
		Notes:    "Recruiter reached out on LinkedIn",
	},
}
