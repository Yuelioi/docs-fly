import plugin from 'tailwindcss/plugin'
import typography from '@tailwindcss/typography'
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
        container: {
            center: true,
            padding: {
                DEFAULT: '1rem',
                sm: '2rem',
                lg: '4rem',
                xl: '5rem',
                '2xl': '6rem'
            }
        },
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
                theme: {
                    'text-base': 'var(--color-text-base)',
                    'text-muted': 'var(--color-text-muted)',
                    'text-inverse': 'var(--color-text-inverse)',

                    primary: 'var(--color-primary)',
                    secondary: 'var(--color-secondary)',
                    error: 'var(--color-error)',
                    warn: 'var(--color-warn)',
                    success: 'var(--color-success)',
                    info: 'var(--color-info)'
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

                    primary: {
                        base: 'var(--color-bg-primary-base)',
                        hover: 'var(--color-bg-primary-hover)'
                    },
                    secondary: {
                        base: 'var(--color-bg-secondary-base)',
                        hover: 'var(--color-bg-secondary-hover)'
                    },
                    error: {
                        base: 'var(--color-bg-error-base)',
                        hover: 'var(--color-bg-error-hover)'
                    },
                    warn: {
                        base: 'var(--color-bg-warn-base)',
                        hover: 'var(--color-bg-warn-hover)'
                    },
                    success: {
                        base: 'var(--color-bg-success-base)',
                        hover: 'var(--color-bg-success-hover)'
                    },
                    info: {
                        base: 'var(--color-bg-info-base-base)',
                        hover: 'var(--color-bg-info-hover)'
                    },

                    tooltip: 'var(--color-tooltip-bg)',
                    backToTop: {
                        base: 'var(--color-bg-back-top-base)',
                        hover: 'var(--color-bg-back-top-hover)'
                    }
                }
            },
            borderColor: {
                theme: {
                    base: 'var(--color-border-base)',
                    primary: 'var(--color-primary)',
                    muted: 'var(--color-border-muted)'
                }
            }
        }
    },
    plugins: [
        plugin(function ({ addComponents, addBase, theme }) {
            addBase({})
        }),
        typography
    ]
}
