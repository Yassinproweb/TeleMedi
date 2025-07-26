/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html",
    "./views/partials/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        w: '#fefeff',
        f: '#f5f5f9',
        d: '#1a1b23',
        t: '#7d8389',
        l: '#98CCFF',
        b: '#387CFF',
      },
    },
  },
  plugins: [],
}

