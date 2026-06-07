// Package spinner holds the animation frames for an in-progress indicator,
// shared by the sequential and concurrent presenters so they stay in sync.
package spinner

// Frames are double-width so the spinner stays aligned with the double-width
// ✅/❌/⏩ status icons.
var Frames = []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}

// Frame returns the i-th animation frame, wrapping around.
func Frame(i int) string {
	return Frames[i%len(Frames)]
}
