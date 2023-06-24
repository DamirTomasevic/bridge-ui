import daisyuiPlugin from 'daisyui';

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ['class', '[data-theme]'],
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        /***************
         * Base colors *
         ***************/

        grey: {
          0: '#FFFFFF',
          5: '#FAFAFA',
          10: '#F3F3F3',
          50: '#E7E7E7',
          100: '#CACBCE',
          200: '#ADB1B8',
          300: '#91969F',
          400: '#767C89',
          500: '#5D636F',
          600: '#444A55',
          700: '#2B303B',
          800: '#191E28',
          900: '#0B101B',
          1000: '#050912',
        },

        pink: {
          0: '#FFFFFF',
          5: '#FFF8FC',
          10: '#FFE7F6',
          50: '#FFC6E9',
          100: '#FF98D8',
          200: '#FF6FC8',
          300: '#FF40B6',
          400: '#E81899',
          500: '#C8047D',
          600: '#9A0060',
          700: '#7D004E',
          800: '#4B002F',
          900: '#240017',
          1000: '#050912',
        },

        red: {
          0: '#FFFFFF',
          5: '#FEF5F5',
          10: '#FFE7E7',
          50: '#FFC5C5',
          100: '#FF9B9C',
          200: '#FD7576',
          300: '#F15C5D',
          400: '#DB4546',
          500: '#CE2C2D',
          600: '#BB1A1B',
          700: '#790102',
          800: '#440000',
          900: '#250000',
          1000: '#050912',
        },

        green: {
          0: '#FFFFFF',
          5: '#F2FFFA',
          10: '#E4FFF4',
          50: '#BFFFE4',
          100: '#89FFCD',
          200: '#65F0B6',
          300: '#47E0A0',
          400: '#2DCA88',
          500: '#19BA76',
          600: '#059458',
          700: '#005E36',
          800: '#00321D',
          900: '#001C10',
          1000: '#050912',
        },

        yellow: {
          0: '#FFFFFF',
          5: '#FFFCF3',
          10: '#FFF6DE',
          50: '#FFEAB5',
          100: '#FFDC85',
          200: '#FFCF55',
          300: '#F8C23B',
          400: '#EBB222',
          500: '#DBA00D',
          600: '#C28B00',
          700: '#775602',
          800: '#382800',
          900: '#201700',
          1000: '#050912',
        },

        /*******************
         * Semantic colors *
         *******************/

        primary: {
          DEFAULT: 'var(--primary-brand)',
          brand: 'var(--primary-brand)',
          content: 'var(--primary-content)',
          link: {
            DEFAULT: 'var(--primary-link)',
            hover: 'var(--primary-link-hover)',
          },
          icon: 'var(--primary-icon)',
          background: 'var(--primary-background)',
          interactive: {
            DEFAULT: 'var(--primary-interactive)',
            hover: 'var(--primary-interactive-accent)',
            accent: 'var(--primary-interactive-accent)',
          },
        },

        secondary: {
          DEFAULT: 'var(--secondary-brand)',
          brand: 'var(--secondary-brand)',
          content: 'var(--secondary-content)',
          icon: 'var(--secondary-icon)',
          interactive: {
            hover: 'var(--secondary-interactive-hover)',
          },
        },

        tertiary: {
          content: 'var(--tertiary-content)',
          interactive: {
            hover: 'var(--tertiary-interactive-hover)',
          },
        },

        'positive-sentiment': 'var(--positive-sentiment)',
        'negative-sentiment': 'var(--negative-sentiment)',
        'warning-sentiment': 'var(--warning-sentiment)',

        'elevated-background': 'var(--elevated-background)',
        'neutral-background': 'var(--neutral-background)',
        'overlay-background': 'var(--overlay-background)',
      },
    },
  },

  plugins: [daisyuiPlugin],

  // https://daisyui.com/docs/config/
  daisyui: {
    darkTheme: 'dark', // name of one of the included themes for dark mode
    base: true, // applies background color and foreground color for root element by default
    styled: true, // include daisyUI colors and design decisions for all components
    utils: true, // adds responsive and modifier utility classes
    rtl: false, // rotate style direction from left-to-right to right-to-left. You also need to add dir="rtl" to your html tag and install `tailwindcss-flip` plugin for Tailwind CSS.
    prefix: '', // prefix for daisyUI classnames (components, modifiers and responsive class names. Not colors)
    logs: false, // Shows info about daisyUI version and used config in the console when building your CSS
    themes: [
      {
        dark: {
          'color-scheme': 'dark',
          '--btn-text-case': 'capitalize',

          '--primary-brand': '#C8047D', // pink-500
          '--primary-content': '#F3F3F3', // grey-10
          '--primary-link': '#FF6FC8', // pink-200
          '--primary-link-hover': '#FFC6E9', // pink-50
          '--primary-icon': '#CACBCE', // grey-100
          '--primary-background': '#0B101B', // grey-900
          '--primary-interactive': '#C8047D', // pink-500
          '--primary-interactive-accent': '#E81899', // pink-400

          '--secondary-brand': '#E81899', // pink-400
          '--secondary-content': '#ADB1B8', // grey-200
          '--secondary-icon': '#444A55', // grey-600
          '--secondary-interactive-hover': '#2B303B', // grey-700

          '--tertiary-content': '#5D636F', // grey-500
          '--tertiary-interactive-hover': '#444A55', // grey-600
          '--tertiary-interactive-accent': '#5D636F', // grey-500

          '--positive-sentiment': '#47E0A0', // green-300
          '--negative-sentiment': '#F15C5D', // red-300
          '--warning-sentiment': '#EBB222', // yellow-400

          '--elevated-background': '#191E28', // grey-800
          '--neutral-background': '#2B303B', // grey-700
          '--overlay-background': 'rgba(12, 17, 28, 0.5)', // grey-900|50%

          // ================================ //

          primary: '#C8047D', // pink-500,
          'primary-focus': '#E81899', // pink-400
          'primary-content': '#F3F3F3', // grey-10

          neutral: '#2B303B', // grey-700
          'neutral-focus': '#444A55', // grey-600
          'neutral-content': '#F3F3F3', // grey-10

          'base-100': '#0B101B', // grey-900
          'base-content': '#F3F3F3', // grey-10

          success: '#00321D', // green-800
          error: '#440000', // red-800
          warning: '#382800', // yellow-800
        },

        light: {
          // TODO: add light theme

          'color-scheme': 'light',
          '--btn-text-case': 'capitalize',

          '--primary-content': '#191E28', // grey-800
          '--primary-background': '#FAFAFA', // grey-5

          // ================================ //

          'base-100': '#FAFAFA', // grey-5
          'base-content': '#191E28', // grey-800

          success: '#E4FFF4', // green-10
          error: '#FFE7E7', // red-10
          warning: '#FFF6DE', // yellow-10
        },
      },
    ],
  },
};
