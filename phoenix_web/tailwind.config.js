/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        cyber: {
          bg: "#080c14",
          panel: "#0f1626",
          border: "#1E293B",
          accent: "#06B6D4", // Cyber Cyan
          danger: "#EF4444", // Cyber Red
          warning: "#F59E0B", // Cyber Amber
          success: "#10B981", // Cyber Emerald
          text: "#F8FAFC",
          muted: "#94A3B8"
        }
      },
      fontFamily: {
        mono: ['JetBrains Mono', 'Fira Code', 'monospace'],
        sans: ['Inter', 'system-ui', 'sans-serif'],
      }
    },
  },
  plugins: [],
}
