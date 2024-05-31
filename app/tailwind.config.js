/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {},
    daisyui: {
        themes: ["cupcake"],
    },
    plugins: [require("daisyui")],
};
