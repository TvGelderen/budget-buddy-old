/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')

module.exports = {
    content: [
        './view/**/*.templ',
        './handler/*.go',
    ],
    theme: {
        colors: {
            ...colors,
            'primary': '#121212',
            'secondary': '#181818',
            'tertiary': '#202020',
            'theme': '#ff7700',
            'theme-secondary': '#ffaa44',
            'theme-tertiary': '#ffaa00',
            'error': 'rgb(185 28 28)',
            'error-background': 'rgb(255 232 234)',
            'success': 'rgb(22 163 74)',
            'success-background': 'rgb(230 255 230)',
            'white': '#ffffff',
            'black': '#000000',
        },
        fontFamily: {
            sans: ['Montserrat', 'sans-serif'],
            serif: ['serif'],
        },
        extend: {},
    },
    plugins: [],
}

