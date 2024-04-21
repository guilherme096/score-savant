/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {
        extend: {
            colors: {
                "primary-darker": "#220151",
            },
        },
    },
    daisyui: {
        themes: [
            {
                dark: {
                    ...require("daisyui/src/theming/themes")["dark"],
                    primary: "#26005c",
                    secondary: "#9501ff",
                    "base-100": "#3c007b",
                },
            },
        ],
    },
    plugins: [require("daisyui")],
};
