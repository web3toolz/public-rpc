import type {GatsbyConfig, PluginRef} from "gatsby";


const plugins: PluginRef[] = [
    "gatsby-plugin-postcss",
    "gatsby-plugin-sitemap",
    {
        resolve: 'gatsby-plugin-manifest',
        options: {
            "icon": "src/images/icon.png"
        }
    },
    "gatsby-plugin-mdx",
    {
        resolve: 'gatsby-source-filesystem',
        options: {
            "name": "pages",
            "path": "./src/pages/"
        },
        __key: "pages"
    },
    {
        resolve: "gatsby-plugin-react-svg",
        options: {
            rule: {
                include: /assets/
            }
        }
    },
]

const gtagId = process.env.GATSBY_GTAG_ID;

if (gtagId) {
    plugins.push({
        resolve: `gatsby-plugin-google-gtag`,
        options: {
            trackingIds: [gtagId],
            pluginConfig: {
                head: true
            },
        },
    },)
}

const config: GatsbyConfig = {
    siteMetadata: {
        title: `Public RPC | Web3toolz`,
        siteUrl: `https://publicrpc.web3toolz.com`,
    },
    plugins: plugins,
};

export default config;
