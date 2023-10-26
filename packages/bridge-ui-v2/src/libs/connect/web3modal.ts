import { createWeb3Modal } from '@web3modal/wagmi';

import { PUBLIC_WALLETCONNECT_PROJECT_ID } from '$env/static/public';
import { chains, getChainImages } from '$libs/chain';
import { wagmiConfig } from '$libs/wagmi';

const projectId = PUBLIC_WALLETCONNECT_PROJECT_ID;
const chainImages = getChainImages();

export const web3modal = createWeb3Modal({
  wagmiConfig,
  projectId,
  chains,
  chainImages,
  themeVariables: {
    '--w3m-font-family': '"Public Sans", sans-serif',
    '--w3m-accent': 'var(--primary-brand)',

    // Body small regular
    // @ts-ignore
    '--w3m-text-small-regular-line-height': '20px',

    // Body regular
    // @ts-ignore
    '--w3m-text-medium-regular-size': '16px',
    '--w3m-text-medium-regular-weight': '400',
    '--w3m-text-medium-regular-line-height': '24px',
    '--w3m-text-medium-regular-letter-spacing': 'normal',

    // Title body bold
    // @ts-ignore
    '--w3m-text-big-bold-size': '18px',
    '--w3m-text-big-bold-weight': '700',
    '--w3m-text-big-bold-line-height': '24px',

    // @ts-ignore
    '--w3m-background-color': 'var(--primary-brand)',
    '--w3m-overlay-background-color': 'var(--overlay-background)',
    '--w3m-background-border-radius': '20px',
    '--w3m-container-border-radius': '0',

    // Unofficial variables
    // @ts-ignore
    '--w3m-color-fg-1': 'var(--primary-content)',
    '--w3m-color-bg-1': 'var(--primary-background)',
    '--w3m-color-bg-2': 'var(--neutral-background)',
    '--w3m-color-overlay': 'none',
    '--w3m-accent-fill-color': 'var(--dark-background)',
  },
  themeMode: (localStorage.getItem('theme') as 'dark' | 'light') ?? 'dark',
});
