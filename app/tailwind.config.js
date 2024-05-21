/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {},
    daisyui: {
        themes: ["cmyk"],
    },
    plugins: [require("daisyui")],
};
