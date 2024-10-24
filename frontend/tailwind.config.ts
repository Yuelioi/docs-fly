import plugin from 'tailwindcss/plugin'
import typography from '@tailwindcss/typography'
import { Config } from 'tailwindcss'

import { addDynamicIconSelectors } from '@iconify/tailwind'
import daisyui from 'daisyui'

export default {
  darkMode: ['selector', '[data-theme="dark"]'],
  content: ['./index.html', './src/**/*.{js,ts,vue,css}'],
  daisyui: {
    themes: [
      {
        light: {
          primary: '#ff6699',
          secondary: '#00aeec',
          accent: '#a36ffd',
          neutral: '#9facc5',
          'neutral-content': '#0f172a',

          'base-100': '#ffffff',
          'base-200': '#f6f8fa',
          'base-300': '#a6adbb',
          'base-content': '#2d314a',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#ffa200',
          error: '#f53135'
        }
      },
      {
        dark: {
          primary: '#ff6699',
          secondary: '#4ac7ff',
          accent: '#a36ffd',
          neutral: '#22212c',
          'neutral-content': '#f8fafc',

          'base-100': '#16181d',
          'base-200': '#2a2e37',
          'base-300': '#2f323c',
          'base-content': '#f8fafc',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#ffa200',
          error: '#f53135'
        }
      }
    ]
  },
  theme: {
    container: {
      center: true,
      padding: {
        DEFAULT: '1rem',
        sm: '2rem',
        md: '2rem',
        lg: '4rem',
        xl: '5rem',
        '2xl': '6rem'
      },
      screens: {
        sm: '100%',
        md: '100%',
        lg: '100%',
        xl: '100%',
        '2xl': '100%'
      }
    },
    extend: {
      screens: {
        xs: '576px',
        sm: '720px',
        md: '960px',
        lg: '1280px',
        xl: '1440px',
        '2xl': '1920px'
      }
    }
  },
  plugins: [
    plugin(function ({ addBase }) {
      addBase({})
    }),
    typography,
    daisyui,
    addDynamicIconSelectors()
  ]
} satisfies Config
