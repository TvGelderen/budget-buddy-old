/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './view/**/*.templ'
    ],
    theme: {
        colors: {
            'primary': '#121212',
            'secondary': '#181818',
            'tertiary': '#202020',
            'theme': '#ff7700',
            'theme-secondary': '#ffaa44',
            'theme-tertiary': '#ffaa00',
        },
        fontFamily: {
            sans: ['Montserrat', 'sans-serif'],
            serif: ['serif'],
        },
        extend: {},
    },
    plugins: [],
}

