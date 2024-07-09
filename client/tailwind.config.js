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
        250: "#EEEEEE",
        300: "#D9D9D9",
        400: "#737373",
        500: "#333333",
      },
      red: "#FF3939",
      success: "#22c55e",
    },
    fontFamily: {
      sans: ["Instrument Sans Variable", "sans-serif"],
    },
    extend: {
      boxShadow: {
        inputShadow: "0 0 32px 0 #633cff40",
        dropdownShadow: "0 0 32px #0000001a",
      },
      backgroundImage: {
        phone: "url(/images/Phone.png)",
      },
    },
  },
  plugins: [],
};
