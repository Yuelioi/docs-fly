import plugin from 'tailwindcss/plugin'

/** @type {import('tailwindcss').Config} */

function withOpacity(variableName) {
    return ({ opacityValue }) => {
        if (opacityValue != null) {
            return `var(${variableName}, ${opacityValue})`
        }
        return `var(${variableName})`
    }
}

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
            },
            textColor: {
                theme: {
                    base: 'var(--color-text-base)',
                    muted: 'var(--color-text-muted)',
                    inverse: 'var(--color-text-inverse)',
                    link: {
                        base: 'var(--color-link-base)',
                        hover: 'var(--color-link-hover)'
                    }
                }
            },
            backgroundColor: {
                theme: {
                    base: 'var(--color-bg-base)',
                    card: 'var(--color-bg-card)',
                    header: 'var(--color-bg-header)',
                    footer: 'var(--color-bg-footer)',
                    button: {
                        primary: 'var(--color-btn-primary)',
                        secondary: 'var(--color-btn-secondary)',
                        danger: 'var(--color-btn-danger)'
                    },
                    tooltip: 'var(--color-tooltip-bg)'
                }
            },
            borderColor: {
                theme: {
                    base: withOpacity('--color-border-base'),
                    muted: withOpacity('--color-border-muted')
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
