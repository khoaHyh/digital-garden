(function () {
  "use strict";

  function getInitialTheme() {
    const savedTheme = localStorage.getItem("theme");
    if (savedTheme != null) {
      return savedTheme;
    }

    return window.matchMedia("(prefers-color-scheme: light)").matches
      ? "light"
      : "dark";
  }

  function applyTheme(theme) {
    const root = document.documentElement;
    root.setAttribute("data-theme", theme);
  }

  function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute("data-theme");
    const newTheme = currentTheme === "light" ? "dark" : "light";

    applyTheme(newTheme);
    localStorage.setItem("theme", newTheme);
  }

  function initTheme() {
    const theme = getInitialTheme();
    applyTheme(theme);

    // Set up theme toggle button
    const toggleButton = document.getElementById("theme-toggle");
    if (toggleButton != null) {
      toggleButton.addEventListener("click", toggleTheme);
    }
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initTheme);
  } else {
    initTheme();
  }
})();
