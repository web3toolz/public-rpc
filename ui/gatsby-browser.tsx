import "./src/styles/global.css"
import '@mantine/core/styles.css';
import * as React from 'react';
import {MantineProvider} from '@mantine/core';
import {theme} from './src/theme';

export const wrapPageElement = ({element}) => {
    return <MantineProvider theme={theme}>{element}</MantineProvider>;
};
