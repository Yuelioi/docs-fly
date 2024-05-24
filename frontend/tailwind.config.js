import plugin from 'tailwindcss/plugin'

/** @type {import('tailwindcss').Config} */

export default {
    darkMode: 'selector',
    content: ['./index.html', './src/**/*.{js,ts,vue}'],
    theme: {
        extend: {
            screens: {
                sm: '576px',
                // => @media (min-width: 640px) { ... }

                md: '960px',
                // => @media (min-width: 768px) { ... }

                lg: '1280px',
                // => @media (min-width: 1024px) { ... }

                xl: '1440px',
                // => @media (min-width: 1280px) { ... }

                '2xl': '1920px'
            },
            colors: {
                white: {
                    light: '#e4e8f7',
                    base: '#fff'
                },
                dark: {
                    light: 'rgb(34 42 59)',
                    base: '#0f172a',
                    extra: 'rgb(7 13 26)'
                }
            }
        }
    },
    plugins: [
        plugin(function ({ addComponents }) {
            addComponents({
                '.btn': {
                    padding: '.5rem 1rem',
                    borderRadius: '.25rem',
                    fontWeight: '600'
                },
                '.btn-blue': {
                    backgroundColor: '#3490dc',
                    color: '#fff',
                    '&:hover': {
                        backgroundColor: '#2779bd'
                    }
                },
                '.btn-red': {
                    backgroundColor: '#e3342f',
                    color: '#fff',
                    '&:hover': {
                        backgroundColor: '#cc1f1a'
                    }
                }
            })
        })
    ]
}
