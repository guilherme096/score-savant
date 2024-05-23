/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {},
    daisyui: {
        themes: ["emerald"],
    },
    plugins: [require("daisyui")],
};
