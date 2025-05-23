package handlers

import (
	"context"
	"digital-garden/internal/templates/pages"
	"net/http"
)

// Home handles the homepage
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get theme from cookie
	theme := "dark" // default
	if cookie, err := r.Cookie("theme"); err == nil {
		theme = cookie.Value
	}

	component := pages.Home(theme)
	component.Render(context.Background(), w)
}

// ToggleTheme handles theme switching via HTMX
func ToggleTheme(w http.ResponseWriter, r *http.Request) {
	// Get current theme from cookie
	currentTheme := "dark" // default
	if cookie, err := r.Cookie("theme"); err == nil {
		currentTheme = cookie.Value
	}
	
	// Toggle theme
	newTheme := "light"
	if currentTheme == "light" {
		newTheme = "dark"
	}
	
	// Set new theme cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "theme",
		Value:    newTheme,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 365, // 1 year
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})
	
	// Return the theme toggle button with updated icon
	w.Header().Set("Content-Type", "text/html")
	
	if newTheme == "light" {
		w.Write([]byte(`<button 
			id="theme-toggle"
			class="p-2 text-secondaryText hover:text-cottonCandyPink transition-colors rounded"
			aria-label="Toggle theme"
			hx-post="/toggle-theme"
			hx-target="#theme-toggle"
			hx-swap="outerHTML"
			hx-trigger="click"
			_="on click add [@data-theme='light'] to <html/>"
		>
			<svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
				<path d="M12 3c-4.97 0-9 4.03-9 9s4.03 9 9 9c.83 0 1.5-.67 1.5-1.5 0-.39-.15-.74-.39-1.01-.23-.26-.38-.61-.38-.99 0-.83.67-1.5 1.5-1.5H16c2.76 0 5-2.24 5-5 0-4.42-4.03-8-9-8z"/>
			</svg>
		</button>`))
	} else {
		w.Write([]byte(`<button 
			id="theme-toggle"
			class="p-2 text-secondaryText hover:text-cottonCandyPink transition-colors rounded"
			aria-label="Toggle theme"
			hx-post="/toggle-theme"
			hx-target="#theme-toggle"
			hx-swap="outerHTML"
			hx-trigger="click"
			_="on click remove [@data-theme] from <html/>"
		>
			<svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
				<path d="M12 7c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58c-.39-.39-1.03-.39-1.41 0-.39.39-.39 1.03 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.41L5.99 4.58zm12.37 12.37c-.39-.39-1.03-.39-1.41 0-.39.39-.39 1.03 0 1.41l1.06 1.06c.39.39 1.03.39 1.41 0 .39-.39.39-1.03 0-1.41l-1.06-1.06zm1.06-10.96c.39-.39.39-1.03 0-1.41-.39-.39-1.03-.39-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06zM7.05 18.36c.39-.39.39-1.03 0-1.41-.39-.39-1.03-.39-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.41s1.03.39 1.41 0l1.06-1.06z"/>
			</svg>
		</button>`))
	}
}
