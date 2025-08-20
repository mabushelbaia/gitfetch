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

// DefaultTheme is a good starting point.
var DefaultTheme = Theme{
	AccentColor: lipgloss.AdaptiveColor{
		Light: "#c19beb", // light mode purple-ish
		Dark:  "#a374d0", // dark mode purple-ish
	},
	SecondaryColor: lipgloss.AdaptiveColor{
		Light: "#B3B3B3", // light mode grey
		Dark:  "#666666", // dark mode grey
	},
	HighlightColor: lipgloss.AdaptiveColor{
		Light: "#FFCC00", // yellow/orange
		Dark:  "#e6b800", // slightly darker yellow
	},
	TextColor: lipgloss.AdaptiveColor{
		Light: "#FFFFFF", // white
		Dark:  "#EEEEEE", // light grey for dark mode
	},
	LinkColor: lipgloss.AdaptiveColor{
		Light: "#43BF6D", // cyan/blue
		Dark:  "#73F59F", // darker cyan/blue
	},
}

// PrintDashboard prints user stats in a pfetch-style layout
func PrintDashboard(user *github.UserInfo, theme ...Theme) {
	// Use DefaultTheme if no theme is provided
	th := DefaultTheme
	if len(theme) > 0 {
		th = theme[0]
	}

	// Define styles using the provided theme.
	keyStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(th.AccentColor).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(th.TextColor).
		Align(lipgloss.Left)

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(th.HighlightColor)

	bioStyle := lipgloss.NewStyle().
		Foreground(th.SecondaryColor).
		Faint(true).
		Width(45).
		PaddingBottom(1)

	urlStyle := lipgloss.NewStyle().
		Foreground(th.LinkColor).
		Underline(true)

	// Helper function to create a formatted line if the value is not empty.
	createInfoLine := func(key, value string, style lipgloss.Style) string {
		if value == "" {
			return ""
		}
		return lipgloss.JoinHorizontal(lipgloss.Left,
			keyStyle.Render(key+":"),
			" "+style.Render(value),
		)
	}

	// Header: Name @login
	header := headerStyle.Render(fmt.Sprintf("%s @%s", user.Name, user.Login))

	// Bio
	var bio string
	if user.Bio != "" {
		bio = bioStyle.Render(user.Bio)
	}

	// Main Stats
	stats := lipgloss.JoinVertical(lipgloss.Left,
		createInfoLine("Followers", fmt.Sprintf("%d", user.Followers), valueStyle),
		createInfoLine("Following", fmt.Sprintf("%d", user.Following), valueStyle),
	)

	// Company, Location, Website
	info := lipgloss.JoinVertical(lipgloss.Left,
		createInfoLine("Company", user.Company, valueStyle),
		createInfoLine("Location", user.Country, valueStyle),
		createInfoLine("Website", user.Website, urlStyle),
	)

	rightSide := lipgloss.JoinVertical(lipgloss.Left, header, bio, stats, "", info)

	// Get your ASCII art here
	// asciiArt := getAsciiArt()

	// TODO: Combine ASCII art and stats
	dashboard := lipgloss.NewStyle().
		Padding(1, 0).
		Render(lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Render(rightSide),
		))

	fmt.Println(dashboard)
}
