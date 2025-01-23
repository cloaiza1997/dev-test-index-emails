/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    screens: {
      sm: '640px',
      md: '768px',
      lg: '1024px',
      xl: '1280px',
      '2xl': '1440px',
    },
    extend: {
      colors: {
        primary: {
          100: '#bdd0fb',
          200: '#b0c7fa',
          300: '#97b6f9',
          400: '#85a9f8',
          500: '#729cf7',
          600: '#608ff6',
          700: '#4d82f5',
          800: '#3b75f3',
          900: '#2867f2',
        },
        secondary: {
          100: '#0e51e6',
          200: '#0c4bd3',
          300: '#0b44c1',
          400: '#09379c',
          500: '#083089',
          600: '#072a77',
          700: '#062364',
          800: '#04163f',
          900: '#01022e',
        },
        background: '#e8efff',
      },
    },
  },
  plugins: [],
}
