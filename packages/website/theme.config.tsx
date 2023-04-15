import { useRouter } from "next/router";
import { useConfig } from "nextra-theme-docs";
import Footer from "./components/Footer";
import ThemedImage from "./components/ThemedImage";

export default {
  banner: {
    key: "banner",
    text: (
      <a
        href="https://taiko.xyz/docs/reference/ethglobal-tokyo-reference"
        target="_blank"
      >
        📌 Please use the <u>ETHTokyo network config</u> for the hackathon →
      </a>
    ),
  },
  chat: {
    link: "https://discord.gg/taikoxyz",
  },
  darkMode: true,
  docsRepositoryBase:
    "https://github.com/taikoxyz/taiko-mono/blob/main/packages/website",
  editLink: {
    text: "Improve this page on GitHub ↗",
  },
  feedback: {
    content: null,
  },
  footer: {
    component: Footer,
  },
  head: () => {
    const { asPath } = useRouter();
    const { frontMatter } = useConfig();
    return (
      <>
        <meta property="og:url" content={`https://taiko.xyz${asPath}`} />
        <meta property="og:title" content={frontMatter.title || "Taiko"} />
        <meta
          property="og:description"
          content={
            frontMatter.description ||
            "A decentralized, Ethereum-equivalent ZK-Rollup."
          }
        />
        <link rel="icon" href="/images/favicon.png" />
      </>
    );
  },
  logo: <ThemedImage />,
  nextThemes: {
    defaultTheme: "light",
  },
  primaryHue: 315,
  project: {
    link: "https://github.com/taikoxyz",
  },
  useNextSeoProps() {
    return {
      titleTemplate: "%s – Taiko",
    };
  },
};
