import daisyuiPlugin from 'daisyui'

/** @type {import('tailwindcss').Config} */
export default {
    darkMode: ['class', '[data-theme="dark"]'],
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            fontFamily: {
                'clash-grotesk': 'ClashGrotesk-Medium',
            },
            fontSize: {
                'h0': ['6.25rem', {
                    lineHeight: '5.313rem'
                }],
                'h4': ['1.375rem', {
                    lineHeight: '1.75rem'
                }],
            },
            backgroundImage: {
                'footer': "url('/bg/footer-gradient.svg')",
                'general': "url('/bg/general-gradient.svg')",
            },
            keyframes: {
                'cell-pulse-animation': {
                    '0%': { opacity: '0' },
                    '50%': { opacity: '1' },
                    '100%': { opacity: '0' },
                },
                'cell-pulse-negative-animation': {
                    '0%': { opacity: '1' },
                    '50%': { opacity: '0' },
                    '100%': { opacity: '1' },
                },
            },
            animation: {
                'cell-pulse-3': 'cell-pulse-animation 3s ease-in infinite',
                'cell-pulse-5': 'cell-pulse-animation 5s ease-in infinite',
                'cell-pulse-7': 'cell-pulse-animation 7s ease-in infinite',
                'cell-pulse-negative-3': 'cell-pulse-negative-animation 3s ease-in infinite',
                'cell-pulse-negative-5': 'cell-pulse-negative-animation 5s ease-in infinite',
                'cell-pulse-negative-7': 'cell-pulse-negative-animation 7s ease-in infinite',
            },
            colors: {
                /***************
                 * Base colors *
                 ***************/

                blue: {
                    0: '#FFFFFF',
                    5: '#F3F8FF',
                    10: '#E7F1FF',
                    50: '#C1DFFF',
                    100: '#8DC4FF',
                    200: '#5AAAFF',
                    300: '#2C8FFF',
                    400: '#006AFF',
                    500: '#0052CC',
                    600: '#003E99',
                    700: '#002966',
                    800: '#001833',
                    900: '#000C0D',
                    1000: '#050912',
                },

                grey: {
                    0: '#FFFFFF',
                    5: '#FAFAFA',
                    10: '#F3F3F3',
                    50: '#E3E3E3',
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
                        accent: 'var(--primary-interactive-accent)',
                        hover: 'var(--primary-interactive-hover)',
                    },
                    border: {
                        DEFAULT: 'var(--primary-border)',
                        dark: 'var(--primary-border-dark)',
                        hover: 'var(--primary-border-hover)',
                        accent: 'var(--primary-border-accent)',
                    },

                    base: {
                        content: 'var(--primary-base-content)',
                        background: 'var(--primary-base-background)',
                    },
                },

                secondary: {
                    DEFAULT: 'var(--secondary-brand)',
                    brand: 'var(--secondary-brand)',
                    content: 'var(--secondary-content)',
                    icon: 'var(--secondary-icon)',
                    interactive: {
                        accent: 'var(--primary-interactive-accent)',
                        hover: 'var(--secondary-interactive-hover)',
                    },
                },

                tertiary: {
                    content: 'var(--tertiary-content)',
                    interactive: {
                        accent: 'var(--tertiary-interactive-accent)',
                    },
                },

                positive: {
                    sentiment: 'var(--positive-sentiment)',
                    background: 'var(--positive-background)',
                },

                negative: {
                    sentiment: 'var(--negative-sentiment)',
                    background: 'var(--negative-background)',
                },

                warning: {
                    sentiment: 'var(--warning-sentiment)',
                    background: 'var(--warning-background)',
                },

                dialog: {
                    background: 'var(--dialog-background)',
                    interactive: {
                        disabled: 'var(--dialog-dialog-interactive-disabled)',
                    },
                },




                'elevated-background': 'var(--elevated-background)',
                'neutral-background': 'var(--neutral-background)',
                'overlay-background': 'var(--overlay-background)',
                'divider-border': 'var(--divider-border)',

                'grey-500-10': 'var(--grey-500-10)',
                'grey-500-20': 'var(--grey-500-20)',

                'text-dark': 'var(--text-dark)',
                'text-light': 'var(--text-light)',


                'interactive-primary-pink': 'var(--interactive-primary-pink)',
                'interactive-primary-accent': 'var(--interactive-primary-accent)',
                'interactive-secondary': 'var(--interactive-secondary)',
                'interactive-tertiary': 'var(--interactive-tertiary)',
                'interactive-accent': 'var(--interactive-accent)',


                'content-primary': 'var(--content-primary)',
                'content-secondary': 'var(--content-secondary)',
                'content-tertiary': 'var(--content-tertiary)',
                'content-link-primary': 'var(--content-link-primary)',
                'content-link-hover': 'var(--content-link-hover)',

                'background-primary': 'var(--background-primary)',
                'background-neutral': 'var(--background-neutral)',
                'background-elevated': 'var(--background-elevated)',

                'icon-primary': 'var(--icon-primary)',
                'icon-secondary': 'var(--icon-secondary)',


                'border-divider-default': 'var(--border-divider-default)',


            },
        },
    },

    plugins: [daisyuiPlugin, require('tailwindcss-image-rendering')()],

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
                    // '--rounded-box': '0.625rem', // 10px

                    '--primary-brand': '#C8047D', // pink-500
                    '--primary-content': '#F3F3F3', // grey-10
                    '--primary-link': '#FF6FC8', // pink-200
                    '--primary-link-hover': '#FFC6E9', // pink-50
                    '--primary-icon': '#CACBCE', // grey-100
                    '--primary-background': '#0B101B', // grey-900
                    '--primary-interactive': '#C8047D', // pink-500
                    '--primary-interactive-accent': '#E81899', // pink-400
                    '--primary-interactive-hover': '#E81899', // pink-400
                    '--primary-border-hover': '#FF6FC8', // pink-200
                    '--primary-border-dark': '#5D636F', // grey-500
                    '--primary-border-accent': '#E81899', // pink-400
                    '--primary-base-background': '#FFFFFF', // grey-0
                    '--primary-base-content': '#191E28', // grey-800

                    '--secondary-brand': '#E81899', // pink-400
                    '--secondary-content': '#ADB1B8', // grey-200
                    '--secondary-icon': '#2B303B', // grey-700

                    '--secondary-interactive-accent': '#2B303B', // grey-700
                    '--secondary-interactive-hover': '#ADB1B8', // grey-200

                    '--tertiary-content': '#5D636F', // grey-500
                    '--tertiary-interactive-accent': '#5D636F', // grey-500
                    '--tertiary-interactive-hover': '#444A55', // grey-600

                    '--positive-sentiment': '#47E0A0', // green-300
                    '--positive-background': '#00321D', // green-800

                    '--negative-sentiment': '#F15C5D', // red-300
                    '--negative-background': '#440000', // red-800

                    '--warning-sentiment': '#EBB222', // yellow-400
                    '--warning-background': '#382800', // yellow-800

                    '--elevated-background': '#191E28', // grey-800
                    '--neutral-background': '#2B303B', // grey-700
                    '--neutral-content': '#2B303B', // grey-800
                    '--neutral-accent': '#2B303B', // grey-700
                    '--overlay-background': 'rgba(12, 17, 28, 0.5)', // grey-900|50%
                    '--overlay-dialog': 'rgba(12, 17, 28, 0.90)', // grey-900|90%
                    '--divider-border': '#444A55', // grey-600
                    '--dialog-background': '#2B303B', // grey-700
                    '--dialog-dialog-interactive-disabled': '#444A55', // grey-600


                    // ==Taikoons Color Customizations==//
                    '--grey-500-10': 'rgba(93, 99, 111, 0.1)', // grey-500, 10% opacity
                    '--grey-500-20': 'rgba(93, 99, 111, 0.2)', // grey-500, 20% opacity

                    '--text-dark': '#f3f3f3',
                    '--text-light': '#444A55',// grey-600

                    '--neutral': '#2B303B', // grey-700

                    // figma's theme
                    '--interactive-primary-pink': '#C8047D', // pink 500
                    '--interactive-primary-accent': '#E81899', // pink-400
                    '--interactive-secondary': '#2b303b', // grey-700
                    '--interactive-tertiary': '#444a55', // grey-600
                    '--interactive-accent': '#5D636F', // grey-500

                    '--content-primary': '#F3F3F3', // grey-10
                    '--content-secondary': '#ADB1B8', // grey-200
                    '--content-tertiary': '#5D636F', // grey-500
                    '--content-link-primary': '#FF6FC8', // pink-200
                    '--content-link-hover': '#FFC6E9', // pink-50

                    '--border-divider-default': '#444A55', // grey-600

                    '--background-primary': '#0B101B', // grey-900
                    '--background-neutral': '#2B303B', // grey-700
                    '--background-elevated': '#191E28', // grey-800

                    '--icon-primary': '#CACBCE', // grey-100
                    '--icon-secondary': '#2B303B', // grey-700
                    // ================================ //

                    primary: '#C8047D', // pink-500,
                    'primary-focus': '#E81899', // pink-400
                    'primary-content': '#F3F3F3', // grey-10

                    secondary: '#E81899', // pink-400
                    // 'secondary-focus': '',
                    'secondary-content': '#ADB1B8', // grey-200

                    neutral: '#2B303B', // grey-700
                    'neutral-focus': '#444A55', // grey-600
                    'neutral-content': '#F3F3F3', // grey-10

                    'base-100': '#0B101B', // grey-900
                    // 'base-200': '',
                    // 'base-300': '',
                    'base-content': '#F3F3F3', // grey-10

                    success: '#00321D', // green-800
                    'success-content': '#47E0A0', // green-300
                    error: '#440000', // red-800
                    'error-content': '#F15C5D', // red-300
                    warning: '#382800', // yellow-800
                    'warning-content': '#EBB222', // yellow-400
                },

                light: {
                    'color-scheme': 'light',
                    '--btn-text-case': 'capitalize',

                    '--primary-brand': '#C8047D', // pink-500
                    '--primary-content': '#191E28', // grey-800
                    '--primary-link': '#C8047D', // pink-500
                    '--primary-link-hover': '#E81899', // pink-400
                    '--primary-icon': '#5D636F', // grey-500
                    '--primary-background': '#FAFAFA', // grey-5
                    '--primary-interactive': '#C8047D', // pink-500
                    '--primary-interactive-accent': '#E81899', // pink-400
                    '--primary-interactive-hover': '#E3E3E3', //grey-50
                    '--primary-border-hover': '#FF6FC8', // pink-200
                    '--primary-border-accent': '#E81899', // pink-400

                    // TODO: these two are yet to be decided
                    '--primary-base-background': '#FFFFFF', // grey-0
                    '--primary-base-content': '#191E28', // grey-800

                    '--secondary-brand': '#E81899', // pink-400
                    '--secondary-content': '#444A55', // grey-600
                    '--secondary-icon': '#2B303B', // grey-700
                    '--secondary-interactive-accent': '#E3E3E3', // grey-50
                    '--secondary-interactive-hover': '##F3F3F3', // grey-10

                    '--tertiary-content': '#91969F', // grey-300

                    // TODO: these two are missing. Remain the same as dark theme
                    '--tertiary-interactive-hover': '#444A55', // grey-600
                    '--tertiary-interactive-accent': '#5D636F', // grey-500

                    '--positive-sentiment': '#005E36', // green-700
                    '--positive-background': '#BFFFE4', // green-50

                    '--negative-sentiment': '#BB1A1B', // red-600
                    '--negative-background': '#FFE7E7', // red-10

                    '--warning-sentiment': '#775602', // yellow-700
                    '--warning-background': '#FFF6DE', // yellow-10

                    '--elevated-background': '#e3e3e3', //#FAFAFA', // grey-5
                    '--neutral-background': '#FFFFFF', //  grey-0
                    '--neutral-content': '#191E28', // grey-800
                    '--neutral-accent': '#e3e3e3', // grey-50
                    '--overlay-background': 'rgba(12, 17, 28, 0.2)', // grey-900|20%
                    '--overlay-dialog': 'rgba(12, 17, 28, 0.9)', // grey-900|20



                    '--dialog-background': '#FFFFFF', // grey-0
                    '--dialog-dialog-interactive-disabled': '#E3E3E3', // grey-50

                    '--divider-border': '#CACBCE', // grey-100
                    // ==Taikoons Color Customizations==//

                    '--grey-500-10': 'rgba(250,250,250,0.5)',
                    '--grey-500-20': 'rgba(250,250,250,0.5)',
                    '--text-dark': '#191e28',
                    '--text-light': '#91969f',
                    '--neutral': '#E3E3E3', // grey-50

                    // figma's theme
                    '--interactive-primary-pink': '#C8047D', // pink 500
                    '--interactive-primary-accent': '#E81899', // pink-400
                    '--interactive-secondary': '#f3f3f3', // grey-10
                    '--interactive-tertiary': '#e3e3e3', // grey-50
                    '--interactive-accent': '#cacbce', // grey-100

                    '--content-primary': '#191E28', // grey-800
                    '--content-secondary': '#444A55', // grey-600
                    '--content-tertiary': '#91969F', // grey-300
                    '--content-link-primary': '#C8047D', // pink-500
                    '--content-link-hover': '#E81899', // pink-400

                    '--border-divider-default': '#CACBCE', // grey-100

                    '--background-primary': '#ffffff', // grey-5
                    '--background-neutral': '#F8f8f8', // grey-50
                    '--background-elevated': '#ffffff', // grey-5

                    '--icon-primary': '#5D636F', // grey-500
                    '--icon-secondary': '#e3e3e3', // grey-50
                    // ================================ //

                    primary: '#C8047D', // pink-500,
                    'primary-focus': '#E81899', // pink-400
                    'primary-content': '#191E28', // grey-800

                    secondary: '#E81899', // pink-400
                    // 'secondary-focus': '',
                    'secondary-content': '#444A55', // grey-600

                    neutral: '#E3E3E3', // grey-50
                    'neutral-focus': '#CACBCE', // grey-100
                    'neutral-content': '#191E28', // grey-800

                    'base-100': '#FAFAFA', // grey-5
                    // 'base-200': '',
                    // 'base-300': '',
                    'base-content': '#191E28', // grey-800

                    success: '#BFFFE4', // green-50
                    'success-content': '#005E36', // green-700
                    error: '#FFE7E7', // red-10
                    'error-content': '#BB1A1B', // red-600
                    warning: '#FFF6DE', // yellow-10
                    'warning-content': '#775602', // yellow-700
                },
            },
        ],
    },
}
