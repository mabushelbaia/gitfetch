package dashboard

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabushelbaia/gitfetch/internal/github"
)

// Theme holds all the color information for a theme.
type Theme struct {
	AccentColor    lipgloss.AdaptiveColor
	SecondaryColor lipgloss.AdaptiveColor
	HighlightColor lipgloss.AdaptiveColor
	TextColor      lipgloss.AdaptiveColor
	LinkColor      lipgloss.AdaptiveColor
}

func PrintDashboard(user *github.UserInfo) {
	// Terminal-based colors (ANSI)t ...
	// 1=red, 2=green, 3=yellow, 4=blue, 5=magenta, 6=cyan, 7=white, 8=bright variants
	theme := Theme{
		AccentColor:    lipgloss.AdaptiveColor{Light: "4", Dark: "4"}, // Blue
		SecondaryColor: lipgloss.AdaptiveColor{Light: "7", Dark: "7"}, // White/Grey
		HighlightColor: lipgloss.AdaptiveColor{Light: "3", Dark: "3"}, // Yellow
		TextColor:      lipgloss.AdaptiveColor{Light: "7", Dark: "7"}, // White
		LinkColor:      lipgloss.AdaptiveColor{Light: "6", Dark: "6"}, // Cyan
	}

	// --- Styles ---
	labelStyle := lipgloss.NewStyle().
		Foreground(theme.AccentColor).
		Bold(true).
		Padding(0, 1)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.SecondaryColor).
		Padding(0, 1).
		MarginRight(2)

	valueStyle := lipgloss.NewStyle().
		Foreground(theme.TextColor)

	bioStyle := lipgloss.NewStyle().
		Foreground(theme.SecondaryColor).
		Faint(true)

	urlStyle := lipgloss.NewStyle().
		Foreground(theme.LinkColor).
		Underline(true)

	headerStyle := lipgloss.NewStyle().
		Foreground(theme.HighlightColor).
		Bold(true)

	// --- Labels (with icons) ---
	labels := lipgloss.JoinVertical(lipgloss.Left,
		labelStyle.Render(" user"),
		labelStyle.Render(" bio"),
		"",
		labelStyle.Render(" Public Repos"), // TODO: Replace With Stars
		labelStyle.Render(" followers"),
		labelStyle.Render(" following"),
		"",                             // blank line
		labelStyle.Render(" company"), // briefcase
		labelStyle.Render(" country"),
		labelStyle.Render(" website"),
	)

	left := boxStyle.Render(labels)

	// --- Right column values ---
	rightValues := lipgloss.JoinVertical(lipgloss.Left,
		headerStyle.Render(fmt.Sprintf("%s (@%s)", user.Name, user.Login)),
		bioStyle.Render(user.Bio),
		"", // spacing line
		valueStyle.Render(fmt.Sprintf("%d", user.PublicRepos)),
		valueStyle.Render(fmt.Sprintf("%d", user.Followers)),
		valueStyle.Render(fmt.Sprintf("%d", user.Following)),
		"", // spacing line
		valueStyle.Render(user.Company),
		valueStyle.Render(user.Country),
		urlStyle.Render(user.Website),
	)

	// align with first label (skip border line)
	right := lipgloss.NewStyle().MarginTop(1).Render(rightValues)

	// --- Final Layout ---
	dashboard := lipgloss.JoinHorizontal(lipgloss.Top, left, right)

	fmt.Println(dashboard)
}
