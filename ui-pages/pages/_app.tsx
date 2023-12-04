import '@/styles/globals.css'
import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';

import type {AppProps} from 'next/app'
import {createTheme, MantineProvider} from '@mantine/core';
import {Notifications} from '@mantine/notifications';


const theme = createTheme({});

export default function App({Component, pageProps}: AppProps) {
    return (
        <MantineProvider theme={theme}>
            <Notifications/>
            <Component {...pageProps} />
        </MantineProvider>
    );
}
