package game

// CheckWin evaluates whether a card has won under the given winning rule.
// winningRule examples: "line", "fullhouse", "fourcorners".
// Kept separate from handlers so win logic is reusable and easy to unit-test.
func CheckWin(card *Card, winningRule string) bool {
	// TODO: implement winner detection for each winning rule
	return false
}

// MarkCell marks a cell on the card that matches the given song ID.
func MarkCell(card *Card, songID int) error {
	// TODO: find matching cell and set Marked = true
	return nil
}
