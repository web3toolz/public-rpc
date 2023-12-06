import '@/styles/globals.css'
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';

import type {AppProps} from 'next/app'
import {createTheme, MantineProvider, MantineColorsTuple} from '@mantine/core';
import {Notifications} from '@mantine/notifications';
import React from "react";
import Script from "next/script";

const gtagId: string | undefined = process.env.NEXT_PUBLIC_GTAG_ID;
function gtagScript(): React.ReactElement {
    if (gtagId) {
        return (
            <>
                <Script id="gtag-source" async src={`https://www.googletagmanager.com/gtag/js?id=${gtagId}`}></Script>
                <Script id="gtag-content">
                    {`
                    window.dataLayer = window.dataLayer || [];
                    function gtag(){dataLayer.push(arguments);}
                    gtag('js', new Date());
                    gtag('config', '${gtagId}');
                `}
                </Script>
            </>
        )
    }
    return <></>;
}

const myColor: MantineColorsTuple = [
    "#ecf5ff",
    "#dee5f3",
    "#bcc9df",
    "#98accc",
    "#7992ba",
    "#6582b2",
    "#5a7aae",
    "#4a6999",
    "#405d8a",
    "#32507c"
];

const theme = createTheme({
    colors: {
        myColor,
    }
});
export default function App({Component, pageProps}: AppProps) {
    return (
        <MantineProvider theme={theme}>
            <Notifications/>
            <Component {...pageProps} />
            {gtagScript()}
        </MantineProvider>
    );
}
