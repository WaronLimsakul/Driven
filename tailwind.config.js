/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: ["internal/templates/*.templ", "internal/templates/*.go"],
  theme: {
    extend: {
      colors: {
        flame: {
          bg: "#0A0A0A",
          surface: "#1A1A1A",
          primary: "#00CCFF",
          hover: "#33D6FF",
          dark: "#0099CC",
          accent: "#00FFFF",
          text: "#E0F7FF",
          muted: "#7A9EAA",
        },
      },
    },
  },
  plugins: [],
};
