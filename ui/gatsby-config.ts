import type {GatsbyConfig, PluginRef} from "gatsby";


const siteTitle: string = `Public RPC for EVM and non-EVM chains | Web3toolz`
const siteDescription: string = `Explore the largest collection of free, public RPC endpoints with broad blockchain support.`


const plugins: PluginRef[] = [
    "gatsby-plugin-postcss",
    "gatsby-plugin-sitemap",
    {
        resolve: 'gatsby-plugin-manifest',
        options: {
            name: siteTitle,
            short_name: "Public RPC",
            start_url: `/`,
            icon: "src/images/icon.png"
        }
    },
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
    {
        resolve: `gatsby-omni-font-loader`,
        options: {
            enableListener: true,
            preconnect: [`https://fonts.googleapis.com`, `https://fonts.gstatic.com`],
            web: [
                {
                    name: `Open Sans`,
                    file: `https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300;0,400;0,500;0,600;0,700;1,300;1,400;1,500;1,600;1,700&display=swap`,
                },
            ],
        },
    },
]

const gtagId = process.env.GATSBY_GTAG_ID;

console.log(gtagId)

if (gtagId) {
    plugins.push({
        resolve: `gatsby-plugin-google-gtag`,
        options: {

            trackingIds: ["G-T30T3RLVK0"],
            pluginConfig: {
                head: true,
                respectDNT: false,
            },
        },
    },)
}

const config: GatsbyConfig = {
    siteMetadata: {
        title: siteTitle,
        description: siteDescription,
        siteUrl: `https://publicrpc.web3toolz.com`,
    },
    plugins: plugins,
};

export default config;
