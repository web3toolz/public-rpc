import {Burger, Group} from '@mantine/core';
import {useDisclosure} from '@mantine/hooks';
import * as classes from './navbar.module.css';
import React from 'react';
import {navigate} from "gatsby";

const links = [
    {link: '/new', label: 'Add new RPC'},
    {link: '/contact', label: 'Contact Us'},
];

export function Navbar(): React.ReactElement {
    const [opened, {toggle}] = useDisclosure(false);

    const items = links.map((link) => (
        <a
            key={link.label}
            className={classes.link}
            onClick={() => navigate(link.link)}
        >
            {link.label}
        </a>
    ));

    return (
        <header className={classes.header}>
            <div className={classes.inner}>
                <Group>
                    <Burger opened={opened} onClick={toggle} size="sm" hiddenFrom="sm"/>
                </Group>

                <Group>
                    <Group ml={50} gap={5} visibleFrom="sm">
                        {items}
                    </Group>
                </Group>
            </div>
        </header>
    );
}

export default Navbar;