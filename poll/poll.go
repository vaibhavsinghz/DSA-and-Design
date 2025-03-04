package main

import (
	"fmt"
	"time"
)

// PrintPollResults prints the results of a poll
func (ps *PollService) PrintPollResults(pollID string) error {
	poll, err := ps.GetPoll(pollID)
	if err != nil {
		return err
	}

	fmt.Printf("Poll: %s - %s\n", poll.Title, poll.Description)
	fmt.Printf("Created by: %s\n", poll.Creator)
	fmt.Printf("Created at: %s\n", poll.CreatedAt.Format(time.RFC1123))
	fmt.Printf("Expires at: %s\n", poll.ExpiresAt.Format(time.RFC1123))
	fmt.Println("------------------------------")

	for _, question := range poll.Questions {
		fmt.Printf("Question %d: %s\n", question.ID, question.Text)

		totalVotes := 0
		for _, option := range question.Options {
			totalVotes += option.Votes
		}

		for _, option := range question.Options {
			percentage := 0.0
			if totalVotes > 0 {
				percentage = float64(option.Votes) / float64(totalVotes) * 100
			}
			fmt.Printf("  - Option %d: %s - %d votes (%.2f%%)\n", option.ID, option.Text, option.Votes, percentage)
		}
		fmt.Println()
	}

	return nil
}

// Main function with example usage
func main() {
	// Create a new poll service
	pollService := NewPollService()

	// Example: Create a poll
	favoriteColorPoll, err := pollService.CreatePoll(
		"poll1",
		"Favorite Colors",
		"What are your favorite colors?",
		"admin",
		[]Question{
			{
				Text: "What is your favorite primary color?",
				Options: []Option{
					{Text: "Red"},
					{Text: "Blue"},
					{Text: "Yellow"},
				},
				MultipleChoice: false,
			},
			{
				Text: "Which colors do you like for your room? (Select all that apply)",
				Options: []Option{
					{Text: "White"},
					{Text: "Black"},
					{Text: "Green"},
					{Text: "Purple"},
				},
				MultipleChoice: true,
			},
		},
		24*time.Hour, // Poll expires in 24 hours
	)

	if err != nil {
		fmt.Printf("Error creating poll: %v\n", err)
		return
	}

	fmt.Println("Poll created successfully!")

	// Example: Vote on the poll
	err = pollService.VoteOnPoll("poll1", "user1", favoriteColorPoll.Questions[0].ID, 2) // User1 votes for Blue
	if err != nil {
		fmt.Printf("Error voting on poll: %v\n", err)
	}

	err = pollService.VoteOnPoll("poll1", "user2", favoriteColorPoll.Questions[0].ID, 1) // User2 votes for Red
	if err != nil {
		fmt.Printf("Error voting on poll: %v\n", err)
	}

	err = pollService.VoteOnPoll("poll1", "user3", favoriteColorPoll.Questions[0].ID, 2) // User3 votes for Blue
	if err != nil {
		fmt.Printf("Error voting on poll: %v\n", err)
	}

	// Multiple choice question voting
	err = pollService.VoteOnPoll("poll1", "user1", favoriteColorPoll.Questions[1].ID, 1) // User1 votes for White
	if err != nil {
		fmt.Printf("Error voting on poll: %v\n", err)
	}

	err = pollService.VoteOnPoll("poll1", "user1", favoriteColorPoll.Questions[1].ID, 3) // User1 also votes for Green
	if err != nil {
		fmt.Printf("Error voting on poll: %v\n", err)
	}

	// Example: Update a poll question
	err = pollService.UpdatePollQuestion(
		"poll1",
		favoriteColorPoll.Questions[0].ID,
		"What is your absolute favorite primary color?",
		[]Option{
			{ID: 1, Text: "Red"},
			{ID: 2, Text: "Blue"},
			{ID: 3, Text: "Yellow"},
			{ID: 4, Text: "Green (not actually primary)"},
		},
		false,
	)
	if err != nil {
		fmt.Printf("Error updating poll question: %v\n", err)
	}

	// Example: Print poll results
	fmt.Println("\nPoll Results:")
	err = pollService.PrintPollResults("poll1")
	if err != nil {
		fmt.Printf("Error printing poll results: %v\n", err)
	}

	// Example: Delete a poll
	// err = pollService.DeletePoll("poll1")
	// if err != nil {
	//     fmt.Printf("Error deleting poll: %v\n", err)
	// } else {
	//     fmt.Println("Poll deleted successfully!")
	// }
}
