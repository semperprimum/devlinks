/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx,vue}"],
  theme: {
    screens: {
      sm: "375px",
      md: "768px",
      lg: "976px",
      xl: "1440px",
    },
    colors: {
      brand: {
        100: "#EFEBFF",
        200: "#BEADFF",
        300: "#633CFF",
      },
      neutral: {
        100: "#FFFFFF",
        200: "#FAFAFA",
        300: "#D9D9D9",
        400: "#737373",
        500: "#333333",
      },
      red: "#FF3939",
    },
    fontFamily: {
      sans: ["Instrument Sans Variable", "sans-serif"],
    },
  },
  plugins: [],
};
