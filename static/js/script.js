// Waits for the DOM to load.
document.addEventListener('DOMContentLoaded', function () {
  const themeToggleButton = document.getElementById('theme-toggle');
  const themeIcon = themeToggleButton.querySelector('.icon');
  // Checks localStorage for the user's saved theme preference and applies it.
  const currentTheme = localStorage.getItem('theme');
  const isDarkModePreferred = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;

  function updateThemeIcon() {
    if (document.documentElement.classList.contains('dark-theme')) {
      themeIcon.classList.remove('sun');
      themeIcon.classList.add('moon');
    } else {
      themeIcon.classList.remove('moon');
      themeIcon.classList.add('sun');
    }
  }

  if (currentTheme) {
    document.documentElement.classList.add(currentTheme);
  } else if (isDarkModePreferred) {
    document.documentElement.classList.add('dark-theme');
  }

  updateThemeIcon();

  // Adds an event listener to the button to toggle the dark-theme class on the :root element.
  themeToggleButton.addEventListener('click', function () {
    if (document.documentElement.classList.contains('dark-theme')) {
      document.documentElement.classList.remove('dark-theme');
      document.documentElement.classList.add('light-theme');
      // Saves the user's current theme preference to localStorage whenever the button is clicked.
      localStorage.setItem('theme', 'light-theme');
    } else {
      document.documentElement.classList.remove('light-theme');
      document.documentElement.classList.add('dark-theme');
      localStorage.setItem('theme', 'dark-theme');
    }
    updateThemeIcon();
  });
});

